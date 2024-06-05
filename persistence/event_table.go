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
	return nil
}

func (e *EventTable) insertEvent(event request.Event) error {
	insertSQL := `INSERT INTO domain(source_id, domain_id, domain_name, matched_at) VALUES (?, ?, ?, ?)`
	statement, err := e.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(e.SourceId, event.Id, event.Name, time.Now())
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (e *EventTable) InsertEvents(events []request.Event) error {
	for _, event := range events {
		err := e.insertEvent(event)
		if err != nil {
			fmt.Printf("Inserting of event: %v, sourceId: %d failed: %v\n", event, e.SourceId, err)
			return err
		}
	}
	return nil
}

func (e *EventTable) GetEvent(eventId string) (*request.Event, error) {
	var rec request.Event
	err := e.DB.QueryRow(`SELECT domain_id, domain_name FROM domain WHERE id = ?`,
		eventId).Scan(&rec.Id, &rec.Name)
	if err != nil {
		fmt.Printf("Selecting of event: %s, sourceId: %d failed: %v\n", eventId, e.SourceId, err)
		return nil, err
	}
	return &rec, nil
}
