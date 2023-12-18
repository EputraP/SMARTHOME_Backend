package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()


	port := "8080"

	if port == "" {
		port = "8080"
	}


	if err:=srv.Run(fmt.Sprintf(":%s", port)); err != nil{
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)

	}
	fmt.Println("Hello, World!")
}