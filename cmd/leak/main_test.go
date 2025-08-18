package main

//
//import (
//	"github.com/xxl6097/go-http/cmd/app/test"
//	"github.com/xxl6097/go-http/pkg/httpserver"
//	"go.uber.org/goleak"
//	"net/http"
//	"testing"
//)
//
//func TestHealthyEndpoint(t *testing.T) {
//	defer goleak.VerifyNone(t) // 泄漏检测
//
//	// 启动测试服务器
//	server := NewServer(":8001")
//	server.Start()
//	defer server.Stop()
//
//	// 发送健康检查请求
//	resp, err := http.Get("http://localhost" + server.server.Addr + "/healthy")
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer resp.Body.Close()
//}
//
//func TestLeakyEndpoint1(t *testing.T) {
//	defer goleak.VerifyNone(t) // 预期此处检测到泄漏
//
//	server := NewServer(":8001")
//	server.Start()
//	defer server.Stop()
//
//	// 触发泄漏端点
//	resp, err := http.Get("http://localhost" + server.server.Addr + "/leaky")
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer resp.Body.Close()
//}
//
//func TestLeakyEndpoint(t *testing.T) {
//	defer goleak.VerifyNone(t) // 预期此处检测到泄漏
//	server := httpserver.New().
//		AddRoute(test.NewRoute(test.NewController())).
//		AddRoute(test.NewWsRoute()).
//		Done(8081)
//
//	defer server.Stop()
//	// 触发泄漏端点
//	resp, err := http.Get("http://localhost:8081/mqtt/test")
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer resp.Body.Close()
//}
//
//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m) // 全局泄漏检测
//}
