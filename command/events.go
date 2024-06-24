package command

import (
	"fmt"

	"github.com/chytilp/unitExplorer/common"
	"github.com/chytilp/unitExplorer/persistence"
	"github.com/chytilp/unitExplorer/request"
)

type ListEvents struct {
	SourceName string
	Config     *common.Config
	DomainId   string
}

func (e *ListEvents) Validate() error {
	return nil
}

func (e *ListEvents) getDomain(sourceId int) (*request.Domain, error) {
	domain, err := GetDomain(sourceId, e.DomainId, e.Config.DatabaseFile)
	if err != nil {
		fmt.Printf("err in getDomain: %v\n", err)
		return nil, err
	}
	return domain, nil
}

func (e *ListEvents) Run() error {
	sender := request.NewSender(e.Config.ApiUrl)
	sourceId := e.Config.FindSourceId(e.SourceName)
	if sourceId == nil {
		return fmt.Errorf("source of name %s was not found", e.SourceName)
	}
	domain, err := e.getDomain(*sourceId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	req, err := request.CreateEventRequest(*sourceId, *domain)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	payload, err := sender.GetEvents(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	for _, event := range payload.Payload {
		fmt.Printf("event -> id: %s, name: %s\n", event.Id, event.Name)
	}
	err = e.save(*sourceId, domain.Id, payload.Payload)
	if err != nil {
		return err
	}
	return nil
}

func (e *ListEvents) save(sourceId int, domainId string, events []request.Event) error {
	db, err := persistence.GetDatabase(e.Config.DatabaseFile)
	if err != nil {
		fmt.Printf("GetDatabase err: %v\n", err)
		return err
	}
	eventTable := persistence.EventTable{
		DB:       db,
		SourceId: sourceId,
		DomainId: domainId,
	}
	err = eventTable.DeleteEvents()
	if err != nil {
		fmt.Printf("DeleteEvents err: %v\n", err)
		return err
	}
	err = eventTable.InsertEvents(events)
	if err != nil {
		fmt.Printf("InsertEvents err: %v\n", err)
		return err
	}
	return nil
}
