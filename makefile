.gen:
	protoc --go_out=rpc/gen --proto_path=rpc/proto rpc/proto/gophkeeper.proto --go_opt=paths=source_relative \
		--go-grpc_out=rpc/gen --proto_path=rpc/proto rpc/proto/gophkeeper.proto --go-grpc_opt=paths=source_relative

	protoc-go-inject-tag -remove_tag_comment -input="./rpc/gen/*.pb.go"

	mockgen -source=rpc/gen/gophkeeper_grpc.pb.go -destination=rpc/mock/gophkeeper_mock.go -package=gophkeepermock

.migrate:
	migrate create -ext sql -dir internal/server/repository/migrations -seq create_gophkeeper_tables

.certs:
	touch cert/server-ext.cnf
	openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout cert/ca-key.pem -out cert/ca-cert.pem -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*"
	@echo "CA's self-signed certificate"
	openssl x509 -in cert/ca-cert.pem -noout -text
	openssl req -newkey rsa:4096 -nodes -keyout cert/server-key.pem -out cert/server-req.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.mipa.com"
	openssl x509 -req -in cert/server-req.pem -days 60 -CA cert/ca-cert.pem -CAkey cert/ca-key.pem -CAcreateserial -out cert/server-cert.pem -extfile cert/server-ext.cnf
	@echo "Server's signed certificate"
	openssl x509 -in cert/server-cert.pem -noout -text
