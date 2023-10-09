package main

import (
	"fmt"

	"example.com/controllers"
	internal "example.com/internal/database"
	services "example.com/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := internal.InitDB()
	notesService := &services.NotesServices{}
	notesService.InitService(db)
	if db == nil {
		fmt.Print("db init error")
	}
	fmt.Print(db)
	NotesController := &controllers.NotesController{}
	NotesController.InitRoutes(r)
	NotesController.InitController(*notesService)
	r.Run(":8000")
}
