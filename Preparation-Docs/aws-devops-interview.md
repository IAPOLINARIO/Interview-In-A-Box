# Interview Questions for Technical Round

## Q1: What is OSI Model and can you explain the layers?

**Answer**:  
The OSI Model (Open Systems Interconnection) is a framework used for understanding how different networking protocols interact across networks. The seven layers are:

1. Physical Layer - Deals with physical medium and transmission of raw bits.
2. Data Link Layer - Responsible for error detection and framing.
3. Network Layer - Routing and forwarding packets.
4. Transport Layer - Ensures data integrity and establishes, maintains, and terminates connections.
5. Session Layer - Manages sessions between applications.
6. Presentation Layer - Transforms data into a format that applications can understand.
7. Application Layer - Interface between the networking stack and the application.

## Q2: What are the differences between EC2 and Lambda?

**Answer**:  
EC2 is a virtual machine in the cloud, whereas Lambda is a serverless compute service. EC2 instances are long-lived, and you pay for uptime. Lambdas are ephemeral and you pay per execution time.

### Q3: What is the difference between Docker Image and Docker Container?

**Answer**:  
A Docker Image is a template that contains the application, libraries, and dependencies. A Docker Container is a running instance of an image.

### Q4: What is Docker Compose and what is it used for?

**Answer**:  
Docker Compose is a tool for defining and running multi-container Docker applications. You can define your application stack in a `docker-compose.yml` file and spin up all services with a single command (`docker-compose up`).

### Q5: How do you define Infrastructure as Code (IaC)?

**Answer**:  
Infrastructure as Code is the practice of managing and provisioning infrastructure through code rather than manual setup.

### Q6: What is a Terraform Provider?

**Answer**:  
A Terraform Provider is a plugin that allows Terraform to manage resources in a particular API, like AWS, Azure, or GCP.

### Q7: What is a VPC in AWS and how is it different from a Subnet?

**Answer**:  
A VPC (Virtual Private Cloud) is a logically isolated section of the AWS Cloud, whereas a Subnet is a range of IPs within a VPC.

### Q8: How does DNS work?

**Answer**:  
Domain Name System (DNS) translates human-readable domain names into IP addresses. It uses a hierarchical lookup system to resolve these names.

### Q9: What is the purpose of Load Balancers?

**Answer**:  
Load Balancers distribute incoming network traffic across multiple servers to ensure high availability and reliability.

### Q10: Explain Idempotence and how it is relevant to RESTful APIs?

**Answer**:  
Idempotence means that an operation can be repeated multiple times and yield the same result. In RESTful APIs, idempotent methods like GET, PUT, and DELETE can be called many times without different outcomes, making it easier to reason about the system.

### Q11: Can you explain the TCP Three-Way Handshake?

**Answer**:  
The TCP Three-Way Handshake is a process used to establish a connection between a client and a server. The steps are:

1. **SYN**: Client sends a TCP packet with the SYN flag set to initiate a connection.
2. **SYN-ACK**: Server responds with a packet with both SYN and ACK flags set, acknowledging the client's request.
3. **ACK**: Finally, the client sends an ACK packet back to the server, and the connection is established.

### Q12: What is CIDR notation and how is it used in networking?

**Answer**:  
CIDR (Classless Inter-Domain Routing) notation is a compact way of specifying IP addresses and subnet masks. It's written as `IP_address/subnet_mask`, like `192.168.1.1/24`. CIDR makes it easier to allocate IP ranges and route traffic.

### Q13: What are some common HTTP status codes and what do they mean?

**Answer**:

- `200 OK`: Successful GET or POST request
- `201 Created`: Resource successfully created
- `400 Bad Request`: Client error
- `401 Unauthorized`: Missing or invalid authentication
- `404 Not Found`: Resource doesn't exist
- `500 Internal Server Error`: Server-side error

### Q14: Explain the difference between TCP and UDP.

**Answer**:  
TCP (Transmission Control Protocol) is connection-oriented and reliable, ensuring data integrity and order. UDP (User Datagram Protocol) is connectionless and faster but doesn't guarantee delivery or order.

### Q15: What is NAT and why is it used?

**Answer**:  
NAT (Network Address Translation) allows private IPs to connect to the Internet. It's used to map a local IP address to a public IP address, allowing multiple devices on a local network to share a single public IP.

### Q16: Describe Amazon S3 and its use cases.

**Answer**:  
Amazon S3 (Simple Storage Service) is a scalable object storage service. Use cases include backup, content distribution, and data archiving.

### Q17: What is an Elastic IP in AWS?

**Answer**:  
Elastic IPs are static, public IPv4 addresses that can be allocated to your AWS account. They're useful for avoiding downtime and are often used in High Availability configurations.

### Q18: Explain AWS RDS and its advantages.

**Answer**:  
AWS RDS (Relational Database Service) is a managed database service that supports multiple database engines like MySQL, PostgreSQL, SQL Server, etc. Advantages include automated backups, updates, and scaling.

### Q19: What are AWS Security Groups?

**Answer**:  
Security Groups act as a virtual firewall for instances in a VPC. They control inbound and outbound traffic based on rules that you define.

### Q20: What is Amazon CloudWatch used for?

**Answer**:  
Amazon CloudWatch is a monitoring service that provides data and insights to monitor applications, understand performance, and optimize resources.

### Q21: What's the difference between EBS and EFS in AWS?

**Answer**:  
EBS (Elastic Block Store) provides block-level storage and is used with EC2 instances. EFS (Elastic File System) is a scalable file storage service for use with multiple EC2 instances.

### Q22: How does AWS Auto Scaling work?

**Answer**:  
AWS Auto Scaling adjusts the number of running instances based on performance metrics or schedules. It helps maintain application availability and optimizes costs.

### Q23: What is AWS Lambda and what are its triggers?

**Answer**:  
AWS Lambda is a serverless computing service. Triggers can include S3 events, SNS notifications, CloudWatch Logs, etc.

### Q24: Explain the use of VPC Endpoints in AWS.

**Answer**:  
VPC Endpoints allow you to connect to AWS services without going through the public Internet, improving security and performance.

### Q25: What are IAM Roles and Policies in AWS?

**Answer**:  
IAM (Identity and Access Management) Roles are sets of permissions that grant access to actions and resources in AWS. Policies are objects that define permissions and can be attached to roles.

### Q26: What is AWS Route 53 and what are its main features?

**Answer**:  
AWS Route 53 is a scalable Domain Name System (DNS) web service designed for domain name registration, DNS routing, and health checking. Main features include latency-based routing, geo-proximity, and weighted round-robin.

### Q27: Can you explain the components of AWS ECS?

**Answer**:  
AWS ECS (Elastic Container Service) has several components:

- **Cluster**: A logical grouping of services and tasks.
- **Task Definitions**: Describe the Docker container and settings for an application.
- **Service**: Manages task definitions to ensure the desired count of tasks are running.
- **Tasks**: Instances of a task definition that run on a cluster.

### Q28: How does AWS EKS differ from ECS?

**Answer**:  
EKS (Elastic Kubernetes Service) is AWS's managed Kubernetes service. While ECS is native to AWS, EKS provides a Kubernetes-native experience. EKS integrates well with Kubernetes tools, but ECS may offer tighter integration with other AWS services.

### Q29: What is AWS API Gateway and what are its key features?

**Answer**:  
API Gateway is a service for creating, deploying, and managing APIs. Key features include rate limiting, caching, data transformation, and authentication.

### Q30: How does AWS Lambda handle scaling?

**Answer**:  
Lambda automatically scales function execution in response to incoming traffic. Each function can trigger independently, scaling precisely with the workload size, and AWS handles all the operational aspects.

### Q31: What is AWS ACM and how does it integrate with other AWS services?

**Answer**:  
AWS Certificate Manager (ACM) is a service that handles SSL/TLS certificates. It can be used to deploy certificates to AWS services like Load Balancers, CloudFront distributions, and API Gateway.

**Answer**:  
Aliases in Route 53 provide a Route 53-specific extension to DNS functionality, allowing you to map your domain to AWS resources like Elastic Load Balancers, S3 buckets, or CloudFront distributions without the need for a CNAME.

### Q33: What are launch types in ECS?

**Answer**:  
ECS supports two launch types: EC2 and Fargate. EC2 allows you to manage your clusters and tasks on EC2 instances. Fargate manages the instances for you, abstracting away a lot of operational overhead.

### Q34: What is VPC peering in the context of EKS?

**Answer**:  
VPC peering allows for networking connection between two VPCs in AWS. This is useful for EKS if you have services in different VPCs that need to communicate with each other.

### Q35: How do you secure APIs in API Gateway?

**Answer**:  
APIs in API Gateway can be secured using various methods like IAM roles, Lambda authorizers, or Cognito user pools. You can also implement API keys and OAuth 2.0 tokens for access control.

### Q36: What are TCP and UDP, and how do they differ?

**Answer**:  
TCP (Transmission Control Protocol) is connection-oriented and reliable. It establishes a dedicated connection and ensures packet ordering. UDP (User Datagram Protocol) is connectionless and does not guarantee reliability or ordering. TCP is generally used for data that needs reliable transport, while UDP is used for quick, "best-effort" type of services.

### Q37: How does HTTPS improve the security over HTTP?

**Answer**:  
HTTPS (HTTP Secure) employs SSL/TLS to encrypt the data between the client and server. This ensures data confidentiality and integrity. It also involves a certificate exchange for establishing the server's authenticity, thus providing an extra layer of security.

### Q38: Can you explain what DNS poisoning is?

**Answer**:  
DNS poisoning refers to the practice of inserting a rogue entry into a DNS cache, so that DNS queries return an incorrect IP address. This is usually done for malicious purposes such as phishing or man-in-the-middle attacks.

### Q39: What is a VPN and how does it work?

**Answer**:  
VPN (Virtual Private Network) creates a secure tunnel between your device and a VPN server, encrypting your data. Your IP address is also masked, improving both security and privacy.

### Q40: What is the ICMP protocol and what is it commonly used for?

**Answer**:  
ICMP (Internet Control Message Protocol) is used for error reporting and diagnostics in an IP network. It's commonly used for utilities like `ping` and `traceroute`.
