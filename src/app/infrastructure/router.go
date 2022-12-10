package infrastructure

import (
	"github.com/gin-contrib/cors"//apiへのアクセスを許可するURL, method(GET, POST), header情報を設定できる
	"github.com/gin-gonic/gin"
	"github.com/higayu624/portfolio_go/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	r := gin.Default()//*Engineインスタンスを生成する(Endpoint, Middlewareを操作できる)

	r.Use(cors.Default())//apiへのアクセスを許可するURL, method(GET, POST), header情報を設定できる

	taskController := controllers.NewTaskController(NewSqlHandler())//構造体を初期化

	//c *gin.Contextはurlのqueryの取得、POSTで送信されたjsonの取得などを行うことができる。
	r.GET("/api/tasks", func(c *gin.Context) {taskController.Index(c)})//第一引数にEndpoint, 第二引数にControllerと呼ばれる関数を指定

	Router = r//おまじない？
}