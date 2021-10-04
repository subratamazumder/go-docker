package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const SERVICE_VERSION = "6.0"
const SERVICE_PORT = 8081

func main() {
	http.HandleFunc("/ep-registration-service/health", handleHealthFunc)
	http.HandleFunc("/ep-registration-service/registrations", handleRegistrationFunc)
	log.Printf("HTTP Go Server is Listening on  %s : %d", getHostName(), SERVICE_PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(SERVICE_PORT), nil))
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
	response["processingNode"] = getHostName()
	response["serviceVersion"] = SERVICE_VERSION
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("JSON marshalling error : %s", err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
	log.Printf("Returning %d from node %s", http.StatusCreated, getHostName())
}
func registerUser() string {
	// add data base logic here
	return uuid.New().String()
}
func getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}
