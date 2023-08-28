variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
}

variable "bucket_acl" {
  description = "Access control list setting for the bucket. (private, public-read, etc.)"
  type        = string
  default     = "private"
}

variable "environment" {
  description = "Environment tag (e.g., dev, prod, staging)"
  type        = string
  default     = "dev"
}

variable "enable_versioning" {
  description = "Boolean to determine if versioning should be enabled on the bucket"
  type        = bool
  default     = true
}
