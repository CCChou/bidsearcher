package bidsearcher

import (
	"errors"
	"io/ioutil"
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
	client     *http.Client
	baseUrl    string
	loginPath  string
	searchPath string
	username   string
	password   string
}

func NewBidSearcher(username string, password string) *BidSearcher {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
	}
	return &BidSearcher{
		client:     client,
		baseUrl:    "https://www.taiwanbuying.com.tw/",
		loginPath:  "MemLoginAction.asp",
		searchPath: "QueryCloseCaseAction.ASP",
		username:   username,
		password:   password,
	}
}

func (b *BidSearcher) Search(keyword string) []*Bid {
	log.Printf("Search %s", keyword)
	err := b.login()
	if err != nil {
		log.Fatal(err)
	}

	var bids []*Bid
	doc, err := b.getDocumentByPost(b.baseUrl+b.searchPath, url.Values{"DataType": {"OBJ"}, "Keyword": {keyword}})
	if err != nil {
		log.Println(b.baseUrl+b.searchPath, err)
	}
	bids = append(bids, b.getBids(doc)...)

	for {
		nextPage, exists := b.getNextPage(doc)
		if !exists {
			break
		}
		doc, err = b.getDocumentByGet(nextPage)
		if err != nil {
			log.Println(nextPage, err)
		}
		bids = append(bids, b.getBids(doc)...)
	}
	return bids
}

func (b *BidSearcher) login() error {
	resp, err := b.client.PostForm(b.baseUrl+b.loginPath, url.Values{"LogID": {b.username}, "PWD": {b.password}, "Submit": {"送出"}})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Log in failed with " + strconv.Itoa(resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(body)
	if strings.Contains(bodyString, "帳號或密碼輸入錯誤") {
		return errors.New("Login failed with wrong username or password")
	}
	return nil
}

func (b *BidSearcher) getDocumentByPost(url string, data url.Values) (*goquery.Document, error) {
	resp, err := b.client.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (b *BidSearcher) getBids(doc *goquery.Document) []*Bid {
	var bids []*Bid
	doc.Find("td[valign=top] > table").Eq(2).Find("td > table").Each(func(i int, s *goquery.Selection) {
		s = s.Find("td").Eq(1)
		html, _ := s.Html()
		text := s.SetHtml(strings.Replace(html, "<br/>", "<br/>\n", -1)).Text()
		bid, err := b.parse(text)
		if err != nil {
			log.Println(err)
		}
		bids = append(bids, bid)
	})
	return bids
}

func (b *BidSearcher) parse(text string) (*Bid, error) {
	lines := strings.Split(text, "\n")
	var cleanLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleanLines = append(cleanLines, line)
		}
	}
	unit := strings.Split(cleanLines[0], ":")[1]
	caseName := strings.Split(cleanLines[1], ":")[1]
	vendor := strings.Split(cleanLines[2], ":")[1]

	var err error
	var award int
	var date time.Time
	if len(cleanLines) > 4 {
		re := regexp.MustCompile(`[\d,/]+`)
		matches := re.FindAllString(cleanLines[3], -1)
		award, err = strconv.Atoi(strings.ReplaceAll(matches[0], ",", ""))
		if err != nil {
			return nil, err
		}
		date, err = time.Parse("2006/1/2", matches[1])
		if err != nil {
			return nil, err
		}
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

func (b *BidSearcher) getNextPage(doc *goquery.Document) (string, bool) {
	nextPage, exists := doc.Find("#Pagers td[align=left] a").First().Attr("href")
	return b.baseUrl + nextPage, exists
}

func (b *BidSearcher) getDocumentByGet(url string) (*goquery.Document, error) {
	resp, err := b.client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
