package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/k-yomo/elastic-cloud-autoscaler/autoscaler"
	"github.com/k-yomo/elastic-cloud-autoscaler/metrics"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	deploymentID := os.Getenv("DEPLOYMENT_ID")
	elasticCloudAPIKey := os.Getenv("EC_API_KEY")

	elasticsearchCloudID := os.Getenv("ELASTICSEARCH_CLOUD_ID")
	elasticsearchUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	elasticsearchPassword := os.Getenv("ELASTICSEARCH_PASSWORD")

	monitoringElasticsearchCloudID := os.Getenv("MONITORING_ELASTICSEARCH_CLOUD_ID")
	monitoringElasticsearchUsername := os.Getenv("MONITORING_ELASTICSEARCH_USERNAME")
	monitoringElasticsearchPassword := os.Getenv("MONITORING_ELASTICSEARCH_PASSWORD")

	ecClient, err := api.NewAPI(api.Config{
		Client:     http.DefaultClient,
		AuthWriter: auth.APIKey(elasticCloudAPIKey),
	})
	if err != nil {
		return err
	}

	esClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID:  elasticsearchCloudID,
		Username: elasticsearchUsername,
		Password: elasticsearchPassword,
	})
	if err != nil {
		return err
	}
	monitoringESClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID:  monitoringElasticsearchCloudID,
		Username: monitoringElasticsearchUsername,
		Password: monitoringElasticsearchPassword,
	})
	if err != nil {
		return err
	}

	esAutoScaler, err := autoscaler.New(&autoscaler.Config{
		ElasticCloudClient:  ecClient,
		DeploymentID:        deploymentID,
		ElasticsearchClient: esClient,
		Scaling: autoscaler.ScalingConfig{
			DefaultMinSizeMemoryGB: int(autoscaler.SixtyFourGiBNodeNumToTopologySize(1)),
			DefaultMaxSizeMemoryGB: int(autoscaler.SixtyFourGiBNodeNumToTopologySize(4)),
			AutoScaling: &autoscaler.AutoScalingConfig{
				MetricsProvider:       metrics.NewMonitoringElasticsearchMetricsProvider(monitoringESClient),
				DesiredCPUUtilPercent: 50,
			},
			ScheduledScalings: []*autoscaler.ScheduledScalingConfig{
				{
					StartCronSchedule: "TZ=UTC 30 * * * *",
					Duration:          30 * time.Minute,
					MinSizeMemoryGB:   int(autoscaler.SixtyFourGiBNodeNumToTopologySize(2)),
					MaxSizeMemoryGB:   int(autoscaler.SixtyFourGiBNodeNumToTopologySize(4)),
				},
			},
			Index:         "elastic-cloud-autoscaler-demo",
			ShardsPerNode: 2,
		},
	})
	if err != nil {
		return err
	}
	for {
		select {
		case <-time.After(1 * time.Minute):
			log.Println("[START] elastic cloud autoscaler")
			scalingOperation, err := esAutoScaler.Run(context.Background())
			if err != nil {
				return err
			}
			if scalingOperation.Direction() != autoscaler.ScalingDirectionNone {
				fmt.Println("==============================================")
				fmt.Println("scaling direction:", scalingOperation.Direction())
				fmt.Println(fmt.Sprintf("topology size updated from: %d => to %d", *scalingOperation.FromTopologySize.Value, *scalingOperation.ToTopologySize.Value))
				fmt.Println(fmt.Sprintf("replica num updated from: %d => to %d", scalingOperation.FromReplicaNum, scalingOperation.ToReplicaNum))
				fmt.Println("reason:", scalingOperation.Reason)
				fmt.Println("==============================================")
			} else {
				fmt.Println("not scaling:", scalingOperation.Reason)
			}
			log.Println("[END] elastic cloud autoscaler")
		}
	}
}
