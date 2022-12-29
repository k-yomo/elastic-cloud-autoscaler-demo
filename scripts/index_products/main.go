package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/refresh"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
	log.Println("indexed products to elasticsearch")
}

func realMain() error {
	ctx := context.Background()

	elasticsearchCloudID := os.Getenv("ELASTICSEARCH_CLOUD_ID")
	elasticsearchUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	elasticsearchPassword := os.Getenv("ELASTICSEARCH_PASSWORD")

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		CloudID:  elasticsearchCloudID,
		Username: elasticsearchUsername,
		Password: elasticsearchPassword,
	})
	if err != nil {
		return err
	}
	productsBulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        esClient,
		Index:         "products",
		FlushInterval: 1 * time.Second,
		OnError: func(ctx context.Context, err error) {
			log.Println(err)
		},
	})
	if err != nil {
		return err
	}
	defer func() {
		productsBulkIndexer.Close(ctx)
		log.Println(fmt.Sprintf("%+v", productsBulkIndexer.Stats()))
	}()

	productsCSV, err := os.Open("products.csv")
	if err != nil {
		return err
	}

	if err := indexProducts(ctx, productsBulkIndexer, csv.NewReader(productsCSV)); err != nil {
		return err
	}
	if isSuccess, err := refresh.New(esClient).Index("products").IsSuccess(ctx); err != nil {
		return err
	} else if !isSuccess {
		return errors.New("refresh products index failed")
	}
	return nil
}

func indexProducts(ctx context.Context, productsBulkIndexer esutil.BulkIndexer, csvReader *csv.Reader) error {
	header, err := csvReader.Read()
	if err != nil {
		return err
	}

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		doc := map[string]interface{}{}
		for i, v := range row {
			// row num
			if i == 0 {
				continue
			}
			doc[header[i]] = v
		}
		docJSON, err := json.Marshal(doc)
		if err != nil {
			return err
		}
		err = productsBulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: doc["product_id"].(string),
			Body:       bytes.NewReader(docJSON),
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
				fmt.Println(item2.Error)
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
