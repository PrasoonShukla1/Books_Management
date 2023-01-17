package books

import (
	"Exercise6/db"
	"Exercise6/models"
	"errors"
	"reflect"
	"testing"
)

func Test_Create(t *testing.T) {
	tc := []struct {
		b   models.Book
		err error
	}{
		{models.Book{ISBN: 10, Title: "Money", Author: "Ambani", Genre: "N-Fiction", Publication: "Pearson", Year_of_publication: 2009}, nil},
		{models.Book{ISBN: 10, Title: "Money", Author: "Ambani", Genre: "N-Fiction", Publication: "Pearson", Year_of_publication: 2009}, errors.New("Error From Exec")},
	}
	conn := db.Connecttomysql()
	d := Handler{Db: conn}
	for i, v := range tc {
		actualError := d.Create(v.b)
		if !reflect.DeepEqual(actualError, v.err) {
			t.Errorf("Test [%d] Failed.Expected:%v,Got:%v", i, v.err, actualError)
		}
	}
}
func Test_Update(t *testing.T) {
	tc := []struct {
		b   models.Book
		err error
	}{
		{models.Book{ISBN: 10}, nil},
		{models.Book{ISBN: 12}, errors.New("Error From Exec")},
	}
	conn := db.Connecttomysql()
	d := Handler{Db: conn}
	for i, v := range tc {
		actualError := d.Create(v.b)
		if !reflect.DeepEqual(actualError, v.err) {
			t.Errorf("Test [%d] Failed.Expected:%v,Got:%v", i, v.err, actualError)
		}
	}
}
func Test_Delete(t *testing.T) {
	tc := []struct {
		b   models.Book
		err error
	}{
		{models.Book{ISBN: 10}, nil},
		{models.Book{ISBN: 12}, errors.New("Error From Exec")},
	}
	conn := db.Connecttomysql()
	d := Handler{Db: conn}
	for i, v := range tc {
		actualError := d.Create(v.b)
		if !reflect.DeepEqual(actualError, v.err) {
			t.Errorf("Test [%d] Failed.Expected:%v,Got:%v", i, v.err, actualError)
		}
	}
}
func Test_Read(t *testing.T) {
	tc := []struct {
		b   models.Book
		err error
	}{
		{models.Book{ISBN: 10}, nil},
		{models.Book{ISBN: 15}, errors.New("Error From Exec")},
	}
	conn := db.Connecttomysql()
	d := Handler{Db: conn}
	for i, v := range tc {
		actualError := d.Create(v.b)
		if !reflect.DeepEqual(actualError, v.err) {
			t.Errorf("Test [%d] Failed.Expected:%v,Got:%v", i, v.err, actualError)
		}
	}
}
