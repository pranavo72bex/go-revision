package services

import (
	"errors"
	"fmt"

	internal "example.com/internal/models"
	"gorm.io/gorm"
)

// connect database here
type NotesServices struct {
	db *gorm.DB
}

func (n *NotesServices) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesServices) GetNotesService(status bool) ([]*internal.Notes, error) {

	var notes []*internal.Notes
	if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
		return nil, err

	}
	return notes, nil

}

func (n *NotesServices) PostNoteServices(title string, status bool) (*internal.Notes, error) {

	note := &internal.Notes{
		Title:  title,
		Status: status,
	}
	if note.Title == "" {
		return nil, errors.New("Title is required")
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return note, nil
}
