package hud_test

import (
	"context"
	"os"
	"testing"

	"github.com/goexl/hud"
	"go.sichuancredit.cn/protocol/storage/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMultipart(t *testing.T) {
	// addr := "192.168.102.225:7987"
	addr := "localhost:9001"
	if connection, de := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials())); nil != de {
		t.Error("无法连接服务器")
	} else if ue := uploadBytes(connection); nil != ue {
		t.Error("上传二进制数组出错")
	}
}

func uploadFile(connection *grpc.ClientConn) (err error) {
	builder := hud.New().Build().Upload()
	client := file.NewRpcClient(connection)
	ctx := context.Background()
	app := 3
	filepath := "D:\\Apps\\ToDesk\\ToDesk.exe"
	filename := "ToDesk.File.exe"
	var id int64 = 0
	_lifecycle := newLifecycle(ctx, client, int64(app), filename, &id)
	err = builder.Filepath(filepath).Multipart().Lifecycle(_lifecycle).Build().Build().Do()

	return
}

func uploadBytes(connection *grpc.ClientConn) (err error) {
	builder := hud.New().Build().Upload()
	client := file.NewRpcClient(connection)
	ctx := context.Background()
	app := 3
	filepath := "D:\\Downloads\\四川金控建设规划咨询项目蓝图报告_v3_0315.docx"
	filename := "ToDesk.Bytes.docx"
	var id int64 = 0
	_lifecycle := newLifecycle(ctx, client, int64(app), filename, &id)
	if content, rfe := os.ReadFile(filepath); nil != rfe {
		err = rfe
	} else {
		err = builder.Bytes(content).Multipart().Lifecycle(_lifecycle).Build().Build().Do()
	}

	return
}
