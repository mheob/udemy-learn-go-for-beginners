package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())

	firstNumber := rand.Intn(8) + 2
	secondNumber := rand.Intn(8) + 2
	subtraction := rand.Intn(8) + 2
	answer := firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)
}
