package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var board = [3][3]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "}}
var currentPlayer = " "
var winner = " "
var xHuman = 0
var yHuman = 0

func main() {
	printboard()
	//Choose wether to be X's or O's
	var team string
	for {
		team = chooseTeam()
		if team == "x" || team == "X" || team == "O" || team == "o" {
			break
		}
		fmt.Println("Please choose either X's or O's")
	}
	//Choose wether to go first or second
	//goFirst asks if you want to go first so yes is first no is second
	//check validity of answer then set the true value of going first
	var first = getYesNo()
	if first == true {
		setPlayer(team)
	} else {
		setPlayer(opposite(team))
	}

	for {
		playGame()
		if playAgain() == false {
			break
		}
		clearBoard()
		printboard()
	}
}

func playGame() {
	for {
		if checkWin() == true {
			break
		}
		askInput()
		printboard()
		changePlayer()
	}
	displayWinner()
}

func choosePlayers() string {
	//ask the player how many people are playing
	//return string playnumber
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many Players?")
	var playerNumber string
	playerNumber, _ = reader.ReadString('\n')
	playerNumber = strings.TrimRight(playerNumber, "\r\n")
	return playerNumber
}

func chooseTeam() string {
	//function to ask user to choose a team
	//return string team
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you wanna be X's or O's?")
	var team string
	team, _ = reader.ReadString('\n')
	team = strings.TrimRight(team, "\r\n")
	return team
}

func goFirst() string {
	//ask the player to answer yes or no if they want to go first
	//return string playerFirst
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you wanna go first? yes or no?")
	var playerFirst string
	playerFirst, _ = reader.ReadString('\n')
	playerFirst = strings.TrimRight(playerFirst, "\r\n")
	return playerFirst
}

func askInput() {
	for {
		fmt.Println("Current Player: " + currentPlayer)
		fmt.Println("Where do you want to go? 1-9")
		var choice int
		fmt.Scan(&choice)
		if choice < 0 && choice > 10 {
			fmt.Println("I said 1-9, try again")
		} else if checkTile(choice) != " " {
			fmt.Println("That spot is taken try again")
		} else {
			placeTile(choice, currentPlayer)
			break
		}
	}

}

func getYesNo() bool {
	reader := bufio.NewReader(os.Stdin)
	var response string
	for {
		response, _ = reader.ReadString('\n')
		response = strings.TrimRight(response, "\r\n")
		if response == "yes" || response == "Yes" || response == "y" || response == "Y" || response == "no" || response == "No" || response == "n" || response == "N" {
			if response == "yes" || response == "Yes" || response == "y" || response == "Y" {
				return true
			} else {
				return false
			}
		} else {
			fmt.Println("Please answer yes or no")
		}
	}
}

func printboard() {
	fmt.Println(board[0][0] + "|" + board[0][1] + "|" + board[0][2])
	fmt.Println("-----------")
	fmt.Println(board[1][0] + "|" + board[1][1] + "|" + board[1][2])
	fmt.Println("-----------")
	fmt.Println(board[2][0] + "|" + board[2][1] + "|" + board[2][2])
}

func clearBoard() {
	//replaces each tile in the board with an empty string
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			board[i][j] = " "
		}
	}
}

func updateBoard(xlocation int, ylocation int, tile string) {
	board[xlocation][ylocation] = tile
}

func placeTile(location int, tile string) {
	xlocation := (location - 1) / 3
	ylocation := (location - 1) % 3
	updateBoard(xlocation, ylocation, tile)
}

func setPlayer(team string) {
	currentPlayer = capitalize(team)
}

func getWinner() string {
	var sum int = 0
	for i := 1; i <= 3; i++ {
		sum = rowSum(i)
		if sum > 2 {
			return "0"
		} else if sum < -2 {
			return "X"
		}
	}
	for i := 1; i <= 3; i++ {
		sum = columnSum(i)
		if sum > 2 {
			return "0"
		} else if sum < -2 {
			return "X"
		}
	}
	for i := 1; i <= 2; i++ {
		sum = diagonalSum(i)
	}
	if sum > 2 {
		return "0"
	} else if sum < -2 {
		return "X"
	}
	return " "
}

func checkWin() bool {
	if getWinner() != " " {
		return true
	} else {
		return false
	}
}

func displayWinner() {
	fmt.Println()
	fmt.Println("Congrats!" + getWinner() + " won the game")
	fmt.Println()
}

func tileNumber(location int) int {
	if board[convertX(location)][convertY(location)] == "O" {
		return 1
	} else if board[convertX(location)][convertY(location)] == "X" {
		return -1
	} else {
		return 0
	}
}

func rowSum(row int) int {
	return tileNumber(row*3) + tileNumber(row*3-1) + tileNumber(row*3-2)
}

func columnSum(column int) int {
	return tileNumber(column) + tileNumber(column+3) + tileNumber(column+6)
}

func diagonalSum(diagonal int) int {
	if diagonal > 0 {
		return tileNumber(7) + tileNumber(5) + tileNumber(3)
	} else {
		return tileNumber(1) + tileNumber(5) + tileNumber(9)
	}
}

func capitalize(letter string) string {
	if letter == "x" || letter == "X" {
		return "X"
	} else {
		return "O"
	}
}

func opposite(letter string) string {
	if letter == "X" {
		return "O"
	} else {
		return "X"
	}
}

func convertX(location int) int {
	return (location - 1) / 3
}

func convertY(location int) int {
	return (location - 1) % 3
}

func checkTile(location int) string {
	return board[convertX(location)][convertY(location)]
}

func changePlayer() {
	currentPlayer = opposite(currentPlayer)
}

func playAgain() bool {
	fmt.Println("Do you want to play again?")
	return getYesNo()
}
