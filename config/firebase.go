package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/DamnDanielV/go-rest/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type client struct{}

// ClientInit inicializa la conexión a firebase
// retorna el cliente con acceso a firestore
func ClientInit() (client *firestore.Client, err error) {
	// Use a service account
	ctx := context.Background()

	opt := option.WithCredentialsFile("/home/daniel/Documentos/go/rest/key/credentials.json")
	config := &firebase.Config{ProjectID: "go-server-a1bc0"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		fmt.Println("error1")
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	client2, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("error2")
		log.Fatalln(err)
		return nil, err
	}
	return client2, nil
}

// CreatePost crea un documento en la colección post en Firestore
// retorna un estructura de tipo PostData (para almacenar en una base de datos local)
func CreatePost(post *entity.PostData) (*entity.PostData, error) {
	client, err := ClientInit()
	if err != nil {
		log.Fatalf("Failed initializing: %v", err)
		return nil, err
	}
	ctx := context.Background()

	_, _, err = client.Collection("posts").Add(ctx, map[string]interface{}{
		"ID":       post.Id,
		"Title":    post.Title,
		"Messagge": post.Messagge,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return nil, err
	}
	defer client.Close()

	return post, nil
}

func GetPosts() (psts []entity.PostData, err error) {
	client, _ := ClientInit()
	ctx := context.Background()
	var posts []entity.PostData
	iter := client.Collection("posts").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		post := entity.PostData{
			Id:       doc.Data()["ID"].(int64),
			Title:    doc.Data()["Title"].(string),
			Messagge: doc.Data()["Messagge"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
