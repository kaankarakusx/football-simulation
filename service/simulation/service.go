package simulation

import (
	"fmt"
	"football-simulation/types"
	"math/rand"
)

type Service struct {
	store types.SimulationStore
}

func NewService(store types.SimulationStore) *Service {
	return &Service{store: store}
}

func (s *Service) GenerateFixture(teams []types.Team) error {
	var matches []types.Match
	numTeams := len(teams)
	// for scalability: If the number of teams is odd, add a dummy team.
	if numTeams%2 != 0 {
		teams = append(teams, types.Team{ID: -1, Name: "BYE"})
		numTeams++
	}

	halfSeasonWeeks := (numTeams - 1)
	weekMatches := numTeams / 2

	// first half of the season
	for week := 0; week < halfSeasonWeeks; week++ {
		for match := 0; match < weekMatches; match++ {
			home := (week + match) % (numTeams - 1)
			away := (numTeams - 1 - match + week) % (numTeams - 1)

			if match == 0 {
				away = numTeams - 1
			}

			matches = append(matches, types.Match{Week: week + 1, Team1ID: teams[home].ID, Team2ID: teams[away].ID, Played: false})
		}
	}

	// second half of the season
	for week := 0; week < halfSeasonWeeks; week++ {
		for match := 0; match < weekMatches; match++ {
			home := (week + match) % (numTeams - 1)
			away := (numTeams - 1 - match + week) % (numTeams - 1)

			if match == 0 {
				away = numTeams - 1
			}

			matches = append(matches, types.Match{Week: week + 1 + halfSeasonWeeks, Team1ID: teams[away].ID, Team2ID: teams[home].ID, Played: false})
		}
	}

	// filter out matches involving the dummy team
	filteredMatches := matches[:0]
	for _, match := range matches {
		if match.Team1ID != -1 && match.Team2ID != -1 {
			filteredMatches = append(filteredMatches, match)
		}
	}

	err := s.store.SaveFixture(filteredMatches)

	if err != nil {
		return fmt.Errorf("could not save filtered matches: %v", err)
	}
	return nil
}

func (s *Service) PlayMatch(team1, team2 types.Team) (team1Score, team2Score int) {
	team1BaseScore := float64(team1.Strength) / float64(40+rand.Intn(31))
	team2BaseScore := float64(team2.Strength) / float64(40+rand.Intn(31))

	team1RandomFactor := rand.Float64() * 2
	team2RandomFactor := rand.Float64() * 2

	team1Score = int(team1BaseScore + team1RandomFactor)
	team2Score = int(team2BaseScore + team2RandomFactor)

	//set a limit of 5 goals to increase realism
	if team1Score > 5 {
		team1Score = 5
	}
	if team2Score > 5 {
		team2Score = 5
	}

	return team1Score, team2Score
}

func (s *Service) CalculateChampionshipOdds(teams []types.Team, matches []types.Match) ([]types.Prediction, error) {
	const simulationCount = 1000
	teamChampionshipCounts := make(map[int]int)

	for i := 0; i < simulationCount; i++ {
		simulatedStandings := make(map[int]int)
		for _, team := range teams {
			simulatedStandings[team.ID] = team.Points
		}

		for _, match := range matches {
			if !match.Played {
				team1 := simulatedStandings[match.Team1ID]
				team2 := simulatedStandings[match.Team2ID]
				team1Score, team2Score := s.PlayMatch(types.Team{Strength: team1}, types.Team{Strength: team2})

				if team1Score > team2Score {
					simulatedStandings[match.Team1ID] += 3
				} else if team2Score > team1Score {
					simulatedStandings[match.Team2ID] += 3
				} else {
					simulatedStandings[match.Team1ID]++
					simulatedStandings[match.Team2ID]++
				}
			}
		}

		var maxPoints int
		var championTeamID int
		for teamID, points := range simulatedStandings {
			if points > maxPoints {
				maxPoints = points
				championTeamID = teamID
			}
		}
		teamChampionshipCounts[championTeamID]++
	}

	var predictions []types.Prediction
	for _, team := range teams {
		championshipCount := teamChampionshipCounts[team.ID]
		odds := float64(championshipCount) / float64(simulationCount) * 100
		predictions = append(predictions, types.Prediction{
			TeamID:           team.ID,
			TeamName:         team.Name,
			ChampionshipOdds: odds,
		})
	}

	return predictions, nil
}
