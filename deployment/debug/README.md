# Debug Setup

To debug the application, you can setup a single-node cluster on your local machine.

For the setup I use a ubuntu machine, when using other distributions or Windows the setup is really simular, the steps for redis installation may vary.

The instructions require you to have the *flutter sdk* and the *go* compiler installed.

### Installation Redis

Install redis with apt
```bash
sudo apt install redis-server
```

### Run Redis Cluster

Configure ACL list with your desired credentials
```bash
nano ./rbac.acl
```

Start redis server
```bash
redis-server ./redis.conf
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
Configure acl list (user must match with the configs)
```bash
nano ./rbac.acl
```

Now you can run start the server/orchestrator with the debug config
```bash
go run ../../orchestrator/main.go ./debug.orchestrator.yaml
```

```bash
go run ../../server/main.go ./debug.server.yaml
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
go build ../../server/main.go -o server
```
```bash
go build ../../orchestrator/main.go -o orchestrator
```

### Debug with mutual ssl connection to redis

In production you should use a mutual ssl connection to the database, to debug it you will can do it like this:

Generate self signed certificate with openssl
```bash
openssl req -x509 -newkey rsa:4096 -keyout priv.key -out pub.crt -days 365
```

Uncomment the lines in the *redis.conf*
```bash
# port 6379
tls-port 6379
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
