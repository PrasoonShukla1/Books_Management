package books

import (
	"Exercise6/models"
	"database/sql"
	"errors"
)

type Handler struct {
	Db *sql.DB
}

func (d Handler) Create(b models.Book) error {
	result, err := d.Db.Exec("INSERT INTO Books(ISBN,title,author,genre,publication,year_of_publication) VALUES (?, ?,?,?,?,?)", b.ISBN, b.Title, b.Author, b.Genre, b.Publication, b.Year_of_publication)
	if err != nil {
		return errors.New("Error From Exec")
	}
	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("Errors from Rows Affected")
	}

	return nil
}
func (d Handler) Update(b models.Book) error {

}
func (d Handler) Delete(b models.Book) error {

}
func (d Handler) Read(b models.Book) error {

}