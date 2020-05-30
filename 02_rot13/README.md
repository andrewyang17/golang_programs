## About
- This program is a rudimentary encryption. 
- It set up a tcp server and listening to port 8080, and what the server does is it takes input from the connection 
and rotate by 13 places for each of the characters.

## Usage
- ROT13 is an example of the encryption algorithm known as a Caesar cipher, attributed to Julius Caesar.
- The story goes that Julius Caesar would send messages to his generals in the field
and he would encrypt those messages by rotating a certain number of characters in the alphabet.

## Command
```
// set up the server and waiting to listen
$ go run main.go

// open a new terminal and use telnet to access the server
// and start writing sentences
$ telnet localhost 8080
```

## Logic
For a efficient execution, the program achieves concurrency by implementing goroutine,
which is a lightweight execution thread running in the background. The program takes input, lowercase them,
and convert into slice of bytes, which is based on ASCII. If the index of the char is less than 109, the index will be
added by 13, else if it's more than 109, the index will be minus by 13.

## Result
```
// Return result
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
wait for my signal to set up the bomb    
wait for my signal to set up the bomb - jnvg-sbe-zl-fvtany-gb-frg-hc-gur-obzo

jnvg-sbe-zl-fvtany-gb-frg-hc-gur-obzo
jnvg-sbe-zl-fvtany-gb-frg-hc-gur-obzo - wait:for:my:signal:to:set:up:the:bomb 

we will be sending supply in the midnight
we will be sending supply in the midnight - jr-jvyy-or-fraqvat-fhccyl-va-gur-zvqavtug 

jr-jvyy-or-fraqvat-fhccyl-va-gur-zvqavtug           
jr-jvyy-or-fraqvat-fhccyl-va-gur-zvqavtug - we:will:be:sending:supply:in:the:midnight 

```