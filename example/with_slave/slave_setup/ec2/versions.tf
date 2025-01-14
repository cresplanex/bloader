terraform {
  required_version = ">= 1.3.2"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.72.0"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "4.0.6"
    }
    time = {
      source  = "hashicorp/time"
      version = "0.12.1"
    }
    http = {
      source  = "hashicorp/http"
      version = "3.4.5"
    }
  }
}
