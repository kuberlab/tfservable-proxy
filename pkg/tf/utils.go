package tf

import (
	"context"
	"fmt"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/example"
	tf "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow_serving/apis"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

type TFFeatureJSON struct {
	Float     *float32   `json:"float",omitempty`
	FloatList *[]float32 `json:"float_list",omitempty`
	Int       *int64     `json:"int",omitempty`
	IntList   *[]int64   `json:"int_list",omitempty`
	Bytes     *[]byte    `json:"bytes",omitempty`
	BytesList *[][]byte  `json:"bytes_list",omitempty`
}

type TFInputJSON struct {
	Dtype tf.DataType  `json:"dtype",omitempty`
	Data  interface{} `json:"data",omitempty`
}

func (t *TFInputJSON) Tensor() (*tf.TensorProto, error) {
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
	TFFeatures  []map[string]TFFeatureJSON `json:"features",omitempty`
	Inputs      map[string]TFInputJSON     `json:"inputs",omitempty`
	OutFilter   []string                   `json:"out_filter",omitempty`
	OutMimeType string                     `json:"out_mime_type",omitempty`
}

func (f *TFFeatureJSON) TFFeature() *example.Feature {
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
		mSpec.Version = &google_protobuf.Int64Value{
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
		messages := [][]byte{}
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
	conn, err := grpc.Dial(servingAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Failed open grpc connection %v", err)
	}
	defer conn.Close()
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
	case tf.DataType_DT_INT64:
		res := make([]interface{}, len(t.Int64Val))
		for i := range res {
			res[i] = t.Int64Val[i]
		}
		return shapeContainer(t.TensorShape.Dim, res)
	case tf.DataType_DT_FLOAT:
		res := make([]interface{}, len(t.FloatVal))
		for i := range res {
			res[i] = t.FloatVal[i]
		}
		return shapeContainer(t.TensorShape.Dim, res)
	case tf.DataType_DT_DOUBLE:
		res := make([]interface{}, len(t.DoubleVal))
		for i := range res {
			res[i] = t.DoubleVal[i]
		}
		return shapeContainer(t.TensorShape.Dim, res)
	case tf.DataType_DT_STRING:
		res := make([]interface{}, len(t.StringVal))
		for i := range res {
			res[i] = t.StringVal[i]
		}
		return shapeContainer(t.TensorShape.Dim, res)

	}
	return nil
}

func shapeContainer(dim []*tf.TensorShapeProto_Dim, data []interface{}) interface{} {
	if len(dim) < 2 {
		return data
	}
	res := []interface{}{}
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

func fillTensor(data interface{}, proto *tf.TensorProto, index int) error {
	switch v := data.(type) {
	case []byte:
		if err := addBytes(proto.Dtype, proto, v); err != nil {
			return err
		}
	case []interface{}:
		if index == 0 {
			proto.TensorShape.Dim = append(proto.TensorShape.Dim, &tf.TensorShapeProto_Dim{
				Size: int64(len(v)),
			})
		}
		for i, v1 := range v {
			switch v2 := v1.(type) {
			case string:
				if err := addString(proto.Dtype, proto, v2); err != nil {
					return err
				}
			case []interface{}:
				fillTensor(v2, proto, i)
			case float64:
				if err := addFloat64(proto.Dtype, proto, v2); err != nil {
					return err
				}
			case int64:
				if err := addInt64(proto.Dtype, proto, v2); err != nil {
					return err
				}
			}
		}
	case string:
		if err := addString(proto.Dtype, proto, v); err != nil {
			return err
		}
	case float64:
		if err := addFloat64(proto.Dtype, proto, v); err != nil {
			return err
		}
	case int64:
		if err := addInt64(proto.Dtype, proto, v); err != nil {
			return err
		}
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
func addString(mtype tf.DataType, proto *tf.TensorProto, v string) error {
	switch mtype {
	case tf.DataType_DT_STRING:
		proto.StringVal = append(proto.StringVal, []byte(v))
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
