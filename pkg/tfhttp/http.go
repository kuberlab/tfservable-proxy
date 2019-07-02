package tfhttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	"github.com/kuberlab/tfservable-proxy/pkg/tf"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	defaultMaxMemory = 50 << 20 // 50 MB
)

type Proxy struct {
	Timeout        time.Duration
	URIPrefix      string
	globalPrefix   string
	DefaultAddress string
	DefaultPort    int

	router *mux.Router
}

func NewProxy(globalPrefix, predictPrefix string, staticRoot string) *Proxy {
	predictPrefix = strings.TrimPrefix(predictPrefix, "/")
	if !strings.HasSuffix(predictPrefix, "/") {
		predictPrefix = predictPrefix + "/"
	}

	globalPrefix = strings.TrimPrefix(globalPrefix, "/")
	globalPrefix = strings.TrimSuffix(globalPrefix, "/")
	if globalPrefix != "" {
		globalPrefix = "/" + globalPrefix + "/"
	} else {
		globalPrefix = "/"
	}

	predictPrefix = globalPrefix + predictPrefix

	p := &Proxy{URIPrefix: predictPrefix, globalPrefix: globalPrefix}

	p.router = mux.NewRouter()
	p.router.PathPrefix(predictPrefix).HandlerFunc(p.PredictHandler)
	p.router.PathPrefix(globalPrefix + "hls/").HandlerFunc(p.ProxyStreams)
	p.router.PathPrefix(globalPrefix + "mjpg/").HandlerFunc(p.ProxyStreams)
	if staticRoot != "" {
		p.router.PathPrefix(globalPrefix).Handler(
			http.FileServer(http.Dir(staticRoot)),
		)
	}

	return p
}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	proxy.router.ServeHTTP(w, req)
}

func (proxy *Proxy) ProxyStreams(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	if strings.HasSuffix(req.URL.Path, "m3u8") || strings.HasSuffix(req.URL.Path, "ts") {
		resp, err := client.Get(fmt.Sprintf("http://localhost:8080%v", req.URL.Path))

		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(resp.StatusCode)
		for k, v := range resp.Header {
			for _, hv := range v {
				w.Header().Add(k, hv)
			}
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
	if strings.HasSuffix(req.URL.Path, "mjpg") {
		splitted := strings.Split(req.URL.Path, "/")
		resp, err := client.Get(fmt.Sprintf("http://localhost:8000/%v", splitted[len(splitted)-1]))

		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(resp.StatusCode)
		for k, v := range resp.Header {
			for _, hv := range v {
				w.Header().Add(k, hv)
			}
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
}

func (proxy *Proxy) PredictHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	status := http.StatusOK
	var returnError error
	defer func() {
		log.Printf("%s -> %d(%v)\n", req.RequestURI, status, time.Since(start))
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
			model.TFFeatures = []map[string]tf.FeatureJSON{{}}
			for k, v := range req.MultipartForm.Value {
				p := strings.Split(k, "_")
				if len(p) > 0 {
					if parser, ok := parsers[p[0]]; ok {
						feature := tf.FeatureJSON{}
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
						feature := tf.FeatureJSON{}
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
						returnError = fmt.Errorf("Unsupported binary hanldler for %s", p[0])
						status = http.StatusBadRequest
						return
					}
				}
			}

		} else {
			model.Inputs = map[string]tf.InputJSON{}
			for k, v := range req.MultipartForm.Value {
				values := make([]interface{}, 0)
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
						model.Inputs[strings.Join(p[1:], "_")] = tf.InputJSON{
							Dtype: dtype,
							Data:  values,
						}
					} // else {
					//	returnError = fmt.Errorf(
					//		"Unsupported handler %v; currently supported: %v",
					//		strings.Trim(p[0], "\n"),
					//		parsersList(),
					//	)
					//	status = http.StatusBadRequest
					//	return
					//}
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
						model.Inputs[strings.Join(p[1:], "_")] = tf.InputJSON{
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
				texts := make([]string, 0)
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
				_ = encoder.Encode(texts)
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
