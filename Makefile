gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-cal-server:
	go run calculator/server/server.go
run-cal-client:
	go run calculator/client/client.go

gen-cont:
	protoc contact/contactpb/contact.proto --go_out=plugin=grpc:.
run-cont-server:
	go run contact/server/server.go
run-cont-client:
	go run contact/client/client.go