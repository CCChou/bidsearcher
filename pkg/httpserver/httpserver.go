package httpserver

import (
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/CCChou/bidsearcher/pkg/bidsearcher"
)

var baseDir string = "files/"
var config *Config

func Serve() {
	InitConfig()
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./frontend/build")))
	mux.HandleFunc("/files/", downloadCsv)
	mux.HandleFunc("/search", search)
	http.ListenAndServe(":8080", mux)
}

func InitConfig() {
	var err error
	config, err = NewConfig()
	if err != nil {
		log.Fatal("Init config failed")
		panic(err)
	}
}

type Response struct {
	Url    string `json:"url"`
	Status string `json:"status"`
}

func downloadCsv(w http.ResponseWriter, r *http.Request) {
	handle := http.StripPrefix("/files/", http.FileServer(http.Dir("./files")))
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	handle.ServeHTTP(w, r)
}

func search(w http.ResponseWriter, r *http.Request) {
	keyword := strings.TrimSpace(r.URL.Query().Get("keyword"))
	b := bidsearcher.NewBidSearcher(config.Username, config.Password)
	bids := b.Search(keyword)
	e := bidsearcher.NewExporter()
	h := fnv.New64a()
	h.Write([]byte(time.Now().UTC().String()))
	path := baseDir + hex.EncodeToString(h.Sum(nil)) + ".csv"
	e.Export(bids, path)

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(&Response{
		Url:    path,
		Status: "OK",
	})
	w.Write(json)
}
