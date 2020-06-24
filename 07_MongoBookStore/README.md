## About
- This program is exactly same as program from 06_postgresBookStore, the only difference is it binds to MongoDB.
- A good practice of code organization leads us to change only two files (config/db.go and books/models.go) and imports
in the previous program. It's awesome, because we swapped out entire database with a little touch.  
- The main focus is to get the idiomatic way of writing CRUD operations in MongoDB with Go.

## Requirements
- MongoDB is installed on local machine and running.
- OPTIONAL: Use systemd to start the mongod server for the ease of development, rather than open a terminal to run mongod, 
and another shell to run mongo.

```
# enter into mongo shell
$ mongo

# create db
$ use bookstore

# create collection books & insert pre-documents
$ db.books.insert([{"isbn":"978-1505255607","title":"The Time Machine","author":"H. G. Wells","price":5.99},{"isbn":"978-1503261960","title":"Wind Sand \u0026 Stars","author":"Antoine","price":14.99},{"isbn":"978-1503261961","title":"West With The Night","author":"Beryl Markham","price":14.99}])

# test
$ show collections
$ db.books.find()

# create user
# it's advised not to include special characters in user password
# otherwise it would get you a gotcha
$ db.createUser(
    {
      user: "<username>",
      pwd: "<user password>",
      roles: [ { role: "readWrite", db: "bookstore" } ]
    }
  )

# authenticate user
$ db.auth("<username>", "user password")

# check status of authentication
$ db.runCommand({connectionStatus : 1})

# go into config directory and replace your credentials in the following statement in db.go
# alternatively you could use environment variables to hide your credentials. 
s, err := mgo.Dial("mongodb://<username>:<user password>@localhost/bookstore")

# run the app
# check localhost:8080 in browser
$ go run main.go

# or using curl
$ curl -i localhost:8080/books
$ curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process

```
