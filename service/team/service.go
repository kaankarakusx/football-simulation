package team

import "football-simulation/types"

type Service struct {
	store types.Teamstore
}

func NewService(store types.Teamstore) *Service {
	return &Service{store: store}
}

func (s *Service) GetTeams() ([]types.Team, error) {

	teams, err := s.store.GetTeams()

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *Service) GetTeamByID(id int) (*types.Team, error) {
	team, err := s.store.GetTeamByID(id)

	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *Service) GetTeamByName(name string) (*types.Team, error) {
	team, err := s.store.GetTeamByName(name)

	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *Service) UpdateTeam(team types.Team) error {
	err := s.store.UpdateTeam(team)

	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTeamStatsReverse(match types.Match) error {
	team1, err := s.GetTeamByID(match.Team1ID)
	if err != nil {
		return err
	}

	team2, err := s.GetTeamByID(match.Team2ID)
	if err != nil {
		return err
	}

	if match.Team1Score > match.Team2Score {
		team1.Wins--
		team1.Points -= 3
		team2.Losses--
	} else if match.Team1Score < match.Team2Score {
		team2.Wins--
		team2.Points -= 3
		team1.Losses--
	} else {
		team1.Draws--
		team2.Draws--
		team1.Points--
		team2.Points--
	}

	team1.GoalsFor -= match.Team1Score
	team1.GoalsAgainst -= match.Team2Score
	team1.GoalsDifference = team1.GoalsFor - team1.GoalsAgainst

	team2.GoalsFor -= match.Team2Score
	team2.GoalsAgainst -= match.Team1Score
	team2.GoalsDifference = team2.GoalsFor - team2.GoalsAgainst

	if err := s.UpdateTeam(*team1); err != nil {
		return err
	}

	if err := s.UpdateTeam(*team2); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateTeamStats(team1, team2 types.Team, team1Score, team2Score int, isUpdate bool) error {

	if team1Score > team2Score {
		team1.Wins++
		team1.Points += 3
		team2.Losses++
	} else if team1Score < team2Score {
		team2.Wins++
		team2.Points += 3
		team1.Losses++
	} else {
		team1.Draws++
		team2.Draws++
		team1.Points++
		team2.Points++
	}

	if !isUpdate {
		team1.Matches++
		team2.Matches++
	}

	team1.GoalsFor += team1Score
	team1.GoalsAgainst += team2Score
	team1.GoalsDifference = team1.GoalsFor - team1.GoalsAgainst

	team2.GoalsFor += team2Score
	team2.GoalsAgainst += team1Score
	team2.GoalsDifference = team2.GoalsFor - team2.GoalsAgainst

	err := s.UpdateTeam(team1)
	if err != nil {
		return err
	}

	err = s.UpdateTeam(team2)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ResetTeams() error {
	err := s.store.ResetTeams()

	if err != nil {
		return err
	}

	return nil
}
