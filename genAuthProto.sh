#这个是创建auth服务的proto文件的脚本，生成对应代码

#设置文件目录
PROTO_PATH=./author/api
GO_OUT_PATH=./author/api/gen/v1
#确保文件目录存在,如果不存在则创建
mkdir -p $GO_OUT_PATH

protoc -I=. --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ./author/auth.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/auth.yaml:$GO_OUT_PATH ./author/auth.proto
 
#pbjs可执行文件目录
PBTS_BIN_DIR=./process_node_file
#输出目录
PBTS_OUT_DIR=./author/output_jsfile 
#确保文件目录存在，如果不存在则创建
mkdir -p $PBTS_OUT_DIR

#-w es6使用es6的语法，auth.proto是当前目录下要有这个文件
$PBTS_BIN_DIR/pbjs -t static -w es6 ./author/auth.proto --no--create --no--encode --no--decode --no--delimited -o $PBTS_OUT_DIR/auth_pb.js
$PBTS_BIN_DIR/pbts -o $PBTS_OUT_DIR/auth_pb.d.ts $PBTS_OUT_DIR/auth_pb.js