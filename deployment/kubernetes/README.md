# Kubernetes Deployment

The kubernetes deployment is handled with HELM and the provided chart.

Dependencies:
 - Metallb (v0.13.12)
 - Longhorn (v1.5.2)
 
The dependencies are included in the repository, so you will not need to install it manually.

For the longhorn block-storage every kubernetes node needs to have *open-iscsi* installed and up and running:

```bash
# Example for arch
sudo pacman -S open-iscsi

sudo systemctl enable iscsid && sudo systemctl start iscsid
```

### Installation

Install the Helm chart with:

```bash
helm install <stage> ./chart --set orchestratorSpeakerIP=<speakerIp> --namespace obstrike-system --create-namespace
```
Usually you want to use a custom *values.yaml* file, you can specify it after the *-f* or *--values* flag. The mandatory flag *orchestratorSpeakerIP* is the IP that is "exposed" by the metallb speaker.

For instructions to the specific attributes, just copy the *values.yaml* file from the *chart* directory, inside you will find any attribute documented and described.
```bash
cp ./chart/values.yaml myvalues.yaml
emacs -nw myvalues.yaml
```

### Initialization (also described in HELM output)

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

### Update dependencies

When you need to update the dependencie-versions you can simply do this by updating the *version* attribute in the *dependencie* section of the *Chart.yaml* file and then running:
```
helm dependencie build
```

Make sure to extensively test the application after updating the dependencies!

### Manifest

When you need a manifest instead of the Helm chart, you can just generate the file from the chart:
```bash
helm template <releaseName> chart --namespace orbstrike-system --set orchestratorSpeakerIP=<speakerIp> orbstrike.yaml
```

Specify all variables you want to customize with either *--set* or with a custom *value.yaml* file (specified by *--values*)
