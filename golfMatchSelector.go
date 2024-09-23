package main

import "fmt"

func players(){
	var playerCount int
	var player1 string
	var player2 string
	var player3 string
	var player4 string
	fmt.Print("How many players?")
	fmt.Scanf("%v", &playerCount)
	
	if playerCount == 4 {
		fmt.Print("Enter Player 1's name:")
		fmt.Scanf("%s", &player1)
		fmt.Print("Enter Player 2's name:")
		fmt.Scanf("%s", &player2)
		fmt.Print("Enter Player 3's name:")
		fmt.Scanf("%s", &player3)
		fmt.Print("Enter Player 4's name:")
		fmt.Scanf("%s", &player4)
		fmt.Printf("Golfers: %s, %s, %s, and %s\n", player1, player2, player3, player4)
	} else if playerCount == 3 {
		fmt.Print("Enter Player 1's name:")
		fmt.Scanf("%s", &player1)
		fmt.Print("Enter Player 2's name:")
		fmt.Scanf("%s", &player2)
		fmt.Print("Enter Player 3's name:")
		fmt.Scanf("%s", &player3)
		fmt.Printf("Golfers: %s, %s, and %s\n", player1, player2, player3)
	} else if playerCount == 2 {
		fmt.Print("Enter Player 1's name:")
		fmt.Scanf("%s", &player1)
		fmt.Print("Enter Player 2's name:")
		fmt.Scanf("%s", &player2)
		fmt.Printf("Golfers: %s, and %s\n", player1, player2)
	} else {
		fmt.Print("Enter player name:")
		fmt.Scanf("%s", &player1)
		fmt.Printf("%s is golfing alone\n", player1)
	}
}



func main(){
	players()
}