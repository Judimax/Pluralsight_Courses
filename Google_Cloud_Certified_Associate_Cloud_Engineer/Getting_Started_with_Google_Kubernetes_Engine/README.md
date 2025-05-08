# Introduction to Google Cloud
* hierachal resources
## Cloud Computing and Google Cloud
* on demand self service
* no human intervential
* global benefit from economies of scale
* pay for what u use or reserve
* GKE  runs containerzed apps managed for you
* Cloud Functions are like webhooks

## Resource Management
* phyiscal and virtual region
* multi-region, region zones,data center
* deploy application along multiple zones
* google responds to users use edge networks
* there are global,regional and zonal resouces
* gke clusters are regional
* you can get an org for free using google cloud identity

## Billing
* there are Rate quota
  * no more 1000 requests per 100 seconds

* __allocation quota__
  * 5 netowrks per project

* protects users from unforseen spikes in usage

## Interacting with Google Cloud
* shared security model your responsiblity are for the resource hierachy

## Computing Options

## Lab Intro

## Pluralsight: Getting Started with GCP and Qwiklabs

## Lab: Accessing the Google Cloud Console and Cloud Shell

## Lab solution

# Introduction to Containers and Kubernetes
* create container using cloud build
## Introduction to Containers
* back then an app had to run on a computer, computer used to do only one thing
* running apps in 1 vm, resources can be used up by applications and starve
* vms are slow to startup
* so just startup the userspace, because you are starting and stopping the process, not the computer itself


## Containers and Container Images
* one container builds the software, another runs the application
* want to store data other than the container itself
* alpine is very small container
* GCR is similar to cloud IAM
* cloud build can get source code from many repo


## Lab Intro
https://app.pluralsight.com/lti-integration/redirect/fb356ff2-aec8-45af-9d93-6fa2d8c8dedb?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fgetting-started-google-kubernetes-engine-10&contextTitle=Getting+Started+with+Google+Kubernetes+Engine

## Lab: Working with Cloud Build

### Lab setup

### Task 1. Confirm that needed APIs are enabled
* check for cloud build and container registry

### Task 2. Building containers with DockerFile and Cloud Build

### Task 3. Building containers with a build configuration file and Cloud Build
* once the image is build head over to the container registry
Navigation > Cloud Build > History
### Task 4. Building and testing containers with a build configuration file and Cloud Build
* run the other lab you can also see the build failing in cloud builld

## Lab Solution

## Introduction to Kubernetes
* use decalaritve config so you always know whats going on
* imperative changes are for quick updates
* portability

## Introduction to Google Kubernetes Engine
* GKE managed services
* hepls deploy manage and scale
* easy to bring kubernetes into the clould
* scale easy, mimanl footprint
  * auto repair, gracefully shutdown node and restarts
  * kubernetes dashboard is hard to make secure


## Compute Options in Detail
* compute engine
  * 160 GPU 3 terabytes of memory
  * put apps behind load balancing and autoscaling
  * lift and shift from on prem to compute
  * long lived,vm
* app engine
  * use code and app engine deploys for you
  * app engine has sdk with visual studio go
  * websites, mobile apps and REST API
* kubernetes engines
  * cluster scalining,
  * builting in with stackdriver, cloud build.
  * admin of container workloads
* cloud run
  * scales up and down to 0
  * built by knative
  * stateless containers
  * build any language
  * stateless containers
* cloud functions
  * only charged when code runs
  * there is free tier
  * microservice, integrate w/ 3rd party apps


# Kubernetes Architecture

## Kubernetes Concepts
* pods can use shared storage
* kubernetes control plane endlessly compares reality to what is declared


## Kubernetes Control Plane
* other vm are nodes
* control plan
  * launches pods
  * kubectl
  * kube api server, also does auth
  * etcd - clusters database
  * kube-scheduler lanuches pods onto nodes
  * kube-controller manager
    * will try to get cluster to desired states
    * 3 nginx objects, can be gather to be called a deployment
  * kube cloud-manager
    * if you lanuch on cloud it gets 3rd party deps to make it work
  * node
    * kubelet is a node
    * kube-proxy had IP tables
* GKE creates nodes, where admins create them
* GCP  never uses a platform older than the CPU
* node pools are a GKE feature
  * subset of nodes that share certain configs
* if a node has 15 GB, there is stuff reserved
  for GKE so things can rn
* 3 nodes plus control plannes in a zonal
  * if a zone goes down use regional
  * cannot convert regional to zonal cluster and vice versa
  * nodes can access the internet w/o external IP address


## Google Kubernetes Engine Concepts
* kubeadm
  * if a node fails or needs maintenacne run kubeadm and thats it
  * 416 vCPU cores
## Kubernetes Object Management
* mainfest
  * yaml file
  * save yaml in version control
  * only one ojbect can have a unique name that can go up to 253
  * you can select resources by its labels
  * pods are supposed to be disopable
  * declaire controller object to manage these ephermal pods
  * can split phyiscal node into vitrual node w/ namespaces
    * allow for resources quota (not GCP quotas)
    * keep everything unique
## A note about Deployments and ReplicaSets

## A note about Services
* Services provide load-balanced access to specified Pods.
  * ClusterIP:
  * NodePort
  * LoadBalancer:

## Controller objects to know about
* ReplicaSets
● Deployments
● Replication Controllers
● StatefulSets
● DaemonSets
● Jobs
## Lab Intro
https://app.pluralsight.com/lti-integration/redirect/5794b7ca-6126-414e-add9-c2da8c57059b?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fgetting-started-google-kubernetes-engine-10&contextTitle=Getting+Started+with+Google+Kubernetes+Engine

## Lab: Deploying Google Kubernetes Engine

### Task 1. Deploy GKE clusters

### Task 2. Modify GKE clusters

### Task 3. Deploy a sample workload
* Kubernetes Workflow -> worloads -> deploy
default deploys 3 pods with a each container running nginx

### Task 4. View details about workloads in the Google Cloud Console

## Lab solution

## Migrate for Anthos introduction
* anthos move non container to container environments
* completed in 10 mins
* one move or steram till app is live

## Migrate for Anthos Architecture
* configure cluster,
  add migration seouce
  generate and review plan
  generate artifacts
  test and deploy

## Migration Path
* must be a gke admin
* must setup network for anthos for compute and GKE
  * install migctl
## Migrate for Anthos Installation

# Introduction to Kubernetes Workloads

## The kubectl Command
* transforms command line to API
* must have proper auth
* must config cluster first
$HOME/.kube/config
* gcloud controls not auth
* kubectl cant create or change clusters


## Deployments
* deploymenets mean desired state of pods
* you can rollback pods
* designed for stateless applications
* replcica set is a controller that makes sure a certain of amnt of pods are runing at any given time
* 3 states
  * progress,complete,failed
    * failed
      * resources quote, permissions, runtime bug


## Ways to Create Deployments


## Services and Scaling

## Updating Deployments

## Blue-Green Deployments
* disadvantage deouble resouces are used

## Canary Deployments
* gradual traffic move from one deployment to antoher
* miminal users see if the canary development is working and if there are issues back to the original
* a/b testing is business prediction based on data
  * effectiveness of functionality in an application
* shadow testing, new version is hidden from user version
* no full rollout occurs until the application meets requirements

## Managing Deployments
* you can pause kubectl rollout

## Lab Intro
https://app.pluralsight.com/lti-integration/redirect/ebbcdd96-488b-4cf2-860e-2d7501c7fe26?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fgetting-started-google-kubernetes-engine-10&contextTitle=Getting+Started+with+Google+Kubernetes+Engine


## Lab: Creating Google Kubernetes Engine Deployments

### Task 1. Create deployment manifests and deploy to the cluster

### Task 2. Manually scale up and down the number of Pods in deployments
unlike vm containers immediately scale up and down

### Task 3. Trigger a deployment rollout and a deployment rollback
* A deployment's rollout is triggered if and only if the deployment's Pod template (that is, .spec.template) is changed
### Task 4. Define the service type in the manifest

### Task 5. Perform a canary deployment
* I could not see the traffic going to the canary deployment because the nginx server was not working

## Pod Networking
* each pods have their own IP
* they connect via root node,
* they connect with the world because the root node has a nic
* everything is a object to kuberentes

## Volumes
* volumne directory accessible to all nodes in a pod
  * some or ephermal
    * configmap user to inject data into application
    * secrets volumne are sensitive
    * dowardAPI containers can learn about their environment so it can make unique names
    * created when a pod is assigned to a node
    * data is safe form container crashers
  * some are persistent
    * exists independent
    * can be created by cluster admin

## Volume Types
*

## The PersistentVolume abstraction
* PV
* PersistentVolumneClaims
  * enable pods access to PersistentVolumns
* managed by kubernetes
* manualy or dyanimically provisioned

## Lab Intro
https://app.pluralsight.com/lti-integration/redirect/e6d20f62-5725-4240-90d7-1c3e21b61cc4?originUrl=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fgetting-started-google-kubernetes-engine-10&contextTitle=Getting+Started+with+Google+Kubernetes+Engine

## Lab: Configuring Persistent Storage for Google Kubernetes Engine

### Lab setup

### Task 1. Create PVs and PVCs

### Task 2. Mount and verify Google Cloud persistent disk PVCs in Pods

### Task 3. Create StatefulSets with PVCs

### Task 4. Verify the persistence of Persistent Volume connections to Pods managed by StatefulSets
