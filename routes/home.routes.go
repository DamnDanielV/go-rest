package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DamnDanielV/go-rest/config"
	"github.com/DamnDanielV/go-rest/entity"
)

var posts []entity.PostData // slice (array dinamico) que contendrá structuras PostData

func GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	posts, err := config.GetPosts()
	_, err = json.Marshal(posts) // codifica a formato JSON
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": err.errors()}`))
		return
	}
	// //vars := mux.Vars(req)          // contiene los parametros del request
	res.WriteHeader(http.StatusOK) // respuesta http con el codigo de estado indicado
	// // fmt.Fprintf(res, "Route: %v\n", vars["text"])
	json.NewEncoder(res).Encode(posts)
	// res.Write(result)

}

func CreatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json") // la respuesta tendra formato JSON

	post := entity.PostData{}
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Fatalf(err.Error())
		res.Write([]byte(`{"error": "decoding the body"}`))
		return
	}
	post.Id = int64(len(posts) + 1)
	config.CreatePost(&post)
	posts = append(posts, post) // adiciona el usuario al slice users
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts) // codifica a formatao JSON

	res.Write(result)
}
