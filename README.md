# makepass
Yet another password generator written in Go.
Uses high-entropy system calls and is pretty configurable.

![GitHub commit activity](https://img.shields.io/github/commit-activity/m/core6quad/makepass)


## usage
this program has some launch parameters avaliable 

"-nosymbols" - Won't include Symbols (!@#...) in the result

"-length number" - controls length of the generated password

"-file path" - saves password in a file instead of logging it

"-simpleout" - just outputs the password and nothing else, useful for being an API