gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-cal-server:
	go run calculator/server/server.go
run-cal-client:
	go run calculator/client/client.go
