# Debug Setup

To debug the application, you can setup a three-node cluster on your local machine.

For the setup I use a ubuntu machine, when using other distributions or Windows the setup is really simular, the steps for redis installation may vary.

The instructions require you to have the *flutter sdk* and the *go* compiler installed.

### Installation Redis

Install redis
**Important**: Make sure to install at least redis 7.0.0, the library used does not support redis-servers below version 7. It will produce weird errors.
```bash
# Arch
sudo pacman -S redis

# Ubuntu/Debian
sudo apt install redis-server
```

### Run Redis Cluster

Configure ACL list with your desired credentials (user must match with the configs)
```bash
nano ./access.acl
```

Start cluster (port 7001,7002,7003)
```bash
./startcluster.sh
```
To setup the cluster and assign the hash-slots you need to execute this command
```bash
redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-replicas 0
```
(Notice that you need to specify at least 2 nodes, you can just specify the local node twice here)

When you want to add another shard master you can do it like this
```bash
redis-cli --cluster add-node 127.0.0.1:7002 127.0.0.1:7001
```
(Notice that when doing this after creation, you will also need to move hash-slots to the shard)

To add a replica to a shard master
```bash
redis-cli --cluster add-node 127.0.0.1:7003 127.0.0.1:7001 --cluster-slave
```


### Run Server Applications

Configure debug orchestrator config
```bash
nano ./debug.orchestrator.yaml
```
Configure debug server config
```bash
nano ./debug.server.yaml
```

Now you can run start the server/orchestrator with the debug config
```bash
(cd ../../orchestrator && go run main.go ../deployment/debug/debug.orchestrator.yaml)
```

```bash
(cd ../../server && go run -tags debug main.go ../deployment/debug/debug.server.yaml)
```

### Run Client Application

For the Flutter-Client application, I recommend installing Android-Studios.

You can install the required packages with
```bash
flutter pub get
```

Run the application with
```bash
flutter run -d linux
```
(windows respectively)


### Rebuild protobuf

When changing the proto API endpoints, you can rebuild them with the provided script

```bash
cd ../../
./generate_proto
```

Notice that the generate proto script should install the protoc compiler aswell as the required go/dart proto plugins. If this fails, you will need to install it manually.

### Build Application

The server applications are fully standalone and can simply be built with
```bash
(cd ../../server && go build -o ../deployment/debug/server)
```
```bash
(cd ../../orchestrator && go main.go -o ../deployment/debug/orchestator)
```

### Debug with mutual ssl connection to redis

In production you should use a mutual ssl connection to the database, to debug it you will can do it like this:

Generate self signed certificate with openssl
```bash
openssl req -x509 -newkey rsa:4096 -keyout priv.key -out pub.crt -days 365
```

Uncomment the lines in the *redis.conf*
```bash
tls-cert-file pub.crt
tls-key-file priv.key
tls-ca-cert-file pub.crt
```

In the *debug.[orchestrator/server].yaml* you will need to input the inner cert block (ignore the PEM markers) to those options
```yaml
db_base64_ssl_certificate: "<inner block of pub.crt>"
db_base64_ssl_privatekey: "<inner block of priv.key>"
db_base64_ssl_ca: "<inner block of pub.crt>"
```

To run the cluster with the tls-port, you will need to change the *--port* flag in the *startcluster.sh* script to *--tls-port*

```bash
#!/bin/bash

redis-server ./redis.conf --tls-port 7001 &
pid1=$!
redis-server ./redis.conf --tls-port 7002 &
pid2=$!

trap "kill $pid1 $pid2" EXIT

redis-server ./redis.conf --tls-port 7003
```
