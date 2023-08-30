## Main Types of Deployments in Kubernetes

### 1. Rolling Updates

**How it works:**  
Rolling updates are the default strategy for updating Pods in a Deployment. Kubernetes gradually replaces the old Pods with new ones.

#### Pros:

- No downtime.
- Easy rollback if something goes wrong.

#### Cons:

- The application must be able to run with old and new versions of Pods during the transition.

#### Steps:

1. New ReplicaSet is created, with desired replicas set to zero.
2. For each new Pod created in the new ReplicaSet, one old Pod is killed.
3. This continues until the new ReplicaSet has all the desired replicas and the old ReplicaSet is scaled down to zero.

### 2. Blue-Green Deployments

**How it works:**  
Two identical environments ("Blue" for the current deployment and "Green" for the new one) are set up.

#### Pros:

- Quick rollback by just switching the service endpoint.
- Easy to do A/B testing.

#### Cons:

- Double the resources are needed as you have two environments running simultaneously.

#### Steps:

1. Deploy the new version to the Green environment.
2. Test the Green environment.
3. Switch the service endpoint from Blue to Green.

### 3. Canary Deployments

**How it works:**  
You release the new version to a subset of your users before you roll it out to everyone.

#### Pros:

- Lower risk as you're only exposing a subset of users to the new version.
- Easy rollback.

#### Cons:

- More complex to set up than rolling updates.

#### Steps:

1. Deploy the canary version alongside the stable version.
2. Gradually route more traffic to the canary.
3. Monitor performance and errors.
4. Roll back or proceed to full deployment.

---

## Kubernetes Objects for Managing Pods

### Deployments

- **What**: Higher-level abstraction over ReplicaSets to provide declarative updates.
- **Use Case**: When you need to manage the stateless application layer and want zero-downtime during updates.

### StatefulSets

- **What**: Manages deployment and scaling of a set of Pods but maintains a sticky identity for each Pod.
- **Use Case**: When you're running databases or any other stateful applications.

### ReplicaSets

- **What**: Ensures that a specified number of replicas for a Pod are running at all times.
- **Use Case**: Use directly when you need simple scaling and don't require updates or rollbacks. Mostly, it's better to use Deployments to manage ReplicaSets.

---

## Key Differences

- **Deployments vs StatefulSets**: Deployments are for stateless apps and offer easier scaling and zero-downtime updates. StatefulSets are for stateful apps, where each Pod needs a unique and persistent identity.
- **Deployments vs ReplicaSets**: Deployments add an abstraction layer over ReplicaSets to provide features like rolling updates and rollbacks, which ReplicaSets don't offer.
- **StatefulSets vs ReplicaSets**: StatefulSets offer unique, persistent identities and ordered, graceful deployment and scaling. ReplicaSets simply make sure a certain number of Pods are running.

## Journey of a Request to a Kubernetes Service

### Step 1: User Sends a Request

- The user types a URL in their browser and hits enter.
- The browser performs a DNS lookup to find the IP address associated with the URL.

### Step 2: Load Balancer

- The request hits a cloud or hardware load balancer.
- The load balancer forwards the request to one of the available Kubernetes Nodes.

### Step 3: NodePort or ClusterIP

- If the service is exposed via NodePort, the Node's IP and port handle the request.
- If the service is exposed via ClusterIP, the request is forwarded internally within the cluster.

### Step 4: kube-proxy

- `kube-proxy` on the Node performs IPTables or IPVS routing.
- It knows which Pods are part of the service due to label selectors.

### Step 5: Reaching the Pod

- The request is routed to one of the available Pods that make up the service.
- If the Pod has more than one container, the request is sent to the appropriate container based on port mapping.

### Step 6: Application Logic

- The container processes the request based on the application logic.
- Any necessary database reads/writes occur here.

### Step 7: Response

- The container generates a response and sends it back up the chain.
- It goes back through `kube-proxy`, the Node, the load balancer, and finally back to the user's browser.

### Protocols and Technologies Involved:

- HTTP/HTTPS for web requests.
- DNS for domain name resolution.
- IPTables/IPVS for packet routing.
- TCP/UDP for the underlying transport layer.
