package proto

// 生成 proto
//go:generate protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../../third_party/pagination --proto_path=../../../../ --proto_path=../../../../comment/service/v1/proto/ --proto_path=../../../../content/service/v1/proto/ --proto_path=../../../../user/service/v1/proto/ --proto_path=../../../../file/service/v1/proto/ --go_out=paths=source_relative:../ ./*.proto

// 生成 proto http
//go:generate protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../../third_party/pagination --proto_path=../../../../ --proto_path=../../../../comment/service/v1/proto/ --proto_path=../../../../content/service/v1/proto/ --proto_path=../../../../user/service/v1/proto/ --proto_path=../../../../file/service/v1/proto/ --go-http_out=paths=source_relative:../ ./*.proto

// 生成 proto errors
//go:generate protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../../third_party/pagination --proto_path=../../../../ --proto_path=../../../../comment/service/v1/proto/ --proto_path=../../../../content/service/v1/proto/ --proto_path=../../../../user/service/v1/proto/ --proto_path=../../../../file/service/v1/proto/ --go-errors_out=paths=source_relative:../ ./*.proto

// 生成 swagger
//go:generate protoc --proto_path=. --proto_path=../../../../../third_party --proto_path=../../../../../third_party/pagination --proto_path=../../../../ --proto_path=../../../../comment/service/v1/proto/ --proto_path=../../../../content/service/v1/proto/ --proto_path=../../../../user/service/v1/proto/ --proto_path=../../../../file/service/v1/proto/ --openapiv2_out ../ --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true ./*.proto
