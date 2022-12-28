output "demo_deployment_id" {
  value = ec_deployment.elastic_cloud_autoscaler_demo.id
}
output "demo_elasticsearch_cloud_id" {
  value = ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch.0.cloud_id
}
output "demo_elasticsearch_username" {
  value = ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch_username
}
output "demo_elasticsearch_password" {
  sensitive = true
  value = ec_deployment.elastic_cloud_autoscaler_demo.elasticsearch_password
}
output "demo_monitoring_elasticsearch_cloud_id" {
  value = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch.0.cloud_id
}
output "demo_monitoring_elasticsearch_username" {
  value = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch_username
}
output "demo_monitoring_elasticsearch_password" {
  sensitive = true
  value = ec_deployment.elastic_cloud_autoscaler_demo_monitoring.elasticsearch_password
}
