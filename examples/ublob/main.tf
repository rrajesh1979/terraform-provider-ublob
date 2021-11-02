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

#resource "ublob_blob" "ublob" {
#  bucket = "rrajesh1979-002"
##  cloud = "AWS"
##  region = "us-east-1"
#}
#
#output "ublob_out" {
#  value = ublob_blob.ublob
#}