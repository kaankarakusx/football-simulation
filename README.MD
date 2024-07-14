# Football Simulation API

## Project Overview

The Football Simulation API simulates a football league with four teams. Users can simulate weekly match results, view league standings, and get championship predictions.

## Setup and Installation Instructions

### Requirements

- Go 1.16+
- PostgreSQL

### Steps

1. Clone the repository:

   ```sh
   git clone https://github.com/kaankarakusx/football-simulation.git
   cd football-simulation
   ```

2. Set up environment variables:

   ```sh
    .env
   ```

3. Create the PostgreSQL database

4. Build and run the application using Makefile:
   ```sh
   make build
   make run
   ```

## API Endpoints

### League Management

- **Restart League**: Resets and restarts the league.

  - URL: `/api/v1/league/restart`
  - Method: `POST`

- **Get Standings**: Returns the current league standings.

  - URL: `/api/v1/league/standings`
  - Method: `GET`

- **Next Week**: Simulates the next week's matches.

  - URL: `/api/v1/league/nextweek`
  - Method: `POST`

- **Play All**: Simulates all remaining matches and determines the champion.
  - URL: `/api/v1/league/playall`
  - Method: `POST`

### Match Management

- **Get Matches Weekly**: Returns that week matches

- URL: `/api/v1/league/weekresults`
- Method: `GET`

- **Get Matches by Week**: Returns matches for a given week.

  - URL: `/api/v1/league/matches/{week}`
  - Method: `GET`

- **Get All Matches**: Returns all matches.

  - URL: `/api/v1/league/matches`
  - Method: `GET`

- **Update Match Results**: Updates the results of a match.
  - URL: `/api/v1/league/match/{id}`
  - Method: `PUT`

### Championship Prediction

- **Get Championship Predictions**: Returns championship odds.
  - URL: `/api/v1/league/predictions`
  - Method: `GET`
