package league

import (
	"errors"
	"football-simulation/types"
)

type Service struct {
	store             types.LeagueStore
	teamService       types.TeamService
	simulationService types.SimulationService
}

func NewService(store types.LeagueStore, simulationService types.SimulationService, teamService types.TeamService) *Service {
	return &Service{
		store:             store,
		teamService:       teamService,
		simulationService: simulationService,
	}
}

func (s *Service) StartLeague() error {

	teams, err := s.teamService.GetTeams()

	if err != nil {
		return err
	}
	err = s.simulationService.GenerateFixture(teams)

	if err != nil {
		return err
	}

	totalWeeks := calculateTotalWeeks(teams)

	league, err := s.store.GetLeagueInfo()

	if err != nil {
		return err
	}

	err = s.store.UpdateLeague(types.League{
		ID:               league.ID,
		Name:             league.Name,
		CurrentWeek:      league.CurrentWeek + 1,
		TotalWeeks:       totalWeeks,
		ChampionTeamName: league.ChampionTeamName,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) NextWeek() ([]types.Match, *types.Team, error) {
	matches, err := s.store.GetAllMatches()
	if err != nil {
		return nil, nil, err
	}

	if len(matches) == 0 {
		//iff there is no match, start the league
		err = s.StartLeague()
		if err != nil {
			return nil, nil, err
		}

	}

	matches, err = s.store.GetMatchesForNextWeek()
	if err != nil {
		return nil, nil, err
	}

	var playedMatches []types.Match

	for _, match := range matches {
		team1, err := s.teamService.GetTeamByID(match.Team1ID)
		if err != nil {
			return nil, nil, err
		}

		team2, err := s.teamService.GetTeamByID(match.Team2ID)
		if err != nil {
			return nil, nil, err
		}

		team1Score, team2Score := s.simulationService.PlayMatch(*team1, *team2)
		match.Team1Score = team1Score
		match.Team2Score = team2Score
		match.Played = true

		err = s.store.SaveMatchResult(match)
		if err != nil {
			return nil, nil, err
		}

		err = s.teamService.UpdateTeamStats(*team1, *team2, team1Score, team2Score, false)
		if err != nil {
			return nil, nil, err
		}
		playedMatches = append(playedMatches, match)
	}

	err = s.store.IncrementWeek()
	if err != nil {
		return nil, nil, err
	}

	remainingMatches, err := s.store.GetMatchesForNextWeek()
	if err != nil {
		return nil, nil, err
	}

	if len(remainingMatches) == 0 {
		standings, err := s.store.GetStandings()
		if err != nil {
			return nil, nil, err
		}

		champion := &standings[0]
		return playedMatches, champion, nil
	}

	return playedMatches, nil, nil

}

func (s *Service) PlayAll() ([]types.MatchResult, *types.Team, error) {
	matches, err := s.GetAllMatches()
	if err != nil {
		return nil, nil, err
	}

	if len(matches) == 0 {
		err = s.StartLeague()
		if err != nil {
			return nil, nil, err
		}
		matches, err = s.GetAllMatches()
		if err != nil {
			return nil, nil, err
		}
	}

	var playedMatches []types.MatchResult

	for _, match := range matches {
		if !match.Played {
			team1, err := s.teamService.GetTeamByName(match.Team1Name)
			if err != nil {
				return nil, nil, err
			}

			team2, err := s.teamService.GetTeamByName(match.Team2Name)
			if err != nil {
				return nil, nil, err
			}

			team1Score, team2Score := s.simulationService.PlayMatch(*team1, *team2)
			match.Team1Score = team1Score
			match.Team2Score = team2Score
			match.Played = true

			matchToSave := types.Match{
				ID:         match.ID,
				Week:       match.Week,
				Team1ID:    team1.ID,
				Team2ID:    team2.ID,
				Team1Score: team1Score,
				Team2Score: team2Score,
				Played:     true,
			}

			err = s.store.SaveMatchResult(matchToSave)
			if err != nil {
				return nil, nil, err
			}

			err = s.teamService.UpdateTeamStats(*team1, *team2, team1Score, team2Score, false)
			if err != nil {
				return nil, nil, err
			}
		}

		playedMatches = append(playedMatches, types.MatchResult{
			ID:         match.ID,
			Week:       match.Week,
			Team1Name:  match.Team1Name,
			Team2Name:  match.Team2Name,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Played:     match.Played,
		})
	}

	//champion
	standings, err := s.store.GetStandings()
	if err != nil {
		return nil, nil, err
	}

	champion := &standings[0]
	return playedMatches, champion, nil
}
func (s *Service) GetWeekResults() ([]types.MatchResult, error) {
	currentWeek, err := s.store.GetCurrentWeek()
	if err != nil {
		return nil, err
	}

	matches, err := s.store.GetMatchesByWeek(currentWeek - 1)
	if err != nil {
		return nil, err
	}

	var weekResults []types.MatchResult

	for _, match := range matches {
		team1, err := s.teamService.GetTeamByID(match.Team1ID)
		if err != nil {
			return nil, err
		}

		team2, err := s.teamService.GetTeamByID(match.Team2ID)
		if err != nil {
			return nil, err
		}

		weekResults = append(weekResults, types.MatchResult{
			ID:         match.ID,
			Week:       match.Week,
			Team1Name:  team1.Name,
			Team2Name:  team2.Name,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Played:     match.Played,
		})
	}

	return weekResults, nil
}

func (s *Service) GetMatchesByWeek(id int) ([]types.MatchResult, error) {

	matches, err := s.store.GetMatchesByWeek(id)
	if err != nil {
		return nil, err
	}

	var weekResults []types.MatchResult

	for _, match := range matches {
		team1, err := s.teamService.GetTeamByID(match.Team1ID)
		if err != nil {
			return nil, err
		}

		team2, err := s.teamService.GetTeamByID(match.Team2ID)
		if err != nil {
			return nil, err
		}

		weekResults = append(weekResults, types.MatchResult{
			ID:         match.ID,
			Week:       match.Week,
			Team1Name:  team1.Name,
			Team2Name:  team2.Name,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Played:     match.Played,
		})
	}

	return weekResults, nil
}

func (s *Service) GetAllMatches() ([]types.MatchResult, error) {
	matches, err := s.store.GetAllMatches()
	if err != nil {
		return nil, err
	}

	var matchResults []types.MatchResult

	for _, match := range matches {
		team1, err := s.teamService.GetTeamByID(match.Team1ID)
		if err != nil {
			return nil, err
		}

		team2, err := s.teamService.GetTeamByID(match.Team2ID)
		if err != nil {
			return nil, err
		}

		matchResults = append(matchResults, types.MatchResult{
			ID:         match.ID,
			Week:       match.Week,
			Team1Name:  team1.Name,
			Team2Name:  team2.Name,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Played:     match.Played,
		})
	}

	return matchResults, nil
}

func (s *Service) UpdateMatch(match types.Match) error {
	existingMatch, err := s.store.GetMatchByID(match.ID)
	if err != nil {
		return err
	}

	if err := s.teamService.UpdateTeamStatsReverse(*existingMatch); err != nil {
		return err
	}

	existingMatch.Team1Score = match.Team1Score
	existingMatch.Team2Score = match.Team2Score
	if err := s.store.UpdateMatch(*existingMatch); err != nil {
		return err
	}

	team1, err := s.teamService.GetTeamByID(existingMatch.Team1ID)
	if err != nil {
		return err
	}

	team2, err := s.teamService.GetTeamByID(existingMatch.Team2ID)
	if err != nil {
		return err
	}

	if err := s.teamService.UpdateTeamStats(*team1, *team2, existingMatch.Team1Score, existingMatch.Team2Score, true); err != nil {
		return err
	}

	return nil
}

func (s *Service) RestartLeague() error {

	err := s.store.ClearFixtures()

	if err != nil {
		return err
	}

	err = s.teamService.ResetTeams()

	if err != nil {
		return err
	}

	league, err := s.store.GetLeagueInfo()

	if err != nil {
		return err
	}

	err = s.store.UpdateLeague(types.League{
		ID:               league.ID,
		Name:             league.Name,
		CurrentWeek:      0,
		TotalWeeks:       0,
		ChampionTeamName: "",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStandings() ([]types.Team, error) {
	teams, err := s.store.GetStandings()

	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (s *Service) GetPredictions() ([]types.Prediction, error) {
	league, err := s.store.GetLeagueInfo()
	if err != nil {
		return nil, err
	}

	if league.CurrentWeek < 4 {
		return nil, errors.New("championship predictions can only be made after week 4")
	}

	teams, err := s.store.GetStandings()
	if err != nil {
		return nil, err
	}

	matchResults, err := s.GetAllMatches()
	if err != nil {
		return nil, err
	}

	matches := make([]types.Match, len(matchResults))
	for i, matchResult := range matchResults {
		team1, err := s.teamService.GetTeamByName(matchResult.Team1Name)
		if err != nil {
			return nil, err
		}
		team2, err := s.teamService.GetTeamByName(matchResult.Team2Name)
		if err != nil {
			return nil, err
		}

		matches[i] = types.Match{
			ID:         matchResult.ID,
			Week:       matchResult.Week,
			Team1ID:    team1.ID,
			Team2ID:    team2.ID,
			Team1Score: matchResult.Team1Score,
			Team2Score: matchResult.Team2Score,
			Played:     matchResult.Played,
		}
	}

	return s.simulationService.CalculateChampionshipOdds(teams, matches)
}
func calculateTotalWeeks(teams []types.Team) int {
	return 2 * (len(teams) - 1)
}
