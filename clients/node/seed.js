const firebase = require('@firebase/testing');

//firebase testing
firebase.initializeAdminApp({projectId: process.env.FIRESTORE_PROJECT_ID});
const firebaseApp = firebase.apps()[0];
const firestore = firebaseApp.firestore();

// firestore seed constant
const LINE_ITEMS = [{
  id: 1,
  product: {
    name: 'Apple',
    cost: 1.20
  },
  quantity: 2,
  total_cost: 2.40,
  supplier: {
    name: 'Apple Factory',
    country: 'Canada'
  }
}];

module.exports = async function seedFirestore() {
  // create / fetch firestore line_items collection
  const collection = firestore.collection('line_items');
  let docs = [];

  // create line_items
  for(let li in LINE_ITEMS) {
    docs.push(collection.doc(li).set(LINE_ITEMS[li]));
  }

  // store line_items
  await docs

  // get line_items collection
  const lineItemsCollection = await collection.get();
  const lineItems = lineItemsCollection.docs.map(li => li.data());

  // log and return line_items collection
  console.log('Seeded lineItems: ', lineItems);
}
