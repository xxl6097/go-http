package test

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
	"go.uber.org/zap"

	"log"
	"net/http"
)

type ClinkWS struct {
	clients map[string]*websocket.Conn
}

// 用于升级HTTP连接到WebSocket连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有的源，生产环境下需要配置跨域策略
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *ClinkWS) onMessageRecv(ws *websocket.Conn) {
	for {
		// 读取消息
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			z.L().Debug("websocket client 断开", zap.Int("messageType", messageType), zap.Error(err))
			//delete(this.clients, ws.RemoteAddr().String())
			break
		} else {
			z.L().Debug("Received", zap.String("RemoteAddr", ws.RemoteAddr().String()), zap.Int("messageType", messageType), zap.ByteString("message", message))
		}

	}
}

func (this *ClinkWS) Send(uuid string, payload []byte) error {
	v, ok := this.clients[uuid]
	if ok && v != nil {
		err := v.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			return err
		}
	}
	return nil
}

// 处理WebSocket连接
func (this *ClinkWS) handleConnections(w http.ResponseWriter, r *http.Request) {
	// 将HTTP请求升级为WebSocket协议
	z.L().Debug("将HTTP请求升级为WebSocket协议", zap.Any("url", r.URL.Query()))
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	uuid := r.URL.Query().Get("uuid")
	z.L().Debug("Client Connected", zap.String("uuid", uuid), zap.Any("r", r))
	this.clients[uuid] = ws
	// 不断读取客户端消息
	this.onMessageRecv(ws)
	delete(this.clients, uuid)
}

func (s ClinkWS) Setup(router *mux.Router) {
	router.HandleFunc("/ws", s.handleConnections).Methods("GET")
}

func NewWsRoute() ihttpserver.IRoute {
	opt := ClinkWS{
		clients: make(map[string]*websocket.Conn),
	}
	return opt
}
