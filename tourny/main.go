package main

import "fmt"

const HOME_TEAM_WON = 1

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}

	fmt.Println("Winner is ", TournamentWinner(competitions, results))

}

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.

	currentBestTeam := ""
	m := make(map[string]int)
	m[currentBestTeam] = 0
	for i, teams := range competitions {
		winner := teams[1]
		if results[i] == HOME_TEAM_WON {
			if val, ok := m[teams[0]]; ok {
				m[teams[0]] = val + 3
			} else {
				m[teams[0]] = 3
			}
			winner = teams[0]
		} else {
			if val1, ok := m[teams[1]]; ok {
				m[teams[1]] = val1 + 3

			} else {
				m[teams[1]] = 3
			}
		}
		if m[winner] > m[currentBestTeam] {
			currentBestTeam = winner
		}
	}
	return currentBestTeam
}
