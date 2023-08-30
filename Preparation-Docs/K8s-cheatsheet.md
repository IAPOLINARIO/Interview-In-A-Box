# Kubernetes Cheat Sheet for Interview

## Table of Contents

1. [Basic Commands](#basic-commands)
2. [Cluster Info](#cluster-info)
3. [Nodes](#nodes)
4. [Pods](#pods)
5. [Services](#services)
6. [Deployments](#deployments)
7. [ConfigMaps & Secrets](#configmaps--secrets)
8. [Networking](#networking)
9. [Volumes & Persistent Storage](#volumes--persistent-storage)
10. [Debugging](#debugging)

---

## Basic Commands

### Install `kubectl`

[[[bash

# macOS

brew install kubectl

# Linux

apt-get install kubectl

# Windows

choco install kubectl
]]]

### Check `kubectl` version

[[[bash
kubectl version]]]

---

## Cluster Info

### Get cluster info

[[[bash
kubectl cluster-info]]]

### Describe the cluster

[[[bash
kubectl describe cluster]]]

---

## Nodes

### List all nodes

[[[bash
kubectl get nodes]]]

### Describe a specific node

[[[bash
kubectl describe node [NODE_NAME]]]]

### Update a node's labels

[[[bash
kubectl label nodes [NODE_NAME] [KEY]=[VALUE]]]]

---

## Pods

### List all Pods

[[[bash
kubectl get pods]]]

### Create a Pod

[[[bash
kubectl create -f [POD_YAML_FILE]]]]

### Describe a specific Pod

[[[bash
kubectl describe pod [POD_NAME]]]]

### Delete a Pod

[[[bash
kubectl delete pod [POD_NAME]]]]

---

## Services

### List all Services

[[[bash
kubectl get svc]]]

### Create a Service

[[[bash
kubectl create -f [SERVICE_YAML_FILE]]]]

### Describe a specific Service

[[[bash
kubectl describe svc [SERVICE_NAME]]]]

---

## Deployments

### List Deployments

[[[bash
kubectl get deployments]]]

### Create a Deployment

[[[bash
kubectl create -f [DEPLOYMENT_YAML_FILE]]]]

### Scale a Deployment

[[[bash
kubectl scale deployment [DEPLOYMENT_NAME] --replicas=[NUM]]]]

---

## ConfigMaps & Secrets

### Create a ConfigMap

[[[bash
kubectl create configmap [CONFIG_MAP_NAME] --from-file=[FILE_PATH]]]]

### Create a Secret

[[[bash
kubectl create secret generic [SECRET_NAME] --from-file=[FILE_PATH]]]]

---

## Networking

### List Network Policies

[[[bash
kubectl get networkpolicies]]]

### Create a Network Policy

[[[bash
kubectl create -f [NETWORK_POLICY_YAML]]]]

---

## Volumes & Persistent Storage

### Create a Persistent Volume

[[[bash
kubectl create -f [PV_YAML_FILE]]]]

### Create a Persistent Volume Claim

[[[bash
kubectl create -f [PVC_YAML_FILE]]]]

---

## Debugging

### Get logs for a Pod

[[[bash
kubectl logs [POD_NAME]]]]

### Exec into a running Pod

[[[bash
kubectl exec -it [POD_NAME] -- /bin/bash]]]

### Get cluster events

[[[bash
kubectl get events]]]

### Debugging commands

[[[bash

# Check current resource consumption

kubectl top node

# Check component statuses

kubectl get componentstatuses
]]]

Good luck, you've got this!
