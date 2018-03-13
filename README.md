# tfservable-proxy

## Install

```
glide install
```

### Run

```
go run  tfproxy.go --port 8082 --timeout 3000
```


curl -X POST -d '{"features":[{"x":{"Float":70},"y":{"Float":50}},{"x":{"Float":50},"y":{"Float":70}}]}' http://127.0.0.1:8082/proxy/localhost/9000/test

curl -X POST -F float_x=1 -F float_y=49 http://127.0.0.1:8082/proxy/localhost/9000/test

 curl -X POST -F byte_image=@Kuberlab_head1.jpg -F raw_input=yes -F out_mime_type=image/png -F out_filters=result http://127.0.0.1:8082/proxy/localhost/9000/styles/transform > result.png