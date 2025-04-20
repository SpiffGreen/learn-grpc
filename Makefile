
gen:
	@protoc \
    --proto_path=protofiles "protofiles/learn.proto" \
    --go_out=shared/genproto/learn --go_opt=paths=source_relative \
    --go-grpc_out=shared/genproto/learn --go-grpc_opt=paths=source_relative