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

resource "ublob_blob" "ublob" {
  bucket = "rrajesh1979-003"
#  cloud = "AWS"
#  region = "us-east-1"
}

output "unbob_out" {
  value = ublob_blob.ublob
}