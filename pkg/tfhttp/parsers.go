package tfhttp

import (
	"strconv"

	tfcore "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/kuberlab/tfservable-proxy/pkg/tf"
)

type propertyParser func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error))
type binaryParser func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{})

var (
	floatFeatureParser = func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
		if feature != nil {
			values := make([]float32, len(data))
			for i, b := range data {
				values[i] = float32(b)
			}
			feature.FloatList = &values
			return tfcore.DataType_DT_FLOAT, nil
		}
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = float32(b)
		}
		return tfcore.DataType_DT_FLOAT, values
	}
	intFeatureParser = func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
		if feature != nil {
			values := make([]int64, len(data))
			for i, b := range data {
				values[i] = int64(b)
			}
			feature.IntList = &values
			return tfcore.DataType_DT_INT64, nil
		}
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = int64(b)
		}
		return tfcore.DataType_DT_INT64, values
	}
	binaryParsers = map[string]binaryParser{
		"float": floatFeatureParser,
		"double": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return floatFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = float64(b)
			}
			return tfcore.DataType_DT_DOUBLE, values
		},
		"int": intFeatureParser,
		"int8": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int8(b)
			}
			return tfcore.DataType_DT_INT8, values
		},
		"int16": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int16(b)
			}
			return tfcore.DataType_DT_INT16, values
		},
		"int32": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int32(b)
			}
			return tfcore.DataType_DT_INT32, values
		},
		"int64": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int64(b)
			}
			return tfcore.DataType_DT_INT64, values
		},
		"uint8": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint8(b)
			}
			return tfcore.DataType_DT_UINT8, values
		},
		"uint16": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint16(b)
			}
			return tfcore.DataType_DT_UINT16, values
		},
		"bytes": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			values := []interface{}{data}
			return tfcore.DataType_DT_STRING, values
		},
		"strings": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			values := []interface{}{data}
			return tfcore.DataType_DT_STRING, values
		},
		"byte": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			return tfcore.DataType_DT_STRING, data
		},
		"string": func(feature *tf.FeatureJSON, data []byte) (tfcore.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			return tfcore.DataType_DT_STRING, data
		},
	}
	parsers = map[string]propertyParser{
		"float": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_FLOAT, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 32); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]float32, 0)
						feature.FloatList = &values
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float32(f), nil
				}
			}
		},
		"double": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_DOUBLE, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]float32, 0)
						feature.FloatList = &values
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float64(f), nil
				}
			}
		},
		"int": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
		"int8": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int8(f), nil
				}
			}
		},
		"int16": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT16, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int16(f), nil
				}
			}
		},
		"int32": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT32, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int32(f), nil
				}
			}
		},
		"int64": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
		"uint8": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return uint8(f), nil
				}
			}
		},
		"uint16": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := make([]int64, 0)
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return uint16(f), nil
				}
			}
		},
		"string": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_STRING, func(val string) (interface{}, error) {
				if feature != nil {
					values := make([][]byte, 0)
					feature.BytesList = &values
					*feature.BytesList = append(*feature.BytesList, []byte(val))
				}
				return []byte(val), nil
			}
		},
		"bool": func(feature *tf.FeatureJSON) (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_BOOL, func(val string) (interface{}, error) {
				b, _ := strconv.ParseBool(val)
				if feature != nil {
					values := make([]bool, 0)
					feature.BoolList = &values
					*feature.BoolList = append(*feature.BoolList, b)
				}
				return b, nil
			}
		},
	}
)

func binaryParsersList() []string {
	res := make([]string, 0)
	for k := range binaryParsers {
		res = append(res, k)
	}
	return res
}

func parsersList() []string {
	res := make([]string, 0)
	for k := range parsers {
		res = append(res, k)
	}
	return res
}
