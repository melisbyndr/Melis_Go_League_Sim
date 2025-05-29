package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	teams := ReadTeams(db) // DB’den veriyi oku
	fmt.Println("👀 getTeams handler çalıştı. Toplam takım sayısı:", len(teams))
	json.NewEncoder(w).Encode(teams)
}

func getMatches(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query(`SELECT home_team, away_team, home_goals, away_goals, played FROM matches`)
	defer rows.Close()

	type MatchRow struct {
		HomeTeam  string `json:"home_team"`
		AwayTeam  string `json:"away_team"`
		HomeGoals int    `json:"home_goals"`
		AwayGoals int    `json:"away_goals"`
		Played    bool   `json:"played"`
	}

	var matches []MatchRow
	for rows.Next() {
		var m MatchRow
		rows.Scan(&m.HomeTeam, &m.AwayTeam, &m.HomeGoals, &m.AwayGoals, &m.Played)
		matches = append(matches, m)
	}
	json.NewEncoder(w).Encode(matches)
}

func playWeekHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekIDStr := vars["id"]
	weekID, err := strconv.Atoi(weekIDStr)
	if err != nil {
		http.Error(w, "Geçersiz hafta numarası", http.StatusBadRequest)
		return
	}

	league.PlayWeek(weekID - 1) // index 0’dan başladığı için -1
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Week %d played.", weekID)
}

func playAllHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(league.Weeks); i++ {
		league.PlayWeek(i)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "All weeks played.")
}

func leagueTableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(league.Teams)
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
	// Tahmin sonuçlarını hesapla
	result := league.SimulateForecast()

	// JSON yanıtını hazırla
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
