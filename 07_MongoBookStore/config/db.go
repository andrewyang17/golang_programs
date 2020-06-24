package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database
var Books *mgo.Collection

func init() {
	var err error
	s, err := mgo.Dial("mongodb://<username>:<user password>@localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")
	fmt.Println("Connected to database.")
}