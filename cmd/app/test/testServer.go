package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Test() {
	// 创建自定义多路复用器
	mux := http.NewServeMux()
	// 注册路由处理函数
	mux.HandleFunc("/", rootHandler)
	// 配置 HTTP 服务器
	server := &http.Server{
		Addr:         ":8080",          // 监听端口
		Handler:      mux,              // 使用自定义多路复用器
		ReadTimeout:  10 * time.Second, // 读超时
		WriteTimeout: 10 * time.Second, // 写超时
		IdleTimeout:  60 * time.Second, // 空闲连接超时
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v\n", err)
	}

	// 设置关闭超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
	log.Println("Server exited")
}

// 根路径处理器
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Welcome to Go HTTP Server!\n")
	logRequest(r)
}

// 记录请求日志
func logRequest(r *http.Request) {
	log.Printf(
		"Request: %s %s - ClientIP: %s",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
	)
}
