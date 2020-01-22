package tfhttp

import (
	"strconv"

	tfcore "github.com/kuberlab/tfservable-proxy/pkg/tensorflow/core/framework"
)

type propertyParser func() (tfcore.DataType, func(val string) (interface{}, error))
type binaryParser func(data []byte) (tfcore.DataType, interface{})

var (
	floatParser = func(data []byte) (tfcore.DataType, interface{}) {
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = float32(b)
		}
		return tfcore.DataType_DT_FLOAT, values
	}
	intParser = func(data []byte) (tfcore.DataType, interface{}) {
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = int64(b)
		}
		return tfcore.DataType_DT_INT64, values
	}
	binaryParsers = map[string]binaryParser{
		"float": floatParser,
		"double": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = float64(b)
			}
			return tfcore.DataType_DT_DOUBLE, values
		},
		"int": intParser,
		"int8": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int8(b)
			}
			return tfcore.DataType_DT_INT8, values
		},
		"int16": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int16(b)
			}
			return tfcore.DataType_DT_INT16, values
		},
		"int32": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int32(b)
			}
			return tfcore.DataType_DT_INT32, values
		},
		"int64": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int64(b)
			}
			return tfcore.DataType_DT_INT64, values
		},
		"uint8": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint8(b)
			}
			return tfcore.DataType_DT_UINT8, values
		},
		"uint16": func(data []byte) (tfcore.DataType, interface{}) {
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint16(b)
			}
			return tfcore.DataType_DT_UINT16, values
		},
		"bytes": func(data []byte) (tfcore.DataType, interface{}) {
			values := []interface{}{data}
			return tfcore.DataType_DT_STRING, values
		},
		"strings": func(data []byte) (tfcore.DataType, interface{}) {
			values := []interface{}{data}
			return tfcore.DataType_DT_STRING, values
		},
		"byte": func(data []byte) (tfcore.DataType, interface{}) {
			return tfcore.DataType_DT_STRING, data
		},
		"string": func(data []byte) (tfcore.DataType, interface{}) {
			return tfcore.DataType_DT_STRING, data
		},
	}
	parsers = map[string]propertyParser{
		"float": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_FLOAT, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 32); err != nil {
					return nil, err
				} else {
					return float32(f), nil
				}
			}
		},
		"double": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_DOUBLE, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 64); err != nil {
					return nil, err
				} else {
					return float64(f), nil
				}
			}
		},
		"int": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return int64(f), nil
				}
			}
		},
		"int8": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return int8(f), nil
				}
			}
		},
		"int16": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT16, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return int16(f), nil
				}
			}
		},
		"int32": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT32, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return int32(f), nil
				}
			}
		},
		"int64": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return int64(f), nil
				}
			}
		},
		"uint8": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return uint8(f), nil
				}
			}
		},
		"uint16": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					return uint16(f), nil
				}
			}
		},
		"string": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_STRING, func(val string) (interface{}, error) {
				return []byte(val), nil
			}
		},
		"bytes": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_STRING, func(val string) (interface{}, error) {
				return []byte(val), nil
			}
		},
		"bool": func() (tfcore.DataType, func(val string) (interface{}, error)) {
			return tfcore.DataType_DT_BOOL, func(val string) (interface{}, error) {
				b, _ := strconv.ParseBool(val)
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
