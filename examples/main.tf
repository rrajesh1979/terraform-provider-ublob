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

#resource "ublob_blob" "ublob" {
#  bucket = "rrajesh1979-005"
#  cloud = "AWS"
#  region = "us-east-2"
#}
#
#output "ublob_out" {
#  value = ublob_blob.ublob
#}

#resource "ublob_blob" "ublob" {
#  bucket = "rrajesh1979-005"
#  cloud = "GCP"
#  region = "asia"
#  project_id = "peer-poc"
#  storage_class = "Standard"
#}
#
#output "ublob_out" {
#  value = ublob_blob.ublob
#}