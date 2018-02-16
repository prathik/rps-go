package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getRandom() string {
	plays := make([]string, 0)
	plays = append(plays, "r", "p", "s")
	rand.Seed(time.Now().UTC().UnixNano())
	return plays[rand.Intn(len(plays))]
}

func getNextRandom(user string) string {
	c_play := getRandom()
	for user == c_play {
		c_play = getRandom()
	}
	return c_play
}

func comp_winner(comp string, user string) bool {
	fmt.Println("Computer played: ", getProperString(comp), " User played: ", getProperString(user))
	if (comp == "r" && user == "s") || (comp == "p" && user == "r") || (comp == "s" && user == "p") {
		return true
	}
	return false
}

func firstRandomPlay() (c_play, u_play string, c_wins bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rock, paper or scissor (r, p, s): ")
	text, _ := reader.ReadString('\n')

	u_play = strings.TrimRight(text, "\n")
	c_play = getNextRandom(u_play)
	c_wins = comp_winner(c_play, u_play)
	return
}

func getCalculatedPlay(p_c_play, p_u_play string, p_c_wins bool) string {
	if p_c_wins {
		return p_u_play
	} else {
		if p_c_play == "r" && p_u_play == "s" {
			return "p"
		}
		if p_c_play == "r" && p_u_play == "p" {
			return "s"
		}
		if p_c_play == "p" && p_u_play == "s" {
			return "r"
		}
		if p_c_play == "p" && p_u_play == "r" {
			return "s"
		}
		if p_c_play == "s" && p_u_play == "r" {
			return "p"
		}
		if p_c_play == "s" && p_u_play == "p" {
			return "r"
		}
	}
	return "p"
}

func getProperString(play string) string {
	switch play {
	case "r":
		return "Rock"
	case "s":
		return "Scissor"
	case "p":
		return "Paper"
	}
	return "Invalid!"
}

func main() {
	c_win := 0
	u_win := 0
	var i_count int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of games to play: ")
	fmt.Scanf("%d", &i_count)

	p_c_play, p_u_play, p_c_wins := firstRandomPlay()

	if p_c_wins {
		c_win = c_win + 1
		fmt.Println("You lost! You played ", getProperString(p_u_play))
	} else {
		u_win = u_win + 1
		fmt.Println("You won! You played ", getProperString(p_u_play))
	}

	i_count = i_count - 1

	for i_count > 0 {

		fmt.Print("Enter rock, paper or scissor (r, p, s): ")
		text, _ := reader.ReadString('\n')

		u_play := strings.TrimRight(text, "\n")
		c_play := getCalculatedPlay(p_c_play, p_u_play, p_c_wins)

		if c_play == u_play {
			c_play = getNextRandom(u_play)
		}

		c_wins := comp_winner(c_play, u_play)

		if c_wins {
			fmt.Println("You lost! You played ", getProperString(u_play))
			c_win = c_win + 1
		} else {
			fmt.Println("You win! You played ", getProperString(u_play))
			u_win = u_win + 1
		}
		i_count = i_count - 1

		p_c_play = c_play
		p_u_play = u_play
		p_c_wins = c_wins
	}

	fmt.Println("Computer Wins: ", c_win)
	fmt.Println("User Wins: ", u_win)

}
