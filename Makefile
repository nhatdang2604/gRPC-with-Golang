gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-server:
	go run calculator/server/server.go
run-client:
	go run calculator/client/client.go