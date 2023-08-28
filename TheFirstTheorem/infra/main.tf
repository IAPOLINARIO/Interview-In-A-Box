terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

data "aws_availability_zones" "available" {}

provider "aws" {
  region = var.aws_region
}

module "storage" {
  source            = "./modules/storage"
  bucket_name       = var.bucket_name
  environment       = var.environment
  enable_versioning = var.enable_versioning
}

resource "aws_s3_object" "foo_json" {
  bucket = module.storage.s3_bucket
  key    = "foo.json"
  content = jsonencode({
    greeting = "I am the Foo"
  })
}

resource "aws_s3_object" "bar_json" {
  bucket = module.storage.s3_bucket
  key    = "bar.json"
  content = jsonencode({
    greeting = "I am the Bar"
  })
}

module "networking" {
  source             = "./modules/networking"
  cidr_block         = var.cidr_block
  availability_zones = data.aws_availability_zones.available.names
}

module "policies" {
  source                     = "./modules/policies"
  lambda_role_name           = var.lambda_role_name
  lambda_s3_read_policy_name = var.lambda_s3_read_policy_name
  s3_bucket_arn              = module.storage.s3_bucket_arn
}

module "lambda" {
  source             = "./modules/lambda"
  function_name      = var.function_name
  handler            = var.handler
  lambda_role_arn    = module.policies.lambda_iam_role_arn
  filename           = var.lambda_filename
  runtime            = var.runtime
  private_subnet_ids = module.networking.private_subnet_ids
  lambda_sg_id       = module.networking.lambda_sg_id
  target_group_arn   = module.networking.target_group_arn
  listener_arn       = module.networking.listener_arn
  path_patterns      = var.path_patterns
}
