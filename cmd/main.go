package main

import "github.com/CCChou/bidsearcher/pkg/bidsearcher"

func main() {
	b := bidsearcher.NewBidSearcher()
	bids := b.Search("考察")
	e := bidsearcher.NewExporter()
	e.Export(bids)
}
