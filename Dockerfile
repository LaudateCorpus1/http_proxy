FROM golang
RUN go get github.com/bountylabs/http_proxy
ENTRYPOINT http_proxy
