package main

import (
	"fmt"
	"log"
	"os"

	dbstore "github.com/EputraP/SMARTHOME_Backend/internal/store/db"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	srv := gin.Default()
	godotenv.Load()
	port := os.Getenv("PORT")
	prepare()
	if port == "" {
		port = "8080"
	}
	
	if err:=srv.Run(fmt.Sprintf(":%s", port)); err != nil{
		log.Println("Error running gin server: ", err)
		log.Fatalln("Error running gin server: ", err)

	}
	fmt.Println("Hello, World!")
}

func prepare(){
	db := dbstore.Get()
	log.Println(db)
}