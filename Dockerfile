FROM golang:latest as build_base
RUN apt-get install -y bash ca-certificates git
RUN mkdir -p $GOPATH/src/github.com/rumsrami/app
WORKDIR $GOPATH/src/github.com/rumsrami/app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder
WORKDIR $GOPATH/src/github.com/rumsrami/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/
RUN chmod +x main

FROM scratch
COPY --from=server_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=server_builder /go/src/github.com/rumsrami/app/assets /app/assets/
COPY --from=server_builder /go/src/github.com/rumsrami/app/main /app/main
ENV port="443"
ENV env="prod"
EXPOSE 443
ENTRYPOINT ["/app/main"]