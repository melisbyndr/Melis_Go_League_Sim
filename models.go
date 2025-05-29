package main

import (
	"fmt"
	"math/rand"
)

type Team struct {
	Name             string
	Defence          int
	MidField         int
	Forward          int
	HomeTeamModifier int
	Played           int
	Wins             int
	Draws            int
	Losses           int
	GoalsFor         int
	GoalsAgainst     int
	Points           int
}

type Match struct {
	HomeTeam  *Team
	AwayTeam  *Team
	HomeGoals int
	AwayGoals int
	Played    bool
}

type League struct {
	Teams []*Team
	Weeks [][]Match
}

type UpdatedPointTeam struct {
	Name        *Team
	TotalPoints int
}

//type ForecastResult struct {
//	TeamName      string  `json:"team_name"`
//	WinChance     float64 `json:"win_chance"`
//	Top2Chance    float64 `json:"top2_chance"`
//	AveragePoints float64 `json:"average_points"`
//}

// This part include the forecast results structs, to be able to use at endpoints
type ForecastResult struct {
	ChampionProbabilities []TeamProbability `json:"champion_probabilities"`
	Top2Probabilities     []TeamProbability `json:"top2_probabilities"`
	AvgPoints             []TeamAvgPoints   `json:"avg_points"`
}

type TeamProbability struct {
	TeamName string  `json:"team_name"`
	Prob     float64 `json:"probability"`
}

type TeamAvgPoints struct {
	TeamName string  `json:"team_name"`
	Points   float64 `json:"points"`
}

func NewLeague() *League { // NewLeague creates a new League and returns its pointer (*League)

	//rand.Seed(time.Now().UnixNano())

	team1 := &Team{Name: "Sirius Stars", Defence: rand.Intn(6) + 5, MidField: rand.Intn(6) + 5, Forward: rand.Intn(6) + 5, HomeTeamModifier: rand.Intn(6) + 5}
	team2 := &Team{Name: "Journey United", Defence: rand.Intn(6) + 5, MidField: rand.Intn(6) + 5, Forward: rand.Intn(6) + 5, HomeTeamModifier: rand.Intn(6) + 5}
	team3 := &Team{Name: "Forecasters City", Defence: rand.Intn(6) + 5, MidField: rand.Intn(6) + 5, Forward: rand.Intn(6) + 5, HomeTeamModifier: rand.Intn(6) + 5}
	team4 := &Team{Name: "Omni FC", Defence: rand.Intn(6) + 5, MidField: rand.Intn(6) + 5, Forward: rand.Intn(6) + 5, HomeTeamModifier: rand.Intn(6) + 5}

	teams := []*Team{team1, team2, team3, team4} // This creates a slice of Team pointers and stores the 4 teams in it.

	fmt.Println("🏁 League Begins-  Teams in the League:")
	for _, team := range teams {
		fmt.Printf("- %s (Defence Strength: %d /MidField Strength: %d /Forward Strength: %d /Home Team Modifier: %d )\n", team.Name, team.Defence, team.MidField, team.Forward, team.HomeTeamModifier)
	}

	return &League{
		Teams: teams,
		Weeks: [][]Match{},
	}

}

type PositionResult struct {
	Minute   int
	Attacker string
	IsGoal   bool
	ScoreA   int
	ScoreB   int
}
