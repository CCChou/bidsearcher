package bidsearcher

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Exporter interface {
	Export(bids []*Bid) error
}

func NewExporter() Exporter {
	return &CsvExporter{}
}

type CsvExporter struct {
}

func (c *CsvExporter) Export(bids []*Bid) error {
	csvFiles, err := os.Create("./bids.csv")
	if err != nil {
		return err
	}
	w := csv.NewWriter(csvFiles)
	err = c.writeHeader(w)
	if err != nil {
		return nil
	}
	for _, bid := range bids {
		record := c.toRecord(bid)
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}

func (c *CsvExporter) writeHeader(w *csv.Writer) error {
	record := []string{"招標單位", "招標案名", "決標廠商", "決標金額", "決標日期"}
	if err := w.Write(record); err != nil {
		return err
	}
	return nil
}

func (c *CsvExporter) toRecord(bid *Bid) []string {
	award := strconv.Itoa(bid.award)
	date := bid.date.Format("2006/1/2")
	return []string{bid.unit, bid.caseName, bid.vendor, award, date}
}
