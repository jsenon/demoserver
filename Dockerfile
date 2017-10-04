FROM alpine:latest

RUN apk add --no-cache bash curl wget
ADD demoserver2 /
CMD ["./demoserver2"]