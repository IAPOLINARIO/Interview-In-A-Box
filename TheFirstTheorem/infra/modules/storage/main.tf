resource "aws_s3_bucket" "data_storage" {
  bucket = var.bucket_name

  tags = {
    Name        = var.bucket_name
    Environment = var.environment
  }
}

resource "aws_s3_bucket_ownership_controls" "bucket_ownership" {
  bucket = aws_s3_bucket.data_storage.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}
