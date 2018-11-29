package tfhttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	tf_core "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/kuberlab/tfservable-proxy/pkg/tf"
)

const (
	defaultMaxMemory = 50 << 20 // 50 MB
)

type TFHttpProxy struct {
	Timeout        time.Duration
	URIPrefix      string
	DefaultAddress string
	DefaultPort    int
}

type propertyParser func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error))
type binaryParser func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{})

var (
	floatFeatureParser = func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
		if feature != nil {
			values := make([]float32, len(data))
			for i, b := range data {
				values[i] = float32(b)
			}
			feature.FloatList = &values
			return tf_core.DataType_DT_FLOAT, nil
		}
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = float32(b)
		}
		return tf_core.DataType_DT_FLOAT, values
	}
	intFeatureParser = func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
		if feature != nil {
			values := make([]int64, len(data))
			for i, b := range data {
				values[i] = int64(b)
			}
			feature.IntList = &values
			return tf_core.DataType_DT_INT64, nil
		}
		values := make([]interface{}, len(data))
		for i, b := range data {
			values[i] = int64(b)
		}
		return tf_core.DataType_DT_INT64, values
	}
	binaryParsers = map[string]binaryParser{
		"float": floatFeatureParser,
		"double": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return floatFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = float64(b)
			}
			return tf_core.DataType_DT_DOUBLE, values
		},
		"int": intFeatureParser,
		"int8": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int8(b)
			}
			return tf_core.DataType_DT_INT8, values
		},
		"int16": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int16(b)
			}
			return tf_core.DataType_DT_INT16, values
		},
		"int32": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int32(b)
			}
			return tf_core.DataType_DT_INT32, values
		},
		"int64": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = int64(b)
			}
			return tf_core.DataType_DT_INT64, values
		},
		"uint8": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint8(b)
			}
			return tf_core.DataType_DT_UINT8, values
		},
		"uint16": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				return intFeatureParser(feature, data)
			}
			values := make([]interface{}, len(data))
			for i, b := range data {
				values[i] = uint16(b)
			}
			return tf_core.DataType_DT_UINT16, values
		},
		"bytes": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			values := []interface{}{data}
			return tf_core.DataType_DT_STRING, values
		},
		"strings": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			values := []interface{}{data}
			return tf_core.DataType_DT_STRING, values
		},
		"byte": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			return tf_core.DataType_DT_STRING, data
		},
		"string": func(feature *tf.TFFeatureJSON, data []byte) (tf_core.DataType, interface{}) {
			if feature != nil {
				bytesData := [][]byte{data}
				feature.BytesList = &bytesData
			}
			return tf_core.DataType_DT_STRING, data
		},
	}
	parsers = map[string]propertyParser{
		"float": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_FLOAT, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 32); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []float32{}
						feature.FloatList = &values
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float32(f), nil
				}
			}
		},
		"double": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_DOUBLE, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []float32{}
						feature.FloatList = &values
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float64(f), nil
				}
			}
		},
		"int": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
		"int8": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_INT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int8(f), nil
				}
			}
		},
		"int16": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_INT16, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int16(f), nil
				}
			}
		},
		"int32": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_INT32, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int32(f), nil
				}
			}
		},
		"int64": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
		"uint8": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return uint8(f), nil
				}
			}
		},
		"uint16": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_UINT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						values := []int64{}
						feature.IntList = &values
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return uint16(f), nil
				}
			}
		},
		"string": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			return tf_core.DataType_DT_STRING, func(val string) (interface{}, error) {
				if feature != nil {
					values := [][]byte{}
					feature.BytesList = &values
					*feature.BytesList = append(*feature.BytesList, []byte(val))
				}
				return []byte(val), nil
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

func (proxy TFHttpProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	status := http.StatusOK
	var returnError error
	defer func() {
		log.Printf("%s->%d(%.2f)\n", req.RequestURI, status, time.Since(start).Seconds())
		if returnError != nil {
			w.WriteHeader(status)
			errStr := returnError.Error()
			if !strings.HasSuffix(errStr, "\n") {
				errStr += "\n"
			}
			w.Write([]byte(errStr))
			logrus.Error(returnError.Error())
		}
	}()
	addr := req.Header.Get("Proxy-Addr")
	sport := req.Header.Get("Proxy-Port")
	port := proxy.DefaultPort
	if sport != "" {
		if v, err := strconv.Atoi(sport); err == nil {
			port = v
		}
	}
	if addr == "" {
		if proxy.DefaultAddress != "" {
			addr = proxy.DefaultAddress
		} else {
			returnError = errors.New("Provide target address using header 'Proxy-Addr'")
			status = http.StatusBadRequest
			return
		}
	}
	addr = fmt.Sprintf("%s:%d", addr, port)

	modelName, modelVersion, modelSinature, err := parseRequestURI(proxy.URIPrefix, req.RequestURI)
	if err != nil {
		returnError = err
		status = http.StatusNotFound
		return
	}

	if req.Method != "POST" {
		returnError = errors.New("Only POST request is supported")
		status = http.StatusMethodNotAllowed
		return
	}

	defer req.Body.Close()
	model := tf.ModelData{}
	if strings.Contains(req.Header.Get("content-type"), "multipart/form-data") {
		err = req.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			returnError = err
			status = http.StatusBadRequest
			return
		}
		if v, ok := req.MultipartForm.Value["out_mime_type"]; ok && len(v) > 0 {
			model.OutMimeType = v[0]
		}
		if v, ok := req.MultipartForm.Value["out_filters"]; ok && len(v) > 0 {
			model.OutFilter = strings.Split(v[0], ",")
		}
		if _, ok := req.MultipartForm.Value["raw_input"]; !ok {
			model.TFFeatures = []map[string]tf.TFFeatureJSON{map[string]tf.TFFeatureJSON{}}
			for k, v := range req.MultipartForm.Value {
				p := strings.Split(k, "_")
				if len(p) > 0 {
					if parser, ok := parsers[p[0]]; ok {
						feature := tf.TFFeatureJSON{}
						_, valueParser := parser(&feature)
						for _, s1 := range v {
							for _, s2 := range strings.Split(s1, ",") {
								if _, err := valueParser(s2); err != nil {
									returnError = err
									status = http.StatusBadRequest
									return
								}
							}
						}
						model.TFFeatures[0][strings.Join(p[1:], "_")] = feature
					}
				}
			}
			for k, fHeader := range req.MultipartForm.File {
				p := strings.Split(k, "_")
				if len(p) > 0 && len(fHeader) > 0 {
					if parser, ok := binaryParsers[p[0]]; ok {
						feature := tf.TFFeatureJSON{}
						file, err := fHeader[0].Open()
						if err != nil {
							returnError = err
							status = http.StatusBadRequest
							return
						}
						defer file.Close()
						data, err := ioutil.ReadAll(file)
						if err != nil {
							returnError = err
							status = http.StatusBadRequest
							return
						}
						parser(&feature, data)
						model.TFFeatures[0][strings.Join(p[1:], "_")] = feature
					} else {
						returnError = fmt.Errorf("Unsupotred binary hanldler for %s", p[0])
						status = http.StatusBadRequest
						return
					}
				}
			}

		} else {
			model.Inputs = map[string]tf.TFInputJSON{}
			for k, v := range req.MultipartForm.Value {
				values := []interface{}{}
				p := strings.Split(k, "_")
				if len(p) > 0 {
					if parser, ok := parsers[p[0]]; ok {
						dtype, valueParser := parser(nil)
						for _, s1 := range v {
							for _, s2 := range strings.Split(s1, ",") {
								if iv, err := valueParser(s2); err != nil {
									returnError = err
									status = http.StatusBadRequest
									return
								} else {
									values = append(values, iv)
								}
							}
						}
						model.Inputs[strings.Join(p[1:], "_")] = tf.TFInputJSON{
							Dtype: dtype,
							Data:  values,
						}
					}
				}
			}

			for k, fHeader := range req.MultipartForm.File {
				p := strings.Split(k, "_")
				if len(p) > 0 && len(fHeader) > 0 {
					if parser, ok := binaryParsers[p[0]]; ok {
						file, err := fHeader[0].Open()
						if err != nil {
							returnError = err
							status = http.StatusBadRequest
							return
						}
						defer file.Close()
						fileData, err := ioutil.ReadAll(file)
						if err != nil {
							returnError = err
							status = http.StatusBadRequest
							return
						}
						dtype, data := parser(nil, fileData)
						model.Inputs[strings.Join(p[1:], "_")] = tf.TFInputJSON{
							Dtype: dtype,
							Data:  data,
						}
					} else {
						returnError = fmt.Errorf(
							"Unsupported binary handler %v; currently supported: %v",
							strings.Trim(p[0], "\n"),
							binaryParsersList(),
						)
						status = http.StatusBadRequest
						return
					}
				}
			}
		}
	} else {
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&model)
		if err != nil {
			returnError = fmt.Errorf("Failed decode input data %v", err)
			status = http.StatusBadRequest
			return
		}
	}
	tContext, _ := context.WithTimeout(context.Background(), proxy.Timeout)
	result, err := tf.CallTF(tContext, addr, modelName, modelVersion, modelSinature, model)
	if err != nil {
		returnError = fmt.Errorf("Failed call %v", err)
		status = http.StatusBadRequest
		return
	}
	if len(model.OutFilter) > 0 {
		filter := map[string]interface{}{}
		for _, k := range model.OutFilter {
			if v, ok := result[k]; ok {
				filter[k] = v
			}
		}
		result = filter
	}
	if strings.HasPrefix(model.OutMimeType, "image/") && len(model.OutFilter) == 1 {
		if v, ok := result[model.OutFilter[0]]; ok {
			switch b := v.(type) {
			case []interface{}:
				if len(b) == 1 {
					switch v := b[0].(type) {
					case []byte:
						w.Write(v)
						return
					default:
						returnError = errors.New("Output must be shape tf.string [1]")
						status = http.StatusBadRequest
					}
				} else {
					returnError = errors.New("Output must be shape [1]")
					status = http.StatusBadRequest
					return
				}
			default:
				returnError = errors.New("Bad out type")
				status = http.StatusBadRequest
				return
			}
		} else {
			returnError = errors.New("Ouput field not found")
			status = http.StatusBadRequest
			return
		}
	} else if model.OutMimeType == "text/plain" && len(model.OutFilter) == 1 {
		if v, ok := result[model.OutFilter[0]]; ok {
			switch b := v.(type) {
			case []interface{}:
				texts := []string{}
				for _, v := range b {
					switch t := v.(type) {
					case []byte:
						texts = append(texts, string(t))
					default:
						returnError = errors.New("Output must be array of bytes")
						status = http.StatusBadRequest
						return
					}
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(texts)
				return
			default:
				returnError = errors.New("Bad out type")
				status = http.StatusBadRequest
				return
			}
		} else {
			returnError = errors.New("Ouput field not found")
			status = http.StatusBadRequest
			return
		}
	} else {
		encoder := json.NewEncoder(w)
		err = encoder.Encode(result)
		if err != nil {
			returnError = fmt.Errorf("Failed encode output data %v", err)
			status = http.StatusBadRequest
			return
		}
	}

}

func parseRequestURI(prefix, uri string) (model string, version int64, signature string, err error) {
	version = -1
	if i := strings.Index(uri, prefix); i < 0 {
		err = fmt.Errorf("Wrong request path, need prefix /%v", prefix)
	} else {
		uri = strings.TrimPrefix(uri[i+len(prefix):], "/")
		p := strings.Split(uri, "/")
		if len(p) < 1 {
			err = fmt.Errorf("Wrong request path, need prefix /%v", prefix)
		}
		model = p[0]
		if len(p) > 1 {
			signature = p[1]
		}
		if len(p) > 2 {
			if version, err = strconv.ParseInt(p[2], 10, 64); err != nil {
				err = fmt.Errorf("Bad version value")
			}
		}

	}
	return
}
