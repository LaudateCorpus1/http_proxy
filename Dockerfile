FROM golang
RUN go build -o /bin/http_proxy main.go
ENTRYPOINT http_proxy
