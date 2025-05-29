package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "league.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	schema, err := os.ReadFile("create_table.sql")
	if err != nil {
		log.Fatal("Failed to read schema.sql:", err)
	}
	if err == nil {
		fmt.Printf("✅ Sql comment is read")
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal("Failed to execute schema:", err)
	}
	if err == nil {
		fmt.Printf("✅ Sql comment is executed")
	}

	fmt.Println("✅ Database initialized.")
	return db
}

func InsertTeams(db *sql.DB, teams []*Team) {
	for _, team := range teams {
		_, err := db.Exec(`
			INSERT INTO teams 
			(name, defence, midfield, forward, hometeammodifier, played, wins, draws, losses, goals_for, goals_against, points)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, team.Name, team.Defence, team.MidField, team.Forward, team.HomeTeamModifier, team.Played, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst, team.Points)

		if err != nil {
			log.Printf("❌ Failed to insert team %s: %v\n", team.Name, err)
		} else {
			fmt.Printf("✅ Inserted team: %s\n", team.Name)
		}
	}
}

//	func ReadTeams(db *sql.DB) []*Team {
//		rows, err := db.Query(`SELECT name, strength, played, wins, draws, losses, goals_for, goals_against, points FROM teams`)
//		if err != nil {
//			log.Fatal("Failed to query teams:", err)
//		}
//		defer rows.Close()
//
//		var teams []*Team
//
//		for rows.Next() {
//			var t Team
//			err := rows.Scan(&t.Name, &t.Strength, &t.Played, &t.Wins, &t.Draws, &t.Losses, &t.GoalsFor, &t.GoalsAgainst, &t.Points)
//			if err != nil {
//				log.Println("Error reading row:", err)
//				continue
//			}
//			teams = append(teams, &t)
//		}
//
//		fmt.Println("\n📋 Teams in Database:")
//		for _, team := range teams {
//			fmt.Printf("- %s (Strength: %d, Points: %d)\n", team.Name, team.Strength, team.Points)
//		}
//
//		return teams
//	}
func RunQuery(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		log.Println("❌ Query error:", err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Println("❌ Failed to get columns:", err)
		return
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	fmt.Println("\n📊 Query Result:")
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		for i, col := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				val = string(b)
			}
			fmt.Printf("%s: %v  ", col, val)
		}
		fmt.Println()
	}
}
func InsertMatches(db *sql.DB, matches []Match) {
	for _, m := range matches {
		_, err := db.Exec(`
			INSERT INTO matches 
			(home_team, away_team, home_goals, away_goals, played)
			VALUES (?, ?, ?, ?, ?)
		`, m.HomeTeam.Name, m.AwayTeam.Name, m.HomeGoals, m.AwayGoals, m.Played)

		if err != nil {
			log.Printf("❌ Failed to insert match: %s vs %s → %v\n", m.HomeTeam.Name, m.AwayTeam.Name, err)
		}
	}
	fmt.Println("✅ All matches inserted.")
}

func ReadMatches(db *sql.DB) {
	rows, _ := db.Query(`SELECT home_team, away_team, home_goals, away_goals, played FROM matches`)
	defer rows.Close()

	fmt.Println("\n📅 Matches in DB:")
	for rows.Next() {
		var home, away string
		var hg, ag int
		var played bool
		rows.Scan(&home, &away, &hg, &ag, &played)
		fmt.Printf("%s %d - %d %s | Played: %v\n", home, hg, ag, away, played)
	}
}

func ReadTeams(db *sql.DB) []*Team {
	rows, err := db.Query(`SELECT name, defence, midfield, forward, hometeammodifier, played, wins, draws, losses, goals_for, goals_against, points FROM teams`)
	if err != nil {
		log.Fatal("Failed to query teams:", err)
	}
	defer rows.Close()

	var teams []*Team

	for rows.Next() {
		var t Team
		err := rows.Scan(&t.Name, &t.Defence, &t.MidField, &t.Forward, &t.HomeTeamModifier, &t.Played, &t.Wins, &t.Draws, &t.Losses, &t.GoalsFor, &t.GoalsAgainst, &t.Points)
		if err != nil {
			log.Println("Error reading row:", err)
			continue
		}
		teams = append(teams, &t)
	}

	return teams
}
