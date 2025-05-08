# Introduction

## Course Introduction

# Introduction to Monitoring in Google Cloud

## Module Overview
* for ops and developers

## Overview of Monitoring Tools
* Logs Viewer is now logs explorer
* Logs API can write to google cloud logs
## Operations-Based Tools
* google collects 1000 different streams of data
* collects data from many details about 3rd party services
* data logs stay up to 30 days max 365000 days
* cloud audit is used for forensics

## Application Performance Management Tools
* debugger
  * does not stop slow down application
  * easily knows which version to work with
* trace
  * analyzes all application traces
* profiler
  * can provide heat map of application w/o slowing it down

## Reading: Product Knowledge Instructions

## Reading: Product Knowledge Worksheet

# Avoiding Customer Pain

## Module Overview
* four golden signals
  * SLI,SLO,SLA to avoid customer pain

## Why Monitor?
* collection,proessing,aggregating and displaying real-time quantitiative data about a system
* create good trust and good products that people can trust
* needs to live on resources that can support capacity
* we want alrrt
* we want data crucial to debugging
* answer questions (what to do for holidays
* good SRE will let the customer know as appropriate
* monitoring is a skill
* KISS
* no single tool can do it all
* monitor from multiple vantage points
* can monitor 15 projects from a 16th workspace
* syptom
  * 500 for people in europe
  * cause
    * database or zone is down
* white box- values based on internals of system
* black box- values based on externals of systems
## Critical Measures
* metrics - might rely on ROI
  SMART
    * specific
      * results in 1 ms
    * Measurable
      *
    * Achiveable
    * Relvaant
    * Time bound
__four golden signals__
* latency, how long to return a result
  * # of requests waiting for a thread
* traffic
  * how muc is hitting your system
* saturation
  * how close a system is at capacity
* errors
  * indicate something is failing


## SLIs, SLOs, SLAs
* __service level indicator__ -
  good evnets/all valid events
* __service level objective__
  * want to achieve 99.9%
  * agreeing on desired relibality
  * product mangers hate this
  * more impt feature is relability(perform applicaton w/ a certain time w/o failing)
  * compensating customers can get problematic customers get upset far before SLA break
  * 100% reliability cant make it happen
  * 90% reliable means 36 day/1yr
  * 99.99% means 5mins/1yr
* you can pledge
ITIL -  princples on how to update your applications
* what should we spend error budget
  * stop making changes until service is back in SLL
  * avoid risky experiements




## Choosing a good SLI
* many users who are unhappy, will stop using the service
*

## Specifying SLIs
* make games
  * in game engine, people can pay for resources via money

## Postmortem Review

## Developing SLOs and SLIs
* SLO is where the user is happy
* its not about reliability its about user happyness
* a lot can happen in a year
* sooner you can gather data the better
* SLI out of your control dont cover, even though its good to know about them

## Lab Intro:  Developing SLOs and SLIs
https://app.pluralsight.com/lti-integration/redirect/b3913730-5414-423f-bbe0-3eb246bcaeee?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Develop SLIs and SLOs
![1680818926316](image/README/1680818926316.png)
### Task 1. Study the Buy In-Game Currency user journey

### Task 2. Design a set of SLIs and SLOs


# Alerting Policies

## Module Overview

## Developing an Alerting Strategy
* window length
  * like to receive when trend is past $1000
* add duration for better precisons
* use multiple condtions for better precision
* can define multiple alerts on multiple channels
* give priorities to your alerts
* pagerduty is an incident repsonse app
* when an incident occur cloud montioring opens it up
* all resouces can be monitored as a unit

## Creating Alerts
* already policy is in JSON or yaml


## Creating Alerting Policies with the CLI
https://app.pluralsight.com/lti-integration/redirect/c6ac6f68-ac21-43a7-9027-e52dcaa40055?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab Intro: Alerting in Google Cloud

## Lab: Alerting in Google Cloud

### Task 1. Download and test a sample app from GitHub

### Task 2. Deploy an application to App Engine

### Task 3. Examine the App Engine logs
* Nav-> GAE -> Versions -> Diagnose -> Tools
* now you can see the requests yr app has made
### Task 4. Create an App Engine latency alert
Monitoring > Metrics explorer
### Task 5. Creating an Alerting Policy with the CLI


## Service Monitoring
* when something fails it seems like many things fail
* windows-based slo are basd
  * if error occur Friday moring, people just know not to use it Friday morning
*

## Lab Intro:  Service Monitoring
https://app.pluralsight.com/lti-integration/redirect/0be59a60-83e4-49d2-b21a-e145f2091d60?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Service Monitoring

### Task 1. Deploy a test application

### Task 2. Use Service Monitoring to create an availability SLO
* When an app throws an errror
Navigation Menu > Error Reporting
Navigation Menu > Services > [GAE instance] > create SLO
Availbity. Request based
  * rolling, period length to 7 days
  * goal 99.5 slow
  * if you click on the SLO  you can see that your are burning through it as were causing more errors
* create SLO alert
* you get to set a percentage, choose notification channel and options on what a person should do next

# Monitoring Critical Systems

## Module Overview
* single pane of glass
* IAM will access to dashboard
* smaller workspaces
* monitoring viewer editor and admin
* services many need to write metrics
  * monotiroing metric writing, (needed by service accts)

## Observability Architecture

## Understanding Dashboards
* create dashboards so you can understand things

## Creating Charts
* create charts with metrics explorer
* there is outlier mode
* stark

## Dashboard Construction
* can put dashboard on an iframe
* crating from monitoring  page or metrics explorer

## Uptime Checks
* org allows for errors
* uptime checks when we find issues
* also known as healthcheck
* easy to create

## Lab Intro: Monitoring and Dashboarding Multiple Projects from a Single Workspace
https://app.pluralsight.com/lti-integration/redirect/3cd35394-9587-48bc-a769-cd74f967f977?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Monitoring and Dashboarding Multiple Projects from a Single Workspace

### Task 1. Configure the resource projects
Label Project ID 1 as Monitoring Project. qwiklabs-gcp-00-77242c05866f
  Uptime check Frontend Servers Uptime Check
Label Project ID 2 as Worker 1. qwiklabs-gcp-03-66e273122e5c
  worker-1-server-vm http://34.69.103.209/
    instance ID 1442103103179643031
Label Project ID 3 as Worker 2. qwiklabs-gcp-03-63a78e81e96e
  worker-2-server-vm 35.194.22.196

### Task 2. Create a Monitoring Workspace and link the two worker projects into it
Monitoring > Overview >ADD GCP projects
* Each Workspace can support up to five-hundred groups and up to six layers of sub-groups.

### Task 3. Create and configure Monitoring groups
* Monitoring > Groups
  * create sub group

### Task 4. Create and test an uptime check
Monitoring > Uptime checks
Metrics Exploer chose
VM iNSTANCES >Uptime Checks > Check pass
  * view via resource ID
go back to alerts to check out the uptime check
### Task 5. Create a custom dashboard
Monitoring > Dashboards
* cloud shell does work will with stress testing, it will think u  are ahacker
# Configuring Google Cloud Services for Observability

## Overview
* integrating mointoring via agents
## Monitoring
* monitoring can access some metrics, because compute engines run on google hardware
* agents can grab even more such as memory usage
* App Engine has agents
* Anthos GKE On-Prem has integrated supor
* look up the monitoring agent
* adding service account
  * greant monnotiring agent
  * pull down cred.json
  * start the application adding to env var

## Logging
* logging agents work in many places

## Baking an Image
* many organization make images by hand
* image should start w/ base OS
* security should remove everything and make it secure'
* hashicorp packer helps automate virtual builds

## Non-VM Resources
* app engine standard,flex support monitoring
* there is also prometheus

## Exposing Custom Metrics
* application user metrics, use SDK to do this
  * OpenCensus custom tracing library
    * export latency and label
    * openCensus needs the creds
  * GCP Cloud Montoring

## Lab Intro: Compute Logging and Monitoring
https://app.pluralsight.com/lti-integration/redirect/efb15de6-51f7-4239-868f-4b508b01a941?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Compute Logging and Monitoring

### Task 1. Set up a VM and a GKE cluster
* so in the Metrics Explorer if you try to get metrics related to nginx nothing happens the metrics agent is missing

### Task 2. Install and use the logging and monitoring agents for Compute Engine
* so after you installed monitoing (stack drier) and logging (google-fluentd) you should see the nginx metrics in the resource manager

### Task 3. Add a service to the GKE cluster and examine its logs and metrics
Monitoring> Dashboards > GKE, as you can see no agents are needed to be installed unlike vms

# Advanced Logging and Analysis

## Module Overview

## Labeling
* use labels
  * apply programmicatly if possible
  * less 63 chars,

## Working with the Log Viewer
* you can choose resouces, log and log severity
* can expand log entires
* download options
* basic filters simple for searching for text
for:bar
  * no ranges, or time stamps
* there are summary lines
* google has log entry data types
* there is advanced mode
  * supports groups, comparison. AND/OR
  *be specifc,
  * limit the range


## Using Log-Based Metrics
* IAM roles
  * Logging config writers
  * Logs Viewers
  * Monitioring Viewer
  * Logging Admin
* over 1,000 predefined metrics
* winston node.js logging lib integrations w/ google cloud
* metric types
  counter
  distribution - (only support 200 buckets)
    * linear,exponetial,histogram
* each label is limit up to 30,000 time series

## Exporting and Analyzing Logs
*  Gcp logging is a collections of services
* There is log router (where should I send logs to)
* there are include and exclude filters
* you get resource usage prediction for your logging
* log sink go to big query, cloud storage and cloud pub/sub
* log sink are write as the same as log inclusions/exclusions
* dataflow is excellent for log stream
* aggregation levels
  * project level
  * folder level
  * organization level log sink
* aggregated sinks can export billing project
  * destination must exist
* bigquery can
  * stream, make insights accessible, can build for AI, provide real-time insights

## Error Reporting
* 3 error values, 'prod','always','never
* you can manually report an error
* errors understand language behind them and joins accordingly

## Lab Intro: Log Analysis
https://app.pluralsight.com/lti-integration/redirect/462c9640-5187-4cc2-b282-8d00e4d225e7?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Log Analysis
* logs come as unstructured text from endpoints cant do much w/ it
### Task 1. Set up a test application, deploy it, and set up a load test VM

### Task 2. Explore the log files for a test application

### Task 3. Create and use a logs-based metric
Logs Viewer
### Task 4. Export application logs to BigQuery

# Monitoring Network Security and Audit Logs

## Module Overview

## VPC Flow Logs
* every 10 seconds packets are collected and logs for analysis and forensics
* google cloud logs viewer allows one to see the logs
* can transform this in data studio

## Firewall Rules Logging
* can audit and analyize firerulles
* its default by disabled
  * watch out it can generate plenty of data


## Cloud NAT Logs
* security, no need for vm to have external IP address
* global
* scale the number of NAT ip addresses it uses

## Packet Mirroring

## Network Intelligence Center
* give topology, connectivity tests, perforamce dashboard and firewall
* topology
  * on graph of everything
*

## Audit Logs
* all services will eventually provide audit logs
* admin activity logs, free, retained for 400 days
* data access, needs to be enable

* log viewers
  * set the name to the resource type which you want to query

## Data Access Logging
* enablked by org,folder,project,serice
* 50c per gigabyte


## Understanding Audit Logs
* helps understand who ran that $40,000 bigquery query

## Best Practices
* logs can grow quite big
* IaC - Terraform, a way to use code to create infrastucture
* be careful who gets access to logs
* scruitize data, contain PII

## Lab Intro:  Cloud Audit Logs
https://app.pluralsight.com/lti-integration/redirect/2e8311bb-4feb-41ef-b265-249c5a37c813?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud
## Lab: Cloud Audit Logs

### Task 1. Enable data access logs on Cloud Storage
 IAM & Admin > Audit Logs. use the filter to get a service and then the logs from certain permissions

### Task 2. Generate some admin and data access activity

### Task 3. Viewing audit logs
* To view the logs, you must have the Cloud Identity and Access Management roles Logging/Logs Viewer or Project/Viewer.

Cloud OverView > Activity

Logs Explorer > Log Name > Audit log > activity > take a log > protoPayload > authenticationInfo (sees who deleted a resources)

## Lab Intro: Optional Lab - Analyzing Network Traffic with VPC Flow Logs
https://app.pluralsight.com/lti-integration/redirect/511633d4-5439-48e0-a5d2-37e32c2880f7?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Analyzing Network Traffic with VPC Flow Logs

### Task 1. Configure a custom network with VPC flow logs
* create a network and vm to that network
* access the apache server on that vm

### Task 2. Create an Apache web server

### Task 3. Verify that network traffic is logged
Logs Explorer > Resouce TYPE > Subnetwork >
  Log Name > compute.googleapis.com/vpc_flows

add your IP address at the end of the query
### Task 4. Export the network traffic to BigQuery to further analyze the logs
More Actions > Create Sink
### Task 5. Add VPC flow log aggregation
* once you generate traffic in bigquery run
```sql
#standardSQL
SELECT
jsonPayload.src_vpc.vpc_name,
SUM(CAST(jsonPayload.bytes_sent AS INT64)) AS bytes,
jsonPayload.src_vpc.subnetwork_name,
jsonPayload.connection.src_ip,
jsonPayload.connection.src_port,
jsonPayload.connection.dest_ip,
jsonPayload.connection.dest_port,
jsonPayload.connection.protocol
FROM
`your_table_id`
GROUP BY
jsonPayload.src_vpc.vpc_name,
jsonPayload.src_vpc.subnetwork_name,
jsonPayload.connection.src_ip,
jsonPayload.connection.src_port,
jsonPayload.connection.dest_ip,
jsonPayload.connection.dest_port,
jsonPayload.connection.protocol
ORDER BY
bytes DESC
LIMIT
15
```

# Managing Incidents

## Module Overview
* things will go wrong, how do we manage user issues

## Incident Management
* you hurt yr ankle (alert)
  * you walk instead of run (triage)
  * incident (you cant walk and need to do something)
* can I safely ignore an alert, most people do not notice alerts
* learn from pilots certain checklist
* respond quickly
* maintain clear chain of command
  * Incident Commander
    * Communications Lead (CL) - Operations Lead()
    * Communications Lead (CL) -users and stakeholders
* train, train, train the team
  * dont do it in the middle of development
* keep a playbook

## Declaring an Incident
* guidelines to when to declare an incident

## Restoring Service
* getting it restored in the priority
  * do we have a hot or cold replacement
    * do we have IaC to spin up a new environment
  * you have to do what it takes to restore the service
  * get to the root cause
    * might be because of documentation

## Preventing Recurrence
* standarding your docs, can get this from pagerduty and atlassian

# Investigating Application Performance Issues

## Module Overview

## Debugger
* debuggser
* there is snapshot and logpoint
* dynamic log messages with log points

## Trace
* collects latency data and displays in GCP,
* capture from all VMS
  * each trace is collected by span
* can figure out if VM is making too many calls to your VM
* trace list windows allows you to view things in detail
* recommend opencensus w/ python
* app engine, cloud  run has default access
* need cloud trace agent role

## Profiler
* profiler helps understand which resouces are consuming the most resouces
![1680965640455](image/README/1680965640455.png)
* make sure Profiler API is enabled, then install googlecloudprofiler

## Lab Intro:  Application Performance Management
https://app.pluralsight.com/lti-integration/redirect/7c811de2-bb93-45a9-9662-16c4cec00b08?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Flogging-monitoring-observability-google-cloud-3&contextTitle=Logging%2C+Monitoring+and+Observability+in+Google+Cloud

## Lab: Application Performance Management

### Task 1. Download a pair of sample apps from GitHub

### Task 2. Deploy the converter application to App Engine

### Task 3. Debug the application
* debugger is on its way out

### Task 4. Adding log data

### Task 5. Fix the bug and deploy a new version

### Task 6. Examine an error report coming out of Cloud Run in Error Reporting

### Task 7. Examine a default and custom trace span
* By default, Trace only captures spans where services call each other. Since the /slow route does a calculation but doesn't call any other services, the telemetry Trace provides is limited. Fix that by adding a couple of manual trace spans.

# Optimizing the Costs of Monitoring

## Module Overview

## The Costs of Monitoring
![1680966714411](image/README/1680966714411.png)
![1680966729404](image/README/1680966729404.png)
* exporting network telemetry always costs
## Bill Estimation
* billing page
* if the sku is 0 it wont appear in the list
* if you are spending money on metric data what I am doing
* filter by SKU
  * you can se daily, monthy, daily cumulative or monthy cumulative
* external metrics bump up the bill


## Cost Control Best Practices
* cloud load balancing
* computer
* when data is trace its packages by spans
* manual spans get charged
  * watch out for loops
* export an exclude
* be selective w/ the logging agent

# Course Resources

## Course Resources
