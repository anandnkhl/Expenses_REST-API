# Expenses_REST-API

Build the autogen.go file to make auto code generating for the CRUD operation in Handler directory
```bash 
$ go build autogen.go 
``` 

```bash 
$ go generate ./... 
```

Compile the file main.go to start the server at port 8080 [change port no. (if required) on main.go line no:33]
```bash
$ go run main.go
```
Make sure the Mongo service is active before running the main.
```bash
$ sudo service mongod start
```
