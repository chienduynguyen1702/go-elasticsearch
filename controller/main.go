package controller

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	ElasticClient *elasticsearch.TypedClient
)

func SetupElasticsearch(es *elasticsearch.TypedClient) {
	ElasticClient = es
}

// func getDocumentIdByIndex(indexName string, modelId string) (string, error) {

// }
