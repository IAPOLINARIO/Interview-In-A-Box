# K8s Interview Guide 🚀

## Table of Contents

1. [What is Kubernetes? Can you explain its architecture?](#what-is-kubernetes-can-you-explain-its-architecture)
2. [What's the difference between a Pod and a Deployment in Kubernetes?](#whats-the-difference-between-a-pod-and-a-deployment-in-kubernetes)
3. [How does Kubernetes handle scaling?](#how-does-kubernetes-handle-scaling)
4. [Explain what a Kubelet is](#explain-what-a-kubelet-is)
5. [Can you describe what a ConfigMap is and how you would use it?](#can-you-describe-what-a-configmap-is-and-how-you-would-use-it)
6. [How do Kubernetes Services work?](#how-do-kubernetes-services-work)
7. [What is a Kubernetes Namespace and why would you use it?](#what-is-a-kubernetes-namespace-and-why-would-you-use-it)
8. [What is the role of etcd in a Kubernetes cluster?](#what-is-the-role-of-etcd-in-a-kubernetes-cluster)
9. [How does Kubernetes manage secrets?](#how-does-kubernetes-manage-secrets)
10. [What's the difference between a StatefulSet and a ReplicaSet?](#whats-the-difference-between-a-statefulset-and-a-replicaset)
11. [What is Helm and why would you use it?](#what-is-helm-and-why-would-you-use-it)
12. [Can you explain Kubernetes RBAC?](#can-you-explain-kubernetes-rbac)
13. [What is ingress in Kubernetes?](#what-is-ingress-in-kubernetes)
14. [What are Kubernetes Operators and how do they work?](#what-are-kubernetes-operators-and-how-do-they-work)
15. [What's a DaemonSet?](#whats-a-daemonset)
16. [Explain the concept of 'affinity' in Kubernetes](#explain-the-concept-of-affinity-in-kubernetes)
17. [How do you monitor a Kubernetes cluster?](#how-do-you-monitor-a-kubernetes-cluster)
18. [What are Kubernetes labels and selectors?](#what-are-kubernetes-labels-and-selectors)
19. [How do you update an application running in a Kubernetes cluster?](#how-do-you-update-an-application-running-in-a-kubernetes-cluster)
20. [What are the security best practices in Kubernetes?](#what-are-the-security-best-practices-in-kubernetes)
21. [Kubernetes Master Services](#kubernetes-master-services)
22. [Kubernetes Node Services](#kubernetes-node-services)

---

## 1. What is Kubernetes? Can you explain its architecture?

Kubernetes is an open-source container orchestration platform that automates various aspects of application deployment, scaling, and management. It allows you to deploy applications in a reliable and scalable way, abstracting the underlying infrastructure.

### Architecture

Kubernetes has a Master-Worker architecture:

- **Master Node (Control Plane)**: Comprises components like `etcd`, `kube-apiserver`, `kube-scheduler`, and `kube-controller-manager`.
  - **etcd**: Distributed key-value store, acts as the source of truth for the cluster.
  - **kube-apiserver**: Entry point for all RESTful commands, exposing Kubernetes API.
  - **kube-scheduler**: Responsible for pod scheduling on nodes.
  - **kube-controller-manager**: Manages the lifecycle of various controllers like Deployment and StatefulSet controllers.
- **Worker Nodes**: These run your applications. Each node runs a special pod called `kubelet` that communicates with the master node, and a container runtime like Docker or containerd.

---

## 2. What's the difference between a Pod and a Deployment in Kubernetes?

- **Pod**: The smallest deployable unit in Kubernetes, a Pod can house one or multiple containers. Containers in the same Pod share the same network IP, port space, and storage, akin to a single machine.
- **Deployment**: Manages the desired state of Pods. You specify the state in a YAML file, and the Deployment controller ensures that the state is maintained, handling updates and rollbacks.

---

## 3. How does Kubernetes handle scaling?

Kubernetes provides two types of scaling:

1. **Manual Scaling**: You can manually change the number of replicas in a Deployment YAML and apply it, or use the `kubectl scale` command.

2. **Automatic Scaling**: Kubernetes provides a resource called `Horizontal Pod Autoscaler` (HPA), which can scale the number of pod replicas dynamically based on metrics like CPU or memory usage.

---

## 4. Explain what a Kubelet is

Kubelet is an agent that runs on each node in a Kubernetes cluster. It interacts with the container runtime to ensure that the containers described in the PodSpec are up and running. If a Pod fails, the Kubelet tries to restart it.

---

## 5. Can you describe what a ConfigMap is and how you would use it?

ConfigMap is a Kubernetes object used to store configuration data in key-value pairs. You can inject these into your Pods as environment variables, command-line arguments, or as files in a volume.

---

## 6. How do Kubernetes Services work?

Kubernetes Services act as a load balancer for Pods. They provide a stable endpoint for other services or for external traffic to connect to. Types of Services include:

- **ClusterIP**: Exposes the service within the cluster only.
- **NodePort**: Exposes the service on a static port on each Node’s IP.
- **LoadBalancer**: Exposes the service externally using a cloud provider’s load balancer.

---

## 7. What is a Kubernetes Namespace and why would you use it?

Namespaces in Kubernetes serve as a method for dividing cluster resources among multiple users, projects, or even departments. They offer a scope for names, segregate resources, and can have their own policies.

---

## 8. What is the role of etcd in a Kubernetes cluster?

`etcd` is a distributed key-value store used to store all data used to manage the cluster. It is a crucial part of the Control Plane and is responsible for storing the configuration data and the state of the system.

---

## 9. How does Kubernetes manage secrets?

Kubernetes Secrets are used for storing sensitive data like passwords, SSH keys, or tokens. Unlike ConfigMaps, Secrets are encrypted at rest and can be accessed only by authorized Pods.

---

## 10. What's the difference between a StatefulSet and a ReplicaSet?

- **StatefulSet**: Used for workloads that require stable network identifiers and persistent storage.
- **ReplicaSet**: Ensures that a certain number of replicas for a Pod are running at any given time but doesn't guarantee the state or uniqueness.

---

---

## 11. What is Helm and why would you use it?

Helm is a package manager for Kubernetes that simplifies the deployment and management of applications. With Helm charts, you can define, install, and upgrade complex Kubernetes applications.

---

## 12. Can you explain Kubernetes RBAC?

Role-Based Access Control (RBAC) in Kubernetes is used for defining what actions a user, or a group of users, can perform. The main components are Roles, RoleBindings, ClusterRoles, and ClusterRoleBindings.

---

## 13. What is ingress in Kubernetes?

Ingress is an API object that provides HTTP and HTTPS routing to services within the cluster. It can provide features like SSL termination, path-based routing, and host-based routing.

---

## 14. What are Kubernetes Operators and how do they work?

Operators are custom controllers that manage complex stateful applications. They extend the Kubernetes API to create, configure, and manage instances of complex stateful applications on behalf of a Kubernetes user.

---

## 15. What's a DaemonSet?

A DaemonSet ensures that a copy of a Pod runs on all nodes in the cluster, or on a subset of nodes based on selection criteria. This is useful for running cluster-wide operations like logging and monitoring.

---

## 16. Explain the concept of 'affinity' in Kubernetes

Affinity rules define how closely related a Pod is to other Pods or to specific nodes. Affinities are set using labels and selectors, and can be either "hard" (required) or "soft" (preferred).

---

## 17. How do you monitor a Kubernetes cluster?

Monitoring can be accomplished using tools like Prometheus for metrics collection, Grafana for visualization, and logging solutions like ELK (Elasticsearch, Logstash, Kibana) or Loki.

---

## 18. What are Kubernetes labels and selectors?

Labels are key-value pairs attached to Kubernetes objects. Selectors are used to filter objects based on their labels, for operations like scaling or rolling updates.

---

## 19. How do you update an application running in a Kubernetes cluster?

Updates can be done using a rolling update strategy, where Pods are gradually replaced with new versions. Alternatively, you can use blue-green or canary deployment strategies to route a subset of users to the new version for testing before full rollout.

---

## 20. What are the security best practices in Kubernetes?

- Limit access using RBAC
- Use Network Policies to restrict traffic between Pods
- Enable PodSecurityPolicies
- Rotate and properly manage secrets
- Enable logging and monitoring for anomalous behavior

---

## 21. Kubernetes Master Services

This section will delve into more details about the services running on the Master node.

---

## 22. Kubernetes Node Services

This section will cover services like kubelet, kube-proxy and more that run on Worker nodes.
