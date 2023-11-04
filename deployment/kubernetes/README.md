# Kubernetes Deployment

To deploy the application on Kubernetes we will need to deploy the orchestrator and server services aswell as the redis-cluster.


### Prerequisites

For the Kubernetes deployment, an installed Kubernetes cluster (at least v1.28.0) is required.

The deployment also uses some other services, you can install them as follows:

#### Metallb

To loadbalance the nodes without external loadbalancer we will use Metallb, you can install it by manifest (v0.13.12):

```bash
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.12/config/manifests/metallb-native.yaml
```

#### Longhorn

As Storageprovider for the redis-cluster we use Longhorn for high available block storage.

To use longhorn you will need open-iscsi installed on every node, on Arch you can do this like so:

```bash
sudo pacman -S open-iscsi

# Enable and run the iscsid service on every node
sudo systemctl enable iscsid && sudo systemctl start iscsid
```

Then you can install the longhorn system with kubectl (v1.5.2):

```bash
kubectl apply -f https://raw.githubusercontent.com/longhorn/longhorn/v1.5.2/deploy/longhorn.yaml
```



### Setup Redis

The redis-cluster needs to be initialized, to do this you will need to call following command from the redis cli:

```bash
redis-cli --cluster create <Pod1>:6379 <Pod2>:6379 <Pod3>:6379 --cluster-replicas 0
```

This needs to be done initialy, you can do it over a automated solution like Ansible or just simply create a port-forwarding and do it manually:

```bash
kubectl port-forward pod/<Pod1> 7001:6379 &
kubectl port-forward pod/<Pod2> 7002:6379 &
kubectl port-forward pod/<Pod3> 7003:6379 &

redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-replicas 0
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


