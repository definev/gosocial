package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/definev/gosocial/internal/model"
)

func (apiCfg apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call GET handler
		apiCfg.handlerRetrievePosts(w, r)
		break
	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)
		break
	case http.MethodDelete:
		apiCfg.handlerDeletePost(w, r)
		break
	default:
		respondWithError(w, http.StatusNotFound, errors.New("method not supported"))
	}
}

func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"user_email"`
		Title     string `json:"title"`
		Body      string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decoder.Decode(&params)

	post, err := apiCfg.dbClient.CreatePost(params.UserEmail, params.Title, params.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, 200, post)
}

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		PostId model.Maybe[int64] `json:"post_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decoder.Decode(&params)
	if params.PostId.Nil {
		respondWithError(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	err := apiCfg.dbClient.DeletePost(params.PostId.Value)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, 200, nil)

}

func (apiCfg apiConfig) handlerRetrievePosts(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.Trim(r.URL.Path, "/postw/")
	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("user email is required"))
		return
	}

	posts, err := apiCfg.dbClient.GetPosts(userEmail)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, 200, posts)
}
