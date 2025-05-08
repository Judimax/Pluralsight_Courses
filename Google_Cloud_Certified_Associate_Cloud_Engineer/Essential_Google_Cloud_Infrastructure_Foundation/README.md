# Introduction
* GCP  is part of a larger ecosystem
* more sup

![1680631129805](image/README/1680631129805.png)
* get familiar with infrastructure
* Compute Engine allows for backend, Iaas
* GKE -backend docker
* App Engine  backend PaaS
* Cloud Function - backend function
* Cloud Run

* Compute engine is main focus
* netowrk,compute,iam,
* terraform Infrastructure
* after enroll in
  * Essential Cloud Infrastructure:
Core Services and Elastic Cloud Infrastructure: Scaling and Automation

# Inteacting with Google Cloud

## Module Overview

## Using Google Cloud
![1680631869583](image/README/1680631869583.png)
* there is app API and admin API's

## Lab Intro: Working with the Google Cloud Console and Cloud Shell

## Getting Started with GCP and Qwiklabs
* make sure of the project ID name and number

## Lab: Working with the Google Cloud Console and Cloud Shell
https://app.pluralsight.com/lti-integration/redirect/6ca1b866-ea8e-4256-ba3d-5d5a124e5b7f?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fessential-google-cloud-infrastructure-foundation-6&contextTitle=Essential+Google+Cloud+Infrastructure%3A+Foundation

Temporary Compute Engine VM
Command-line access to the instance via a browser
5 GB of persistent disk storage ($HOME dir)
Pre-installed Cloud SDK and other tools
gcloud: for working with Compute Engine and many Google Cloud services
gsutil: for working with Cloud Storage
kubectl: for working with Google Kubernetes Engine and Kubernetes
bq: for working with BigQuery
Language support for Java, Go, Python, Node.js, PHP, and Ruby
Web preview functionality
Built-in authorization for access to resources and instances


```sh
# to make a bucket in cloud shell
gsutil mb gs://<BUCKET_NAME>

# to copy a bucket
gsutil cp [MY_FILE] gs://[BUCKET_NAME]

# to see all available regions in google cloud
gcloud compute regions list
```
us-west4
files persists in cloud shell but not env vars




## Lab Review: Working with the Google Cloud Console and Cloud Shell

## Lab Intro: Infrastructure Preview
https://app.pluralsight.com/lti-integration/redirect/17c5f429-a622-49cf-a3a8-61d0397129db?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fessential-google-cloud-infrastructure-foundation-6&contextTitle=Essential+Google+Cloud+Infrastructure%3A+Foundation
## Lab: Infrastructure Preview
* go to marketplace and find jenkins by bitname
* for one VM instance is $20.00
* for sustained use it could be $13 a month
* deploymenet manager setups the deployment

admin user user
admin pass - b4pfgkQbP9hB
* its using resouces from Rutime Configurator service which is in beta
* go to compute engine go to the external url and type in the user name and password

## Lab Review: Infrastructure Preview

## Demo: Projects
* recoures are relate via project

## Module Review

# Virtual Networks
## Module Overview


## Virtual Private Cloud
* Overview
* Region - speciifc location where you run yr project
* Pop - Points of Prencese - where google is connected to the internet
![1680638927258](image/README/1680638927258.png)
* you chose a network type

## Projects, networks, and subnetworks
* default quota for a network is 5 networks
* default - allowed ingress
* auto mode- subnets get auto created
* custom-
  decide which regions and subnets you wants
  * custom cannot change to auto

* if 2 vm are on the same network they can communciate via internal IP, else they communicate via exteral IP
* every subnet has 4 virtual machines in the IP address range
* new subnets must not overlay
![1680639158921](image/README/1680639158921.png)

## Demo: Expand a Subnet
* /29 address creates 8 address 4 used by GCP
*

## IP addresses
![1680640051574](image/README/1680640051574.png)
* GAE and GKE ussually gets these
* to bring yr own IP you need a /24 block or larger
* you only have a limit of 15000 IP addresses becayse if hardware limitations

## Demo: Internal and external IP

## Mapping IP addresses
* external address is unknown to the VM
![1680640257069](image/README/1680640257069.png)
* external IP is maintain in a domain table
* 100%, uptime SLA for domain managed in GCP
* configure multiple


## IP addresses for default domains
* [All IP addresses owned by google](https://s2.pluralsight.com/links/09cb7d9d-302b-4464-8853-1d15e9e60cc2_f3a07a842139ce70d69a89ca7af548f0.pdf)


## Routes and firewall rules
* every netwrok has routes and a default route,
* firewall rules must allow packets
* a route is create when a network, on subnet is created (allowing VM to communicate)
![1680641064832](image/README/1680641064832.png)
![1680641089671](image/README/1680641089671.png)
* can use Egress to protect against internal VM
![1680641148974](image/README/1680641148974.png)

## Pricing
__ingress__ - inboard traffic
__egress__ - outbound traffic
![1680641184735](image/README/1680641184735.png)
![1680641201529](image/README/1680641201529.png)
* charged more for not use static IP adddrees
* use GCP pricing calculator

  * email or save it
  * can see pricing basd on regions


## Lab: VPC Networking
* In other words, without a VPC network, you cannot create VM instances, containers, or App Engine applications.
*  VPC network as similar to a physical network, except that it is virtualized within Google Cloud.
* create a VPC in this diagram
![1680641441647](image/README/1680641441647.png)
* the VPC here as 25 subnets
* Each subnet is associated with a Google Cloud region and a private RFC 1918 CIDR block for its internal IP addresses range and a gateway.
* each subnet has its own route and there is one for the gateway to the internet, probably just a way of seperating data
* each network has a virtual firewall
* auto VPC auto create subnets, but you dont have full control, subnets are even created in any region
* when you make a vpc a network must exist in the VPC that you have made
* external IP are ephermeral
* ssh connection work because of the firewall rule on the VPC networok and. Compute Engine adds generated key to project or instance metdata, or if you use OS Login, the key gets stored on the user acct
* convert to custom mode so new subnets dont get generated with new regions, see how the subnets overlap in CIDR ranges we want custom because we dont want that
```sh
# to create a VPC from cloud shell
gcloud compute networks create privatenet --subnet-mode=custom

# to create a subnet
gcloud compute networks subnets create privatesubnet-us --network=privatenet --region=us-east4 --range=172.16.0.0/24

# to list available networks
gcloud compute networks list

# to list available subnets
gcloud compute networks subnets list --sort-by=NETWORK

# to create a firewall rule
gcloud compute firewall-rules create privatenet-allow-icmp-ssh-rdp --direction=INGRESS --priority=1000 --network=privatenet --action=ALLOW --rules=icmp,tcp:22,tcp:3389 --source-ranges=0.0.0.0/0

# to create a vm with a certain netwrok and subnet
gcloud compute instances create privatenet-us-vm --zone=us-east4-a --machine-type=e2-micro --subnet=privatesubnet-us --image-family=debian-11 --image-project=debian-cloud --boot-disk-size=10GB --boot-disk-type=pd-standard --boot-disk-device-name=privatenet-us-vm

# list all vm instances by zone
gcloud compute instances list --sort-by=ZONE
```

## Lab Review: VPC Networking

## Common network designs
* putting resournces in different regions helps with avaliblity
* you can increase zones while keeping vm to same subnet
* try to keep cloud vm to internal IP only best practice

## Lab Intro: Implement Private Google Access and Cloud NAT
Configure a VM instance that doesn't have an external IP address
Connect to a VM instance using an Identity-Aware Proxy (IAP) tunnel
Enable Private Google Access on a subnet
Configure a Cloud NAT gateway
Verify access to public IP addresses of Google APIs and services and other connections to the internet

## Lab: Implement Private Google Access and Cloud NAT
* create your netwrok and firewall
* this is the netwrok CIDR - 35.235.240.0/20
*  IAP connections come from a specific set of IP addresses (35.235.240.0/20).
* The default setting for a VM instance is to have an ephemeral external IP address. This behavior can be changed with a policy constraint at the organization or project level.
* Cloud IAP enables context-aware access to VMs via SSH and RDP without bastion hosts.
```sh
# ssh into machine via IAP
gcloud compute ssh [NAME OF VM] --zone [ZONE OF VM] --tunnel-through-iap

# if any issues you may need to activate services use the --troubleshoot flag w/ or w/o the tunnel flag

```



### To connect via private google access
bucket name: lab-bucket-1123123qdf
* pinguing google from the internal machine should not work because it doesnt have external IP

* cant get a bucket to vm with missing external IP  buecause private google access is not enabled on the subnet you can enable it by editing the subnet and finding the option
lab-bucket-1sadadad
* cant also do apt-get update

### Configure a Cloud NAT gateway
* by configuring cloud NAT to the vm with missing IP, now it has access to the internet

### Task 4. Configure and view logs with Cloud NAT Logging
![1680652387629](image/README/1680652387629.png)
![1680652714349](image/README/1680652714349.png)

cloud NAT logs are generated
When a network connection using NAT is created.
When a packet is dropped because no port was available for NAT.

then connect via iap and run some network commands

## Lab Review: Implement Private Google Access and Cloud NAT

## Module Review

# Virtual Machines
## Module Overview

## Compute Engine
* max is 100 gpbs which is 224 cores
* there is storage

## Demo: Create a VM
* different locations means different monthly cost
* can take a snapshot of the vm, move to a

## VM access and lifecycle
![1680653956871](image/README/1680653956871.png)
* if vm crashes you can restart it by default
* when you provision a vm there is cost associated
![1680654042586](image/README/1680654042586.png)


## Lab Intro: Creating Virtual Machines

## Lab: Creating Virtual Machines

### Task one Create a Utility virtual machine
name utility-machine-0
* Notice that you cannot change the machine type, the CPU platform, or the zone.
* You can add additional disks and you can also determine whether the boot disk is deleted when the instance is deleted.
* cant create image from running boot disk
*  disable Delete boot disk when instance is deleted
* cant do non-preemptible instance into a preemptible one. This choice must be made at VM creation
* these vm can fail due to outages
* changes expesically network changes can take several minutes to implement

### Create a windows virtual machine
* its cheaper to run windows machines than linux machines
* there is no ssh for windows vm just RDP
* you can install RDP client from the chrome web store

u: student_02_035cb6350
p: H*s6e:q&ha(wFCO

### Task 3. Create a custom virtual machine
```sh
# To see information about unused and used memory and swap space on your custom VM, run the following command:
free

# To see ram details
sudo dmidecode -t 17

# To verify the number of processors, run the following command:
nproc

# To see details about the CPUs installed on your VM, run the following command:
lscpi


```

## Lab Review: Creating Virtual Machines

## Compute options
![1680656082324](image/README/1680656082324.png)
![1680656116771](image/README/1680656116771.png)
![1680656155631](image/README/1680656155631.png)
![1680656205654](image/README/1680656205654.png)

* .9 - 6.5 GB per vCPU
* you can get more via extended CPU



## Compute pricing
* you get charged per one minute 30 sec 1 min 1 min 1 one sec
* preemptible run for only 24 hours
* spot instances do not have a limit
* if you use vm for 100$ of the month you get 30% discount
![1680656378544](image/README/1680656378544.png)

## Special compute configurations
* try to make your vmrun on preemtible vm
  * no live migrate, no auto restart

__spot vm__
  compute engine may not stop spot vms for a certain peroid of time

__sole tenat__ -are like aws dedicated EC2 instances

__shield vm__ -verified your machine is not compromised by root kits

__confidentials vm__ allow you to encrpy data in use
  * in line memory encrpytion does not compromise performace

## Images
* there are public and custom images,
* SQL Server images are chaged per min after 10 mins
__machine image__ for many system maintenance
  * they are still in beta

## Disk options
__boot disk__ - vm comes w/ a single root persistant disk
  hdd and sdd based on cost and perfomance
  can convert dis to read only so you dont have to replicate data
![1680656791092](image/README/1680656791092.png)
GCP encrpyts all data at rest

![1680657276421](image/README/1680657276421.png)
![1680657293807](image/README/1680657293807.png)

## Common Compute Engine actions
* snapshots can be used to migrate data between zones
* snaptshots can be restored to a new persistent disl
* you can only grow disk not shrink

## Lab Intro: Working with Virtual Machines

## Lab: Working with Virtual Machines

### Task 1 Create the VM
* setup game server
* also attach high performance 50-GB persistent ssd to the instance
* to get a static external IP
  > external IP -> create new ip address and give it a name

### Task 2 Prepare the data disk
```sh
# to format a disk
sudo mkfs.ext4 -F -E lazy_itable_init=0,\
lazy_journal_init=0,discard \
/dev/disk/by-id/google-minecraft-disk

# to mount
sudo mount -o discard,defaults /dev/disk/by-id/google-minecraft-disk /home/minecraft
```

### Task 3. Install and run the application
```sh
sudo apt-get update
sudo apt-get install -y default-jre-headless
cd /home/minecraft
sudo wget https://launcher.mojang.com/v1/objects/d0d0fe2b1dc6ab4c65554cb734270872b72dadd6/server.jar

# init minecraft server
sudo java -Xmx1024M -Xms1024M -jar server.jar nogui

# you need to accept the eula by editing eula.txt from false to true
sudo apt-get install -y screen
# bring task to background so it still runs even if you log out
sudo screen -S mcs java -Xmx1024M -Xms1024M -jar server.jar nogui
# ctrl a ctrl d to detach
# to reattach
sudo screen -r mcs
```

### Task 4. Allow client traffic
* make a firewall rule with all access 0.0.0.0/0 to TCPport 25565 where minecraft server runs
* grap the external IP
* use https://mcsrvstat.us/ to see if the minecraft server is running

### Task 5. Schedule regular backups
bucket name abraca2134d-minecraft-backup

```sh
#!/bin/bash
screen -r mcs -X stuff '/save-all\n/save-off\n'
/usr/bin/gsutil cp -R ${BASH_SOURCE%/*}/world gs://${YOUR_BUCKET_NAME}-minecraft-backup/$(date "+%Y%m%d-%H%M%S")-world
screen -r mcs -X stuff '/save-on\n'

# to make a cron job'
sudo crontab -e

# FILE crontab create a backup every 4 hours
0 */4 * * * /home/minecraft/backup.sh
```
* Cloud Storage offers the Object Lifecycle Management feature to set a Time to Live (TTL) for objects,

### Task 6. Server maintenance
![1680659307263](image/README/1680659307263.png)


## Lab Review: Working with Virtual Machines

## Module Review

## Course Review

## Next Course: Essential Cloud Infrastructure: Core Services
