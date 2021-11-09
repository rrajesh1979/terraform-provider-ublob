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


<h3>Experimental Terraform provider (custom) </h3>
---
<p>One provider to rule them all! Create an AWS S3 bucket, Azure Storage container, GCP cloud storage bucket all using this provider.</p>

Menu
----

- [Pre-requisites](#pre-requisites)
- [Example configurations](#example-configurations)
- [Future enhancements](#future-enhancements)
- [Usage](#usage)

Gitpod 
----
- Environment with dependencies pre-installed and ready for you to develop
- Visual Studio Code with required extensions preloaded 

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/rrajesh1979/terraform-provider-ublob)


Pre-requisites to build your own environment
----
1. Terraform
2. Go SDK
3. AWS CLI
4. GCP CLI
5. AZURE CLI
```bash
brew install hashicorp/tap/terraform
brew install go
brew install awscli
brew install --cask google-cloud-sdk
brew update && brew install azure-cli
```

Example configurations
----
AWS
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

GCP
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

AZURE
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

Usage
----

Cloud provider specific configuration
- AWS
```bash
aws configure
```

- GCP
```bash
gcloud auth application-default login
```

- AZURE
```bash
az login
az account set --subscription ""
##Setup environment variables
export TF_VAR_storage_account_key="<<Your storage account key>>"
```

```bash
#usage
make install
cd examples
terraform init && terraform apply --auto-approve
```

Future enhancements
----
- This is an experimental Terraform provider that allows users to create a storage bucket in any of the cloud platforms - AWS, GCP and Azure
- This Alpha release is to build a working version of the provider.
- The current release is a basic version that only supports
  - Create resource
  - Destroy resource
- ###ToDo
  - Functional
  - [ ] Enhanced Read
  - [ ] Implement Update, Import
  - [ ] Support for fine-grained attributes for storage bucket - ex. permissions, versioning etc. 
  - Non-Functional
  - [ ] Add test cases
  - [ ] Code refactoring
  - [ ] Improve exception handling
  - [ ] Better logging and error handling