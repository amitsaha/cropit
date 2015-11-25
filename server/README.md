### Running the server

```
go run server.go
```

### Client example

http --form POST http://127.0.0.1:9090 image@~/work/golang/src/github.com/amitsaha/cropit/test_images/cat1.jpg w=4000 h=6000 --output cropped.jpg
