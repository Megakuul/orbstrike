# Kubernetes Deployment

The kubernetes deployment is handled with HELM and the provided chart.

Dependencies:
 - Metrics-Server
 - Metallb (>0.13.12)
 - Longhorn (>1.5.2)


Orbstrike requires the following dependencies to be installed on the cluster:

## Prerequisites

### Metrics-Server

Some Kubernetes-Cluster have *metrics-server* enabled by default, if you don't have it you can install it like this:
```bash
# Check if metrics server is enabled
kubectl get deployment metrics-server -n kube-system

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

### Metallb

If you don't have metallb on your cluster, you can install it with:
```bash
helm repo add metallb https://metallb.github.io/metallb && helm repo update
helm install metallb metallb/metallb --namespace metallb-system --create-namespace
```

### Longhorn

If you don't have longhorn on your cluster, you can install it with:
```bash
helm repo add longhorn https://charts.longhorn.io && helm repo update
helm install longhorn longhorn/longhorn --namespace longhorn-system --create-namespace
```

For the longhorn installation, you will also need to have the *iscsid* daemon running on every node, you can set it up like this:
```bash
# Example for arch
sudo pacman -S open-iscsi

sudo systemctl enable iscsid && sudo systemctl start iscsid
```


## Orbstrike-System Installation

Install the Helm chart with:

```bash
helm install <stage> ./orbstrike-chart --set metallbNamespace="metallb-system",orchestratorSpeakerIP=<speakerIp> --namespace orbstrike-system --create-namespace
```
Note that *speakerIp* must be in CIDR format (e.g. 10.100.32.10/24).

Usually you want to use a custom *values.yaml* file, you can specify it after the *-f* or *--values* flag. The mandatory flag *orchestratorSpeakerIP* is the IP that is "exposed" by the metallb speaker.

For instructions to the specific attributes, just copy the *values.yaml* file from the *orbstrike-chart* directory, inside you will find any attribute documented and described.
```bash
cp ./orbstrike-chart/values.yaml myvalues.yaml
emacs -nw myvalues.yaml
```

## Initialization (also described in HELM output)

The redis-cluster needs to be initialized, to do this you will need to call following command from the redis cli:

```bash
# Connect to any database container
kubectl exec -it -n orbstrike-system pod/<Pod1> -- sh
# Create the initial cluster
redis-cli --cluster create <Pod1>:6379 <Pod2>:6379 <Pod3>:6379 --cluster-replicas 0
```

When you want to add another shard master you can do it like this
```bash
redis-cli --cluster add-node <Pod1>:6379 <Pod4>:6379
```
(Notice that when doing this after creation, you will also need to move hash-slots to the shard)

To add a replica to a shard master
```bash
redis-cli --cluster add-node <Pod1>:6379 <Pod4>:6379 --cluster-slave
```

**Important**:
Unfortunaly REDIS does not support connection of cluster nodes through dns-names. It does however support automatically changing IPs of other cluster nodes when single nodes change the IP.

In practice that leads to cluster-corruption when all nodes are restarted at the same time (e.g. using *helm delete <stage>* and then reapplying the chart without deleting the persistent-volumes).

When pods are recreated they aquire a different IP, thats why the *nodes.conf* on the REDIS containers will contain IPs from the cluster. If all of these IPs are not reachable the cluster is corrupted and will not work.

To not use a weird workaround, I decided to just ignore this. As long as not all pods are down at the same time (also when applying a update with Helm or when rollout restarting) this should not cause any issues.
