package main

import "fmt"

func main() {

	//competionsInput := [][]string{{"Html", "C#"}, {"C#", "Python"}, {"Python", "Html"}}
	//resultsInput := []int{0, 0, 1}

	competionsInput := [][]string{{"Html", "Java"}, {"Java", "Python"}, {"Python", "Html"}}
	resultsInput := []int{0, 1, 1}

	fmt.Print(TournamentWinner(competionsInput, resultsInput))

}

// TournamentWinner checks the tournament champion
func TournamentWinner(competitions [][]string, results []int) string {

	tournamentTable := make(map[string]int)
	tournamentWinner := ""
	maxPoints := 0
	winner := ""

	for i := 0; i < len(competitions); i++ {

		homeTeam := competitions[i][0]
		awayTeam := competitions[i][1]
		winner = ""

		if results[i] == 0 {
			winner = awayTeam
		} else {
			winner = homeTeam
		}

		tournamentTable[winner] += 3

		if tournamentTable[winner] > maxPoints {
			maxPoints = tournamentTable[winner]
			tournamentWinner = winner
		}
	}

	return tournamentWinner
}
