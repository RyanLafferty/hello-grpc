# hello-grpc

## Technologies

* docker, docker-compose
* grpc, protoc, protoc-gen-go, grpc_tools_ruby_protoc
* firestore, cloud-firestore-emulator
* node, express, @grpc/proto-loader, @firebase/testing
* golang, gorilla/mux, google.golang.org/grpc, cloud.google.com/go/firestore
* ruby, sinatra, logging, grpc, google-cloud-firestore

## Usage

### Prerequisites

#### Install GRPC

TODO

#### Install Golang

```bash
# Ubuntu
./scripts/go-ubuntu-installation.sh

# MacOS
./scripts/go-macos-installation.sh
```

#### Install Ruby

```bash
# Ubuntu
TODO

# MacOS
./scripts/ruby-macos-installation.sh
```

#### Install GRPC Golang Tools

```bash
./scripts/go-grpc-installation.sh
```

#### Install GRPC Ruby Tools

```bash
./scripts/ruby-grpc-installation.sh
```

#### Generate Protocol Buffers

```bash
./scripts/generate-protobufs.sh
```

### Build & Run

```bash
# build stack
docker-compose build

# run stack
docker-compose up
```

### Access

```bash
curl localhost:<CLIENT_PORT>/<SERVICE>/<SERVER>/<PAYLOAD>
```

#### Client Ports (CLIENT_PORT)

```bash
# node
3000

# golang
3001

# ruby
3002
```

#### Services (SERVICE)

```bash
# Greeter.SayHello
hello

# Store.GetLineItem
item
```

#### Servers (SERVER)

```bash
# node
node

# golang
go

# ruby
ruby
```

#### Payloads (PAYLOAD)

The current firestore seed method will only insert one line item at index `0`.
When attempting to get a line item ensure that the document exists.
In our case we would only access line item `0`.

```bash
# Greeter.SayHello (string)
Ryan

# Store.GetLineItem (unsigned integer)
0
```

#### Examples

##### Greeter.SayHello

```bash
# node client, node server
curl localhost:3000/hello/node/Ryan

# golang client, golang server
curl localhost:3001/hello/go/Ryan

# ruby client, ruby server
curl localhost:3002/hello/ruby/Ryan
```

##### Store.GetLineItem

```bash
# node client, node server
curl localhost:3000/item/node/0

# golang client, golang server
curl localhost:3001/item/go/0

# ruby client, ruby server
curl localhost:3002/item/ruby/0
```

## Sources

### Protocol Buffers

* https://developers.google.com/protocol-buffers/docs/proto3

### grpc

* https://www.grpc.io/docs/tutorials/basic/node/
* https://grpc.io/docs/tutorials/basic/go/
* https://grpc.io/docs/tutorials/basic/ruby/

### Firestore

* https://firebase.google.com/docs/firestore/quickstart

### Sinatra

* http://sinatrarb.com/
