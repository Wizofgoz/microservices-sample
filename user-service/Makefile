build:
	protoc -I. --go_out=plugins=micro:. \
      proto/user/user.proto
	docker build -t microservices-service-user .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 microservices-service-user