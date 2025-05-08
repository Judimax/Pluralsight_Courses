

# Identity and Access Mamangement (IAM)

## Module Overview
![1680701060180](image/README/1680701060180.png)

## Identity and Access Management
![1680701011450](image/README/1680701011450.png)
![1680701031214](image/README/1680701031214.png)


## Organization
__project creator role__ - organization role
super admins acct can control organization
  * be POC
__org admin__
define  IAM
determine hierachy or resource folder

* folders are used to model different legal entities within the organizatii

__folder role__ admin,creater,viewer

## Roles
__basic roles__ -available roles in GCP
  * owner, editor and viewer, billing admins
__predefined roles__
  * GCP collection of permission
  permission [service.resource.verb]
  * network admin cant touch firewall rules and SSL certs

* __iam custom roles__


## Demo: Custom roles
* you can create rolor from selection or create from scratch
* custom roles dont get updates if they were created from a template0

## Members
* service account belong to application
* __google group__
  * collection of google accounts
* __workspace__
  * Cloud Identity to create and add users
* if you move a project to a different folder, the project inherits the proicles of the new folder
* use role recommendations
* __IAM conditions__ configure permissions based on conditions


## Service Accounts
* program account to auth to the service account
* user created, custom, built
* all project come with google api service account for all the interal work that google has to

__default compute enginer serivce account__
  * gets editor roles
  * __scopes__ are use to see if auth ids are authorizations
  * use IAM roles to specifiy scopes and permissions for service accounts
* service account user roles, allows people to impersonate a service
* there are google-managed and user-managed service accounts
* google does not save user private keys
* consider short lived service acct credentials or impersonation

## IAM best practices
* use porjects to close resources in the same trust boundary
* grant roles to groups instead of
* give service acct that clearly displays its name
* use cloud level IAP so resources are protected by roles and not firewalls

## Lab Intro: Exploring IAM
* w/o project owner role you cannot edit any of the role

### Task 2. Explore the IAM console

### Task 3. Prepare a resource for access testing
bucket - jabaad-sdahde-2
user 2 can see it

### Task 4. Remove project access
* user 2 it will be a while but they wont be a ble to see the bucket

### Task 5. Add storage access
* grant access and give the user the storage object viewer role

### Task 6. Set up the Service Account User
* create the account
* give it role
* give users access to this role

#### Grant everyone at altostrat.com the service account user roles
* manage peromssions > principals,altostrat.com > role, service account user
* now you can make a vm and then choose the read-object-buckets only with that role

### Task 7. Explore the Service Account User role
* when you ssh in to the vm the vm can only use what the service account can do
so it can only read storage buckets so if you try
```sh
gcloud compute instances list
```
this fails because the serice acct does not have the the right permis
## Getting Started with GCP and Qwiklabs

## Lab: Exploring IAM

## Lab Review: Exploring IAM

## Module Review

# Storage and Database Services

## Module Overview
* data needs to be stored whether database or object store
## Cloud Storage
![1680705787890](image/README/1680705787890.png)
![1680705814812](image/README/1680705814812.png)
* cloud memorystore (memory store)
* cloud storage  cant index
* standard storage - hot data
  * store dual region, multi region
* nearline
  * cheaper less effcient
  * 30 days
* coldline
  * 90 minimum days
* acrhive
  * 365 days mins

* availibility
  * dont lose data

* you can have classes on each object
* there is object lifecycle
* ACL
  * who has and level access to buckets
* signed url
  * time based access, no need to have google account
  * must expire after a resonalble amnt of time


## Cloud Storage Features
* can sync a vm with a bucket
* objects are immutable
* object versions
### Object lifecycle
* examples
  * downgrade storage of older year
  * key only 3 most recent version
  * changes occurs in async
  * takes 24 hrs to update object lifecycle
* pub/sub notifications for webhoooks, faster, and cheats

### Data import service
__transfer__applicace transter petabytees to google storage

### Strong global consistency
* upload and delete never gets 404
* list returns the button

## Choosing a storage class
cosider region for your users, consider storage access for the storage types

## Filestore
* filestore can scale
* EDA filestore can support for manufacturing
* genomics processing for pharamceturicals can scale like crazy

## Lab Intro: Cloud Storage

## Lab: Cloud Storage



### Task 1. Preparation
bucket-kasdopyls-tab

### Task 2. Access control lists (ACLs)
```sh
# To get the default access list that's been assigned to setup.html, run the following command:

gsutil acl get gs://$BUCKET_NAME_1/setup.html  > acl.txt
cat acl.txt

# set to private
gsutil acl set private gs://$BUCKET_NAME_1/setup.html
gsutil acl get gs://$BUCKET_NAME_1/setup.html  > acl2.txt
cat acl2.txt

# set to pulbic
gsutil acl ch -u AllUsers:R gs://$BUCKET_NAME_1/setup.html
gsutil acl get gs://$BUCKET_NAME_1/setup.html  > acl3.txt
cat acl3.txt
```
### Task 3. Customer-supplied encryption keys (CSEK)
```sh
# to generate AES-256 base-64 key.
```sh
python3 -c 'import base64; import os; print(base64.encodebytes(os.urandom(32)))'
# output
b'AvArkydwY5gfuD742t7LDN+KhRPIQy7H6lPeDDFdxUk=\n'
the key is AvArkydwY5gfuD742t7LDN+KhRPIQy7H6lPeDDFdxUk=

# create boto file
gsutil config -n
vim .boto

# FILE __.boto__
Before:
# encryption_key=
After:
encryption_key=tmxElCaabWvJqR7uXEWQF39DhWTcDvChzuCmpHe6sb0=

# upload the rest of the files to the bucjket and pull them down to see if they made it back


```

### Task 4. Rotate CSEK keys
```sh
Before:
encryption_key=2dFWQGnKhjOcz4h0CudPdVHLG2g+OoxP8FQOIKKTzsg=
# decryption_key1=
After:
# encryption_key=2dFWQGnKhjOcz4h0CudPdVHLG2g+OoxP8FQOIKKTzsg=
decryption_key1=2dFWQGnKhjOcz4h0CudPdVHLG2g+OoxP8FQOIKKTzsg=

# generate a new key and replacy the encrpytion_key entry
gsutil rewrite -k gs://$BUCKET_NAME_1/setup2.html

# reencypt a file
gsutil rewrite -k gs://$BUCKET_NAME_1/setup2.html

try to pull again if it fails that means the file was not reencrpyted wit hthe new key
```
### Task 5. Enable lifecycle management

```sh
# to view current lifecycle
gsutil lifecycle get gs://$BUCKET_NAME_1

# to set a lifecycle
# FILE __life.json__
{
  "rule":
  [
    {
      "action": {"type": "Delete"},
      "condition": {"age": 31}
    }
  ]
}
gsutil lifecycle set life.json gs://$BUCKET_NAME_1



```

### Task 6. Enable versioning
```sh
# get versioning
gsutil versioning get gs://$BUCKET_NAME_1

# set versioning
gsutil versioning set on gs://$BUCKET_NAME_1

# edit the file and upload w/ versioning option
gsutil cp -v setup.html gs://$BUCKET_NAME_1

# list all versions of a file
gsutil ls -a gs://$BUCKET_NAME_1/setup.html

# you can download the file by these listed versions as well
```

### Task 7. Synchronize a directory to a bucket
```sh
# make the directroy
mkdir firstlevel
mkdir ./firstlevel/secondlevel
cp setup.html firstlevel
cp setup.html firstlevel/secondlevel

# sync first level (just a fancy word for copy)
gsutil rsync -r ./firstlevel gs://$BUCKET_NAME_1/firstlevel

```

### Task 8. Cross-project sharing
in a different project
bucket-2-kalyodipsis-23
create IAM service acct
click on the service acct -> KEYS
save as credentials.json

create a vm in project 1 and ssh to  it
```sh
# to use the service acct thatn can access bucket 2 in project 2
# to do additonal work in the other project update the permissions and roles on the service act
gcloud auth activate-service-account --key-file credentials.json

```
### Task 9. Review


## Lab Review: Cloud Storage

## Cloud SQL
* has MySQL,Postgresql or MSSQL
* can connect
* 64tb, 60,000 IOPS. 624 GB of ram
* many verions
* write is replicated before transaction is confirmed as committed
* consider cloud spanner for horizontal scalalbility
* connected via interal IP is performance recommeded
* memorystore is best for real time apps like games

## Lab Intro: Cloud SQL

## Lab: Implementing Cloud SQL




### Task 1. Create a Cloud SQL database
* SQL -> Create instance

password /1_|F_2KBe<Is;"9
* Shared-core machines are good for prototyping, and are not covered by Cloud SLA.
* each vCPU can do 250 MB/s can have cores up to  theoretical maximum of 2000 MB/s
* performance sensitive make sure u have enough memory
* HDD is cheaper, SSD is best choice
* storage capacity affects throughput
*  Setting your storage capacity too low without enabling an automatic storage increase can cause your instance to lose its SLA.
* you also want private IP  you should be accessing db from internal vm for best practice, automatically allocate an IP range for your db
### Task 2. Configure a proxy on a virtual machine
* When your application does not reside in the same VPC connected network and region as your Cloud SQL instance, use a proxy to secure its external connection
SQL make a compute instance
* it takes a couple of minutes to make a SQL compute instance
```sh
# ssh into vm wordpress-proxy

# download cloud sql proxy and make it executable
wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy && chmod +x cloud_sql_proxy

# instance connection name qwiklabs-gcp-04-8aa74db0c4b5:us-central1:wordpress-db

save it to SQL_CONNECTION env var

# to activate a proxy conection
./cloud_sql_proxy -instances=$SQL_CONNECTION=tcp:3306 &
```

### Task 3. Connect an application to the Cloud SQL instance
```sh
# to get the external IP of your vm
curl -H "Metadata-Flavor: Google" http://169.254.169.254/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip && echo
```

### Task 4. Connect to Cloud SQL via internal IP
from the wordpress exteranl url provide the cloud SQL internal IP because you have the proxy running

### Task 5. Review

### End your lab

## Lab Review: Cloud SQL

## Cloud Spanner
* financial and retial applications can rely on this
![1680713240135](image/README/1680713240135.png)
* the info is synchronous across database
## Firestore
* run sophisciated queires no l;oss performanece
* backward cmpatiable with cloud datastore
* do have to worry about disaster
* realtime apps
* datastore for server
* realstore for more and web apps
* for constant changing schema use firestore

## Cloud Bigtable
* supports hbase api
* columns get a column family
* if a cell does not contain any data it doesnt take up any spa
* index based on alphabet
* very high volume of write
* read write latency of 10 m/s
* wide column DB
## Memorystore
* fully managed regis service
* instaces up to 300 GB
easy life & shift
* sub milisecond latency

## Module Review

# Resources Management

## Module Overview

## Resource Manager
* hieracy manager organization > folder> project>resource
* billing policy acculmate bottom up
* IAM policy acculmate top down
* resources are
  global
    images
  regional
    ip
  zonal
    disks


## Quotas
* 5 admin actions per sections
* only 24 CPU per region
* quoates help with billing spike
  * question your selection
## Labels
* key value pairs to organize your resoures
* they are not tags
  * tags apply to instances and are for instances
![1680713834330](image/README/1680713834330.png)

## Billing
* set a budget
* use data studio to understand your information
## Demo: Billing Administration
* click billing
* click transctions
* you can export as a file or export to bigquery

## Lab Intro: Examining Billing Data with BigQuery

## Lab: Examining Billing data with BigQuery


### Task 1. Use BigQuery to import data
bigquery new table -> import the csv from from a sample cloud storage object into a  table
### Task 2. Examine the table
* we download imported billing data into the project
* we see for 9.7 Mbytes of data there was no charge


### Task 3. Compose a simple query

### Task 4. Analyze a large billing dataset with SQL
```sql
SELECT
  product,
  resource_type,
  start_time,
  end_time,
  cost,
  project_id,
  project_name,
  project_labels_key,
  currency,
  currency_conversion_rate,
  usage_amount,
  usage_unit
FROM
  `cloud-training-prod-bucket.arch_infra.billing_data`

-- to understand which service is costing the most
SELECT
  product,
  COUNT(*) AS billing_records
FROM
  `cloud-training-prod-bucket.arch_infra.billing_data`
GROUP BY
  product
ORDER BY billing_records DESC
```

### Task 5. Review

## Lab Review: Examining Billing Data with BigQuery

## Module Review

# Resource Monitoring

## Module Overview
* relies on stackdriver

## Google Cloud's Operations Suite
* only pay for what you use in stackdriver
* monitoing, logging, error reporting
* parterns with splunk,pagerdupty

## Monitoring
* SRE allows things to be highly scalable
* Workspace is the root entity that holds monitor and config info
* a workspace you can view all gcp and aws accts
* you can get many metrics and charts
* create alreating policies when your service goes down
* alert symptions not causes
* avoid noise
* there are uptime checks
* you can create custom metics
* mointoring is at the bases of site reliablilty

## Lab Intro: Resource Monitoring
https://app.pluralsight.com/lti-integration/redirect/4b9b7644-9eed-4b0b-b5de-8cd16d19c98d?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fessential-google-cloud-infrastructure-core-services-11&contextTitle=Essential+Google+Cloud+Infrastructure%3A+Core+Services


## Lab: Resource Monitoring

### Task 1. Create a Cloud Monitoring workspace
navigation > monitoring

### Task 2. Custom dashboards
also check out metrics explorer so you dont have to create custom solutions
### Task 3. Alerting policies
Monitoring > Alerting > Create Policy
* there is also multi condition trigger
* you also need a notification channel

### Task 4. Resource groups
* Like Azure Resource Group you create groups from monitoring

### Task 5. Uptime monitoring
Mointoring  > Uptime checks
* check the uptime task
* its like a healthcheck

### Task 6. Disable the alert

### Task 7. Review
## Lab Review: Resource Monitoring
https://app.pluralsight.com/lti-integration/redirect/f08b4c22-e43e-4f17-ba92-e08e34a7ee8a?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fessential-google-cloud-infrastructure-core-services-11&contextTitle=Essential+Google+Cloud+Infrastructure%3A+Core+Services

## Logging
* can connect to GCP and AWS,
* can do in CLI or SDK
* can analyze logs in bigquery

## Error Reporting
* can do App Engine
* can doing java,node,python,php,ruby

## Tracing
* displays data in near real time

## Debugging
* inspect an application in realtime, without stopping or slowing it down
* can do breakpoints and supports

## Lab Intro: Error Reporting and Debugging

## Lab: Error Reporting and Debugging
### Task 1. Create an application
```sh
mkdir appengine-hello
cd appengine-hello
gsutil cp gs://cloud-training/archinfra/gae-hello/* .

dev_appserver.py $(pwd)

# in cloud shell use the <> icon to open the app via web preview

# close and eploy the app to app engine
gcloud app deploy app.yaml


# to view the app
gcloud app browse


```


### Task 2. Explore Cloud Error Reporting

```sh
# use sed to replace text in files (in this case were causing an error)
sed -i -e 's/webapp2/webapp22/' main.py

# deploy to app engine quietly
gcloud app deploy app.yaml --quiet

# then go to error reporting
* kuberenetes, app engine and compute engine support error reporting
```

### Task 3. Review

## Lab Review: Error Reporting and Debugging

## Module Review

## Course Review

## Next Course: Elastic Cloud Infrastructure: Scaling and Automation
