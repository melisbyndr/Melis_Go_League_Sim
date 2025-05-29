package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB
var league *League

func main() {

	fmt.Println(" 🚀 We are ready to go!!")

	db = InitDB()

	league = NewLeague()

	InsertTeams(db, league.Teams)

	ReadTeams(db)

	league.GenerateFixtures()

	league.PlayAllWeeks()

	PrintTable(league)

	r := mux.NewRouter()

	// ndpoint tanımlamları
	r.HandleFunc("/teams", getTeams).Methods("GET")
	r.HandleFunc("/matches", getMatches).Methods("GET")
	r.HandleFunc("/play/week/{id}", playWeekHandler).Methods("POST")
	r.HandleFunc("/play/all", playAllHandler).Methods("POST")
	r.HandleFunc("/table", leagueTableHandler).Methods("GET")
	r.HandleFunc("/forecast", forecastHandler).Methods("GET")

	fmt.Println("-------Server running at http://localhost:8080")

	defer db.Close()

	http.ListenAndServe(":8080", r)

}
