variable "cidr_block" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "main_route_cidr" {
  description = "CIDR block for the main route"
  type        = string
  default     = "0.0.0.0/0"
}

variable "availability_zones" {
  description = "List of availability zones"
  type        = list(string)
}
/* 
variable "alb_security_group_id" {
  description = "Security group ID for the ALB"
  type        = string
}
 */
