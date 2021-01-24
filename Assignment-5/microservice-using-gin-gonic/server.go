package main
import (
"github.com/gin-gonic/gin"
"gitlab.com/pragmaticreviews/golang-gin-poc/service"
"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
)


var(
videoService service.VideoService =service.New()
videoController controller.VideoController =controller.New(videoService)

)

func main(){
server := gin.New()
server.Use(gin.Recovery())
server.Use(gin.Logger())

server.GET("/videos", func(ctx *gin.Context){
ctx.JSON(200, videoController.FindAll())
})

server.POST("/videos", func(ctx *gin.Context){
ctx.JSON(200, videoController.Save(ctx))
})
server.Run(":8091")
}