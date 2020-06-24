package books

import (
	"book/07_MongoBookStore/config"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"errors"
	"strconv"
)

type Book struct {
	Isbn string
	Title string
	Author string
	Price float32
}

func AllBooks() ([]Book, error) {
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400, bad request")
	}
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400, bad request")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. not acceptable")
	}
	bk.Price = float32(f64)

	// insert values
	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500, internal server error," + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400, bad request")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. not acceptable")
	}
	bk.Price = float32(f64)

	// insert values
	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400, bad request")
	}

	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500, internal server error")
	}
	return nil
}