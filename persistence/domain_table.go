package persistence

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/chytilp/unitExplorer/request"
)

type DomainTable struct {
	DB       *sql.DB
	SourceId int
}

func (d *DomainTable) DeleteDomains() error {
	deleteSQL := `DELETE FROM domain WHERE source_id = ?`
	statement, err := d.DB.Prepare(deleteSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(d.SourceId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (d *DomainTable) insertDomain(domain request.Domain) error {
	insertSQL := `INSERT INTO domain(source_id, domain_id, domain_name, matched_at) VALUES (?, ?, ?, ?)`
	statement, err := d.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(d.SourceId, domain.Id, domain.Name, time.Now())
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func (d *DomainTable) InsertDomains(domains []request.Domain) error {
	for _, domain := range domains {
		err := d.insertDomain(domain)
		if err != nil {
			fmt.Printf("Inserting of domain: %v, sourceId: %d failed: %v\n", domain, d.SourceId, err)
			return err
		}
	}
	return nil
}

func (d *DomainTable) GetDomain(domainId string) (*request.Domain, error) {
	var rec request.Domain
	err := d.DB.QueryRow(`SELECT domain_id, domain_name FROM domain WHERE domain_id = ?`,
		domainId).Scan(&rec.Id, &rec.Name)
	if err != nil {
		fmt.Printf("Selecting of domain: %s, sourceId: %d failed: %v\n", domainId, d.SourceId, err)
		return nil, err
	}
	return &rec, nil
}
