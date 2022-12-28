
provider "ec" {
}

provider "elasticstack" {
  elasticsearch {
    username  = ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch_username
    password  = ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch_password
    endpoints = [ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch.0.https_endpoint]
  }
}

provider "elasticstack" {
  alias = "monitoring"
  elasticsearch {
    username  = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch_username
    password  = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch_password
    endpoints = [ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch.0.https_endpoint]
  }
}
