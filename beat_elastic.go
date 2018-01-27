package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"os"
)


// Identifier returns a string which identifies the object
type Identifier interface {
	Identify() string
}

func getElasticsearchUrl() string {
	// ElasticsearchURL is where to find elastic
	var elasticsearchURL= os.Getenv("ELASTIC_URL")
	if elasticsearchURL == "" {
		mlog.Printf("INFO: Using http://localhost:9200 as default ELASTIC_URL.")
		elasticsearchURL = "http://localhost:9200"
	} else {
		mlog.Printf("INFO: Using " + elasticsearchURL + " as ELASTIC_URL.")
	}
	return elasticsearchURL
}

func writeBulk(bulk, indx, typ string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/_bulk", getElasticsearchUrl(), indx, typ)
	req, err := http.NewRequest("POST", url, strings.NewReader(bulk))
	req.Header.Add("Content-Type", "application/octet-stream")
	if err != nil {
		mlog.Fatalln("Error creating request: %s", err)
		return nil, err
	}
	hc := http.Client{}
	res, err := hc.Do(req)
	if err != nil {
		mlog.Fatalln("Error sending request: %s", err)
		return nil, err
	}
	return res, nil
}

// WriteRecords write any identifyable object to Elasticsearch.
// Assumes that object can be marshalled to JSON
func WriteRecords(recs []Identifier, indx, typ string) (*http.Response, error) {
	bulk := ""
	for _, rec := range recs {
		header := fmt.Sprintf("{ \"index\" : { \"_id\" : \"%s\" } }", rec.Identify())
		data, err := json.Marshal(rec)
		if err != nil {
			mlog.Fatalln("Error marshaling issue with ID: %d", rec.Identify())
			continue
		}
		bulk = fmt.Sprintf("%s%s\n%s\n", bulk, header, data)
	}
	mlog.Printf("JSON:\n%s", bulk)
	return writeBulk(bulk, indx, typ)
}

// WriteJSON write any identifyable object to Elasticsearch.
// Assumes that object can be marshalled to JSON
func WriteJSON(recs map[string]string, indx, typ string) (*http.Response, error) {
	bulk := ""
	for id, json := range recs {
		header := fmt.Sprintf("{ \"index\" : { \"_id\" : \"%s\" } }", id)
		bulk = fmt.Sprintf("%s%s\n%s\n", bulk, header, json)
	}
	mlog.Printf("JSON:\n%s", bulk)
	return writeBulk(bulk, indx, typ)
}
