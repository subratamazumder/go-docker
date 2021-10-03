package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"os"
)

func main() {
	http.HandleFunc("/ep-registration-service/health", handleHealthFunc)
	http.HandleFunc("/ep-registration-service/registrations", handleRegistrationFunc)
	log.Printf("HTTP Go Server is Listening on  %s : 8081",getHostName())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleHealthFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	log.Printf("Returning %d", http.StatusOK)
}
func handleRegistrationFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	response := make(map[string]string)
	response["registrationId"] = registerUser()
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("JSON marshalling error : %s", err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
	log.Printf("Returning %d", http.StatusCreated)
}
func registerUser() string {
	// add data base logic here
	return uuid.New().String()
}
func getHostName() string{
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}
