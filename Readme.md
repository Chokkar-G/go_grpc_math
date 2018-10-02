
Math App - Go Grpc

To Build proto file:

go get -u github.com/Chokkar-G/go_grpc_math/...

go_grpc_math\mathpb>protoc math.proto --go_out=plugins=grpc:.

client.exe -Operation=add -Param1=10 -Param2=20
