# Interconnecting Networks
*  there are many ways to connect to GCP
## Cloud VPN
* provide  99.9 SLA
* ikev1 & v2 ciphhers
![1680726677981](image/README/1680726677981.png)
* __ha VPN__ - very secure
* 99.99 service availability
![1680726807010](image/README/1680726807010.png)
* when configuing HA VPN, one option is transit gateway
* cloud VPN  supports dynamic and static routes
* for expanding nets , a BGP session is made
## HA VPN

## Lab Intro: Configuring Google Cloud HA VPN
* HA VPN is a high-availability (HA) Cloud VPN solution that lets you securely connect your on-premises network to your VPC network through an IPsec VPN connection in a single region
* HA VPN gateways have two interfaces, each with its own public IP address
* When HA VPN is configured with two tunnels, Cloud VPN offers a 99.99% service availability uptime.


## Lab: Configuring Google Cloud HA VPN

### Task 1. Set up a Global VPC environment

### Task 2. Set up a simulated on-premises environment

### Task 3. Set up an HA VPN gateway

### Task 4. Create two VPN tunnels
* add two tunnels (IP addr) from the gateway VPM (the 2 in this lab)
* connect from interface0 - interface0 and interface1 - interface1
* need 2 tunnels for each VPN

### Task 5. Create Border Gateway Protocol (BGP) peering for each tunnel

### Task 6. Verify router configurations

### Task 7. Verify and test the configuration of HA VPN tunnels

### Task 8. (Optional) Clean up lab environment
so to create a fake HA VPN between google cloud and on prem
1 VPC -> many Subnets -> many Firewall Rules -> many VM's -> 2 VPN gateways -> 2 Router ->  2 or 4 Router intrerfaces -> 4 BGP peers -> 4 VPN Tunnel

## Cloud Interconnect and Peering
* different,
layer 3 is public internet
layer 2


![1680728763510](image/README/1680728763510.png)

## Cloud Interconnect
* dedicated interconnect can be used to transer plenty of data over the the network
* partner interconnect- work with a provider, who has physical network with google
  * can be configured for 99.99 sla
IPsec < partner < dedicated
![1680728904015](image/README/1680728904015.png)

## Peering
* Direct Peering is done by exchanging BGP requreies Edge PoP
* Edge PoP
* if you are not near a PoP, use Carrier Perring

## Choosing a connection
* ![1680729044223](image/README/1680729044223.png)
* use direct connect if you dont have sophsiticated stuff

## Shared VPC and VPC Peering
* shared VPC means several projects share the same VPC
* designate a project as the host project
* VPC peering allows for RFC1918 IP
* Shared VPC is centrailzied
* VPC networking is decentrialized,

# Load Balancing and Autoscaling
* distribute compute over region

## Managed instance groups
* managed instance can scale to all instance in the group
* auto restart due to crash
* to create managed instance group, create instance template
* decide when you want to autoscle
## Autoscaling and health checks
* vm graph from stackdriver help understand how autoscale should work

## Overview of HTTP(S) load balancing
* instance or on anycast IP addres,
* supports IPv4 - v6
* can configure URL maps
![1680729658466](image/README/1680729658466.png)
* route is not based on health but based on best servce that can do the job
* there is heath check, sesesion affinity
* if all backend in a region is fully used new requests are routed to the next router
* takes several minuites to propagate changes


## Example: HTTP load balancer
* for example balance based on content
  * video backend handle videos
  * all else handles all else

## HTTP(S) load balancing
* you need SSL cert for the load balancer
* NEG - network endpoint groups
  * zonal,internet,serveless, hybrid connectivity
  * hosted outside of google cloud


## Lab Intro: Configuring an HTTP Load Balancer with Autoscaling
https://app.pluralsight.com/lti-integration/redirect/50288ddd-8fec-499b-889f-1580746410d3?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Felastic-google-cloud-infrastructure-scaling-automation-10&contextTitle=Elastic+Google+Cloud+Infrastructure%3A+Scaling+and+Automation

## Lab: Configuring an HTTP Load Balancer with Autoscaling
### Task 1. Configure a health check firewall rule

### Task 2. Create a NAT configuration using Cloud Router
* use nat for vm with no external IP adddresses to connect to the internet

### Task 3. Create a custom image for a web server
* ensure that when you create the VM you dont allow the boot disk to be deleted on deletion,
```sh
# to ensure that a service starts on runtine
sudo update-rc.d apache2 enable
```
* delete the vm but keep the disk to create a new image from that disk

### Task 4. Configure an instance template and create instance groups

### Task 5. Configure the HTTP load balancer
Network Services -> Load Balancer -> http load balancer
* HTTP(S) load balancing supports both IPv4 and IPv6 addresses for client traffic. Client IPv6 requests are terminated at the global load balancing layer and then proxied over IPv4 to your backends.
![1680733185396](image/README/1680733185396.png)
* This configuration means that the load balancer attempts to keep each instance of us-central1-mig at or below 50 requests per second (RPS).
![1680733334805](image/README/1680733334805.png)
* This configuration means that the load balancer attempts to keep each instance of europe-west1-mig at or below 80% CPU utilization.


### Task 6. Stress test the HTTP load balancer
```sh
# check to see if the load balancer is up
LB_IP=[LB_IP_v4]
while [ -z "$RESULT" ] ;
do
  echo "Waiting for Load Balancer";
  sleep 5;
  RESULT=$(curl -m1 -s $LB_IP | grep Apache);
done

# go to http://[LB_IP_v4]

# Navigation -> Networking -> Load Balancing -> Monitoring
```
### Task 7. Review
* inc compute engine create instance templates
you need to create an image keep the boot disk >
create health check> create instance template > create instance group
> create load balancers
  > attach instance group & health cheks to load balancer


## Lab Review: Configuring an HTTP Load Balancer with Autoscaling

## Cloud CDN
* caches content
* setup with your load balancer
* be sure not to cache dynamic html or api responses

## SSL proxy load balancing
* supports IPv4 -v6
* cert mgnt, security patching, ssl polices, inteligent routing
* choose SSL or TCP
* terminates at the load balancing layer

## TCP proxy load balancing
* terminates at the load balancing layer

## Network load balancing
* regional non-proxy
* only work with vms in the same region
* uses forwarding rules
* target pool resoouces instances, work with forwarding rules
  * each target pool can have one health check and 50 instances

## Internal load balancing
* if you dont need external IP, you have lower latency
* uses google andromeda (virtualization spec)
* database tier and netwrok tier never get exposed

## Lab Intro: Configuring an Internal Load Balancer
https://app.pluralsight.com/lti-integration/redirect/80d399cc-38b6-4f02-8df7-cb5ba016717f?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Felastic-google-cloud-infrastructure-scaling-automation-10&contextTitle=Elastic+Google+Cloud+Infrastructure%3A+Scaling+and+Automation

## Lab: Configuring an Internal Load Balancer
* Easy to deploy network threat detection with Google Cloud IDS from Cloud NAT

### Task 1. Configure internal traffic and health check firewall rules

### Task 2. Create a NAT configuration using Cloud Router

### Task 3. Configure instance templates and create instance groups
Hostname</h2>Server Hostname:
 instance-group-2-q5wp
 <h2>Server Location</h2>Region and Zone: us-central1-b

### Task 4. Configure the internal load balancer
* they use TCP load balancing

### Task 5. Test the internal load balancer

## Lab Review: Configuring an Internal Load Balancer

## Choosing a load balancer
* HTTPS TCP and Netowrk supprt IPv6
* load balancer acts as a proxy to conver IPv6 to IPv4
* Internal HTTP is beta

# Infrastructure Automation

## Terraform
* its good to know which code created which vm
* IAC
  * quick provisions and removing on infrastrcutre
  * part of a CI/CD pipeline
* terrafrom lets you provisons google cloud resources
  * hasicorp is a language lets you do this
  * terraform deploys resources in parallel
  * resources are infrastucture object
![1680749305464](image/README/1680749305464.png)
```sh
# makes sure google cloud init is enabled
terraform init

# start your infrastucture
terraform apply
```
## Lab Intro: Automating the Infrastructure of Networks Using Terraform
https://app.pluralsight.com/lti-integration/redirect/c6fa12a9-5dfc-43e3-9081-b4cf0cfaa043?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Felastic-google-cloud-infrastructure-scaling-automation-10&contextTitle=Elastic+Google+Cloud+Infrastructure%3A+Scaling+and+Automation

## Lab: Automating the Deployment of Infrastructure Using Terraform
* terraform comes with a cloud instances
* you can use terraform module registry
[Terraform templates](https://registry.terraform.io/browse/providers)
### Task 1. Set up Terraform and Cloud Shell

### Task 2. Create mynetwork and its resources

### Task 3. Verify your deployment

### Task 4. Review

## Lab Review: Automating the Infrastructure of networks using Terraform

## Google Cloud Marketplace
* 3rd party service from other who developed everything already on terraform
* go to market place and

## Demo: Launch on Google Cloud Marketplace
* if we need to deploy infrastructe in a meanigful way use teraform

# Managed Services

## BigQuery
* servless petabyte scale database warehouse
*

## Dataflow
* wide varierty of data processing patterns
* servless
* batch and stream processing
* tied to stackdriver so you can monitor quality services

## Dataprep
* exploring and clean data for machine learning big data prcessing

## Dataproc
* if you depend on apache or hands on use dataproc
* if you want hands off use dataflow

## Demo: Dataproc
* its a cluster of vm machines
* you submit jobs to dataproc cluster for it to do

## Module Review

## Course Series Review
