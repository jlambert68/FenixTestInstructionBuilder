
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun

compileProto:
	@echo "Compile proto file..."
	sed -i 's/\/\/import "common_grpc_api.proto";/import "common_grpc_api.proto";/' ..\/grpc_api\/qml_server_grpc_api\/qml_grpc_server.proto
	sed -i 's/import "..\/common_grpc_api\/common_grpc_api.proto";/\/\/import "..\/common_grpc_api\/common_grpc_api.proto";/' ..\/grpc_api\/qml_server_grpc_api\/qml_grpc_server.proto

	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I ../grpc_api/qml_server_grpc_api -I ../grpc_api/common_grpc_api --go_out=plugins=grpc:../grpc_api/qml_server_grpc_api ../grpc_api/qml_server_grpc_api/qml_grpc_server.proto

	sed -i 's/import "common_grpc_api.proto";/\/\/import "common_grpc_api.proto";/' ..\/grpc_api\/qml_server_grpc_api\/qml_grpc_server.proto
	sed -i 's/\/\/import "..\/common_grpc_api\/common_grpc_api.proto";/import "..\/common_grpc_api\/common_grpc_api.proto";/' ..\/grpc_api\/qml_server_grpc_api\/qml_grpc_server.proto

	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I ../grpc_api/common_grpc_api -I ../grpc_api/common_grpc_api --go_out=plugins=grpc:../grpc_api/common_grpc_api ../grpc_api/common_grpc_api/common_grpc_api.proto
