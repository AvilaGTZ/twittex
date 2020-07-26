package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN Es el objeto de conexcion a la BD*/
var MongoCN = ConectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://grayefx:a43056043A@cluster0.shdj5.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD Es la funcion que me permite conectar a la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client

}

/*ChequeoConnection es el Ping a la base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
