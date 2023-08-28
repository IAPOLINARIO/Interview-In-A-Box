# REST API with S3 Backend via Terraform

This project provides a simple REST API written in Python that returns JSON objects from an S3 bucket when specific endpoints are called.

## Infrastructure Architecture and Components

This infrastructure is designed with a focus on security, scalability, and maintainability. Here's a breakdown of the architecture and how the components interact:

### S3 Bucket

- **Purpose:** Stores two JSON files (`foo.json` and `bar.json`).
- **Ownership Control:** The bucket is set to have the bucket owner preferred, ensuring that the AWS account that creates the resource becomes the owner.

### Virtual Private Cloud (VPC)

- **CIDR Block:** `10.0.0.0/16`
- **Purpose:** Provides an isolated environment to deploy resources and run the application securely.

### Subnets

- **Private Subnets:** These subnets do not have direct access to the public internet. Lambda functions reside here to maintain a secure environment.
- **Public Subnets:** Have direct access to the public internet via the Internet Gateway. Resources like the ALB (Application Load Balancer) and NAT Gateway are deployed here.

## Internet Gateway & Route Tables

- Ensures that resources in the VPC can communicate with the public internet. It's attached to the VPC and directs traffic from the VPC to the internet and vice-versa.

## NAT Gateway

- Used to allow resources in the private subnet, like the Lambda function, to initiate outbound connections to the internet for updates or to reach other AWS services.

## Security Groups

- **Lambda Security Group:** Allows the Lambda function to initiate outbound connections.
- **ALB Security Group:** Allows inbound HTTP traffic on port 80 and all outbound traffic.

## Application Load Balancer (ALB)

- Serves as the entry point to the application. It distributes incoming traffic across multiple targets.
- **Target Group:** Points to the Lambda function, ensuring that the ALB can route requests to it. It is set to use the HTTP protocol on port 80.

## AWS Lambda

- **Handler:** `lambda.lambda_handler`
- **Runtime:** Python 3.7
- **Purpose:** When triggered, it fetches the required JSON file from the S3 bucket and returns its content. This function runs within the private subnet but can access the internet through the NAT Gateway. The VPC configuration ensures this function is not directly accessible from the public internet.

## IAM Role & Policies

- **Role:** `lambda_role`
- A role with permissions to access other AWS services on behalf of the Lambda function.
- Policies attached:
  - Allow the Lambda function to create, read, and manage network interfaces within VPC.
  - Read objects from the S3 bucket.
  - Create and manage logs in CloudWatch.

## ALB Listener & Rule

- The listener checks for incoming requests on port 80 and has a rule to forward requests to the target group, which routes it to the Lambda function. This means that when someone accesses the ALB on `/health` or any `/api/*` path, it triggers our Lambda function, which then fetches the corresponding JSON file from the S3 bucket.

## Integration & Communication Flow:

1. A user sends a request to the ALB URL (e.g., `http://alb-url/api/foo`).
2. The ALB receives the request and forwards it to the target group, which triggers the Lambda function.
3. The Lambda function processes the request, fetching the necessary JSON file from the S3 bucket.
4. The response, containing the JSON file content, is sent back to the user through the ALB.

By separating the application and infrastructure into specific resources, we ensure scalability, easy debugging, and maintenance. Each component can be modified, scaled, or replaced independently without affecting the entire system.

## Benefits of this Solution

1. **Lambda Backend**:

   - **Efficiency**: AWS Lambda functions provide a serverless backend. This means no infrastructure management, which reduces complexity and costs.
   - **Scalability**: Lambda scales automatically by running code in response to each trigger. Your triggers can be an uploaded JSON to an S3 bucket, an HTTP request via API Gateway, or even a modification in a DynamoDB table.

2. **Using ALB instead of API Gateway**:

   - **Flexibility**: ALBs provide more flexibility in terms of routing traffic based on HTTP headers, methods, and paths.
   - **Cost-Effective**: Depending on the traffic pattern, using an ALB can be more cost-effective than using API Gateway, especially for high-traffic applications.
   - **Integrated with ECS**: If in the future you decide to containerize applications, ALB provides native support for routing traffic to ECS tasks.

3. **Modular Terraform Code**:

   - **Maintainability**: Splitting Terraform configurations into modules makes it easier to maintain and understand.
   - **Reusability**: Modules can be reused across different environments or AWS accounts.
   - **Isolation**: Changes can be made to a module without affecting others.

4. **Why Terraform**:

   - **Declarative Syntax**: Define and provide data center infrastructure using a declarative configuration language.
   - **Infrastructure as Code (IaC)**: This means that the infrastructure is represented as code, which can be versioned and audited.
   - **Provider Ecosystem**: Terraform can be used with multiple cloud providers, allowing for a multi-cloud strategy if needed.
   - **State Management**: Terraform creates a state file to keep track of your infrastructure's actual state.

5. **Private Networking**:

   - **Security**: By deploying the backend into a private network (like a VPC), you ensure that it is isolated from the public internet, mitigating a wide range of potential security threats.

6. **Publicly Accessible Endpoint**:

   - **Accessibility**: Even though our backend is in a private network, it is accessible via a public endpoint, ensuring that clients can still reach our API without compromising on security.

## Development Task

The primary task was to create a simple REST API that fetches data from an S3 bucket. AWS Lambda was chosen as the backend for this task.

Data from the S3 bucket is served when specific API endpoints are called. For instance, a GET request to `/api/foo` might return a JSON object like:

```json
{
  "greeting": "I am the Foo"
}
```

## How to Deploy

- Ensure you have Terraform installed.
- Navigate to the project directory.
- Initialize the Terraform project:

```
terraform init
```

- Apply the Terraform configurations:

```
terraform apply
```

## Future Work

As with any project, there's always room for improvements. Here are some considerations for the future:

- Integrate a CI/CD pipeline for seamless deployments.
- Add monitoring and logging for better observability.
- Implement more advanced security measures like VPC endpoints for S3, tighter security group rules, etc.

## Conclusion

Considering the constraints of having only 8 hours to build this project from scratch, the outcome closely aligns with best practices in infrastructure development and cloud-native architecture. The design prioritizes scalability, maintainability, and security, showcasing efficient time management and focused engineering.
