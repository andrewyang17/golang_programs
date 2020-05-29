## About
- This program takes files from command-line arguments and looks for adjacent duplicates lines.
- There are two text files which contain a list of movies respectively,  that can be used in running the program.

## Command
```
go run main.go <...*.txt>
go run main.go example1.txt example2.txt
```

## Logic
Instead of operating in a "streaming mode", a better approach is to read the entire input into memory in one big gulp,
split it into lines all at once, then process the lines. 

## Result
```
// Return result
// on left: the number of duplication
// on right: the name of movie

3       Reversal of Fortune
3       Road House
2       Rock, The
3       Rising Sun
3       Replacements, The
3       Resident Evil: Apocalypse
3       Ritz, The
2       Robinson Cruesoeland
```