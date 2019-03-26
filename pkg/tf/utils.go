package tf

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unicode"

	"github.com/Sirupsen/logrus"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/example"
	tf "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow_serving/apis"
	googleproto "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

const (
	MaxMsgLength = 1024 * 1024 * 64 // 64 MB
)

type FeatureJSON struct {
	Float     *float32   `json:"float,omitempty"`
	FloatList *[]float32 `json:"float_list,omitempty"`
	Int       *int64     `json:"int,omitempty"`
	IntList   *[]int64   `json:"int_list,omitempty"`
	Bytes     *[]byte    `json:"bytes,omitempty"`
	BytesList *[][]byte  `json:"bytes_list,omitempty"`
	BoolList  *[]bool    `json:"bool_list,omitempty"`
}

type InputJSON struct {
	Dtype tf.DataType `json:"dtype,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (t *InputJSON) Tensor() (*tf.TensorProto, error) {
	proto := &tf.TensorProto{
		Dtype:       t.Dtype,
		TensorShape: &tf.TensorShapeProto{},
	}

	err := fillTensor(t.Data, proto, 0)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

type ModelData struct {
	TFFeatures  []map[string]FeatureJSON `json:"features,omitempty"`
	Inputs      map[string]InputJSON     `json:"inputs,omitempty"`
	OutFilter   []string                 `json:"out_filter,omitempty"`
	OutMimeType string                   `json:"out_mime_type,omitempty"`
}

func (f *FeatureJSON) TFFeature() *example.Feature {
	if f.Float != nil {
		return &example.Feature{
			Kind: &example.Feature_FloatList{
				FloatList: &example.FloatList{
					Value: []float32{*f.Float},
				},
			},
		}
	} else if f.FloatList != nil {
		return &example.Feature{
			Kind: &example.Feature_FloatList{
				FloatList: &example.FloatList{
					Value: *f.FloatList,
				},
			},
		}
	} else if f.Int != nil {
		return &example.Feature{
			Kind: &example.Feature_Int64List{
				Int64List: &example.Int64List{
					Value: []int64{*f.Int},
				},
			},
		}
	} else if f.IntList != nil {
		return &example.Feature{
			Kind: &example.Feature_Int64List{
				Int64List: &example.Int64List{
					Value: *f.IntList,
				},
			},
		}
	} else if f.Bytes != nil {
		return &example.Feature{
			Kind: &example.Feature_BytesList{
				BytesList: &example.BytesList{
					Value: [][]byte{*f.Bytes},
				},
			},
		}
	} else if f.BytesList != nil {
		return &example.Feature{
			Kind: &example.Feature_BytesList{
				BytesList: &example.BytesList{
					Value: *f.BytesList,
				},
			},
		}
	}
	return nil
}
func CallTF(ctx context.Context, servingAddr string, model string, version int64, signature string, modelData ModelData) (map[string]interface{}, error) {
	mSpec := &apis.ModelSpec{
		Name:          model,
		SignatureName: signature,
	}
	if version > 0 {
		mSpec.Version = &googleproto.Int64Value{
			Value: version,
		}
	}
	feedData := map[string]*tf.TensorProto{}
	if len(modelData.Inputs) > 0 {
		for n, v := range modelData.Inputs {
			t, err := v.Tensor()
			if err != nil {
				return nil, fmt.Errorf("Invalid input '%s' %v", n, err)
			}
			feedData[n] = t
		}
	} else if len(modelData.TFFeatures) > 0 {
		codec := encoding.GetCodec("proto")
		if codec == nil {
			return nil, fmt.Errorf("Codec for proto not found")
		}
		messages := make([][]byte, 0)
		for _, f := range modelData.TFFeatures {
			tfFeatures := map[string]*example.Feature{}
			for k, v := range f {
				fv := v.TFFeature()
				if fv == nil {
					return nil, fmt.Errorf("value for %s is empty", k)
				}
				tfFeatures[k] = fv
			}
			exp := &example.Example{
				Features: &example.Features{
					Feature: tfFeatures,
				},
			}
			msg, err := codec.Marshal(exp)
			if err != nil {
				return nil, fmt.Errorf("Failed encode fetaure %v", err)
			}
			messages = append(messages, msg)
		}
		feedData["examples"] = &tf.TensorProto{
			StringVal: messages,
			Dtype:     tf.DataType_DT_STRING,
			TensorShape: &tf.TensorShapeProto{
				Dim: []*tf.TensorShapeProto_Dim{
					{
						Size: int64(len(messages)),
					},
				},
			}}
	}
	req := &apis.PredictRequest{
		ModelSpec: mSpec,
		Inputs:    feedData,
	}
	conn, err := grpc.Dial(
		servingAddr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxMsgLength)),
	)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("Failed open grpc connection %v", err)
	}
	client := apis.NewPredictionServiceClient(conn)
	resp, err := client.Predict(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Predict call failed %v", err)
	}
	result := map[string]interface{}{}
	for k, v := range resp.Outputs {
		result[k] = tensor2Go(v)
	}
	return result, nil
}

func tensor2Go(t *tf.TensorProto) interface{} {
	switch t.Dtype {
	case tf.DataType_DT_BOOL:
		var arr interface{} = t.BoolVal
		if t.BoolVal == nil {
			arr = t.TensorContent
		}
		res := makeRes(arr, reflect.TypeOf(true), t.TensorShape.Dim)
		switch shaped := res.(type) {
		case []interface{}:
			return shapeContainer(t.TensorShape.Dim, shaped)
		}
		return res
	case tf.DataType_DT_INT64:
		var arr interface{} = t.Int64Val
		if t.Int64Val == nil {
			arr = t.TensorContent
		}
		res := makeRes(arr, reflect.TypeOf(int64(1)), t.TensorShape.Dim)
		switch shaped := res.(type) {
		case []interface{}:
			return shapeContainer(t.TensorShape.Dim, shaped)
		}
		return res
	case tf.DataType_DT_INT32:
		var arr interface{} = t.IntVal
		if t.IntVal == nil {
			arr = t.TensorContent
		}
		res := makeRes(arr, reflect.TypeOf(int32(1)), t.TensorShape.Dim)
		switch shaped := res.(type) {
		case []interface{}:
			return shapeContainer(t.TensorShape.Dim, shaped)
		}
		return res
	case tf.DataType_DT_FLOAT:
		var arr interface{} = t.FloatVal
		if t.FloatVal == nil {
			arr = t.TensorContent
		}
		res := makeRes(arr, reflect.TypeOf(float32(1)), t.TensorShape.Dim)
		switch shaped := res.(type) {
		case []interface{}:
			for i := range shaped {
				if math.IsInf(float64(mustFloat32(shaped[i])), -1) {
					shaped[i] = float32(-math.MaxFloat32)
				} else if math.IsInf(float64(mustFloat32(shaped[i])), 1) {
					shaped[i] = float32(math.MaxFloat32)
				}
			}
			return shapeContainer(t.TensorShape.Dim, shaped)
		}
		return res
	case tf.DataType_DT_DOUBLE:
		var arr interface{} = t.DoubleVal
		if t.DoubleVal == nil {
			arr = t.TensorContent
		}
		res := makeRes(arr, reflect.TypeOf(float64(1)), t.TensorShape.Dim)
		switch shaped := res.(type) {
		case []interface{}:
			for i := range shaped {
				if math.IsInf(float64(mustFloat32(shaped[i])), -1) {
					shaped[i] = float32(-math.MaxFloat32)
				} else if math.IsInf(float64(mustFloat32(shaped[i])), 1) {
					shaped[i] = float32(math.MaxFloat32)
				}
			}
			return shapeContainer(t.TensorShape.Dim, shaped)
		}
		return res
	case tf.DataType_DT_STRING:
		res := make([]interface{}, len(t.StringVal))
		for i := range res {
			pureString := true
			for j := range t.StringVal[i] {
				if t.StringVal[i][j] > unicode.MaxASCII {
					pureString = false
					break
				}
			}
			if pureString {
				res[i] = string(t.StringVal[i])
			} else {
				res[i] = t.StringVal[i]
			}
		}
		if len(t.TensorShape.Dim) > 0 {
			return shapeContainer(t.TensorShape.Dim, res)
		} else {
			if len(res) > 0 {
				return res[0]
			}
		}
	}
	return nil
}

func makeRes(arr interface{}, type_ reflect.Type, dim []*tf.TensorShapeProto_Dim) interface{} {
	res := make([]interface{}, 0)
	plain := len(dim) < 1

	// First, convert byte array to numeric array if it takes place.
	switch v := arr.(type) {
	case []byte:
		targetValue := reflect.MakeSlice(reflect.SliceOf(type_), len(v)/int(type_.Size()), len(v)/int(type_.Size()))
		targetArr := targetValue.Interface()
		buf := bytes.NewBuffer(v)
		err := binary.Read(buf, binary.LittleEndian, targetArr)
		if err != nil {
			logrus.Warnf("Error decoding bytes to array: %v", err)
		}
		arr = targetArr
	}

	switch v := arr.(type) {
	case []bool:
		for _, el := range v {
			res = append(res, el)
		}
	case []int32:
		for _, el := range v {
			res = append(res, el)
		}
	case []int16:
		for _, el := range v {
			res = append(res, el)
		}
	case []int8:
		for _, el := range v {
			res = append(res, el)
		}
	case []int64:
		for _, el := range v {
			res = append(res, el)
		}
	case []float32:
		for _, el := range v {
			res = append(res, el)
		}
	case []float64:
		for _, el := range v {
			res = append(res, el)
		}
	case []uint:
		for _, el := range v {
			res = append(res, el)
		}
	case []uint16:
		for _, el := range v {
			res = append(res, el)
		}
	case []uint32:
		for _, el := range v {
			res = append(res, el)
		}
	case []uint64:
		for _, el := range v {
			res = append(res, el)
		}
	case []byte:

	}

	if plain && len(res) == 1 {
		return res[0]
	}

	return res
}

func shapeContainer(dim []*tf.TensorShapeProto_Dim, data []interface{}) interface{} {
	if len(dim) < 2 {
		return data
	}
	res := make([]interface{}, 0)
	last := len(dim) - 1
	l := int(dim[last].Size)
	for i := 0; i < len(data); i += l {
		res = append(res, data[i:i+l])
	}
	if last > 1 {
		return shapeContainer(dim[0:last], res)
	}
	return res
}

func fillBaseTensor(data interface{}, proto *tf.TensorProto) error {
	switch v := data.(type) {
	case float32:
		return addFloat32(proto.Dtype, proto, v)
	case float64:
		return addFloat64(proto.Dtype, proto, v)
	case int8:
		return addInt64(proto.Dtype, proto, int64(v))
	case int16:
		return addInt64(proto.Dtype, proto, int64(v))
	case int32:
		return addInt64(proto.Dtype, proto, int64(v))
	case int:
		return addInt64(proto.Dtype, proto, int64(v))
	case int64:
		return addInt64(proto.Dtype, proto, int64(v))
	case uint8:
		return addInt64(proto.Dtype, proto, int64(v))
	case uint16:
		return addInt64(proto.Dtype, proto, int64(v))
	case uint32:
		return addInt64(proto.Dtype, proto, int64(v))
	case uint64:
		return addInt64(proto.Dtype, proto, int64(v))
	case uint:
		return addInt64(proto.Dtype, proto, int64(v))
	case []byte:
		return addBytes(proto.Dtype, proto, v)
	case string:
		return addString(proto.Dtype, proto, v)
	case bool:
		return addBool(proto.Dtype, proto, v)
	}

	return fmt.Errorf("Unsupported type: %v", reflect.TypeOf(data))
}

func fillTensor(data interface{}, proto *tf.TensorProto, nesting int) error {
	switch v := data.(type) {
	case []interface{}:
		if len(proto.TensorShape.Dim) <= nesting {
			proto.TensorShape.Dim = append(proto.TensorShape.Dim, &tf.TensorShapeProto_Dim{
				Size: int64(len(v)),
			})
		}
		for _, v1 := range v {
			switch v2 := v1.(type) {
			case []interface{}:
				_ = fillTensor(v2, proto, nesting+1)
			default:
				if err := fillBaseTensor(v2, proto); err != nil {
					return err
				}
			}
		}
	default:
		return fillBaseTensor(v, proto)
	}
	return nil
}

func addInt64(mtype tf.DataType, proto *tf.TensorProto, v int64) error {
	switch mtype {
	case tf.DataType_DT_DOUBLE:
		proto.DoubleVal = append(proto.DoubleVal, float64(v))
	case tf.DataType_DT_FLOAT:
		proto.FloatVal = append(proto.FloatVal, float32(v))
	case tf.DataType_DT_INT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT32:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT64:
		proto.Int64Val = append(proto.Int64Val, int64(v))
	case tf.DataType_DT_UINT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_UINT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	default:
		return fmt.Errorf("can't convert int64 to tf:%v", mtype)
	}
	return nil
}

func addFloat64(mtype tf.DataType, proto *tf.TensorProto, v float64) error {
	switch mtype {
	case tf.DataType_DT_DOUBLE:
		proto.DoubleVal = append(proto.DoubleVal, float64(v))
	case tf.DataType_DT_FLOAT:
		proto.FloatVal = append(proto.FloatVal, float32(v))
	case tf.DataType_DT_INT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT32:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT64:
		proto.Int64Val = append(proto.Int64Val, int64(v))
	case tf.DataType_DT_UINT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_UINT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	default:
		return fmt.Errorf("can't convert float64 to tf:%v", mtype)
	}
	return nil
}

func addFloat32(mtype tf.DataType, proto *tf.TensorProto, v float32) error {
	switch mtype {
	case tf.DataType_DT_DOUBLE:
		proto.DoubleVal = append(proto.DoubleVal, float64(v))
	case tf.DataType_DT_FLOAT:
		proto.FloatVal = append(proto.FloatVal, float32(v))
	case tf.DataType_DT_INT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT32:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_INT64:
		proto.Int64Val = append(proto.Int64Val, int64(v))
	case tf.DataType_DT_UINT8:
		proto.IntVal = append(proto.IntVal, int32(v))
	case tf.DataType_DT_UINT16:
		proto.IntVal = append(proto.IntVal, int32(v))
	default:
		return fmt.Errorf("can't convert float32 to tf:%v", mtype)
	}
	return nil
}

func addBool(mtype tf.DataType, proto *tf.TensorProto, v bool) error {
	switch mtype {
	case tf.DataType_DT_BOOL:
		proto.BoolVal = append(proto.BoolVal, v)
	default:
		return fmt.Errorf("can't convert bool to tf:%v", mtype)
	}
	return nil
}

func addString(mtype tf.DataType, proto *tf.TensorProto, v string) error {
	switch mtype {
	case tf.DataType_DT_STRING:
		//proto.TensorShape.Dim = append(proto.TensorShape.Dim, &tf.TensorShapeProto_Dim{
		//	Size: 1,
		//})
		bts, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			proto.StringVal = append(proto.StringVal, []byte(v))
		} else {
			proto.StringVal = append(proto.StringVal, bts)
		}
	default:
		return fmt.Errorf("can't convert string to tf:%v", mtype)
	}
	return nil
}

func addBytes(mtype tf.DataType, proto *tf.TensorProto, v []byte) error {
	switch mtype {
	case tf.DataType_DT_STRING:
		proto.StringVal = append(proto.StringVal, v)
	default:
		return fmt.Errorf("can't convert string to tf:%v", mtype)
	}
	return nil
}

func mustFloat32(i interface{}) float32 {
	f, ok := i.(float32)
	if ok {
		return f
	} else {
		return 0
	}
}

func mustFloat64(i interface{}) float64 {
	f, ok := i.(float64)
	if ok {
		return f
	} else {
		return 0
	}
}
