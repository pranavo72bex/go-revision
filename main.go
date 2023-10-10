package main

import (
	"fmt"
	"log"

	"example.com/controllers"
	internal "example.com/internal/database"
	services "example.com/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	r := gin.Default()
	db := internal.InitDB()
	notesService := &services.NotesServices{}
	notesService.InitService(db)
	if db == nil {
		fmt.Print("db init error")
	}
	NotesController := &controllers.NotesController{}
	NotesController.InitRoutes(r)
	NotesController.InitController(*notesService)

	//manual dependency injection
	// t := Title("hello")
	// p := NewPublisher(&t)
	// m := NewMainService(p)
	// m.Run()

	//with uber
	fx.New(
		fx.Provide(NewMainService),
		fx.Provide(fx.Annotate(NewPublisher, fx.As(new(IPublisher)))),
		fx.Provide(func() *Title {
			t := Title("hello fx")
			return &t
		}),
		fx.Invoke(func(service *MainService) {
			service.Run()
		}),
	).Run()
	r.Run(":8000")
}

type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("main program")
}

// Dependency

type IPublisher interface {
	Publish()
}
type Publisher struct {
	title *Title
}

func NewPublisher(title *Title) *Publisher {
	return &Publisher{title: title}
}

func (publisher *Publisher) Publish() {
	log.Print("publisher check ", *publisher.title)
}

// dependency
type Title string
