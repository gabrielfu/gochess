package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	gochess "gochess/internal"

	"github.com/inancgumus/screen"
	"github.com/urfave/cli/v2"
)

func runCli(ctx *cli.Context) error {
	flip := ctx.Bool("flip")
	eval := ctx.Bool("eval")
	whiteEngine := ctx.Bool("white")
	blackEngine := ctx.Bool("black")
	flag.Parse()
	return runGame(flip, eval, whiteEngine, blackEngine)
}

func runGame(flip bool, eval bool, whiteEngine bool, blackEngine bool) error {
	screen.Clear()
	reader := bufio.NewReader(os.Stdin)
	var errMsg string

	g := gochess.NewGame()
	depth := 4

	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println(gochess.Banner())
		fmt.Println()
		if flip && g.Board().Turn() == gochess.BLACK {
			fmt.Println(g.VisualizeFlipped())
		} else {
			fmt.Println(g.Visualize())
		}

		if eval {
			evalScore := gochess.Evaluate(g.Board(), g.Status())
			bar := gochess.EvaluationBar(evalScore, 18)
			fmt.Println()
			fmt.Printf("%s %.2f\n", bar, float32(evalScore)/100)
		}

		fmt.Println()
		fmt.Println(g.PGN())

		if g.Status() == gochess.WhiteWon {
			fmt.Printf("\033[0;32mGame over! White won.\033[0;39m\n")
			break
		} else if g.Status() == gochess.BlackWon {
			fmt.Printf("\033[0;32mGame over! Black won.\033[0;39m\n")
			break
		} else if g.Status() == gochess.Draw {
			fmt.Printf("\033[0;32mGame over! %s.\033[0;39m\n", g.DrawReason())
			break
		}

		fmt.Printf("\033[0;31m%s\033[0;39m\n", errMsg)

		playerTurn := (g.Turn() == gochess.WHITE && !whiteEngine) || (g.Turn() == gochess.BLACK && !blackEngine)
		if playerTurn {
			fmt.Println("Input a move in Standard Algebraic Notation, or 'undo' to undo a move.")
			fmt.Printf("Your move (%s): ", g.Turn())

			input, err := reader.ReadString('\n')
			if err != nil {
				errMsg = "Error reading input: " + err.Error()
				continue
			}
			input = input[:len(input)-1]

			if input == "undo" {
				// undo a full move
				// undo twice because the last move was the opponent's
				if err := g.Undo(); err != nil {
					errMsg = err.Error()
					continue
				}

				if err := g.Undo(); err != nil {
					errMsg = err.Error()
				}
				continue
			}

			move, err := gochess.ParseSAN(input, g.Board())
			if err != nil {
				errMsg = err.Error()
				continue
			}

			if err := g.Move(move); err != nil {
				errMsg = err.Error()
				continue
			}
		} else {
			fmt.Print("Engine is thinking...")
			result := gochess.Search(g.Board(), depth)
			if result.Move() == nil {
				errMsg = "Internal Error: Engine could not find a move!"
				break
			}
			fmt.Println(result.Move().ToSAN(g.Board()))
			g.Move(result.Move())
			time.Sleep(50 * time.Millisecond)
		}

		// reset error message
		errMsg = ""
	}

	fmt.Printf("\033[0;31m%s\033[0;39m\n", errMsg)
	fmt.Print("Press Enter key to exit...")
	reader.ReadByte()
	return nil
}

func main() {
	app := &cli.App{
		Name:  "gochess",
		Usage: "gochess",
		Commands: []*cli.Command{
			{
				Name:  "cli",
				Usage: "play the game on CLI",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "flip", Aliases: []string{"f"}, Usage: "flip the board at Black's turn"},
					&cli.BoolFlag{Name: "eval", Aliases: []string{"e"}, Usage: "show the evaluation bar"},
					&cli.BoolFlag{Name: "white", Aliases: []string{"w"}, Usage: "white will be played by an engine"},
					&cli.BoolFlag{Name: "black", Aliases: []string{"b"}, Usage: "black will be played by an engine"},
				},
				Action: runCli,
			},
			{
				Name:  "ui",
				Usage: "play the game on a graphic UI",
				Action: func(ctx *cli.Context) error {
					return errors.New("ui not implemented")
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
