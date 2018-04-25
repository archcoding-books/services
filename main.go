package main

import (
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
	"os/exec"
)

var(
	port int
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	flag.IntVar(&port, "port", 9090, "http listen port") 
	r := gin.Default()
	r.GET("/sync", func(c *gin.Context) {
		go func(){
			exec.Command("/archcoding_books/sync.sh")
		}()
		c.JSON(200, gin.H{
			"message": "syncing OK",
		})
	})
	r.Run(fmt.Sprintf(":%d", port)) 
}
