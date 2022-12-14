gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
run-cal-server:
	go run calculator/server/server.go
run-cal-client:
	go run calculator/client/client.go

gen-contact:
	protoc contact/contactpb/contact.proto --go_out=plugins=grpc:.
run-contact-server:
	go run contact/server/server.go contact/server/models.go
run-contact-client:
	go run contact/client/client.go

gen-gateway:
	protoc gateway/gatewaypb/gateway.proto --go_out=plugins=grpc:.
gen-gateway-stub:
	protoc -I . \
		--go_out ./ --go_opt paths=source_relative \
		--go-grpc_out ./ --go-grpc_opt paths=source_relative \
		gateway/gatewaypb/gateway.proto
gen-rp:
	protoc-gen-grpc-gateway
update-gateway-stub:
	protoc -I . --grpc-gateway_out ./ \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		gateway/gatewaypb/gateway.proto
run-gateway-server:
	go run gateway/server/server.go
run-gateway-rproxy:
	go run gateway/proxy/proxy.go