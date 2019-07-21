package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type ApiServer struct {
	engine *gin.Engine
}

func (server *ApiServer) register() {

	apiServerInit(server.engine)
}

//路由注册
func apiServerInit(r *gin.Engine) {

	//swagger
	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.GET("/hello",Hello)

	v1 := r.Group("/api/v1")
	v1.GET("/record/:userId", record)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (server *ApiServer) Start() {
	server.register()
	server.engine.Run(":5000")
}

//create server
func New() *ApiServer {
	return &ApiServer{
		engine: gin.Default(),
	}
}
