FROM golang:1.17-alpine as builder

RUN apk --no-cache add ca-certificates git
WORKDIR /workdir

COPY ./../../../go.mod .
COPY ./../../../go.sum .

RUN go mod download

COPY ./../../../. .

RUN go build ./services/login/server/main.go

# Run container
FROM alpine:latest

#RUN services/register/server/apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /workdir/main .
COPY --from=builder /workdir/services/login/server/demo.rsa .
COPY --from=builder /workdir/services/login/server/demo.rsa.pub .

EXPOSE "50051"

CMD ["./main"]