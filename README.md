# demoserver

This is a demo server which print it's version and it's IP on HTTP Server and provide API

## Build

Change version print in web/web/go `hello, Im Service version 1.X`

```
go build src/demoserver.go
```

Check name of your binairies against name used in dockerfile


### Create Container

```
docker build -t jsenon/mydemoserver:vX.X .
```

### Launch Container

```
docker run -d -p 80:9010 jsenon/mydemoserver:vX.X
```

Server HTTP in container serve on port 9010

### Func

- Print IP
- Print Version
- An API is available for healthy /healthy/am-i-up and /healthy/about


### TIPS

If you receive `cannot execute binary file: Exec format error`

Compile with  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build src/demoserver.go`

And build `docker build -t jsenon/demoserver:v1.3 .`

