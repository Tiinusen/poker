package generate

//build +ignore
//go:generate glob protoc -Ithird_party/proto/ --proto_path=api/proto/ --go_out=plugins=grpc,paths=source_relative,import_path=proto:./pkg/proto/ api/proto/poker/*.proto
