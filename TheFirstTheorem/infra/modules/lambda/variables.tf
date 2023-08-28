variable "function_name" {}
variable "handler" {}
variable "lambda_role_arn" {}
variable "filename" {}
variable "runtime" {}
variable "private_subnet_ids" {
  type = list(string)
}
variable "lambda_sg_id" {}
variable "target_group_arn" {}
variable "listener_arn" {}
variable "path_patterns" {
  type = list(string)
}
