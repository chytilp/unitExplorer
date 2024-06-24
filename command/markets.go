package command

import (
	"fmt"

	"github.com/chytilp/unitExplorer/common"
	"github.com/chytilp/unitExplorer/persistence"
	"github.com/chytilp/unitExplorer/request"
)

type ListMarkets struct {
	SourceName string
	Config     *common.Config
	DomainId   string
	EventId    string
}

func (m *ListMarkets) Validate() error {
	return nil
}

func (m *ListMarkets) getDomain(sourceId int) (*request.Domain, error) {
	domain, err := GetDomain(sourceId, m.DomainId, m.Config.DatabaseFile)
	if err != nil {
		fmt.Printf("err in getDomain: %v\n", err)
		return nil, err
	}
	return domain, nil
}

func (m *ListMarkets) getEvent(sourceId int) (*request.Event, error) {
	db, err := persistence.GetDatabase(m.Config.DatabaseFile)
	if err != nil {
		fmt.Printf("GetDatabase err: %v\n", err)
		return nil, err
	}
	eventTable := persistence.EventTable{
		DB:       db,
		SourceId: sourceId,
		DomainId: m.DomainId,
	}
	event, err := eventTable.GetEvent(m.EventId)
	if err != nil {
		fmt.Printf("eventTable.GetEvent err: %v\n", err)
		return nil, err
	}
	return event, nil
}

func (m *ListMarkets) Run() error {
	sender := request.NewSender(m.Config.ApiUrl)
	sourceId := m.Config.FindSourceId(m.SourceName)
	if sourceId == nil {
		return fmt.Errorf("source of name %s was not found", m.SourceName)
	}
	domain, err := m.getDomain(*sourceId)
	if err != nil {
		fmt.Printf("getDomain err: %v\n", err)
		return err
	}
	event, err := m.getEvent(*sourceId)
	if err != nil {
		fmt.Printf("getEvent err: %v\n", err)
		return err
	}
	req, err := request.CreateMarketRequest(*sourceId, *domain, *event)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	payload, err := sender.GetMarkets(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	for _, market := range payload.Payload {
		fmt.Printf("market -> id: %s, name: %s\n", market.Id, market.Name)
	}
	return nil
}
