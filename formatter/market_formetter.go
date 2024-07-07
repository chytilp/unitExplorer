package formatter

import (
	"fmt"

	"github.com/chytilp/unitExplorer/request"
)

type MarketFormatter struct {
	filters *[]request.Base
	results *[]request.Market
}

func NewMarketFormatter() *MarketFormatter {
	return &MarketFormatter{}
}

func (m *MarketFormatter) SetFilters(filters []request.Base) {
	m.filters = &filters
}

func (m *MarketFormatter) SetResults(results []request.Market) {
	m.results = &results
}

func (e *MarketFormatter) Print() error {
	fmt.Print("filter etc.")
	return nil
}
