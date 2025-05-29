package main

import (
	"fmt"
	"sort"
)

func PrintTable(l *League) {
	// Takımları puana göre sırala , sonra Goals Diff.  , sonra atılan gol
	sort.Slice(l.Teams, func(i, j int) bool {
		a := l.Teams[i]
		b := l.Teams[j]

		if a.Points != b.Points {
			return a.Points > b.Points // daha çok puan
		}
		// GoalsFor - GoalsAgainst
		goalDiffA := a.GoalsFor - a.GoalsAgainst
		goalDiffB := b.GoalsFor - b.GoalsAgainst

		if goalDiffA != goalDiffB {
			return goalDiffA > goalDiffB
		}

		// Atılan gol
		return a.GoalsFor > b.GoalsFor
	})

	/*
		MP - Match Played
		W - Wins
		D- Draws
		L - Losses
		GF - Goals For
		GA - Goals Against
		GD- Goal Difference
		Pts- Points --> W*3 + D*1 + L*0
	*/

	// Tablo başlığı
	fmt.Printf("\n%-15s MP  W  D  L  GF  GA  GD  Pts\n", "Team")
	fmt.Println("-------------------------------------------------")

	// Her takımın satırı
	for _, team := range l.Teams {
		goalDiff := team.GoalsFor - team.GoalsAgainst
		fmt.Printf("%-15s %2d  %2d %2d %2d  %2d  %2d  %2d  %3d\n",
			team.Name, team.Played, team.Wins, team.Draws, team.Losses,
			team.GoalsFor, team.GoalsAgainst, goalDiff, team.Points)
	}
}
