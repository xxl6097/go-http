package test

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
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
			glog.Println("websocket client 断开:", messageType, err)
			//delete(this.clients, ws.RemoteAddr().String())
			break
		} else {
			glog.Printf("Received:%s %d %s\n", ws.RemoteAddr().String(), messageType, message)
			//this.clients[ws.RemoteAddr().String()] = ws
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
	glog.Error("将HTTP请求升级为WebSocket协议", r.URL.Query())
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	uuid := r.URL.Query().Get("uuid")
	glog.Println("Client Connected", uuid, r)
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
