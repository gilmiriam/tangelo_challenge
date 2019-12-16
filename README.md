# FLATTEN API
## How to use 
You many different ways the API
- Terminal
```sh
$ curl -X GET 'http://localhost:8080/flat?key=\[1,2,3,\[2\]\]'
```
- Postman
```Plaintext
localhost:8080/flat?key=[1,2,3,[2]]
```

- With the go test tool
```sh
$ make test.go
```