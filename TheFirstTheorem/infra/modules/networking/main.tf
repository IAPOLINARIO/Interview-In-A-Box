resource "aws_vpc" "iapolinario_vpc" {
  cidr_block = var.cidr_block
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.iapolinario_vpc.id
}

resource "aws_route_table" "main" {
  vpc_id = aws_vpc.iapolinario_vpc.id

  route {
    cidr_block = var.main_route_cidr
    gateway_id = aws_internet_gateway.igw.id
  }
}

resource "aws_main_route_table_association" "main" {
  vpc_id         = aws_vpc.iapolinario_vpc.id
  route_table_id = aws_route_table.main.id
}

resource "aws_subnet" "private_subnet" {
  count                   = length(var.availability_zones)
  vpc_id                  = aws_vpc.iapolinario_vpc.id
  cidr_block              = "10.0.${count.index}.0/24"
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = false
}

resource "aws_subnet" "public_subnet" {
  count                   = length(var.availability_zones)
  vpc_id                  = aws_vpc.iapolinario_vpc.id
  cidr_block              = "10.0.${count.index + length(var.availability_zones)}.0/24"
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = true
}

resource "aws_eip" "nat_eip" {}

resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.nat_eip.id
  subnet_id     = aws_subnet.public_subnet[0].id
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.iapolinario_vpc.id

  route {
    cidr_block     = var.main_route_cidr
    nat_gateway_id = aws_nat_gateway.nat_gateway.id
  }
}

resource "aws_route_table_association" "private_subnet_association" {
  count          = length(aws_subnet.private_subnet)
  subnet_id      = aws_subnet.private_subnet[count.index].id
  route_table_id = aws_route_table.private.id
}

resource "aws_security_group" "alb_sg" {
  vpc_id = aws_vpc.iapolinario_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lb" "api_alb" {
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = aws_subnet.public_subnet[*].id
}

resource "aws_security_group" "lambda_sg" {
  vpc_id = aws_vpc.iapolinario_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

/* resource "aws_lb_target_group" "api_tg" {
  name     = "api-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.iapolinario_vpc.id
}
 */
resource "aws_lb_target_group" "api_tg" {
  vpc_id      = aws_vpc.iapolinario_vpc.id
  protocol    = "HTTP"
  port        = 80
  target_type = "lambda"
}

resource "aws_lb_listener" "api_listener" {
  load_balancer_arn = aws_lb.api_alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.api_tg.arn
  }
}
