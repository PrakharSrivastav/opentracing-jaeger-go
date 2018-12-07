# Setup

Run the below docker image for jager-tracer backend

```
docker run \
  --rm \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 16686:16686 \
  jaegertracing/all-in-one:1.7 \
  --log-level=debug
```


jaeger ui should be avialable at : http://localhost:16686