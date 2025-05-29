package main

import "fmt"

/* 	Each team should play only one match per week.

--- Maximum of 2 matches can be played each week.

--- Each team plays every other team twice (home and away).

--- With 4 teams: total matches = 4 × 3 = 12
---  6 weeks in total. */

/* How match struct is constructed
type Match struct {
	HomeTeam  *Team
	AwayTeam  *Team
	HomeGoals int
	AwayGoals int
	Played    bool
}
*/

func (l *League) GenerateFixtures() { //allows modifying the original league's data
	var allMatches []Match

	for i := 0; i < len(l.Teams); i++ {
		for j := 0; j < len(l.Teams); j++ {
			if i != j {
				allMatches = append(allMatches, Match{
					HomeTeam: l.Teams[i],
					AwayTeam: l.Teams[j],
					Played:   false,
				})
			}
		}
	}

	for i, match := range allMatches {
		fmt.Printf("Match %d: %s vs %s Played: %t  \n ", i+1, match.HomeTeam.Name, match.AwayTeam.Name, match.Played)
	}

	var weeks [][]Match

	for len(allMatches) > 0 {
		week := []Match{}
		used := make(map[string]bool) // Create a map to track which teams have already played this week

		for i := 0; i < len(allMatches); {
			match := allMatches[i]
			homeName := match.HomeTeam.Name
			awayName := match.AwayTeam.Name

			if !used[homeName] && !used[awayName] { // If this team  hasnt played yet;
				week = append(week, match)
				used[homeName] = true
				used[awayName] = true
				allMatches = append(allMatches[:i], allMatches[i+1:]...) // Remove the match  i from the allMatches
			} else {
				i++
			}

			if len(week) == 2 { // If there are 2 matches played this week, stop adding any games
				break
			}
		}

		weeks = append(weeks, week)
	}

	l.Weeks = weeks
}
