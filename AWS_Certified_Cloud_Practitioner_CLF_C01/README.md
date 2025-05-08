AWS Certified Cloud Practitioner Services
        |
        |__ Compute
              |__ Amazon Elastic Compute Cloud (EC2) - Provides resizable compute
              capacity in the cloud.
              |__ Amazon Elastic Container Service (ECS) - A fully-managed
              container orchestration service.
              |__ AWS Lambda - Runs code in response to events and automatically
              manages the compute resources.
              |__ Amazon Lightsail - Provides a simple way to get started with AWS
              for small-scale projects.
        |
        |__ Storage
        |       |__ Amazon Simple Storage Service (S3) - Object storage service that
                offers industry-leading scalability, data availability, security, and performance.
               |__ Amazon Elastic Block Store (EBS) - Provides block-level storage
                volumes for use with Amazon EC2 instances.
               |__ Amazon Elastic File System (EFS) - Provides scalable file
                storage for use with Amazon EC2 instances.
                |__ Amazon Glacier - A secure, durable, and low-cost storage service
                for data archiving and long-term backup.
        |
        |__ Database
                |__ Amazon Relational Database Service (RDS) - A managed database
                service that makes it easier to set up, operate, and scale a relational database in the cloud.
               |__ Amazon DynamoDB - A fast and flexible NoSQL database service for
                any scale.
                |__ Amazon ElastiCache - A web service that makes it easy to deploy,
                operate, and scale an in-memory cache in the cloud.
        |
        |__ Networking
                |__ Amazon Virtual Private Cloud (VPC) - A secure and isolated cloud
                 network that you can customize.
                |__ Amazon Route 53 - A highly scalable and available Domain Name
                System (DNS) web service.
                |__ Elastic Load Balancing - Automatically distributes incoming
                 application traffic across multiple targets.
                |__ Amazon API Gateway - A fully managed service that makes it easy
                for developers to create, publish, maintain, monitor, and secure APIs at any scale.
        |
        |__ Security, Identity & Compliance
                |__ AWS Identity and Access Management (IAM) - A web service that
                helps you securely control access to AWS resources.
                |__ AWS Certificate Manager (ACM) - A service that lets you easily
                provision, manage, and deploy Secure Sockets Layer/Transport Layer Security (SSL/TLS) certificates for use with AWS services.
                |__ AWS Organizations - A web service that enables you to
                consolidate multiple AWS accounts into an organization that you create and centrally manage.
                |__ AWS CloudTrail - A web service that records AWS API calls for
                your account and delivers log files to you.
                |__ Amazon GuardDuty - A threat detection service that continuously
                monitors for malicious activity and unauthorized behavior to protect your AWS accounts and workloads.
        |
        |__ Management & Governance
              |__ AWS CloudFormation - A service that helps you model and set up
                 your AWS resources so you can spend less time managing those resources and more time focusing on
                your applications that run in AWS.
               |__ Amazon CloudWatch - A monitoring service for AWS resources and
                  the applications you run on AWS.
               |__ AWS Systems Manager - A management service that helps you
                 automatically collect software inventory, apply OS patches, create system images, and configure
                 operating systems and applications at scale.
               |__ AWS Trusted Advisor - An online resource that helps you reduce
                 cost, increase performance,
                | |__ AWS Config - A fully managed service that provides you with an AWS resource inventory, configuration history, and configuration change notifications to enable security and governance.
                | |__ AWS CloudTrail - A web service that records AWS API calls for your account and delivers log files to you.
                | |__ Amazon CloudWatch Events - A service that enables you to respond to state changes in your AWS resources.
                |
                |__ AWS Service Catalog
                |__ AWS Resource Groups
        |__ Notification
                | |__ Amazon Simple Notification Service (SNS) - A highly available, durable, and scalable notification service that enables you to publish and receive messages from applications or services.
                | |__ Amazon Simple Queue Service (SQS) - A fully managed message queuing service that enables you to decouple and scale microservices, distributed systems, and serverless applications.
                | |__ Amazon Simple Workflow Service (SWF) - A fully managed workflow service for building scalable, resilient applications.
                |
        |__ Developer Tools
                | |__ AWS CodeCommit - A fully-managed source control service that makes it easy for companies to host secure and highly scalable private Git repositories.
                | |__ AWS CodePipeline - A fully managed continuous delivery service that helps you automate your release pipelines for fast and reliable application and infrastructure updates.
                | |__ AWS CodeBuild - A fully managed build service that compiles source code, runs tests, and produces software packages that are ready to deploy.
                | |__ AWS CodeDeploy - A fully managed deployment service that automates software deployments to a variety of compute services such as Amazon EC2, AWS Fargate, AWS Lambda, and your on-premises servers.
                | |__ AWS X-Ray - A service that helps developers analyze and debug distributed applications, such as those built using a microservices architecture.
                |
        |__ Analytics
                |__ Amazon Athena - An interactive query service that makes it easy to analyze data in Amazon S3 using standard SQL.
                |__ Amazon EMR - A managed Hadoop framework that makes it easy to process large amounts of data using open-source tools such as Apache Spark, Hadoop, and Hive.
                |__ Amazon CloudSearch - A managed search service that makes it easy to add search capabilities to your website or application.
                |__ Amazon Kinesis - A platform for streaming data on AWS, allowing you to ingest, process, and analyze real-time, streaming data.
                |__ Amazon QuickSight - A cloud-native business intelligence service that makes it easy to build visualizations, perform ad-hoc analysis, and quickly get business insights from your data.
                |__ Amazon Redshift - A fast, fully-managed, petabyte-scale data warehouse that makes it simple and cost-effective to analyze all your data using standard SQL and your existing Business Intelligence (BI) tools.
       |__ Artificial Intelligence
                |__ Amazon Rekognition - A fully managed service that makes it easy to add image and video analysis to your applications.
                |__ Amazon SageMaker - A fully-managed service that enables developers and data scientists to quickly and easily build, train, and deploy machine learning models at any scale.
                |__ Amazon Polly - A service that turns text into lifelike speech.
                |__ Amazon Lex - A service for building conversational interfaces into any application using voice and text.
                |__ Amazon Textract - A fully managed service that automatically extracts text and data from scanned documents.
                |__ Amazon Transcribe - An automatic speech recognition (ASR) service that makes it easy to add speech-to-text capability to your applications.


AWS Well-Architected Framework
        |
        |__ Operational Excellence
        |       |__ Focuses on the ability to run and monitor systems to deliver business value and to continually improve supporting processes and procedures.
        |
        |__ Security
        |       |__ Focuses on protecting information and systems.
        |
        |__ Reliability
        |       |__ Focuses on the ability of a system to recover from infrastructure or service disruptions, dynamically acquire computing resources to meet demand, and mitigate disruptions such as misconfigurations or transient network issues.
        |
        |__ Performance Efficiency
        |       |__ Focuses on using computing resources efficiently to meet system requirements, and maintaining that efficiency as demand changes and technologies evolve.
        |
        |__ Cost Optimization
        |       |__ Focuses on avoiding unnecessary costs, selecting the most appropriate and right-sized resources, analyzing and reducing licensing costs, and using a consumption model to save money on unutilized resources.
        |


AWS Trusted Advisor
        |
        |__ Cost Optimization
        |       |__ Right Sizing - Recommends resizing Amazon EC2 instances to save money.
        |       |__ Reserved Instance Optimization - Recommends purchasing reserved instances to save money.
        |       |__ Savings Plans - Recommends using Savings Plans to save money on EC2 usage.
        |       |__ EC2 Instance Underutilized - Recommends identifying and shutting down underutilized EC2 instances to save money.
        |       |__ Idle Load Balancer - Recommends identifying and shutting down idle load balancers to save money.
        |
        |__ Performance
        |       |__ High Utilization Amazon EC2 Instances - Recommends identifying and addressing instances that have high CPU usage to improve performance.
        |       |__ Unassociated Elastic IP Addresses - Recommends identifying and releasing unused Elastic IP addresses to improve performance.
        |       |__ Underutilized Amazon EBS Volumes - Recommends identifying and addressing underutilized EBS volumes to improve performance.
        |       |__ Amazon RDS Idle DB Instances - Recommends identifying and shutting down idle RDS instances to improve performance.
        |
        |__ Security
        |       |__ Security Group - Specific Ports Unrestricted - Recommends restricting access to specific ports in security groups to improve security.
        |       |__ IAM Use - Recommends using Identity and Access Management (IAM) to improve security.
        |       |__ MFA on Root Account - Recommends enabling multi-factor authentication (MFA) on the root account to improve security.
        |       |__ RDS Public Access - Recommends disabling public access to RDS instances to improve security.
        |       |__ S3 Bucket Permissions - Recommends restricting access to Amazon S3 buckets to improve security.
        |
        |__ Fault Tolerance
        |       |__ Load Balancer Optimization - Recommends optimizing load balancers to improve fault tolerance.
        |       |__ Auto Scaling Group Recommendations - Recommends using Auto Scaling groups to improve fault tolerance.
        |       |__ Amazon EBS Snapshots - Recommends taking regular Amazon EBS snapshots to improve fault tolerance.
        |       |__ Multi-AZ Amazon RDS - Recommends using Multi-AZ Amazon RDS to improve fault tolerance.
        |
        |__ Service Limits
                |__ Service Limit Increase - Recommends requesting limit increases for AWS services when needed.


| Plan Name | Cost | Support Level | Response Time | Support Channels | Features |
| --------- | ---- | ------------ | ------------ | ---------------- | -------- |
| Basic | Free | N/A | N/A | Documentation, whitepapers, support forums | AWS documentation and general guidance. |
| Developer | Starts at $29/month | Level 1 | 12 hours | Email | Technical assistance during business hours. |
| Business | Starts at $100/month | Level 2 | 1 hour | Phone, email, chat | Technical support 24/7, Trusted Advisor, Infrastructure Event Management. |
| Enterprise | Custom pricing | Level 3 | 15 minutes | Phone, email, chat | Dedicated technical account manager, Infrastructure Event Management, personalized support team. |
| AWS Infrastructure Event Management | Starts at $100/month | N/A | N/A | Phone | Technical support for unexpected infrastructure events. |

https://blog.wimwauters.com/devops/2021-11-01-flask_githubactions_elasticbeanstalk/
