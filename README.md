# UBLOB - Universal blob storage

<p>
<img alt="GitHub" src="https://img.shields.io/github/license/rrajesh1979/terraform-provider-ublob">
<img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/rrajesh1979/terraform-provider-ublob">
<a href="https://circleci.com/gh/rrajesh1979/terraform-provider-ublob/tree/master"><img src="https://circleci.com/gh/rrajesh1979/terraform-provider-ublob/tree/master.svg?style=svg"></a>
<a href="https://goreportcard.com/report/github.com/rrajesh1979/terraform-provider-ublob"><img src="https://goreportcard.com/badge/github.com/rrajesh1979/terraform-provider-ublob" alt="Go Report Card"></a>
<a href="https://codecov.io/github/rrajesh1979/terraform-provider-ublob"><img src="https://codecov.io/github/rrajesh1979/terraform-provider-ublob/branch/master/graph/badge.svg?token=ER2FNUMIUV" alt="Codecov branch"></a>
<a href="https://codeclimate.com/github/rrajesh1979/terraform-provider-ublob/maintainability"><img src="https://api.codeclimate.com/v1/badges/186b72a6bed912c8a8ba/maintainability" /></a>
<a href="https://codeclimate.com/github/rrajesh1979/terraform-provider-ublob/test_coverage"><img src="https://api.codeclimate.com/v1/badges/186b72a6bed912c8a8ba/test_coverage" /></a>
<a href="https://img.shields.io/github/contributors/rrajesh1979/terraform-provider-ublob"><img alt="GitHub commits" src="https://img.shields.io/github/contributors/rrajesh1979/terraform-provider-ublob"></a>
<a href="https://img.shields.io/github/commit-activity/w/rrajesh1979/terraform-provider-ublob"><img alt="GitHub contributors" src="https://img.shields.io/github/commit-activity/w/rrajesh1979/terraform-provider-ublob"></a>
</p>


<h3 align="center">Experimental Terraform provider (custom) </h3>
<p align="center">One provider to rule them all! Create an AWS S3 bucket, Azure Storage container, GCP cloud storage bucket all using this provider.</p>

Menu
----

- [Pre-requisites](#pre-reqs)
- [Examples](#examples)
- [Future enhancements](#future)

Pre-requisites
----


Examples
----
###AWS
```terraform
resource "ublob_blob" "aws_blob" {
  bucket = "rrajesh1979-007"
  cloud = "AWS"
  region = "us-east-2"
}

output "ublob_aws_out" {
  value = ublob_blob.aws_blob
}
```

###GCP
```terraform
resource "ublob_blob" "gcp_blob" {
  bucket = "rrajesh1979-001"
  cloud = "GCP"
  region = "asia"
  project_id = "peer-poc"
  storage_class = "Standard"
}

output "ublob_gcp_out" {
  value = ublob_blob.gcp_blob
}
```

###AZURE
```terraform
variable "storage_account_key" {
  type = string
}

resource "ublob_blob" "az_blob" {
  bucket = "tf-blob-2"
  cloud = "AZURE"
  region = "us"
  storage_account = "rrajesh1979"
  storage_account_key = var.storage_account_key
}

output "ublob_az_out" {
  value = ublob_blob.az_blob
}
```

Future enhancements
----
