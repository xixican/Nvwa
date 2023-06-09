package nvwa_http

import (
	"Nvwa/web"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

func (h *HttpServer) StartServer(address string) {
	router := gin.Default()
	agentRoute := &web.AgentRoute{}
	agentRoute.InitAgentRoute(router)
	router.Run(address)
	//http.ListenAndServe(":8890", router)
}
