# Welcome to Preparing for the Associate Cloud Engineer Exam

## Welcome and Getting Started Guide!

# About the Associate Cloud Engineer Certification

## About Google Cloud certification
* only worry about techincal stuff

## Devising a study strategy, part 1

## Devising a study strategy, part 2
* online practice labs and quests, google labs
google.quikalbs.com

## Devising a study strategy, part 3

## Are you ready?

## Taking the exam

## Resource links for Module 1

# Setting up a Cloud Solutions Environment


* org ,folder,project,resources
* project name and ID cant be change
* decide who on your team show be org
* resources permission inherit permission
  * child cannot take away from parent
* there are basic, predefined an custom
* basic
  * owner,editor,viewer, billing adiminstrations
* predinfed roles exist so people know what to do

* REST
* multi cloud for all info
  * can get to root analysis
  * number of days by type

* Creating billing account
  * must be billing admin
  * set budget alert must be billing admin

* sdk
  * set of tools to manage GCP applications
  * also available as a docker image

## Planning and Configuring a Cloud Solution
* budgeting using the pricing calculator
* relational for ecommerce, CMS or financial for rigid strcuture

* cloud datastore, bigtable
  * add new fields for featires

* handle steraming for realtime data, based on SQL

* Cloud Storage
  * Standard Storage,Nearline Stroage(30 days),Coldline Storage (90 days), Archive Storage (365 days)
  * load balnacing
    * global, regional, external internal, SSL,TCP
    * global had ipv6 termination, regional has ipv4

* network load balancer is regional and non-proxied load balancer
Lab: Setting Up Network and HTTP Load Balancers [ACE]
https://app.pluralsight.com/lti-integration/redirect/9eeef07c-a1f7-424c-9c04-438fe08bc7c0?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fpreparing-associate-cloud-engineer-examination

## Deploying and Implementing a Cloud Solution
* 2 types of vm goroups, managed and un-managed
  * managed just means idenicatal vms
* pod is smallest thing you can deploy on contaies
  * unique IP ports and instructions
* App Engine
  *standard and flxible
  * free daily
  * Also provides testing to develop locally
  * cant install 3rd party steps

* Cloud Functions
  * billed by use
* Clound Spanner got 10,240 mib per row
* VPC
* if you want to get started with google ASAP using cloud marketplace
  * doesnt update sotware after you deploy to marketplace
* Deployment Manager is IaaS
  * create .yaml or python

### Lab
https://app.pluralsight.com/lti-integration/redirect/b5c3fa10-b455-4a9b-af78-7fe5e9a420af?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fpreparing-associate-cloud-engineer-examination

## Ensuring the Sucessful Operation of a Cloud Solution
* google has vm images
* custom images only available to the project
* snapshots works in an incremental manner
### Deploy to GKE
https://app.pluralsight.com/lti-integration/redirect/fbe38edc-9239-4246-b8e9-0d6dc0d939c9?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fpreparing-associate-cloud-engineer-examination

* app.yaml
  * scaling type and instances

* you can cofigure lifecycle rules on a bucket
* lifecycle actions
  *delete
  *setstrogeclass
* lifecycle conditions
  * age
  * createdbfeore
  * isLive

* can remove subnet if no other instances are using it
* from auto to custom, or auto the broadest CIDR range is /16

## Configuring Access and Security
* accounts
  * ggogle acount
  * service, google workspace domain, google group
  * cloud idendity doesnt have gsuite

in the Manage Resouces you can see groups and resources and projects select a project and add an IAM role

### Create a custom role
IAM > Roles
* refer basic and predefined roles

* Service Account belongs to a VM
* scopes and permissions must agree our else the feature wont run
* service acct permission can change w/o recreateing vms
Admin Activity,System Events, Data Events

### SRE w/ cloud monitoring
https://app.pluralsight.com/lti-integration/redirect/3652957c-7ca9-4352-8604-4923c6fbae03?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fpreparing-associate-cloud-engineer-examination

