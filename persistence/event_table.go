package persistence

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	"github.com/chytilp/unitExplorer/request"
	_ "github.com/mattn/go-sqlite3"
)

type EventTable struct {
	DB       *sql.DB
	SourceId int
	DomainId string
}

func (e *EventTable) DeleteEvents() error {
	deleteSQL := `DELETE FROM event WHERE source_id = ? AND domain_id = ?`
	statement, err := e.DB.Prepare(deleteSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(e.SourceId, e.DomainId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	competitionTable := e.createCompetitionTable()
	err = competitionTable.DeleteCompetitions()
	if err != nil {
		log.Fatalln(err.Error())
	}
	competitorTable := e.createCompetitorTable()
	err = competitorTable.DeleteCompetitors()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (e *EventTable) createCompetitionTable() CompetitionTable {
	return CompetitionTable{
		DB:       e.DB,
		SourceId: e.SourceId,
		DomainId: e.DomainId,
	}
}

func (e *EventTable) createCompetitorTable() CompetitorTable {
	return CompetitorTable{
		DB:       e.DB,
		SourceId: e.SourceId,
		DomainId: e.DomainId,
	}
}

func (e *EventTable) insertCompetition(event request.Event) (int64, error) {
	competitionTable := e.createCompetitionTable()
	id, err := competitionTable.InsertCompetition(event.Competition)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return id, nil
}

func (e *EventTable) insertCompetitors(event request.Event) ([]int64, error) {
	competitorTable := e.createCompetitorTable()
	id1, err := competitorTable.InsertCompetitor(event.Competitors[0])
	if err != nil {
		log.Fatalln(err.Error())
	}
	id2, err := competitorTable.InsertCompetitor(event.Competitors[1])
	if err != nil {
		log.Fatalln(err.Error())
	}
	return []int64{id1, id2}, nil
}

func (e *EventTable) insertEvent(event request.Event) error {
	competitionId, err := e.insertCompetition(event)
	if err != nil {
		log.Fatalln(err.Error())
	}
	competitorIds, err := e.insertCompetitors(event)
	if err != nil {
		log.Fatalln(err.Error())
	}
	insertSQL := `INSERT INTO event(source_id, domain_id, id, name, full_data, type, start_date, 
		competition_id, competitor1_id, competitor2_id, matched_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := e.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(e.SourceId, e.DomainId, event.Id, event.Name, event.FullData, event.Type,
		event.StartDate, competitionId, competitorIds[0], competitorIds[1], time.Now())
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (e *EventTable) InsertEvents(events []request.Event) error {
	for _, event := range events {
		err := e.insertEvent(event)
		if err != nil {
			fmt.Printf("Inserting of event: %v, sourceId: %d, domainId: %s failed: %v\n", event, e.SourceId,
				e.DomainId, err)
			return err
		}
	}
	return nil
}

func (e *EventTable) GetEvent(eventId string) (*request.Event, error) {
	var rec request.Event
	var competitionId, competitorId1, competitorId2 int64

	err := e.DB.QueryRow(`SELECT id, name, full_data, type, start_date, competition_id, competitor1_id, 
	    competitor2_id FROM event WHERE id = ?`, eventId).Scan(&rec.Id, &rec.Name, &rec.FullData, &rec.Type,
		&rec.StartDate, &competitionId, &competitorId1, &competitorId2)
	if err != nil {
		fmt.Printf("Selecting of event: %s, sourceId: %d failed: %v\n", eventId, e.SourceId, err)
		return nil, err
	}
	competitionTable := e.createCompetitionTable()
	competition, err := competitionTable.GetCompetition(competitionId)
	if err != nil {
		fmt.Printf("Selecting of competition: %d failed: %v\n", competitionId, err)
		return nil, err
	}
	rec.Competition = *competition
	competitorTable := e.createCompetitorTable()
	competitors, err := competitorTable.GetCompetitors([]int64{competitorId1, competitorId2})
	if err != nil {
		fmt.Printf("Selecting of competitors: %d, %d failed: %v\n", competitorId1, competitorId2, err)
		return nil, err
	}
	rec.Competitors = competitors
	return &rec, nil
}
