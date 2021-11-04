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

#output "myblob" {
#  value = module.myblob.ublob_meta
#}
#
#resource "ublob_blob" "aws_blob" {
#  bucket = "rrajesh1979-007"
#  cloud = "AWS"
#  region = "us-east-2"
#}
#
#output "ublob_aws_out" {
#  value = ublob_blob.aws_blob
#}
#
#resource "ublob_blob" "gcp_blob" {
#  bucket = "rrajesh1979-001"
#  cloud = "GCP"
#  region = "asia"
#  project_id = "peer-poc"
#  storage_class = "Standard"
#}
#
#output "ublob_gcp_out" {
#  value = ublob_blob.gcp_blob
#}
#
#variable "storage_account_key" {
#  type = string
#}
#
#resource "ublob_blob" "az_blob" {
#  bucket = "tf-blob-2"
#  cloud = "AZURE"
#  region = "us"
#  storage_account = "rrajesh1979"
#  storage_account_key = var.storage_account_key
#}
#
#output "ublob_az_out" {
#  value = ublob_blob.az_blob
#}