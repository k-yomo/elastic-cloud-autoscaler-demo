resource "elasticstack_elasticsearch_index" "elastic_cloud_autoscaler_demo" {
  name = "elastic-cloud-autoscaler-demo"

  number_of_shards   = 2
  number_of_replicas = 1
  refresh_interval   = "10s"

  mappings = jsonencode({
    dynamic = "true"
  })
}
