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
	var first bool
	var firstS string
	for {
		firstS = goFirst()
		if firstS == "yes" || firstS == "Yes" || firstS == "y" || firstS == "Y" || firstS == "no" || firstS == "No" || firstS == "n" || firstS == "N" {
			if firstS == "yes" || firstS == "Yes" || firstS == "y" || firstS == "Y" {
				first = true
				break
			} else {
				first = false
				break
			}
			if first == true {
				break
			}
		}
	}

	if first == true{
		setPlayer(team)
	}else {
		setPlayer(opposite(team))
	}

	for {
		playGame()
		if (playAgain() == false){
			break
		}
	}



}

//function to ask user to choose a team
//return string team
func chooseTeam() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you wanna be X's or O's?")
	var team string
	team, _ = reader.ReadString('\n')
	team = strings.TrimRight(team, "\r\n")
	return team
}

func capitalize(letter string) string{
	if letter == "x" || letter == "X"{
		return "X"
	}else {
		return "O"
	}
}

func opposite(letter string) string{
		if letter == "X"{
			return "O"
		}else {
			return "X"
		}
}

//ask the player how many people are playing
//ToDo: give option of 0,1,2
//return string playnumber
func choosePlayers() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many Players?")
	var playerNumber string
	playerNumber, _ = reader.ReadString('\n')
	playerNumber = strings.TrimRight(playerNumber, "\r\n")
	return playerNumber
}

//ask the player to answer yes or no if they want to go first
//return string playerFirst
func goFirst() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you wanna go first? yes or no?")
	var playerFirst string
	playerFirst, _ = reader.ReadString('\n')
	playerFirst = strings.TrimRight(playerFirst, "\r\n")
	return playerFirst
}

func clearBoard() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			board[i][j] = " "
		}
	}

}

func playGame() {
	for {
		if checkWin() == true{
			break
		}
		askInput()
		printboard()
		changePlayer()
	}
	displayWinner()
}

func updateBoard(xlocation int, ylocation int, tile string) {
	board[xlocation][ylocation] = tile
}

func printboard() {
	fmt.Println(board[0][0] + "|" + board[0][1] + "|" + board[0][2])
	fmt.Println("-----------")
	fmt.Println(board[1][0] + "|" + board[1][1] + "|" + board[1][2])
	fmt.Println("-----------")
	fmt.Println(board[2][0] + "|" + board[2][1] + "|" + board[2][2])
}

func placeTile(location int, tile string) {
	xlocation := (location - 1) / 3
	ylocation := (location -1) % 3
	updateBoard(xlocation, ylocation, tile)
}

func setPlayer(team string){
	currentPlayer = capitalize(team)
}

func changePlayer(){
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else if currentPlayer== "O" {
		currentPlayer = "X"
	}
}

func checkWin() bool {
	return false
}

func playAgain() bool {
	return false
}

func askInput() {
	for{
		fmt.Println("Current Player: " + currentPlayer)
		fmt.Println("Where do you want to go? 1-9")
		var choice int
		fmt.Scan(&choice)
		if choice < 0 && choice > 10 {
				fmt.Println("I said 1-9, try again")
		}else if checkTile(choice) != " "{
			fmt.Println("That spot is taken try again")
		}else {
			placeTile(choice, currentPlayer)
			break
		}
	}

}

func checkTile(location int) string{
	xlocation := (location - 1) / 3
	ylocation := (location -1) % 3
	return board[xlocation][ylocation]
}

func error() {}

func checkPlayAgain() {}

func displayWinner() {}
