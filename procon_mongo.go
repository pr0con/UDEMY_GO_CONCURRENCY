package main

import(
	//Native
	"fmt"
	"context"
	
	//3rd party
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	//Our Other Packages
	"github.com/pr0con/go_private_modules/procon_config"
)

type key string

const (
	HostKey     = key("hostKey")
	UsernameKey = key("usernameKey")
	PasswordKey = key("passwordKey")
	DatabaseKey = key("databaseKey")	
)

var ctx context.Context;
var client *mongo.Client;

func init() {
	ctx = context.Background()	
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	ctx = context.WithValue(ctx, HostKey, procon_config.MongoHost)
	ctx = context.WithValue(ctx, UsernameKey, procon_config.MongoUser)
	ctx = context.WithValue(ctx, PasswordKey, procon_config.MongoPassword)
	ctx = context.WithValue(ctx, DatabaseKey, procon_config.MongoDb)

	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value(UsernameKey).(string),
		ctx.Value(PasswordKey).(string),
		ctx.Value(HostKey).(string),
		ctx.Value(DatabaseKey).(string),
	)
	clientOptions := options.Client().ApplyURI(uri)	
	
	var err error
	client, err = mongo.Connect(ctx, clientOptions)	
	
	err = client.Ping(ctx, nil)
	if err != nil { fmt.Println(err); } else { fmt.Println("Mongo Connected"); }
}

func main() {
	fmt.Println("test")
}