.PHONY: gen
gen:
	find proto -name "*.proto" | xargs protoc \
		-I proto \
		-I third-party/googleapis \
		--go_out=./gen/go --go_opt=paths=source_relative \
		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./gen/go --grpc-gateway_opt=paths=source_relative
