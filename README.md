# Melis_Go_League_Sim

Football League Simulation - Technical Documentation
General Overview
This project simulates a football league with 4 teams, where each team plays against others in both home and away games (double round-robin format). The project includes:
•	Match simulation based on team attributes.
•	Weekly and full-season simulation.
•	Forecasting of championship probabilities based on Monte Carlo simulation.
•	REST API endpoints to manage and view the league.
•	SQLite database for storing match and team data.
•	Postman support for API testing.
Folder Structure
go-league-sim/
├── go_basics/               # Personal Go le arning playground (excluded from main-project)
│   ├── go_basics.go 
│   └── simulation copy.go
├── create_table.sql         # SQL schema for teams, matches, scoreboard
├── database.go              # DB init, insert, read, and helper functions
├── endpoints.go             # REST API endpoint handlers (e.g. /teams, /forecast)
├── fixture.go               # League fixture creation logic
├── forecast.go              # Monte Carlo simulation to predict final results
├── go.mod / go.sum          # Go module dependencies
├── league.db                 # SQLite database file
├─ main.go                  # Entry point, server config, routing
├── models.go                # Structs: Team, Match, League, PositionResult, etc.
├── simulation.go            # Match simulation logic (position-based)
├── table.go                 # Table sorting, printing, and team ranking logic





League Management and Fixture Generation
Below are the ground rules to generate a match plan for each week. These are implemented in the folders models.go and fixture.go
-	 Each team should play only one match per week.
-	Maximum of 2 matches can be played each week.
-	 Each team plays every other team twice (home and away).
-	 With 4 teams: total matches = 4 × 3 = 12
-	6 weeks in total.  

Details of Simulation
-	Each team is created with random attributes: Defence Strength, MidField Strength, Forward Strength, and Home Team Modifier.
-	Every team plays one match per week, for a total of 6 matches per team (12 matches in total).
-	Each match consists of 5–15 simulated positions (goal chances).

-	For each position:
o	MidField strength determines which team attacks.
o	Forward vs Defence strength (scaled by random values) determines goal success.
o	If the attacker is the home team, HomeTeamModifier boosts their chance.
-	Successful goals update GoalsFor and GoalsAgainst of both teams.

-	At the end of each match:
o	Win = 3 points, Draw = 1 point per team.
o	Played, Wins, Draws, Losses fields are updated.
o	All data updates apply directly to Match and Team structs.

•	PlayWeek(n) simulates a single week.
•	PlayAllWeeks() simulates the full season.
•	Standings calculated from each team's wins, draws, losses.
 Forecasting and Championship Estimation 
•	Simulates all remaining (unplayed) matches 1000 times using the same rules as in actual match simulations.
•	For each iteration, deep copies of the current state are used to avoid altering real results.
•	After each iteration:
o	Final league standings are calculated.
o	Championship wins and top-2 finishes are recorded for each team.
•	At the end of 1000 runs:
o	Championship Probability = number of wins / 1000.
o	Top 2 Probability = number of top-2 appearances / 1000.
o	Average Points = sum of final points across all simulations / 1000.
Database (database.go, create_table.sql)
•	Tables: teams, matches, scoreboard.
•	Insert and query functions for both teams and matches.
•	Used to persist data between server sessions.
API Endpoınts
Method	Endpoint	Description
GET	/teams	List all teams
GET	/matches	List all matches
POST	/play/week/{id}	Play week id
POST	/play/all	Play all remaining weeks
GET	/table	Get current league standings
GET	/forecast	Run 1000 simulations and get forecast

Deployment(Not Yet Completed)
Postman
•	Project is fully testable locally on Postman.
Render/Fly.io (Future Plan)
•	Will deploy as a Web Service.

Full Console Output:
 We are ready to go!!
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.

 League Begins-  Teams in the League:
- Sirius Stars (Defence Strength: 6 /MidField Strength: 8 /Forward Strength: 8 /Home Team Modifier: 10 )
- Journey United (Defence Strength: 10 /MidField Strength: 5 /Forward Strength: 9 /Home Team Modifier: 5 )
- Forecasters City (Defence Strength: 7 /MidField Strength: 6 /Forward Strength: 8 /Home Team Modifier: 9 )
- Omni FC (Defence Strength: 9 /MidField Strength: 9 /Forward Strength: 9 /Home Team Modifier: 6 )

✅ Inserted team: Sirius Stars
✅ Inserted team: Journey United
✅ Inserted team: Forecasters City
✅ Inserted team: Omni FC

Match 1: Sirius Stars vs Journey United Played: false  
 Match 2: Sirius Stars vs Forecasters City Played: false  
 Match 3: Sirius Stars vs Omni FC Played: false  
 Match 4: Journey United vs Sirius Stars Played: false  
 Match 5: Journey United vs Forecasters City Played: false  
 Match 6: Journey United vs Omni FC Played: false  
 Match 7: Forecasters City vs Sirius Stars Played: false  
 Match 8: Forecasters City vs Journey United Played: false  
 Match 9: Forecasters City vs Omni FC Played: false  
 Match 10: Omni FC vs Sirius Stars Played: false  
 Match 11: Omni FC vs Journey United Played: false  
 Match 12: Omni FC vs Forecasters City Played: false  
 

--- Week 1 ---
Sirius Stars 3 - 2 Journey United
Forecasters City 3 - 3 Omni FC
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

--- Week 2 ---
Sirius Stars 5 - 5 Forecasters City
Journey United 4 - 4 Omni FC
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

--- Week 3 ---
Sirius Stars 6 - 5 Omni FC
Journey United 4 - 2 Forecasters City
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

--- Week 4 ---
Journey United 3 - 1 Sirius Stars
Omni FC 5 - 1 Forecasters City
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

 Week 4:
Kalan maç sayısı: 4

 Şampiyonluk Olasılıkları:
- Sirius Stars: 8.6%
- Journey United: 1.1%
- Forecasters City: 0.1%
- Omni FC: 90.2%

 İlk 2'ye Girme Olasılıkları:
- Sirius Stars: 56.8%
- Journey United: 12.3%
- Forecasters City: 32.9%
- Omni FC: 98.0%

Ortalama Puanlar:
- Sirius Stars: 8.69 puan
- Journey United: 7.38 puan
- Forecasters City: 5.94 puan
- Omni FC: 10.68 puan

--- Week 5 ---
Forecasters City 3 - 2 Sirius Stars
Omni FC 12 - 2 Journey United
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

 Week 5:
Kalan maç sayısı: 2

 Şampiyonluk Olasılıkları:
- Sirius Stars: 5.6%
- Journey United: 0.5%
- Forecasters City: 0.0%
- Omni FC: 93.9%

 İlk 2'ye Girme Olasılıkları:
- Sirius Stars: 9.1%
- Journey United: 18.4%
- Forecasters City: 73.1%
- Omni FC: 99.4%

 Ortalama Puanlar:
- Sirius Stars: 7.24 puan
- Journey United: 7.39 puan
- Forecasters City: 7.52 puan
- Omni FC: 10.70 puan

--- Week 6 ---
Forecasters City 6 - 3 Journey United
Omni FC 4 - 3 Sirius Stars
✅ Sql comment is read✅ Sql comment is executed✅ Database initialized.
✅ All matches inserted.

 Week 6:
Kalan maç sayısı: 0

 Şampiyonluk Olasılıkları:
- Sirius Stars: 0.0%
- Journey United: 0.0%
- Forecasters City: 0.0%
- Omni FC: 100.0%

 İlk 2'ye Girme Olasılıkları:
- Sirius Stars: 0.0%
- Journey United: 0.0%
- Forecasters City: 100.0%
- Omni FC: 100.0%

 Ortalama Puanlar:
- Sirius Stars: 7.00 puan
- Journey United: 7.00 puan
- Forecasters City: 8.00 puan
- Omni FC: 11.00 puan





Team            MP  W  D  L  GF  GA  GD  Pts
-------------------------------------------------
Omni FC          6   3  2  1  33  19  14   11
Forecasters City  6   2  2  2  20  22  -2    8
Sirius Stars     6   2  1  3  20  22  -2    7
Journey United   6   2  1  3  18  28  -10    7
-------Server running at http://localhost:8080
