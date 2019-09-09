package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a, b, c, d, e, f, g, h, i = " ", " ", " ", " ", " ", " ", " ", " ", " "
	var s = " | "
	var line = "-----------"
	var line1 = a + s + b + s + c
	var line2 = d + s + e + s + f
	var line3 = g + s + h + s + i

	fmt.Println(line1)
	fmt.Println(line)
	fmt.Println(line2)
	fmt.Println(line)
	fmt.Println(line3)

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
	for (
		firstS = goFirst()
		if firstS == "yes" || firstS == "Yes" || firstS == "y" || firstS == "Y" || firstS == "no" || firstS == "No" || firstS == "n" || firstS == "N" {
			if firstS == "yes" || firstS == "Yes" || firstS == "y" || firstS == "Y" {
				first = true
				break
			}
			else{
				first = false
				break
			}
		}
	)
	
	//

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


