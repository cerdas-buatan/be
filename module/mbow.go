package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kimseokgis/backend-ai/config"
	"github.com/kimseokgis/backend-ai/helper"
	"github.com/kimseokgis/backend-ai/model"
	"net/http"
	"strings"
)

func ChatPredictRegex(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	chat := new(model.Chats)
	token := r.Header.Get("login")
	if token == "" {
		resp.Message = "token is empty"
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	}
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}
	key := keys[0]
	fmt.Println(key)
	decoder, err := helper.DecodeGetUser(config.PublicKey, token)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = false
		helper.WriteJSON(w, http.StatusBadRequest, resp)
		return
	}
	db := helper.SetConnection()
	fmt.Println(decoder)

	_, err = helper.FindUserByUsername(db, decoder)
	if err != nil {
		resp.Message = fmt.Sprintf("Data tidak ditemukan : %s\n"+
			"Username: %s\n", err.Error(), decoder)
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	if strings.Contains(key, "_") {
		key = strings.Replace(key, "_", " ", -1)
	}
	fmt.Printf("%+v\n", key)
	reply, score, err := helper.QueriesDataRegexpALL(db, context.TODO(), key)
	if err != nil {
		resp.Message = "Waduh, maaf saya tak mengerti apa yang kau tanyakan. Coba susun ulang pertanyaan:)"
		resp.Status = false
		chat.Responses = resp.Message
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	chat.IdChats = reply.ID.Hex()
	chat.Message = reply.Question
	chat.Responses = reply.Answer
	chat.Score = score
	defer db.Client().Disconnect(context.Background())
	helper.WriteJSON(w, http.StatusOK, chat)
	return
}

func ChatPredict(w http.ResponseWriter, r *http.Request) {
	resp := new(model.Credential)
	chat := new(model.Chats)
	mess := new(model.ChatRequest)
	err := json.NewDecoder(r.Body).Decode(&mess)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	}
	token := r.Header.Get("secret")
	if token == "" {
		resp.Message = "secret is empty"
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotAcceptable, resp)
		return
	}
	fmt.Println(mess)
	db := helper.SetConnection()
	//fmt.Println(decoder)

	_, err = helper.QueriesSecret(db, context.TODO(), token)
	if err != nil {
		resp.Message = fmt.Sprintf("Data tidak ditemukan : %s\n", err.Error())
		resp.Status = false
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}
	fmt.Printf("%+v\n", mess.Messages)
	reply, score, err := helper.QueriesDataRegexpALL(db, context.TODO(), mess.Messages)
	if err != nil {
		resp.Message = "Aduh aduh aduhhhaiii, aku ga ngerti nihh coba nanya yang lain dongg biar aku ngertiin kamu..."
		resp.Status = false
		chat.Responses = resp.Message
		helper.WriteJSON(w, http.StatusNotFound, resp)
		return
	}

	chat.IdChats = reply.ID.Hex()
	chat.Message = reply.Question
	chat.Responses = reply.Answer
	chat.Score = score
	if reply.Answer == "" {
		chat.Message = "Adududududu aku ga ngerti nih kakak"
	}
	defer db.Client().Disconnect(context.Background())

	helper.WriteJSON(w, http.StatusOK, chat)
	return
}
