resource "aws_iam_role" "lambda_role" {
  name = var.lambda_role_name
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Effect = "Allow",
      },
    ]
  })
}

resource "aws_iam_policy" "lambda_s3_read_policy" {
  name        = var.lambda_s3_read_policy_name
  path        = "/"
  description = "IAM policy to allow Lambda to read from S3 bucket"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListObjects",
          "s3:ListObjectsV2",
          "s3:ListBucket"
        ],
        Resource = [
          var.s3_bucket_arn,
          "${var.s3_bucket_arn}/*"
        ],
        Effect = "Allow",
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_s3_read_policy_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_s3_read_policy.arn
}

resource "aws_iam_role_policy" "lambda_vpc_policy" {
  name = "lambda_vpc_policy"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface"
        ],
        Resource = "*",
        Effect   = "Allow",
      }
    ]
  })
}

resource "aws_iam_role_policy" "lambda_logging_policy" {
  name = "lambda_logging_policy"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "arn:aws:logs:*:*:*",
        Effect   = "Allow",
      }
    ]
  })
}
