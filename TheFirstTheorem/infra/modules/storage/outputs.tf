output "s3_bucket_arn" {
  value       = aws_s3_bucket.data_storage.arn
  description = "ARN of the S3 bucket"
}

output "s3_bucket_id" {
  value       = aws_s3_bucket.data_storage.id
  description = "ID of the S3 bucket"
}

output "s3_bucket" {
  value       = aws_s3_bucket.data_storage.bucket
  description = "ID of the S3 bucket"
}
