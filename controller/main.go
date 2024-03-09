package controller

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	ElasticClient *elasticsearch.Client
)

func SetupElasticsearch(es *elasticsearch.Client) {
	ElasticClient = es
}
