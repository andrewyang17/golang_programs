## About
- This program takes the program in 04_cookieSession to the next level by incorporating the basic functionality of
login, sign up, logout, and clear sessions.
- A local server will be created in port 8080 after running the program.
- This is just to enhancing the fundamental concept and to learn the best go practice.

## Command
```
$ go get github.com/satori/go.uuid
$ go get golang.org/x/crypto/bcrypt
$ go run *.go 
```

## Logic
- "/bar" url is only accessible by users with role "007".
- The session length is set for 30 minutes in cookie's maxAge.
- Every 24 hours, a goroutine will be created to check and clean the sessions in runtime.