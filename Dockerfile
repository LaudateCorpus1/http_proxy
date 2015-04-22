FROM golang
RUN go install github.com/bountylabs/http_proxy
ENTRYPOINT http_proxy
