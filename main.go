package main

import (
	     "fmt"
	     "net/http"
	     "log"
	     "context"
	     "time"
	     "go.mongodb.org/mongo-driver/mongo/options"
	     "go.mongodb.org/mongo-driver/bson"
         "go.mongodb.org/mongo-driver/mongo"
	     "encoding/json"
         
	    
      
)
	
type User struct {
	ID string   
	Name string           
	Password string       
	DOB string            
	Phone string          
	Email string           
}

var client *mongo.Client

func homepage(w http.ResponseWriter,r *http.Request){
	if(r.Method=="GET"){ 
		fmt.Println("WELCOME PAGE")}
	
	
}

func createuser(w http.ResponseWriter,r *http.Request){
	if(r.Method=="POST"){
		w.Header().Add("content-type" , "application/json")
      var user User
		json.NewDecoder(r.Body).Decode(&user)
		collection := client.Database("Infected").Collection("Users")
	    ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)
		result, _ := collection.InsertOne(ctx, user)
		json.NewEncoder(w).Encode(result)
	}
	
}

func main(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Sasuke:Uchiha@cluster0.wskwz.mongodb.net/<dbname>?retryWrites=true&w=majority"))  
	if(err!=nil){
		log.Fatal(err)
	}
	ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if(err!=nil){
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println(databases)
	http.HandleFunc("/",homepage)
	http.HandleFunc("/users",createuser)
 		     
	
   log.Fatal(http.ListenAndServe(":1112",nil))
}
