const express = require('express');
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const seedFirestore = require('./seed');

// constants
const PROTO_PATH = __dirname + '/protocol_buffers/helloworld/helloworld.proto';
const STORE_PROTO_PATH = __dirname + '/protocol_buffers/store/store.proto';
const OPTIONS = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
};

// protocol buffers
const packageDefinition = protoLoader.loadSync(PROTO_PATH, OPTIONS);
const storePackageDefinition = protoLoader.loadSync(STORE_PROTO_PATH, OPTIONS);
const helloProtoDescriptor = grpc.loadPackageDefinition(packageDefinition).helloworld;
const storeProtoDescriptor = grpc.loadPackageDefinition(storePackageDefinition).store;

// server setup
const app = express();
const port = 8080;

// express routes
app.get('/hello/:server/:name', function (req, res) {
  const server = req.params.server;
  const name = req.params.name;
  const client = new helloProtoDescriptor.Greeter(`${server}-grpc-server:50051`, grpc.credentials.createInsecure());

  client.sayHello({name: name}, function(err, response) {
    console.log('Greeting:', response.message);
    res.send(response.message);
  });
});

app.get('/item/:server/:id', function (req, res) {
  const server = req.params.server;
  const id = req.params.id;
  const client = new storeProtoDescriptor.Store(`${server}-grpc-server:50051`, grpc.credentials.createInsecure());

  client.getLineItem({id: id}, function(err, response) {
    console.log('Fetched line item:', response);
    res.send({...response, client: 'node'});
  });
});

// seed firestore
seedFirestore();

// start express server
app.listen(port, () => console.log(`Example app listening on port ${port}!`));
