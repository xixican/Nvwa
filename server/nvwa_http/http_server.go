package nvwa_http

import (
	"Nvwa/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

func (h *HttpServer) StartServer(address string) {
	router := gin.Default()
	// 跨域中间件
	router.Use(cors.Default())
	//初始化router
	agentRoute := &web.AgentRoute{}
	agentRoute.InitAgentRoute(router)
	router.Run(address)
	//http.ListenAndServe(":8890", router)
}
