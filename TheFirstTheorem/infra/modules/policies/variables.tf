variable "lambda_role_name" {
  description = "The name for the lambda IAM role."
  type        = string
}

variable "lambda_s3_read_policy_name" {
  description = "The name for the lambda S3 read IAM policy."
  type        = string
}

variable "s3_bucket_arn" {
  description = "The ARN of the S3 bucket that Lambda needs to access."
  type        = string
}
