package seeders

import (
	"km-kelas-e/config"
	m_articles "km-kelas-e/model/articles"
)

func SetArticle() {
	config.DB.Create(&m_articles.Article{
		Status:  true,
		Title:   "This is title article",
		Content: "This is content article",
	})
}
