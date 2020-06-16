

package routers

import (
	"webst/servers/websocket"
)

// Websocket 路由
func WebsocketInit() {
	websocket.Register("login", websocket.LoginController)
	websocket.Register("heartbeat", websocket.HeartbeatController)
}
