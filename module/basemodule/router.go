package basemodule

import (
	"github.com/gin-gonic/gin"
)

func WebInit() {
	engine := gin.Default()
	wechatStart(engine)
	chat(engine)
	command(engine)
	test(engine)
	by(engine)
	if err := engine.Run(":7777"); err != nil {
		return
	}
}
func by(engine *gin.Engine) {
	engine.GET("/by", func(context *gin.Context) {
		context.JSON(200, "")
	})
}

func test(engine *gin.Engine) {
	group := engine.Group("/test")
	group.GET("/hello", func(context *gin.Context) {
		context.JSON(200, "Hello World!")
	})
}

func wechatStart(engine *gin.Engine) {
	group := engine.Group("/wechat")
	group.GET("/start", func(c *gin.Context) {
		value := c.Query("echostr")
		c.JSON(200, value)
	})
}

func chat(engine *gin.Engine) {
	group := engine.Group("/chat")
	group.POST("/chat", func(context *gin.Context) {

	})
}

func command(engine *gin.Engine) {
	group := engine.Group("/cmd")
	group.GET("/exec", func(context *gin.Context) {
		context.Query("k")
		context.JSON(200, "ok")
	})
}
