
output "alb_dns_name" {
  description = "The DNS name of the Application Load Balancer"
  value       = module.networking.alb_dns_name
}

output "vpc_id" {
  value       = module.networking.vpc_id
  description = "ID of the VPC created"
}

output "s3_bucket_arn" {
  value       = module.storage.s3_bucket_arn
  description = "ARN of the created S3 bucket"
}

output "lambda_iam_role_arn" {
  value       = module.policies.lambda_iam_role_arn
  description = "ARN of the Lambda IAM Role"
}

output "lambda_s3_read_policy_arn" {
  value       = module.policies.lambda_s3_read_policy_arn
  description = "ARN of the Lambda S3 Read Policy"
}
