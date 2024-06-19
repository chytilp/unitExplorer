package persistence

import (
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/chytilp/unitExplorer/request"
)

type CompetitionTable struct {
	DB       *sql.DB
	SourceId int
	DomainId string
}

func (c *CompetitionTable) DeleteCompetitions() error {
	deleteSQL := `DELETE FROM competition WHERE source_id = ? AND domain_id = ?`
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

func (c *CompetitionTable) InsertCompetition(competition request.Competition) (int64, error) {
	insertSQL := `INSERT INTO competition(id, name, source_id, domain_id) VALUES (?, ?, ?, ?)`
	statement, err := c.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := statement.Exec(competition.Id, competition.Name, c.SourceId, c.DomainId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	newId, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return newId, nil
}

func (c *CompetitionTable) GetCompetition(id int64) (*request.Competition, error) {
	var rec request.Competition
	err := c.DB.QueryRow(`SELECT competition_id, id, name FROM competition WHERE competition_id = ?`,
		id).Scan(&rec)
	if err != nil {
		return nil, err
	}
	return &rec, nil
}
