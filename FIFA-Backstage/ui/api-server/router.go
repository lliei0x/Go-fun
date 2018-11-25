package api_server

import (
	"fmt"
	"net/http"

	"leeif.me/Go-fun/FIFA-Backstage/ui/api-server/matches"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type APIServer struct {
	engine *gin.Engine
}

func (a *APIServer) registry() {
	APIServerInit(a.engine)

}

func (a *APIServer) init() {

}

func (a *APIServer) Start() {
	a.registry()
	a.engine.Run(":8080")
}

func New() *APIServer {
	return &APIServer{
		engine: gin.Default(),
	}
}

type Welcome struct {
	Greet string `json:"greet" binding:"required"`
	Words string `json:"words" binding:"required"`
}

func APIServerInit(r *gin.Engine) {
	// docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/welcome", func(context *gin.Context) {
		var welcome Welcome
		if err := context.ShouldBindJSON(&welcome); err == nil {
			if welcome.Greet == "FIFA-World-Cup" && welcome.Words == "Hello World" {
				context.SecureJSON(
					http.StatusOK,
					gin.H{
						"status": fmt.Sprintf("%s : %s", welcome.Words, welcome.Greet),
					},
				)
			} else {
				context.JSON(
					http.StatusAccepted,
					gin.H{
						"err": err.Error(),
					},
				)
			}
		}
	})
	v1 := r.Group("/v1/api")
	indexRegistry(v1)
	matchesRegistry(v1)

}

func indexRegistry(r *gin.RouterGroup) {
	r.GET("", HelloWorld)

}

func HelloWorld(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello World! FIFA world Cup 2018"},
	)
}

func matchesRegistry(r *gin.RouterGroup) {
	r.GET("/matches/:matchID", matches.HandlerGetMatchesByID)
}
