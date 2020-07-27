package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// 允许等待的写入时间
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// 最大的连接ID，每次连接都加1 处理
var maxConnId int64

// 客户端读写消息
type wsMessage struct {
	// websocket.TextMessage 消息类型
	messageType int
	data        []byte
}

// ws 的所有连接
// 用于广播
var wsConnAll map[int64]*wsConnection

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有的CORS 跨域请求，正式环境可以关闭
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println("升级为websocket失败", err.Error())
		return
	}
	defer wsSocket.Close()
	maxConnId++
	// TODO 如果要控制连接数可以计算，wsConnAll长度
	// 连接数保持一定数量，超过的部分不提供服务
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnId,
	}
	wsConnAll[maxConnId] = wsConn
	log.Println("当前在线人数", len(wsConnAll))

	// 处理器,发送定时信息，避免意外关闭
	go wsConn.processLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

// 启动程序
func StartWebsocket(addrPort string) {
	wsConnAll = make(map[int64]*wsConnection)
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(addrPort, nil)
}

func main() {
	StartWebsocket("127.0.0.1:8000")
}
