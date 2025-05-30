package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// connection 结构体用于管理单个WebSocket连接
type connection struct {
	ws   *websocket.Conn // WebSocket连接对象
	sc   chan []byte     // 发送消息的通道
	data *Data           // 连接相关的数据
}

// writer 负责向WebSocket连接写入消息
func (c *connection) writer() {
	defer c.ws.Close()
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
}

// 存储在线用户列表
var user_list = []string{}

// reader 负责从WebSocket连接读取消息
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.u <- c // 发生错误时注销连接
			break
		}

		// 解析接收到的JSON消息
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login": // 处理登录消息
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			send_data, _ := json.Marshal(c.data)
			h.b <- send_data
		case "user": // 处理用户消息
			c.data.Type = "user"
			send_data, _ := json.Marshal(c.data)
			h.b <- send_data
		case "logout": // 处理登出消息
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			send_data, _ := json.Marshal(c.data)
			h.b <- send_data
			h.u <- c
		default:
			fmt.Println("----------- default -----------")
		}
	}
}

// WebSocket升级器配置
var wu = &websocket.Upgrader{
	ReadBufferSize:  512, // 读取缓冲区大小
	WriteBufferSize: 512, // 写入缓冲区大小
	CheckOrigin: func(r *http.Request) bool { // 允许所有来源
		return true
	},
}

// WebSocket处理函数
func myWs(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade failed:", err)
		return
	}

	// 创建新的连接
	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}
	h.r <- c // 注册新连接
	go c.writer()
	c.reader()

	// 延迟处理用户断开连接的情况
	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		send_data, _ := json.Marshal(c.data)
		h.b <- send_data
		h.u <- c
	}()
}

// 从切片中删除指定用户
func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}

	var n_slice []string
	for i := range slice {
		if slice[i] == user && i == count {
			// 如果要删除的用户在切片末尾
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}

	// fmt.Println(n_slice)
	return n_slice
}
