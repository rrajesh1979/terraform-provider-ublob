terraform {
  required_providers {
    ublob = {
      version = "0.0.2"
      source  = "hashicorp.com/rrajesh1979/ublob"
    }
  }
}

data "ublob_meta" "all" {}

output "ublob_meta" {
  value = data.ublob_meta.all
}