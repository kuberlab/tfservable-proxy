package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	tf_core "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/kuberlab/tfservable-proxy/pkg/tf"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	prefix  = "proxy"
	timeout = 300
	port    int
)

const (
	defaultMaxMemory = 50 << 20 // 50 MB
)

func main() {
	flag.IntVar(&port, "port", 8082, "Proxy port")
	flag.IntVar(&timeout, "timeout", 300, "Timeout for model call in sec")
	flag.Parse()
	http.HandleFunc("/", proxy)
	log.Printf("Listen on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

type propertyParser func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error))

var (
	parsers = map[string]propertyParser{
		"float": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []float32{}
			feature.FloatList = &values
			return tf_core.DataType_DT_FLOAT, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 32); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float32(f), nil
				}
			}
		},
		"double": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []float32{}
			feature.FloatList = &values
			return tf_core.DataType_DT_DOUBLE, func(val string) (interface{}, error) {
				if f, err := strconv.ParseFloat(val, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.FloatList = append(*feature.FloatList, float32(f))
					}
					return float64(f), nil
				}
			}
		},
		"int": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []int64{}
			feature.IntList = &values
			return tf_core.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
		"int8": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []int64{}
			feature.IntList = &values
			return tf_core.DataType_DT_INT8, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int8(f), nil
				}
			}
		},
		"int16": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []int64{}
			feature.IntList = &values
			return tf_core.DataType_DT_INT16, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int16(f), nil
				}
			}
		},
		"int32": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []int64{}
			feature.IntList = &values
			return tf_core.DataType_DT_INT32, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int32(f), nil
				}
			}
		},
		"int64": func(feature *tf.TFFeatureJSON) (tf_core.DataType, func(val string) (interface{}, error)) {
			values := []int64{}
			feature.IntList = &values
			return tf_core.DataType_DT_INT64, func(val string) (interface{}, error) {
				if f, err := strconv.ParseInt(val, 10, 64); err != nil {
					return nil, err
				} else {
					if feature != nil {
						*feature.IntList = append(*feature.IntList, int64(f))
					}
					return int64(f), nil
				}
			}
		},
	}
)

func proxy(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	status := http.StatusOK
	var returnError error
	defer func() {
		fmt.Printf("%s->%d(%.2f)\n", req.RequestURI, status, time.Since(start).Seconds())
		if returnError != nil {
			w.WriteHeader(status)
			w.Write([]byte(returnError.Error()))
		}
	}()
	addr := req.Header.Get("PROXY_ADDR")
	if req.Method != "POST" {
		returnError = errors.New("Only POST request is supported")
		status = http.StatusMethodNotAllowed
		return
	}
	defer req.Body.Close()
	r, err := parseRequestURI(req.RequestURI)
	if err != nil {
		returnError = err
		status = http.StatusBadRequest
		return
	}
	model := tf.ModelData{}
	if strings.Contains(req.Header.Get("content-type"), "multipart/form-data") {
		err = req.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			returnError = err
			status = http.StatusBadRequest
			return
		}
		fmt.Printf("%v\n", req.MultipartForm.Value)
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
					if p[0] == "byte" {
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
						bytes := [][]byte{data}
						feature.BytesList = &bytes
						model.TFFeatures[0][strings.Join(p[1:], "_")] = feature
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
							Data:  &values,
						}
					}
				}
			}

			for k, fHeader := range req.MultipartForm.File {
				p := strings.Split(k, "_")
				if len(p) > 0 && len(fHeader) > 0 {
					if p[0] == "byte" {
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
						model.Inputs[strings.Join(p[1:], "_")] = tf.TFInputJSON{
							Dtype: tf_core.DataType_DT_STRING,
							Data:  data,
						}
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
	tContext, _ := context.WithTimeout(context.Background(), time.Duration(time.Duration(timeout)*time.Second))
	result, err := tf.CallTF(tContext, addr, r.model, r.version, r.signature, model)
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

type request struct {
	model     string
	version   int64
	signature string
}

func parseRequestURI(uri string) (e request, err error) {
	if i := strings.Index(uri, prefix); i < 0 {
		err = errors.New("Bad request path")
	} else {
		uri = strings.TrimPrefix(uri[i+len(prefix):], "/")
		p := strings.Split(uri, "/")
		if len(p) < 1 {
			err = errors.New("Bad request path")
		}
		e.model = p[0]
		if len(p) > 1 {
			e.signature = p[1]
		}
		if len(p) > 2 {
			if e.version, err = strconv.ParseInt(p[2], 10, 64); err != nil {
				err = fmt.Errorf("Bad version value")
				return
			}
		}

	}
	return
}
