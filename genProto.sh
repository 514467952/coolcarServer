

function genProto {
    DOMAIN=$1
    SKIP_GETWAY=$1
    #设置文件目录
    PROTO_PATH=./${DOMAIN}/api
    GO_OUT_PATH=./${DOMAIN}/api/gen/v1
    #确保文件目录存在,如果不存在则创建
    mkdir -p $GO_OUT_PATH

    protoc -I=. --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ./${DOMAIN}/${DOMAIN}.proto
    
    # blob服务没有对外暴露，不提供yaml文件
    if [ $SKIP_GETWAY ]; then
        return
    fi

    protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml:$GO_OUT_PATH ./${DOMAIN}/${DOMAIN}.proto
 
    #pbjs可执行文件目录
    PBTS_BIN_DIR=./process_node_file
    #输出目录
    PBTS_OUT_DIR=./${DOMAIN}/output_jsfile 
    #确保文件目录存在，如果不存在则创建
    mkdir -p $PBTS_OUT_DIR
    #-w es6使用es6的语法，auth.proto是当前目录下要有这个文件
    $PBTS_BIN_DIR/pbjs -t static -w es6 ./${DOMAIN}/${DOMAIN}.proto --no--create --no--encode --no--decode --no--delimited --force-number -o $PBTS_OUT_DIR/${DOMAIN}_pb.js
    $PBTS_BIN_DIR/pbts -o $PBTS_OUT_DIR/${DOMAIN}_pb.d.ts $PBTS_OUT_DIR/${DOMAIN}_pb.js
}

genProto rental
genProto blob 2