package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

const SERVICE_VERSION = "12.0"
const SERVICE_PORT = 8081
const DYNAMO_USER_PROFILE_TABLE = "eprescription-user-profile"

func main() {
	http.HandleFunc("/ep-registration-service/health", handleHealthFunc)
	http.HandleFunc("/ep-registration-service/registrations", handleRegistrationFunc)
	log.Printf("HTTP Go Server is Listening on  %s : %d", getHostName(), SERVICE_PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(SERVICE_PORT), nil))
}

func handleHealthFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.WriteHeader(http.StatusOK)
	log.Printf("Returning %d", http.StatusOK)
}
func setupCORSHeaderResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func handleRegistrationFunc(w http.ResponseWriter, r *http.Request) {
	//for browser client
	setupCORSHeaderResponse(&w,r)
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Printf("Request received from %s", r.RemoteAddr)
	statusCode := http.StatusCreated
	response := make(map[string]string)
	response["processingNode"] = getHostName()
	response["serviceVersion"] = SERVICE_VERSION
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	var regRequest RegistrationRequest
	body, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(body, &regRequest); err != nil {
		log.Printf("Body parse error, %v", err)
		statusCode = http.StatusBadRequest // Return 400 Bad Request.
	} else {
		registrationId, err := registerUser(regRequest)
		if err != nil {
			log.Printf("Registration error, %v", err)
			statusCode = http.StatusInternalServerError // Return 500 Internal Server Error.
		} else {
			response["registrationId"] = registrationId
			statusCode = http.StatusCreated // Return 201 Created.
		}
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("JSON marshalling error : %s", err)
	}
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
	log.Printf("Returning %d from node %s", statusCode, getHostName())
}
func registerUser(regRequest RegistrationRequest) (string, error) {
	profileId := uuid.New().String()
	return putItem(profileId, regRequest.FirstName, regRequest.LastName)
}
func getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

type Item struct {
	ProfileId string
	FisrtName string
	LastName  string
}

type RegistrationRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func putItem(profileId string, firstName string, lastName string) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	item := Item{
		ProfileId: profileId,
		FisrtName: firstName,
		LastName:  lastName,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Printf("Got error marshalling new user profile: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(DYNAMO_USER_PROFILE_TABLE),
	}
	log.Println("Item is being added for profile_id ", item.ProfileId)
	log.Println(input)
	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return "", err
	}
	return profileId, nil
}
