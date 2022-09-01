package main
 
/*
    Here are the CRUD Operations related to MongoDB.
*/
 
import (
    "context"
    "log"
 
    "go.mongodb.org/mongo-driver/bson/primitive"
 
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 
// getConn fonk mongoDB conn
func getConn() (*mongo.Client, context.Context) {
 
    /*
        NewClient creating new object with docker
        and erorr handling
    */
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://scoth:tiger@localhost:27017")) // ilgili bağlantı bilgisini kullanarak yeni bir client nesnesi oluşturuldu
    ctx := context.Background()
 
    err = client.Connect(ctx)
    if err != nil {
        log.Printf("Bağlanırken hata alındı: %v", err)
    }
 
    // pinging mongoDB
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Printf("Ping yapılamadı: %v", err)
    }
 
    return client, ctx 
}
 

func AddQuote(q *quote) (primitive.ObjectID, error) {
    log.Println("Add Quote")
    log.Println(q)
 
    client, ctx := getConn()
    defer client.Disconnect(ctx)  
    q.ID = primitive.NewObjectID() 
 
    
    result, err := client.Database("bookworms").Collection("quotes").InsertOne(ctx, q)
    if err != nil { 
        log.Printf("Alıntı eklenmeye çalışırken hata oluştu %v", err)
        return primitive.NilObjectID, err
    }
    id := result.InsertedID.(primitive.ObjectID)
    return id, nil 
}
 

func GetQuotes() ([]*quote, error) {
    var quotes []*quote 
 
    client, ctx := getConn()     
    defer client.Disconnect(ctx) 
 
    db := client.Database("bookworms")            
    collection := db.Collection("quotes")         
    cursor, err := collection.Find(ctx, bson.D{}) 
    if err != nil {                               
        return nil, err
    }
    defer cursor.Close(ctx)       
    err = cursor.All(ctx, "es) // Koleksiyonu quotes'a alıyoruz
    if err != nil {                
        log.Println(err)
        return nil, err
    }
    return quotes, nil 
}