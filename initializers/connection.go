package initializers

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

// create a new connection to elasticsearch
func NewConnection() *elasticsearch.Client {

	// Es config
	cfg := elasticsearch.Config{
		CloudID: os.Getenv("CLOUD_ID"),
		APIKey:  os.Getenv("API_KEY"),
	}

	// Connect to Elasticsearch
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
	return es
}
