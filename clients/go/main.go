package main

import (
	"context"
	"log"
	"time"
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	grpcserver "./lib/grpcserver"
	store "./lib/grpcserver/store"
	storeSerializers "./lib/serializers/store"
	helloPb "./helloworld"
	storePb "./store"
)

// endpoint handlers
func hello(res http.ResponseWriter, req * http.Request) {
	// get vars from url
	vars := mux.Vars(req)

	address := vars["server"] + "-grpc-server:50051"
	conn := grpcserver.Connect(address)
	if conn == nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer conn.Close()
	c := helloPb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := vars["name"]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloPb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	
	log.Printf("Greeting: %s", r.GetMessage())
	fmt.Fprintf(res, "%s\n", r.GetMessage())
}

func getLineItem(res http.ResponseWriter, req * http.Request) {
	// get vars from url
	vars := mux.Vars(req)
	log.Printf("Attempting to fetch: %s", vars["id"])

	id, err := strconv.ParseInt(vars["id"], 10, 16)
	id32 := uint32(id)

	if err == nil {
		address := vars["server"] + "-grpc-server:50051"
		conn := grpcserver.Connect(address)
		if conn == nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer conn.Close()
		client := storePb.NewStoreClient(conn)
		if client == nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		lineItem := store.GetLineItem(client, id32)
		if lineItem == nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		// set headers
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)

		// return line item
		err = json.NewEncoder(res).Encode(storeSerializers.SerializeLineItem(lineItem))
		if err != nil {
			log.Printf("Error: Failed to encode line item: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		return
	}

	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(res, "Error: Failed to parse id\n")
}

func main() {
	// create router
	router := mux.NewRouter()

	// define enpoints
	router.HandleFunc("/hello/{server}/{name}", hello).Methods("GET")
	router.HandleFunc("/item/{server}/{id}", getLineItem).Methods("GET")

	// assign router & start server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
