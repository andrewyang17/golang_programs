## About
- This program is nothing but a very simple bookstore app which bind to Postgres database.
- The main focus is to get the idiomatic way of writing CRUD operations in Postgres with Go.

## Requirements
- PostgreSQL is installed on local machine and running.

```
# get the Postgres's driver
$ go get github.com/lib/pq

# login into your postgres or as Superuser
# create a db
$ CREATE DATABASE bookstore;

# create user
$ CREATE USER <username> with PASSWORD '<user password>';

# grant privileges
$ GRANT ALL PRIVILEGES ON DATABASE bookstore to <username>;

# alter to superuser
$ ALTER USER <username> WITH SUPERUSER;

# switch into the bookstore database
$ \c bookstore

# create table
$ CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

# insert pre-records
$ INSERT INTO books (isbn, title, author, price) VALUES
  ('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
  ('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
  ('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

# exit the postgres from terminal
$ \q
$ exit

# go into config directory and replace your credentials in the following statement in db.go
# alternatively you could use environment variables to hide your credentials. 
DB, err = sql.Open("postgres", "postgres://<username>:<user's password>@localhost/bookstore?sslmode=disable")

# run the app
# check localhost:8080 in browser
$ go run main.go

# or using curl
$ curl -i localhost:8080/books
$ curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process

```
