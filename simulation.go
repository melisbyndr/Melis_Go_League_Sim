package main

import (
	"fmt"
	"math/rand"
)

// max fonksiyonu (Go'da yok)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Atağı yapacak takımı belirleyen fonksiyon
func decideAttackingTeam(homeTeam, awayTeam *Team) *Team {
	randomHome := rand.Intn(10) + 1
	randomAway := rand.Intn(10) + 1
	scoreHome := homeTeam.MidField * randomHome
	scoreAway := awayTeam.MidField * randomAway

	if scoreHome > scoreAway {
		return homeTeam
	} else if scoreAway > scoreHome {
		return awayTeam
	} else {
		if rand.Intn(2) == 0 {
			return homeTeam
		}
		return awayTeam
	}
}

// Gol olup olmadığını belirleyen fonksiyon
func isGoal(attacker, defender *Team, isHome bool) bool {
	randomForward := rand.Intn(10) + 1
	randomDefense := rand.Intn(10) + 1
	scoreForward := attacker.Forward * randomForward
	scoreDefense := defender.Defence * randomDefense

	if isHome {
		scoreForward = scoreForward * max(1, attacker.HomeTeamModifier)
	}

	if scoreForward > scoreDefense {
		return true
	} else if scoreDefense > scoreForward {
		return false
	}
	return rand.Intn(2) == 0
}

// Pozisyon simülasyonunu çalıştıran fonksiyon
func simulateMatch(m *Match) {
	if m.Played {
		return
	}

	// Maça özel geçici gol sayıları
	homeGoals := 0
	awayGoals := 0

	// Pozisyon sayısı ve dakikaları belirle
	positionCount := rand.Intn(11) + 5 // 5-15 positions

	// Her pozisyonu simüle et
	for i := 0; i < positionCount; i++ {
		attacker := decideAttackingTeam(m.HomeTeam, m.AwayTeam)
		var defender *Team
		isHome := false

		if attacker == m.HomeTeam {
			defender = m.AwayTeam
			isHome = true
		} else {
			defender = m.HomeTeam
		}

		if isGoal(attacker, defender, isHome) {
			if attacker == m.HomeTeam {
				homeGoals++
			} else {
				awayGoals++
			}
		}
	}

	// Maç sonuçlarını güncelle
	m.HomeGoals = homeGoals
	m.AwayGoals = awayGoals
	m.Played = true

	// Takım istatistiklerini güncelle
	m.HomeTeam.GoalsFor += homeGoals
	m.HomeTeam.GoalsAgainst += awayGoals
	m.AwayTeam.GoalsFor += awayGoals
	m.AwayTeam.GoalsAgainst += homeGoals

	m.HomeTeam.Played++
	m.AwayTeam.Played++

	if homeGoals > awayGoals {
		m.HomeTeam.Wins++
		m.AwayTeam.Losses++
		m.HomeTeam.Points += 3
	} else if homeGoals < awayGoals {
		m.AwayTeam.Wins++
		m.HomeTeam.Losses++
		m.AwayTeam.Points += 3
	} else {
		m.HomeTeam.Draws++
		m.AwayTeam.Draws++
		m.HomeTeam.Points += 1
		m.AwayTeam.Points += 1
	}
}

// Haftayı oynat
func (l *League) PlayWeek(week int) {
	if week < 0 || week >= len(l.Weeks) {
		fmt.Println("Geçersiz hafta numarası.")
		return
	}

	fmt.Printf("\n--- Week %d ---\n", week+1)
	for i := range l.Weeks[week] {
		match := &l.Weeks[week][i]
		simulateMatch(match)
		fmt.Printf("%s %d - %d %s\n", match.HomeTeam.Name, match.HomeGoals, match.AwayGoals, match.AwayTeam.Name)
	}

	// Maçları veritabanına kaydet
	db := InitDB()
	defer db.Close()
	InsertMatches(db, l.Weeks[week])
}

// Tüm haftaları oynat
func (l *League) PlayAllWeeks() {
	for i := range l.Weeks {
		l.PlayWeek(i)
		if i > 2 {
			fmt.Printf("\n📅 Week %d:\n", i+1)
			l.SimulateForecast()

		}
	}
}

/*
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Atağı yapacak takımı belirleyen fonksiyon
// MidField puanları oranında random seçim
func decideAttackingTeamPos(homeTeam, awayTeam Team) string {
	randomHome := rand.Intn(10) + 1
	randomAway := rand.Intn(10) + 1
	scoreHome := homeTeam.MidField * randomHome
	scoreAway := awayTeam.MidField * randomAway

	//fmt.Printf("\nPosition Calculation:\n")
	//fmt.Printf("%s: MidField=%d * Random=%d = %d\n", homeTeam.Name, homeTeam.MidField, randomHome, scoreHome)
	//fmt.Printf("%s: MidField=%d * Random=%d = %d\n\n", awayTeam.Name, awayTeam.MidField, randomAway, scoreAway)
	//
	var attackingTeam string
	if scoreHome > scoreAway {
		attackingTeam = homeTeam.Name
	} else if scoreAway > scoreHome {
		attackingTeam = awayTeam.Name
	} else {
		if rand.Intn(2) == 0 {
			attackingTeam = homeTeam.Name
		} else {
			attackingTeam = awayTeam.Name
		}
	}
	//fmt.Printf("Attacking Team: %s\n\n", attackingTeam)
	//time.Sleep(1 * time.Second)
	return attackingTeam
}

// max fonksiyonu (Go'da yok)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Gol olup olmadığını belirleyen fonksiyon
// Forward ve Defence puanları oranında random bir algoritma
// Daha karmaşık yapmak için bu kısmı değiştirebilirsin
func isGoalPos(forward, defense int, isHome bool, homeTeamModifier int) bool {
	randomForward := rand.Intn(10) + 1
	randomDefense := rand.Intn(10) + 1
	scoreForward := forward * randomForward
	scoreDefense := defense * randomDefense

	if isHome {
		scoreForward = scoreForward * max(1, homeTeamModifier)
	}

	//fmt.Printf("Goal Calculation: Forward: %d * Random= %d = %d, Defense: %d * Random= %d = %d\n\n", forward, randomForward, scoreForward, defense, randomDefense, scoreDefense)

	isGoal := false
	if scoreForward > scoreDefense {
		isGoal = true
	} else if scoreDefense > scoreForward {
		isGoal = false
	} else {
		isGoal = rand.Intn(2) == 0
	}

	if isGoal {
		//fmtfmt.Printf("GOAL!     Forward: %d  Defense: %d\n\n\n\n", scoreForward, scoreDefense)
	} else {
		//fmtfmt.Printf("No goal!  Forward: %d  Defense: %d\n\n\n\n", scoreForward, scoreDefense)
	}
	//time.Sleep(5 * time.Second)
	return isGoal
}

// Pozisyon simülasyonunu çalıştıran fonksiyon
func simulateMatch(m *Match) {
	//rand.Seed(time.Now().UnixNano())

	if m.Played {
		return // Maç zaten oynandıysa bir daha oynama
	}

	positionCount := rand.Intn(10) + 5 // 5-15 positions
	minutes := make([]int, positionCount)
	usedMinutes := make(map[int]bool)
	for i := 0; i < positionCount; i++ {
		for {
			minute := rand.Intn(90) + 1
			if !usedMinutes[minute] {
				minutes[i] = minute
				usedMinutes[minute] = true
				break
			}
		}
	}
	sort.Ints(minutes)

	home := m.HomeTeam
	away := m.AwayTeam

	//fmt.Printf("\n****Number of positions: %d****\n", positionCount)
	//fmt.Printf("%s: Defence= %d, MidField= %d, Forward= %d\n", home.Name, home.Defence, home.MidField, home.Forward)
	//fmt.Printf("%s: Defence= %d, MidField= %d, Forward= %d\n\n", away.Name, away.Defence, away.MidField, away.Forward)
	//fmt.Println("Positions:")

	for i := 0; i < positionCount; i++ {
		//fmt.Printf("\n<<< Position %d (Minute %d) >>>\n", i+1, minutes[i])
		attacker := decideAttackingTeamPos(*home, *away)
		var attackTeam, defendTeam *Team
		var isHome bool
		if attacker == home.Name {
			attackTeam = home
			defendTeam = away
			isHome = true
		} else {
			attackTeam = away
			defendTeam = home
			isHome = false
		}

		goal := isGoalPos(attackTeam.Forward, defendTeam.Defence, isHome, attackTeam.HomeTeamModifier)
		if goal {
			attackTeam.GoalsFor++
			//fmt.Printf("%s attack - GOAL!  (%d. minute)\n\n", attackTeam.Name, minutes[i])
		} else {
			//fmt.Printf("%s attack - No goal  (%d. minute)\n\n", attackTeam.Name, minutes[i])
		}

		//fmt.Printf("-------  %s %d - %d %s -------\n", home.Name, home.GoalsFor, away.GoalsFor, away.Name)
		//fmt.Printf("=============================================================================\n")
		//time.Sleep(1 * time.Second)
	}

	//m.HomeTeam.GoalsFor = home.GoalsFor
	//m.AwayGoals = away.GoalsFor
	m.Played = true

	//fmt.Printf(finalResult)

	// Goller
	m.HomeTeam.GoalsFor += home.GoalsFor
	m.HomeTeam.GoalsAgainst += away.GoalsFor
	m.AwayTeam.GoalsFor += away.GoalsFor
	m.AwayTeam.GoalsAgainst += home.GoalsFor

	// Oynanan maç sayısı
	m.HomeTeam.Played++
	m.AwayTeam.Played++

	// Sonuca göre puan ve istatistikler
	if home.GoalsFor > away.GoalsFor {
		m.HomeTeam.Wins++
		m.AwayTeam.Losses++
		m.HomeTeam.Points += 3
	} else if home.GoalsFor < away.GoalsFor {
		m.AwayTeam.Wins++
		m.HomeTeam.Losses++
		m.AwayTeam.Points += 3
	} else {
		m.HomeTeam.Draws++
		m.AwayTeam.Draws++
		m.HomeTeam.Points += 1
		m.AwayTeam.Points += 1
	}
}

// simulate a speciifc week
func (l *League) PlayWeek(week int) {

	db := InitDB()
	defer db.Close()

	if week < 0 || week >= len(l.Weeks) {
		fmt.Println("Geçersiz hafta numarası.")
		return
	}

	fmt.Printf("\n--- Week %d ---\n", week+1)
	for i := range l.Weeks[week] {
		match := &l.Weeks[week][i]
		simulateMatch(match)
		fmt.Printf("%s %d - %d %s\n", match.HomeTeam.Name, match.HomeGoals, match.AwayGoals, match.AwayTeam.Name)

		InsertMatches(db, l.Weeks[week])

	}
}

// simulate all weeks
func (l *League) PlayAllWeeks() {

	db := InitDB()
	defer db.Close()

	for i := range l.Weeks {
		l.PlayWeek(i)
		if i > 2 {
			fmt.Printf("\n📅 Week %d:\n", i+1)
			l.SimulateForecast()
		}
	}
}
*/
