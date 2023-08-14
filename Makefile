gen: 
	protoc --go_out=. --go-grpc_out=. proto/processor_message.proto

clean: 
	rm pb/*.go

run: 
	go run main.go