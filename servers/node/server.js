const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const firebase = require('@firebase/testing');

//firebase testing
firebase.initializeAdminApp({projectId: process.env.FIRESTORE_PROJECT_ID});
const firebaseApp = firebase.apps()[0];
const firestore = firebaseApp.firestore();

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

// service methods
function sayHello(call, callback) {
  callback(null, {message: 'Hello ' + call.request.name});
}

async function getLineItem(call, callback) {
  const id = call.request.id;

  try {
    const lineItem = await firestore.collection('line_items').doc(id.toString(10)).get();
    let lineItemData = await lineItem.data();
    lineItemData['source'] = 'node';

    callback(null, lineItemData);
  } catch (error) {
    console.error('Encountered error: ', error);
    callback(null, null);
  }
}

// grpc server
const server = new grpc.Server();
server.addService(helloProtoDescriptor.Greeter.service, {sayHello: sayHello});
server.addService(storeProtoDescriptor.Store.service, {getLineItem: getLineItem});
server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
server.start();
