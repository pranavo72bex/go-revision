package controllers

import (
	"strconv"

	"example.com/services"
	"example.com/utils"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NotesServices services.NotesServices
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesServices) {
	notes := router.Group("/notes")
	notes.POST("/", n.PostNotes())
	notes.GET("/", n.GetNotes())
	n.NotesServices = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		status := ctx.Query("status")

		stausFromStringToBool, err := strconv.ParseBool(status)
		if err != nil {
			utils.HandleError(ctx, 400, err)
			return
		}

		notes, err := n.NotesServices.GetNotesService(stausFromStringToBool)

		if err != nil {
			utils.HandleError(ctx, 400, err)
			return
		}
		ctx.JSON(200, gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) PostNotes() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	return func(ctx *gin.Context) {
		var noteBody NoteBody

		if err := ctx.BindJSON(&noteBody); err != nil {
			utils.HandleError(ctx, 400, err)
			return
		}
		note, err := n.NotesServices.PostNoteServices(noteBody.Title, noteBody.Status)
		if err != nil {
			utils.HandleError(ctx, 400, err)
		}
		ctx.JSON(200, gin.H{
			"notes": note,
		})
	}
}
