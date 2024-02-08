package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BoardSize = 3
)

var (
	board  [BoardSize][BoardSize]string
	player = "X"
)

func main() {
	initializeBoard()
	printBoard()

	for {
		fmt.Printf("Player %s, enter your move (row,col): ", player)
		row, col := getPlayerMove()
		if makeMove(row, col, player) {
			printBoard()
			if checkWin(player) {
				fmt.Printf("Player %s wins!\n", player)
				return
			} else if checkDraw() {
				fmt.Println("It's a draw!")
				return
			}
			switchPlayer()
		} else {
			fmt.Println("Invalid move, try again.")
		}
	}
}

func initializeBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	for i, row := range board {
		for j, cell := range row {
			fmt.Print(cell)
			if j < BoardSize-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < BoardSize-1 {
			fmt.Println("-----")
		}
	}
}

func getPlayerMove() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return -1, -1
	}
	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || row < 0 || row >= BoardSize || col < 0 || col >= BoardSize {
		return -1, -1
	}
	return row, col
}

func makeMove(row, col int, player string) bool {
	if board[row][col] == " " {
		board[row][col] = player
		return true
	}
	return false
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}

func checkWin(player string) bool {
	for i := 0; i < BoardSize; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func checkDraw() bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
