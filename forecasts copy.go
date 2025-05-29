package main

import (
	"fmt"
	"sort"
)

// League yapısı ve diğer tanımlar burada olmalı

// Takımları derin kopyalama
func DeepCopyTeams(original []*Team) []*Team {
	copy := make([]*Team, len(original))
	for i, t := range original {
		newTeam := *t // Değer kopyası
		copy[i] = &newTeam
	}
	return copy
}

// Maçları derin kopyalama
func DeepCopyMatches(original []Match, teamCopy []*Team) []Match {
	// Takım adına göre takımı bulmak için map
	teamMap := make(map[string]*Team)
	for _, t := range teamCopy {
		teamMap[t.Name] = t
	}

	copied := make([]Match, len(original))
	for i, m := range original {
		copied[i] = Match{
			HomeTeam:  teamMap[m.HomeTeam.Name],
			AwayTeam:  teamMap[m.AwayTeam.Name],
			Played:    false,
			HomeGoals: 0,
			AwayGoals: 0,
		}
	}
	return copied
}

func (l *League) SimulateForecast() ForecastResult {
	// 1. Kalan maçları topla
	var remainingMatches []Match
	for _, week := range l.Weeks {
		for _, match := range week {
			if !match.Played {
				remainingMatches = append(remainingMatches, match)
			}
		}
	}
	fmt.Printf("Kalan maç sayısı: %d\n", len(remainingMatches))

	// 2. Takımların mevcut durumunu kopyala
	originalTeams := DeepCopyTeams(l.Teams)

	// 3. İstatistikler için yapılar
	championCounts := make(map[string]int)
	top2Counts := make(map[string]int)
	pointSums := make(map[string]int)

	// 4. 1000 simülasyon çalıştır
	for i := 0; i < 1000; i++ {
		// Her simülasyon için orijinal durumdan kopya oluştur
		simTeams := DeepCopyTeams(originalTeams)
		simMatches := DeepCopyMatches(remainingMatches, simTeams)

		// Kalan tüm maçları simüle et
		for j := range simMatches {
			simulateMatch(&simMatches[j])
		}

		// Takımları sırala (puan > averaj > atılan gol)
		sort.Slice(simTeams, func(i, j int) bool {
			if simTeams[i].Points != simTeams[j].Points {
				return simTeams[i].Points > simTeams[j].Points
			}
			diffI := simTeams[i].GoalsFor - simTeams[i].GoalsAgainst
			diffJ := simTeams[j].GoalsFor - simTeams[j].GoalsAgainst
			if diffI != diffJ {
				return diffI > diffJ
			}
			return simTeams[i].GoalsFor > simTeams[j].GoalsFor
		})

		// Şampiyonu kaydet
		championCounts[simTeams[0].Name]++

		// İlk 2'yi kaydet
		for k := 0; k < 2 && k < len(simTeams); k++ {
			top2Counts[simTeams[k].Name]++
		}

		// Puan toplamlarını güncelle
		for _, team := range simTeams {
			pointSums[team.Name] += team.Points
		}
	}

	// 5. Sonuçları hesapla ve göster
	fmt.Println("\n🏆 Şampiyonluk Olasılıkları:")
	for _, team := range originalTeams {
		prob := float64(championCounts[team.Name]) / 10.0 // % için 1000/100 = 10
		fmt.Printf("- %s: %.1f%%\n", team.Name, prob)
	}

	fmt.Println("\n🥈 İlk 2'ye Girme Olasılıkları:")
	for _, team := range originalTeams {
		prob := float64(top2Counts[team.Name]) / 10.0
		fmt.Printf("- %s: %.1f%%\n", team.Name, prob)
	}

	fmt.Println("\n📊 Ortalama Puanlar:")
	for _, team := range originalTeams {
		avg := float64(pointSums[team.Name]) / 1000.0
		fmt.Printf("- %s: %.2f puan\n", team.Name, avg)
	}

	// 6. JSON formatında sonuçları hazırla
	result := ForecastResult{
		ChampionProbabilities: make([]TeamProbability, 0, len(originalTeams)),
		Top2Probabilities:     make([]TeamProbability, 0, len(originalTeams)),
		AvgPoints:             make([]TeamAvgPoints, 0, len(originalTeams)),
	}

	for _, team := range originalTeams {
		result.ChampionProbabilities = append(result.ChampionProbabilities, TeamProbability{
			TeamName: team.Name,
			Prob:     float64(championCounts[team.Name]) / 10.0,
		})

		result.Top2Probabilities = append(result.Top2Probabilities, TeamProbability{
			TeamName: team.Name,
			Prob:     float64(top2Counts[team.Name]) / 10.0,
		})

		result.AvgPoints = append(result.AvgPoints, TeamAvgPoints{
			TeamName: team.Name,
			Points:   float64(pointSums[team.Name]) / 1000.0,
		})
	}

	return result
}
