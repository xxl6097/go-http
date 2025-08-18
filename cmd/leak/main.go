package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

// Server 结构体管理 HTTP 服务状态
type Server struct {
	server     *http.Server
	shutdownWG sync.WaitGroup
}

// NewServer 创建并配置 HTTP 服务器
func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 注册路由处理器
	mux.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/leaky", leakyHandler) // 故意泄漏的端点

	return &Server{server: srv}
}

// 模拟泄漏的处理器：启动永不退出的 Goroutine
func leakyHandler(w http.ResponseWriter, r *http.Request) {
	//ch := make(chan struct{})
	//go func() {
	//	ch <- struct{}{} // 永久阻塞（无接收方）
	//}()
	w.Write([]byte("Leaky endpoint triggered"))
}

// Start 异步启动 HTTP 服务
func (s *Server) Start() {
	s.shutdownWG.Add(1)
	go func() {
		defer s.shutdownWG.Done()
		log.Printf("Server starting on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
}

// Stop 安全关闭服务（含超时控制）
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	s.shutdownWG.Wait()
	log.Println("Server stopped")
}
