package book

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)


type Book struct {
	gorm.Model
	ID 	  uuid.UUID  `gorm:"type:uuid;"`
	Title string    `json:"name"`
	Author string   `json:"author"`
	Rating int 		`json:"rating"`
}

type Books struct {
	Books []Book `json:"books"`
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid version 4
	book.ID = uuid.New()
	return
}

