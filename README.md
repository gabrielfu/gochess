# gochess

A toy chess game playable against human and engine on terminal.

## Usage

```shell
go run cmd/main.go cli

Options:
    -h, --help   help for cli
    -e, --eval   show evaluation bar under the chessboard
    -f, --flip   flip the chessboard at black's turn
    -w, --white  White will be played by engine
    -b, --black  Black will be played by engine
```

## Interface

A chessboard will be displayed on the terminal. 
If `-e` is specified, an evaluation bar will be shown below it.
The current PGN is also displayed.

![Interface](interface.png)

When it is your turn, you can input one of the following to the prompt:
- A move in [Standard Algebraic Notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess))
- `undo`: This will undo a full move (i.e., your opponent's and your move)
