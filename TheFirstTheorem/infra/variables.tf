variable "aws_region" {
  description = "The AWS region"
  default     = "us-west-2"
}

variable "bucket_name" {
  description = "Name of the S3 bucket"
  default     = "my-unique-bucket-name"
}

variable "environment" {
  description = "Environment name (e.g., production, staging)"
  default     = "production"
}

variable "enable_versioning" {
  description = "Flag to enable versioning on S3 bucket"
  default     = true
}

variable "function_name" {
  description = "Name of the Lambda function"
  default     = "api_function"
}

variable "handler" {
  description = "Handler for the Lambda function"
  default     = "lambda.lambda_handler"
}

variable "lambda_filename" {
  description = "Path to the Lambda function zip"
  default     = "./lambda.zip"
}

variable "runtime" {
  description = "Runtime for the Lambda function"
  default     = "python3.7"
}

variable "path_patterns" {
  description = "List of path patterns for the ALB listener rule"
  default     = ["/health", "/api/*"]
}

variable "cidr_block" {
  description = "CIDR block for the VPC"
  default     = "10.0.0.0/16"
}

variable "lambda_role_name" {
  description = "Name of the IAM role for Lambda"
  default     = "my-lambda-role"
}

variable "lambda_s3_read_policy_name" {
  description = "Name of the S3 read policy for Lambda"
  default     = "my-lambda-s3-read-policy"
}
