output "vpc_id" {
  value       = aws_vpc.iapolinario_vpc.id
  description = "ID of the created VPC"
}

output "private_subnet_ids" {
  value       = aws_subnet.private_subnet[*].id
  description = "IDs of the private subnets"
}

output "public_subnet_ids" {
  value       = aws_subnet.public_subnet[*].id
  description = "IDs of the public subnets"
}

output "alb_id" {
  value       = aws_lb.api_alb.id
  description = "ID of the Application Load Balancer"
}

output "lambda_sg_id" {
  value       = aws_security_group.lambda_sg.id
  description = "The ID of the lambda security group"
}

output "alb_dns_name" {
  description = "The DNS name of the Application Load Balancer"
  value       = aws_lb.api_alb.dns_name
}

output "target_group_arn" {
  value       = aws_lb_target_group.api_tg.arn
  description = "The ARN of the target group for ALB"
}

output "listener_arn" {
  value       = aws_lb_listener.api_listener.arn
  description = "The ARN of the listener for ALB"
}

output "alb_sg_id" {
  value       = aws_security_group.alb_sg.id
  description = "ID of the security group for ALB"
}

