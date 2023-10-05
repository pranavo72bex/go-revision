package controllers

import (
	"example.com/services"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NotesServices services.NotesServices
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesServices) {
	notes := router.Group("/notes")
	notes.POST("/", n.GetNotes())
	n.NotesServices = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	return func(ctx *gin.Context) {
		var noteBody NoteBody

		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.NotesServices.PostNoteServices(noteBody.Title, noteBody.Status)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
		}
		ctx.JSON(200, gin.H{
			"notes": note,
		})
	}
}
