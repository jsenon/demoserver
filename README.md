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
docker build -t jsenon/demoserver:1.X .
```

### Launch Container

```
docker run -d -p 80:9010 jsenon/demoserver:1.X
```

Server HTTP in container serve on port 9010

### Func

- Print IP
- Print Version
- An API is available for healthy /healthy/am-i-up and /healthy/about

