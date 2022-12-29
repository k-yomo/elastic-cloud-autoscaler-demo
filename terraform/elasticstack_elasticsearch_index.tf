resource "elasticstack_elasticsearch_index" "products" {
  name = "products"

  // The products index size is about 2G, so 2 shards won't be needed
  // but we set 2 just for demo purpose
  number_of_shards   = 2
  number_of_replicas = 0
  refresh_interval   = "-1"

  analysis_analyzer = jsonencode({
    "simple_analyzer": {
      "tokenizer": "lowercase",
      "filter": []
    }
  })

  mappings = jsonencode({
    "dynamic": "true",
    "dynamic_templates": [
      {
        "all_text": {
          "match_mapping_type": "string",
          "mapping": {
            "copy_to": "_all",
            "type": "text"
          }
        }
      }
    ]
  })
}
