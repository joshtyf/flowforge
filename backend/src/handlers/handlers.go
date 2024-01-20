package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/////////////////// Handlers ///////////////////

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if serverHealthy() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server working!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server not working!"))
	}
}

func GetAllServiceRequest(w http.ResponseWriter, r *http.Request) {
	allsr, err := getAllServiceRequest()
	if err != nil {
		JSONError(w, NewHttpError(err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allsr)
	w.WriteHeader(http.StatusOK)
}

func GetServiceRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	sr, err := getServiceRequest(requestId)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSONError(w, NewHttpError(err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(sr)
	w.WriteHeader(http.StatusOK)
	return
}

func CreateServiceRequest(w http.ResponseWriter, r *http.Request) {
	sr := NewServiceRequest()
	err := json.NewDecoder(r.Body).Decode(&sr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSONError(w, NewHttpError(err), http.StatusBadRequest)
		return
	}
	err = createServiceRequest(sr)
	if err != nil {
		JSONError(w, NewHttpError(err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

/////////////////// Helper Functions ///////////////////

func serverHealthy() bool {
	// TODO: Include database health check

	return true
}

func getDatabase() (database *mongo.Database, err error) {
	log.Println("function called")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Println("database connection error", err)
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("Successfully connected and pinged.")

	dbName := os.Getenv("MONGO_DATABASE")
	log.Println("DBNAME:", dbName)
	database = client.Database(dbName)

	log.Println(dbName, database.Name())
	return database, nil
}

func createServiceRequest(sr *ServiceRequest) (err error) {
	db, err := getDatabase()
	result, err := db.Collection("service_requests").InsertOne(context.TODO(), sr)
	if err != nil {
		return
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return
}

func getServiceRequest(hexId string) (sr *ServiceRequest, err error) {
	db, err := getDatabase()
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: id}}
	err = db.Collection("service_requests").FindOne(context.TODO(), filter).Decode(&sr)
	if err != nil {
		return nil, err
	}
	return
}

func getAllServiceRequest() (allSr []*ServiceRequest, err error) {
	db, err := getDatabase()
	cursor, err := db.Collection("service_requests").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var elem *ServiceRequest
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		allSr = append(allSr, elem)
	}
	return
}

func JSONError(w http.ResponseWriter, err *HttpError, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
