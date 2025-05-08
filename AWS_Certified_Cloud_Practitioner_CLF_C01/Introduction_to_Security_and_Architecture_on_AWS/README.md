# AWS Architecture Core Concepts
*
## Secrutiy and Architives
* sending unsolocted mass emails, hosting or distrubting harmful content
* least priviliege access
## Shared Repsonibility Model
* AWS responsiblity of cloud
* customer responsiblit in cloud
![1679923753806](image/README/1679923753806.png)
## AWS Well-architected Framework
![1679924123304](image/README/1679924123304.png)

## High-availability and Fault Tolerance
* everything fails all the time,
* EC2 we have to worry about availaiblity
* SQS holds data, route 53 unhealthy route and reroute
## Compliance
* PCI-DSS,HIPAA,SOC 1 SOC2 SOC3, FedRamp,ISO 27018
* AWS Config,Artifacts
![1679924464860](image/README/1679924464860.png)
* aws artifacts sign nondisclouse to get agreements

# AWS Identities and User Manamgenent

## AWS IAM
* this is free, auth and auth, identify federation (external IDP)
* iam types
  * users,group consider aws mgnt tool,role

__policies__
JSON DOC permissions for an AWS IAM identity, custom or managed by AWS,
![1679924858448](image/README/1679924858448.png)
### Best Practices
* MFA, least priviligees access
## Creating and Managing IAM Users
* IAM Service ->
  Create User
    Change passwords
  Add user to group, choose a role and permission to user
     can serach for s3 policies and see the json representation

  User groups
    create new group
      name,give permisions
    Users tab
      add user to group

* users can be a part of several groups

## Enabling MFA

  Security credentials -> MFA Device
  Users
    secruity credentials
      MFA

## Amazon Cognito
* IDP for your customer
  * UI components for many applications,security capaiblityies
  * Configure cognito and s3 so users only have access to their photos

# Data Archicture on AWS

## Integrating On-Premise Data
__AWS Stoage Gateway__
* cloud storage to local netowrk
* file gateway,
  tape gateway -  in IT backups on tape devices
  volume gateway - iscui volumnes
__AWS Data Sync__

## Processing Data
* __aws glue__ - ETL
  * serverless model,
* __emr__ - big data cloud processing using popular tools
  *
![1679925655515](image/README/1679925655515.png)
* __data pipeline__
  * can integrate on prem like aws glue
## Analyizing Data
* __athena__
  * enables querying of large amount data
  * write queries using standard SQL
  * based on query size
* __quicksight__ -bi visualization tool
  * per-user per-session
  * standard and enterprise
* __cloudsearch__
  * large data available to users to search
  * change based on infrastucutre
  * search tons of pdfs docs




## Integrating AI-ml
* __rekognition__ - image and video reconginition service
  * objects and actions in images and video
  * facial analysis
* __translate__
  fully managed service for translation of text
  supports 54 lanaguage
* __transcribe__
  * audio and sppech and convert to text

# Disaster Recovery on AWS

## Disaster Recovery Architecture
![1679926125515](image/README/1679926125515.png)
### Pilot Light
* key infrastucture in the cloud

### Warm Standby
* everything running not at full capacity

### Full site
* full site is running

## Selecting a Disaster REcovery Architecutre
RTO - we need 8 hours to get back an hour
RP0 - we lose a certain amount of time


# Architecting Applications on Amazon EC2

## Scaling EC2 Infrastructure
* horizontal is much better than veritacal
* __auto scaling__
  * 2019 serer needs certain features
  * minum max, and desired instance number
  * performs health check
  * exists in on ore more availablilty zones
![1679926663598](image/README/1679926663598.png)
__secrets manager__
  * integrates, RDS,DocumentDB,and Redsift
  * fine-grained access control

## Controlling Access to EC2 Instances
* EC2 security groups
  * firewall
  * works at instance level
* Network ACL
  * work at subnet
  * both allow and deny traffic,
  * custom acl deines all trafic until rules are added
* AWS VPN
  * single machineto VPN
  * differs from difect connecttravels on internet but encrypted
    site to site

    client


## Protecting Infrastructure from Attacks
* __aws shield__
  * keep running, knows when a DDos is happening

* __macie__
  * ML personal information, that can get leaks

* __inspecter__
  * scanning for best practices keeping machines up to date
    * network reacilbility
    * host

## Deploying Pre-defined Solutions
__aw service catalog__
  * appliocaions for your organization
__marketplace__ - aws approved services from 3rd party providers

## Deveploer Tools
* __AWSCodeCommit__ - git
  * iam policies
* __AWSCodebuild__
  * building appl and output artifiacs
* __AWS CodeDeploy__
  * for testing makes a pipeline as well
* ___CodeStar__ whole ci/cd

# The Certification EXAM

## The Exam
* multiple choice and answer
* score 700+
## Registering for the Exam
* register for exam
* certification cloud practicioner
* login to certfiaction exam

## Studying and Taking the Exam
* take time to analyze the intent,review what is required, skip a question 

## Important next steps
