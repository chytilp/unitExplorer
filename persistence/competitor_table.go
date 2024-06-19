package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/chytilp/unitExplorer/request"
)

type CompetitorTable struct {
	DB       *sql.DB
	SourceId int
	DomainId string
}

func (c *CompetitorTable) DeleteCompetitors() error {
	deleteSQL := `DELETE FROM competitor WHERE source_id = ? AND domain_id = ?`
	statement, err := c.DB.Prepare(deleteSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(c.SourceId, c.DomainId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (c *CompetitorTable) InsertCompetitor(competitor request.Competitor) (int64, error) {
	insertSQL := `INSERT INTO competitior(id, name, source_id, domain_id) VALUES (?, ?, ?, ?)`
	statement, err := c.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := statement.Exec(competitor.Id, competitor.Name, c.SourceId, c.DomainId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	newId, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return newId, nil
}

func (c *CompetitorTable) GetCompetitors(ids []int64) ([]request.Competitor, error) {
	idParams := fmt.Sprintf("%d,%d", ids[0], ids[1])
	rows, err := c.DB.Query(`SELECT competitor_id, id, name FROM competitor WHERE competitor_id IN (?)`, idParams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	competitors := make([]request.Competitor, 0, 2)

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
