package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

// Run the game in the background
func (game *Game) Rounds() {
	for {
		select {
		case round := <-game.RoundChan:
			game.Round.RoundNumber = game.Round.RoundNumber + round
			game.RoundChan <- 0
		case msg := <-game.DisplayChan:
			fmt.Println(msg)
			game.DisplayChan <- ""
		}
	}
}

// Clear the terminal screen
func (game *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else { // linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Print information at the start of the game
func (game *Game) PrintIntro() {
	game.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
----------------------
Game is played for three rounds, and best of three wins the game. Good luck!
`)
	<-game.DisplayChan
}

// Play a single round of the game
func (game *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	game.DisplayChan <- fmt.Sprintf(`
Round %d
--------
`, game.Round.RoundNumber)
	<-game.DisplayChan

	game.DisplayChan <- "Please enter rock, paper, or scissors -> "
	<-game.DisplayChan
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	switch strings.ToLower(playerChoice) {
	case "rock":
		playerValue = ROCK
		break
	case "paper":
		playerValue = PAPER
		break
	case "scissors":
		playerValue = SCISSORS
		break
	default:
	}

	game.DisplayChan <- ""
	<-game.DisplayChan

	game.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-game.DisplayChan

	switch computerValue {
	case ROCK:
		game.DisplayChan <- "Computer chose ROCK"
		<-game.DisplayChan
		break
	case PAPER:
		game.DisplayChan <- "Computer chose PAPER"
		<-game.DisplayChan
		break
	case SCISSORS:
		game.DisplayChan <- "Computer chose SCISSORS"
		<-game.DisplayChan
		break
	default:
	}

	game.DisplayChan <- ""
	<-game.DisplayChan

	if playerValue == computerValue {
		game.DisplayChan <- "It's a draw!"
		<-game.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		default:
			game.DisplayChan <- "Invalid choice!"
			<-game.DisplayChan
			return false
		}
	}

	return true
}

// Computer wins game
func (game *Game) computerWins() {
	game.Round.ComputerScore++
	game.DisplayChan <- "Computer wins!"
	<-game.DisplayChan
}

// Player wins game
func (game *Game) playerWins() {
	game.Round.PlayerScore++
	game.DisplayChan <- "Player wins!"
	<-game.DisplayChan
}

// Print the summary of the game at the end
func (game *Game) PrintSummary() {
	game.DisplayChan <- fmt.Sprintf(`
Final score
-----------
Game is played for three rounds, and best of three wins the game. Good luck!
Player: %d/3, Computer %d/3
`, game.Round.PlayerScore, game.Round.ComputerScore)
	<-game.DisplayChan

	if game.Round.PlayerScore > game.Round.ComputerScore {
		game.DisplayChan <- "Player wins game!"
		<-game.DisplayChan
		// fmt.Println("Player wins game!")
	} else {
		game.DisplayChan <- "Computer wins game!"
		<-game.DisplayChan
		// fmt.Println("Computer wins game!")
	}

	game.DisplayChan <- ""
	<-game.DisplayChan
}
