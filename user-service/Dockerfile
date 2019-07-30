FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /go/src/github.com/wizofgoz/microservices-sample/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -o microservices-service-user -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/wizofgoz/microservices-sample/user-service .

CMD ["./microservices-service-user"]