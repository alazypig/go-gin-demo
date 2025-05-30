package main

import "encoding/json"

// hub 结构体用于管理WebSocket连接
type hub struct {
	c map[*connection]bool // 保存所有活动的WebSocket连接
	b chan []byte          // 广播消息的通道
	r chan *connection     // 注册新连接的通道
	u chan *connection     // 注销连接的通道
}

// 创建一个全局的hub实例
var h = hub{
	c: make(map[*connection]bool),
	b: make(chan []byte),
	r: make(chan *connection),
	u: make(chan *connection),
}

// hub的主运行循环，处理连接的注册、注销和消息广播
func (h *hub) run() {
	for {
		select {
		case c := <-h.r: // 处理新连接注册
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String() // 设置客户端IP
			c.data.Type = "handshake"              // 设置消息类型为握手
			c.data.UserList = user_list            // 设置在线用户列表
			send_data, _ := json.Marshal(c.data)   // 将数据转换为JSON
			c.sc <- send_data                      // 发送给新连接的客户端

		case c := <-h.u: // 处理连接注销
			if _, ok := h.c[c]; ok {
				delete(h.c, c) // 从活动连接map中删除
				close(c.sc)    // 关闭发送通道
			}

		case data := <-h.b: // 处理广播消息
			for c := range h.c { // 遍历所有活动连接
				select {
				case c.sc <- data: // 尝试发送消息
				default: // 如果发送失败
					delete(h.c, c) // 删除该连接
					close(c.sc)    // 关闭其发送通道
				}
			}
		}
	}
}
