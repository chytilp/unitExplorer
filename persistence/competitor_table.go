package persistence

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/chytilp/unitExplorer/request"
)

type CompetitorTable struct {
	DB *sql.DB
}

func (c *CompetitorTable) GetCompetitors(ids []int64) ([]request.Competitor, error) {
	idParams := ""
	rows, err := c.DB.Query(`SELECT competitor_id, id, name FROM competitor WHERE competitor_id IN (?)`, idParams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []request.Competitor

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var competitor request.Competitor
		if err := rows.Scan(&competitor.CompetitorId, &competitor.Id, &competitor.Name); err != nil {
			return competitors, err
		}
		competitors = append(competitors, competitor)
	}
	return competitors, nil
}
