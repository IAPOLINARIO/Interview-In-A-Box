output "lambda_iam_role_arn" {
  value       = aws_iam_role.lambda_role.arn
  description = "ARN of the Lambda IAM Role"
}

output "lambda_s3_read_policy_arn" {
  value       = aws_iam_policy.lambda_s3_read_policy.arn
  description = "ARN of the Lambda S3 Read Policy"
}
