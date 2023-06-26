package response

import (
	"time"

	m_articles "km-kelas-e/model/articles"
)

// actually not needing to this because data model with response is same
type Article struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromModel(model m_articles.Article) Article {
	return Article{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Status:    model.Status,
		Title:     model.Title,
		Content:   model.Content,
	}
}

func FromModelSlice(model []m_articles.Article) []Article {
	var artArray []Article
	for key := range model {
		artArray = append(artArray, FromModel(model[key]))
	}
	return artArray
}
