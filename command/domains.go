package command

import (
	"fmt"

	"github.com/chytilp/unitExplorer/common"
	"github.com/chytilp/unitExplorer/persistence"
	"github.com/chytilp/unitExplorer/request"
)

type ListDomains struct {
	SourceName string
	Config     *common.Config
}

func (d *ListDomains) Validate() error {
	return nil
}

func (d *ListDomains) Run() error {
	sender := request.NewSender(d.Config.ApiUrl)
	sourceId := d.Config.FindSourceId(d.SourceName)
	if sourceId == nil {
		return fmt.Errorf("source of name %s was not found", d.SourceName)
	}
	req, err := request.CreateDomainRequest(*sourceId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	payload, err := sender.GetDomains(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	for _, domain := range payload.Payload {
		fmt.Printf("domain -> id: %s, name: %s\n", domain.Id, domain.Name)
	}
	err = d.save(*sourceId, payload.Payload)
	if err != nil {
		return err
	}
	fmt.Println("successfuly saved to database")
	return nil
}

func (d *ListDomains) save(sourceId int, domains []request.Domain) error {
	db, err := persistence.GetDatabase(d.Config.DatabaseFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	domainTable := persistence.DomainTable{
		DB:       db,
		SourceId: sourceId,
	}
	err = domainTable.DeleteDomains()
	if err != nil {
		fmt.Printf("err during delete domains: %v\n", err)
		return err
	}
	err = domainTable.InsertDomains(domains)
	if err != nil {
		fmt.Printf("err during insert domains: %v\n", err)
		return err
	}
	return nil
}
