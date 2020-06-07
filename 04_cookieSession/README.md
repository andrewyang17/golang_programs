## About
- This program is to demonstrate how to create and retrieve cookie, assign uuid into it as a value, 
and using it in a session. 
- A local server will be created in port 8080 after running the program.
- This is just to enhancing the fundamental concept and to learn the best go practice.

## Command
```
$ go get github.com/satori/go.uuid
$ go run main.go 
```

## Logic
- A form submission is used to create value of type "user", assign the field username
  into the cookie's value, and then assign the value of type "user" as a value to the key username
  in map dbUsers.
- Once the form is submitted, a cookie will be created, where you can check in ->
developer tools -> Application -> Cookies -> http://localhost:8080 -> name = ("session"),
value = (uuid)
- Without submitting the form, which means the cookie doesn't exist in client's browser,
users can't access to "/bar" url, they will be redirected to index by 303 (StatusSeeOther).

