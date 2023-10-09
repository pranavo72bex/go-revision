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

func (n *NotesController) InitController(notesService services.NotesServices) *NotesController {
	n.NotesServices = notesService
	return n
}
func (n *NotesController) InitRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.POST("/", n.PostNotes())
	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetSingleNotes())
	notes.DELETE("/:id", n.DeleteNote())
	notes.PUT("/", n.UpdateNote())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		status := ctx.Query("status")
		var actualStatus *bool
		if status != "" {
			aS, err := strconv.ParseBool(status)
			actualStatus = &aS
			if err != nil {
				utils.HandleError(ctx, 400, err)
			}
		}

		notes, err := n.NotesServices.GetNotesService(actualStatus)

		if err != nil {
			utils.HandleError(ctx, 400, err)
			return
		}
		ctx.JSON(200, gin.H{
			"notes": notes,
		})
	}
}
func (n *NotesController) GetSingleNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			utils.HandleError(ctx, 400, err)
		}

		notes, err := n.NotesServices.GetSingleNotesService(noteId)

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

func (n *NotesController) DeleteNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			utils.HandleError(ctx, 400, err)
		}
		err = n.NotesServices.DeleteNoteServices(noteId)
		if err != nil {
			utils.HandleError(ctx, 400, err)
		}
		ctx.JSON(200, gin.H{
			"message": "note has been deleted",
		})
	}
}

func (n *NotesController) UpdateNote() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var noteBody NoteBody

		if err := ctx.BindJSON(&noteBody); err != nil {
			utils.HandleError(ctx, 400, err)
			return
		}
		note, err := n.NotesServices.UpdateNoteServices(noteBody.Title, noteBody.Status, noteBody.Id)
		if err != nil {
			utils.HandleError(ctx, 400, err)
		}
		ctx.JSON(200, gin.H{
			"notes": note,
		})
	}
}
