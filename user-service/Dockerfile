FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/erikperttu/shippy/user-service

COPY . .

RUN go get -v -u github.com/golang/dep/cmd/dep
RUN dep init -v && dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build  -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/erikperttu/shippy/user-service/user-service .

CMD ["./vessel-service"]