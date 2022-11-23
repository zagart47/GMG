FROM golang:1.19-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main main.go

FROM scratch
COPY server.crt /
COPY server.key /
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
EXPOSE 80