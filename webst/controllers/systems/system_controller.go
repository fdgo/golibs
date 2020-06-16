package systems

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webst/common"
	"webst/controllers"
	"webst/servers/websocket"
	"runtime"
)

// 查询系统状态
func Status(c *gin.Context) {

	isDebug := c.Query("isDebug")
	fmt.Println("http_request 查询系统状态", isDebug)

	data := make(map[string]interface{})

	numGoroutine := runtime.NumGoroutine()
	numCPU := runtime.NumCPU()

	// goroutine 的数量
	data["numGoroutine"] = numGoroutine
	data["numCPU"] = numCPU

	// ClientManager 信息
	data["managerInfo"] = websocket.GetManagerInfo(isDebug)

	controllers.Response(c, common.OK, "", data)
}
