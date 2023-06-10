package main

import (
	handler "go-aerospike/handle"
	"log"
	"net/http"
	"time"

	_ "github.com/aerospike/aerospike-client-go"
	gin "github.com/gin-gonic/gin"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request from:", r.RemoteAddr, r.Header)
	time.Sleep(10 * time.Second)
	w.Write([]byte("ok"))
}

func main() {

	// client, err := as.NewClient("localhost", 3000)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/log", handler.StatementLog())
	// r.POST("/insert", handler.Insert(client))
	// r.GET("/get", handler.Get(client))
	// r.POST("delete", handler.Delete(client))
	// r.POST("/statement", handler.Statement(client))
	// r.POST("/statement-filter", handler.StatementFilerBinName(client))
	r.Run(":3000")
}
