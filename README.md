# tfservable-proxy

## Install

```
glide install
```

### Run

```
go run  tfproxy.go --port 8082 --timeout 3000
```


curl -X POST -d '{"features":[{"x":{"Float":70},"y":{"Float":50}},{"x":{"Float":50},"y":{"Float":70}}]}' -H 'PROXY_ADDR: test.21-kuberlab-demo' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/api/v2/tfproxy/test

curl -X POST -F float_x=1 -F float_y=49 http://127.0.0.1:8082/proxy/localhost/9000/test

curl  -X POST -F float_x=1 -F float_y=49 -H 'PROXY_ADDR: test.21-kuberlab-demo' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/api/v2/tfproxy/test

 curl -X POST -F byte_image=@Kuberlab_head1.jpg -F raw_input=yes -F out_mime_type=image/png -F out_filters=result -H 'PROXY_ADDR: styles-serving-train-4.21-kuberlab-demo' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/api/v2/tfproxy/styles/transform > result.png


 curl -X POST -F byte_inputs=@pet2.jpg -F raw_input=yes -H 'PROXY_ADDR: object-detection.21-kuberlab-demo' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/api/v2/tfproxy/object-detection-pets

  curl -X POST -F uint8_inputs=@pet2.jpg -F raw_input=yes -H 'PROXY_ADDR: 127.0.0.1' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/proxy/object-detection-pets

  curl -X POST -F strings_inputs=@pet2.jpg -F raw_input=yes -H 'PROXY_ADDR: 127.0.0.1' -H 'PROXY_PORT: 9000' http://127.0.0.1:8082/proxy/object-detection-pets