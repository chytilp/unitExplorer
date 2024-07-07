package formatter

import (
	"fmt"

	"github.com/chytilp/unitExplorer/request"
)

type DomainFormatter struct {
	filters *[]request.Base
	results *[]request.Domain
}

func NewDomainFormatter() *DomainFormatter {
	return &DomainFormatter{}
}

func (d *DomainFormatter) SetFilters(filters []request.Base) {
	d.filters = &filters
}

func (d *DomainFormatter) SetResults(results []request.Domain) {
	d.results = &results
}

func (d *DomainFormatter) Print() error {
	if d.filters == nil {
		return fmt.Errorf("no filters assign")
	}
	if d.results == nil {
		return fmt.Errorf("no results assign")
	}
	for _, domain := range *d.results {
		fmt.Printf("domain -> id: %s, name: %s\n", domain.Id, domain.Name)
	}
	return nil
}
