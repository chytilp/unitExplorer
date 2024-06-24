package command

import (
	"fmt"

	"github.com/chytilp/unitExplorer/persistence"
	"github.com/chytilp/unitExplorer/request"
)

type Command interface {
	Validate() error
	Run() error
}

func GetDomain(sourceId int, domainId string, databaseFile string) (*request.Domain, error) {
	db, err := persistence.GetDatabase(databaseFile)
	if err != nil {
		fmt.Printf("GetDatabase err: %v\n", err)
		return nil, err
	}
	domainTable := persistence.DomainTable{
		DB:       db,
		SourceId: sourceId,
	}
	domain, err := domainTable.GetDomain(domainId)
	if err != nil {
		fmt.Printf("domainTable.GetDomain err: %v\n", err)
		return nil, err
	}
	return domain, nil
}
