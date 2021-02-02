
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun

compileProto:
	@echo "Compile proto file..."
#~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I ../grpc_api/backend_server_grpc_api --go_out=plugins=grpc:../grpc_api/backend_server_grpc_api ../grpc_api/backend_server_grpc_api/backend_grpc_server.proto
	sed -i 's/\/\/import "grpc_common.proto";/import "grpc_common.proto";/' ..\/grpc_api\/backend_server_grpc_api\/backend_grpc_server.proto
	sed -i 's/import "..\/common_grpc_api\/grpc_common.proto";/\/\/import "..\/common_grpc_api\/grpc_common.proto";/' ..\/grpc_api\/backend_server_grpc_api\/backend_grpc_server.proto

	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I ../grpc_api/backend_server_grpc_api -I ../grpc_api/common_grpc_api --go_out=plugins=grpc:../grpc_api/backend_server_grpc_api ../grpc_api/backend_server_grpc_api/backend_grpc_server.proto

	sed -i 's/import "grpc_common.proto";/\/\/import "grpc_common.proto";/' ..\/grpc_api\/backend_server_grpc_api\/backend_grpc_server.proto
	sed -i 's/\/\/import "..\/common_grpc_api\/grpc_common.proto";/import "..\/common_grpc_api\/grpc_common.proto";/' ..\/grpc_api\/backend_server_grpc_api\/backend_grpc_server.proto

