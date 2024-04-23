package main

type Hub struct {
	clients          map[*Client]bool // 保存所有的client
	broadcast        chan []byte      // 广播数据
	clientRegister   chan *Client     // client注册channel
	clientUnRegister chan *Client     // client注销channel
}

func (hub *Hub) Run() {
	for {
		select {
		// 处理注册的client
		case client := <-hub.clientRegister:
			hub.clients[client] = true
		// 处理注销的client
		case client := <-hub.clientUnRegister:
			// 如果存在该client，就删除
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				// 删除client后，关闭client的chan管道
				close(client.send)
			}
		// 从广播中获取数据，如果有的话就广播给所有client
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					delete(hub.clients, client)
					// 删除client后，关闭client的chan管道
					close(client.send)
				}
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		clients:          make(map[*Client]bool),
		broadcast:        make(chan []byte),
		clientRegister:   make(chan *Client),
		clientUnRegister: make(chan *Client),
	}
}
