package Models

type Book struct {
	Id				uint		`json:"id"`
	Title			string	`json:"name"`
	PageCount int			`json:"page_count"`
}

func (b *Book) TableName() string {
	return "book"
}