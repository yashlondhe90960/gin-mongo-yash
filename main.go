package main

import "github.com/gin-gonic/gin"

var{
	server *gin.Engine

	userservice services.UserService
	usercontroller controllers.UserController
	ctx context.Context
	usercollection *mongo.Collection
	mongoclient *mongo.Client
	err error

}
func init(){
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
}
func main() {}
