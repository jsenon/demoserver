# demoserver

This is a demo server which print it's version and it's IP on HTTP Server

## Build

Change version print in web/web/go `hello, Im Service version 1.X`

```
go build src/demoserver.go
```


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

