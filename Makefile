gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-cal-server:
	go run calculator/server/server.go
run-cal-client:
	go run calculator/client/client.go

gen-contact:
	protoc contact/contactpb/contact.proto --go_out=plugins=grpc:.
run-contact-server:
	go run contact/server/server.go
run-contact-client:
	go run contact/client/client.go