resource "aws_lambda_function" "api_function" {
  function_name = var.function_name
  handler       = var.handler
  role          = var.lambda_role_arn
  filename      = var.filename
  runtime       = var.runtime

  vpc_config {
    subnet_ids         = var.private_subnet_ids
    security_group_ids = [var.lambda_sg_id]
  }
}

resource "aws_lb_target_group_attachment" "lambda_attachment" {
  target_group_arn = var.target_group_arn
  target_id        = aws_lambda_function.api_function.arn

  depends_on = [aws_lambda_function.api_function]
}



resource "aws_lambda_permission" "allow_alb" {
  statement_id  = "AllowExecutionFromALB"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api_function.arn
  principal     = "elasticloadbalancing.amazonaws.com"
}

resource "aws_lb_listener_rule" "api_rule" {
  listener_arn = var.listener_arn

  action {
    type             = "forward"
    target_group_arn = var.target_group_arn
  }

  condition {
    path_pattern {
      values = var.path_patterns
    }
  }
}
