terraform {
  required_providers {
    ec = {
      source  = "elastic/ec"
      version = "0.5.0"
    }

    elasticstack = {
      source  = "elastic/elasticstack"
      version = "0.5.0"
    }
  }
  required_version = "= 1.3.6"
}
