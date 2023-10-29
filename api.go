package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddr string
}

type BankAPIFun func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func convertBankAPIFunToHandlerFunc(apiFunc BankAPIFun) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			//TODO: handle the proper status codes later
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/account", convertBankAPIFunToHandlerFunc(s.handleGetAccount))

	router.Mount("/v1", v1Router)

	log.Println("🚀 Server listening at: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, NewAccount("Daniel", "Ezihe"))
	//return WriteJson(w, http.StatusOK, &Account{})
	//return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
