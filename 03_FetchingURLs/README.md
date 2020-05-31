## About
- This program takes urls from command-line arguments and discards the responses but reports the size and elapsed time for each one.
- It fetches many urls, all concurrently, using goroutine and channel, so that the process
will take no longer than the longest fetch rather than the sum of all the fetch times.

## Command
```
$ go run main.go <... urls>
$ go run main.go https://www.google.com https://godoc.org https://www.youtube.com
```

## Knowledge
- A goroutine is a concurrent function execution.
- A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.
- The function main runs in a goroutine and the go statement creates additional goroutines.
 

## Result
```
// Return result
// Time taken, size and url

0.14s   12308 https://www.google.com
0.31s    7292 https://godoc.org
0.83s  362890 https://www.youtube.com
0.83s elapsed

```