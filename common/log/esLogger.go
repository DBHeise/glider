package log

import (
	"context"
	stdlog "log"

	"github.com/olivere/elastic"
)

var (
	esClient  *elastic.Client
	esContext context.Context
	esIndex   string
	esType    string
)

// ESLog - Log data to ES
func ESLog(data map[string]interface{}) {
	go func(blob map[string]interface{}) {
		logBlob(blob)
	}(data)
}

// InitESLogger - setup the ElasticSearch logger
func InitESLogger(esurl string, esindex string, estype string) {

	esContext = context.Background()
	esIndex = esindex
	esType = estype

	var err error
	esClient, err = elastic.NewClient(elastic.SetURL(esurl))
	if err != nil {
		stdlog.Fatal(err)
	}

}

func logBlob(blob map[string]interface{}) {
	if esClient != nil {
		go func() {
			_, err := esClient.Index().
				Index(esIndex).
				Type(esType).
				BodyJson(blob).
				Do(esContext)
			if err != nil {
				stdlog.Printf("Error saving to ElasticSearch: %s", err)
			}
		}()
	}

}
