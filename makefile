.gen:
	protoc --go_out=rpc/gen --proto_path=rpc/proto rpc/proto/gophkeeper.proto --go_opt=paths=source_relative \
		--go-grpc_out=rpc/gen --proto_path=rpc/proto rpc/proto/gophkeeper.proto --go-grpc_opt=paths=source_relative

	protoc-go-inject-tag -remove_tag_comment -input="./rpc/gen/*.pb.go"

	mockgen -source=rpc/gen/gophkeeper_grpc.pb.go -destination=rpc/mock/gophkeeper_mock.go -package=gophkeepermock

.migrate:
	migrate create -ext sql -dir internal/server/repository/migrations -seq create_gophkeeper_tables