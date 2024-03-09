package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

// create a new connection to elasticsearch
func NewConnection() *elasticsearch.Client {

	// Es config
	cfg := elasticsearch.Config{
		// Addresses: []string{
		// 	os.Getenv("ELASTICSEARCH_URL"),
		// },
		// Username: os.Getenv("ELASTICSEARCH_USERNAME"),
		// Password: os.Getenv("ELASTICSEARCH_PASSWORD"),

		CloudID: os.Getenv("CLOUD_ID"),
		APIKey:  os.Getenv("API_KEY"),
	}

	// Connect to Elasticsearch
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
	fmt.Println("Connected to Elasticsearch!")
	return es
}
