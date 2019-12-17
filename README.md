# FLATTEN API
## How to use 
You have many different ways to run the API
- Terminal
```sh
$ curl -X GET 'http://localhost:8080/flat?key=\[1,2,3,\[2\]\]'
```
- Postman
```Plaintext
localhost:8181/flat?key=[1,2,3,[2]]
```

- With Makefile commands
```sh
$ make run
$ make testApi
```

* [Decisions made](Resolution.md)