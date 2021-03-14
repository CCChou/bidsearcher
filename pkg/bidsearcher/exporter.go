package bidsearcher

type Exporter interface {
	export(bids []*Bid) error
}

type CsvExporter struct {
}

func (c *CsvExporter) export(bids []*Bid) error {
	return nil
}
