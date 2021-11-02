terraform {
  required_providers {
    ublob = {
      version = "0.0.2"
      source  = "hashicorp.com/rrajesh1979/ublob"
    }
  }
}

provider "ublob" {}

module "myblob" {
  source = "./ublob"
}

output "myblob" {
  value = module.myblob.ublob_meta
}