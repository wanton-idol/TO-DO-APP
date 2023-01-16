package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/wanton-idol/TO-DO-APP/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	CreateDBInstance()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		// fmt.Errorf("error loading .env file: %v", err)
		fmt.Println(fmt.Errorf("error loading .env file: %v", err))
	}
}

func CreateDBInstance() {
	connectionString := os.Getenv("DB_URI")

	dbName := os.Getenv("DB_NAME")

	collectionName := os.Getenv("DB_COLLECTION_NAME")

	cilentOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), cilentOptions)
	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to MongoDB: %v", err))
	}

	//Checking connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(fmt.Errorf("error in connection checking: %v", err))
	}

	fmt.Println("Hurray! Connection established...")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(collectionName))

	fmt.Println("Collection instance succesfully created!!")

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDo
	_ = json.NewDecoder(r.Body).Decode(&task)
	ID := insertOneTask(task)
	json.NewEncoder(w).Encode(ID)

}

func TaskCompleted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskCompleted(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tasks := getAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTasks()
	json.NewEncoder(w).Encode(count)
}

//MongoDB Helper Functions

func insertOneTask(task models.ToDo) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a task: ", inserted.InsertedID)
	return inserted
}

func taskCompleted(task string) {
	fmt.Println(task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated count: ", updated.ModifiedCount)
}

func undoTask(task string) {
	fmt.Println(task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated count: ", updated.ModifiedCount)
}

func deleteTask(task string) {
	fmt.Println(task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", deleted.DeletedCount)

}

func getAllTasks() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []primitive.M

	for cursor.Next(context.Background()) {
		var task bson.M
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	return tasks
}

func deleteAllTasks() int64 {
	delete, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted tasks count: ", delete.DeletedCount)
	return delete.DeletedCount
}
