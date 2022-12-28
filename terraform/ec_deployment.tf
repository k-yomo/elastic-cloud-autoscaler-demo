resource "ec_deployment" "elastic_cloud_autoscaler_demo" {
  name = "elastic-cloud-autoscaler-demo"

  region                 = "gcp-asia-northeast1"
  version                = "8.5.3"
  deployment_template_id = "gcp-cpu-optimized-v3"

  elasticsearch {
    topology {
      id         = "hot_content"
      size       = "1g"
      zone_count = 2
    }
  }

  kibana {
    topology {
      size = "1g"
    }
  }

  observability {
    deployment_id = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.id
  }

  lifecycle {
    // Elastic Cloud Autoscaler will manage topology size
    ignore_changes = [elasticsearch.0.topology.0.size]
  }
}

resource "ec_deployment" "elastic_cloud_autoscaler_demo_monitoring" {
  name = "elastic-cloud-autoscaler-demo-monitoring"

  region                 = "gcp-asia-northeast1"
  version                = "8.5.3"
  deployment_template_id = "gcp-cpu-optimized-v3"

  elasticsearch {
    topology {
      id         = "hot_content"
      size       = "1g"
      zone_count = 1
    }
  }

  kibana {
    topology {
      size = "1g"
    }
  }
}
