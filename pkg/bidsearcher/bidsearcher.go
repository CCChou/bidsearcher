package bidsearcher

import (
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type BidSearcher struct {
	client *http.Client
}

func NewBidSearcher() *BidSearcher {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
	}
	return &BidSearcher{client: client}
}

func (b *BidSearcher) Search(keywork string) []Bid {
	err := b.login()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := b.client.PostForm("https://www.taiwanbuying.com.tw/QueryCloseCaseAction.ASP", url.Values{"DataType": {"OBJ"}, "Keyword": {"機油"}})
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bids []*Bid
	doc.Find("td[valign=top] > table").Eq(2).Find("td > table").Each(func(i int, s *goquery.Selection) {
		s = s.Find("td").Eq(1)
		html, _ := s.Html()
		text := s.SetHtml(strings.Replace(html, "<br/>", "<br/>\n", -1)).Text()
		bid, err := b.parse(text)
		if err != nil {
			log.Fatal(err)
		}
		bids = append(bids, bid)

	})
	return nil
}

func (b *BidSearcher) login() error {
	resp, err := b.client.PostForm("https://www.taiwanbuying.com.tw/MemLoginAction.asp", url.Values{"LogID": {"cindy886637"}, "PWD": {"Inipass07242020"}, "Submit": {"送出"}})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Log in failed with " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}

func (b *BidSearcher) parse(text string) (*Bid, error) {
	lines := strings.Split(text, "\n")
	unit := strings.Split(lines[0], ":")[1]
	caseName := strings.Split(lines[1], ":")[1]
	vendor := strings.Split(lines[2], ":")[1]

	re := regexp.MustCompile(`[\d,/]+`)
	matches := re.FindAllString(lines[3], -1)
	award, err := strconv.Atoi(strings.ReplaceAll(matches[0], ",", ""))
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("2006/1/2", matches[1])
	if err != nil {
		return nil, err
	}

	bid := &Bid{
		unit:     strings.TrimSpace(unit),
		caseName: strings.TrimSpace(caseName),
		vendor:   strings.TrimSpace(vendor),
		award:    award,
		date:     date,
	}
	return bid, nil
}
