# Understanding Cloud Computing
* create aws account
* use AWS budgets to set  budget alerts
* set a budget and you will be notified if you get to 85% of that budget
* securing funding for inital infrastuctior
* took 5 months to build infrastucture
* 1 months to scale up and down
* data centers
  * large up-front investment,
  * forcasting is difficul
  * slow to deploy
  * maintaining is expensive
  * responible for all secuirty and compliance

## Types of Cloud Computing
* pay as you go pricing
* IaaS - run server in the cloud
  * configure to do exactly what we want
* SaaS - no configuring server
* PaaS- deploy own app on it
![1679777128057](image/README/1679777128057.png)


### Cloud Deployment
* Public Cloud
* Private Cloud
  * solutions from vmware

# AWS Global Infrastructure

## AWS Regions and Availability

* each region has a cluster of data centers
* avaliability zones has one or more data center
  * exists to make sure the region does not go down
* regions have several avaliability zones
![1679787015409](image/README/1679787015409.png)
* us-east 2a 2b is availaiblity zone the letters for each availability zone
![1679830087868](image/README/1679830087868.png)
* wavelength zones, effeciectly application that run on 5G networks

## AWS Edge Location
* over 410 different location

## Visualizing AWS Global Infrasture
* https://aws.amazon.com/about-aws/global-infrastructure/
* you can hover over the map to see the availaibility zones
* amazon cdn - aws edge locations
* high uptime- AWS availbility zones


# Understanding Cloud Economics
* __Capilizeing Expenditure__ - getting things started
* __Operatized Expendituire__ - costs to keep things going
![1679831819333](image/README/1679831819333.png)

* __AWS Cost Explorer__ -
  * gives recomenndations for cost optimization
  * can be accessed via API
  * provides predictions for the next 3 months of cost
  * provides breakdowns including
    * by service
    * by cost tag
* __AWS Budgets__
* __AWS Pricing Calucator__
  * 5 websites 2  server 500 gb of data
* __AWS Migration Hub__
  * compare costs from own data center to cloud

  * dont use AWS TCO Calculator, AWS Simple Monthly Calculator

## Using the AWS Pricing Calcuator
* https://calculator.aws/#/addService

## Reviewing Costs with the Cost Explorer
* cuurent monthly costs as well as forcasested month-end costs

## Supporting AWS Infrastucture
*  you can create a resource tag for each depeartment

## Support AWS Infrastructure
* AWS Support - provided in different tiers
* AWS Trusted Advisor- best practices, all customer get 7 core
  * AWS Developer Support
  * AWS Enterprise Support- $15,000 a month
![1679837274092](image/README/1679837274092.png)
```ps
* aws s3 ls
```

# Compute Services

## EC2
* server on
* make good choice on instance types
* always use EBS for the work ur going to do, shut down data stays
* reserved instances, discounts when you can commit for a specific period of time
* there is standard and convertible
* if you have batch processing where it starts and stops use spot instances
* if you have compliance use Dedicated Host
* if you are leveraging lambda
* predictable but not steady workload purchase a Scheduled Reserved Instance

![1679856899062](image/README/1679856899062.png)

## Reserved Instance EC2 Pricing
![1679856858878](image/README/1679856858878.png)
* plenty of saving if you can commit to longer times


![1679857719554](image/README/1679857719554.png)

## Launching EC2 instances
![1679858247611](image/README/1679858247611.png)
![1679858266712](image/README/1679858266712.png)
* there are many AMI to choose from
![1679858877125](image/README/1679858877125.png)
* Be careful with AWS Marketplace AMI's
* some are free tier eligble
* choose instance type
* firewall
  * only certian people can log into the server
* user data is just a list of startup commands for the instance
* we get a name for the instance
* we get public ipv4 for the dns we just lanuched
* to terminate the instance
  instance state -> terminate

## Elastic Beanstalk
* like EC2 its PaaS
* lot of backend but  load balancing is handled by beanstalk
* does monitoring, deployment,scaling,& EC2 Customization

### Use Cases
* deploy app with minimal knowledge of other services
* few customization
* reduce all overall maintenance needed

## Lanucghing app on Elastic Beanstalk
[tutorials](docs.aws.amazon.com/elasticbeanstalk/latest/dg/tutorials.html)
* tags are essential to categorize different costs within the application
* many defaults
![1679919460184](image/README/1679919460184.png)
* we can see load balancer configuration
* there are logs, health and monoitoring and logs
* to delete -[APP ENV NAME] -> actions -> teminate env

## AWS Lambda
* 128 MB - 3008 MB
* event driven
* only charged for execution time
* reduced maintenance
* enables fault tolerance
* scales based on demand
* price based on usage not size like EC2

## Container Services
* AWS App Runner
### Container Orchestration Services
* ECS
* EKS - kubernetes services

* EC2 - not fun
* AWS Fargate - scales container resources for you
![1679919835230](image/README/1679919835230.png)

# Content and Network Delivery Services

## Amazon VPC and Direct Connection
* logically isolated section of aws cloud where u can define and configue
* supports IPv4 and IPv6
* private, public subnets
* connect to other VPC,
* keep traffic within VPC

### AWS Direct Connect
* easy to establish dedicated network connection from data center to AWS


## Amazon Route 53
* DNS,highly available
* send the to server based on country
* changes take hours to propagate
* can configure failover

## Elastic Load Balancing
* ALB -application load balancing
* NLB- network load balancing
* t3 instance, m4large
* horizontal scaling, add more serve to handle load

## Amazon CloudFront and API Gateway
* configure to send info closer to them
* AWS edge locations
* AWS Edge locations
* AWS Shield WAF
### API Gateway
* API management services

## AWS Global Acceletor
* improces netowrk by 60%
* reolsution route through AWS network
* AWS network is used over public internet
* IP clears cacahe
![1679921064456](image/README/1679921064456.png)

# File Storage Services

## Amazon S3
* store files as objects  in buckets'
* storage classs
* multiple availizability zones
* URL
* expire after a certain period of time

###
![1679921156793](image/README/1679921156793.png)
### s3 Intelligent Tiering Storage Class
* move throgh access and archivesd classes
* S3 lifecycle policies, versioning, experiation
__s3 transfer acceleration__ -uploads data faster

## Hosting a Website on Amazon
* you can also encrpy the bucket
* can choose storage tier
* permissions -> read access-> save permissions

## S3 Glacier
* archived storage
![1679921639104](image/README/1679921639104.png)
* instant reterival in miliseconds
* flexible reteriveal - hours to mins
  * flexible retrieval - backup disater recovery
  * archived data for several to 10 years

## Elastic Block Store

### EBS Elastic Block store
* connected to single EC2 isntance
*snapshots, redudancy, multiple volume drives, (HDD,SSH)
![1679921780070](image/README/1679921780070.png)

### Elastic File System
* NFS file system
* standard,infrequent,
configure  lifecycle
network file system

#### Amazon FSx
* windows file server

## AWS Snowball
* snoball petabyte device transfer,
* snowmobile exabytes - large piece of data
* devices are ship[ed]


# Database Services and Utilities

## RDS
* provisiong, patching, backup and recovery
* multiple AZ
* launches into a VPC
* Amazon Aurora is a Postgres MYsql
__Amazon DMS__ - migrate data

## Amazon DynamoDB
* NoSQL service
* key-value & document DB
* automated scaling
* accelator DAX
* 10  trillion requests per day
* scale w/o execcsive maintenance,
  serverless
  * low latency
  * data models w/o blob storage

## Elasticcache
* memcached & redis
* im memory services between application and netwok

## Redshift
* data warehouse for analysis of
* gets VPC
* spectrum can query exabytes of data

## Additional Database Services
* DocumentDB- mongo
* Nepture
  * graph db , mapping of mapping like socal networks
* MemmoryDB
  * redis compatible api
* Timestream
  * sensor data, capturing data at a specific time

# App Integration Service

## AWS Messaging Services
* __amazon sns__
* pub/sub
*organize info on topics
* integrate w/ multiple AWS services
* messages are short lived
![1679922428580](image/README/1679922428580.png)

__sqs (simple queue services)__
stored for 14 days
256 kb payload
* use FIFO queue
__fault tolerant__ an aspect of our service goes down we have queues to still hold on to the data so the data does not get lost
![1679922595590](image/README/1679922595590.png)
* use amazon sqs

## AWS Step Functions
* manage business workflows, worlflowks done with amazon states language
* integrates w/ everything
![1679922682148](image/README/1679922682148.png)

# Management and Governance Services

## AWS CloudTrail
* log & monitor account activity throughout infrastructure
* add to cloudwatch
* logs events to regions which they occur
* best practice to turn on on every account
* compliance,forensic,operational,troubleshooting

## Amazon CloudWatch & Config
* __cloudwatch__
monitoring & mgnt
collects logs, metrics,events from services,enables alarms based on metrics
visualization metrics, custom dashboards

* __config__
monitors resource config against desired config
* conformance packs for compliance standards including PCI-DSS

## AWS Systems Manager
* managing tools to manage infrastructureo, enables automation tasks, secure way to access servers for AWS credentials, stores commonly used passwords

## AWS CloudFormation
* custom application, insteading of going into cli and doing everyhting yourself make a mistake
manage service dependecies, drift detection, someone makes everything globally available we can know about it,
have production and testing

## AWS OpsWoks
chef and puppet, for cloud and onprem
* AWS OpsWorks for Chef Automate, AWS OpsWorks for Puppet Enterprise AWS OpsWorks Stacks


## AWS Orgnaizations & Control Tower
* manage multiple account under master acct, provuides orgs with availity to leverage considated billing

 * __control tower__ centrializers users across all AWS accounts,
 AWS new acct template, guardrails and no need to log into every account
