package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var queue Topic

// Queues holds all the queues for a given instance
var Queues = make(map[string]*Topic)

// Message :Struct for storing JSON message payloads
type Message struct {
	Msg   string `json:"msg"`
	ID    int    `json:"id"`
	Topic string `json:"topic"`
}

func check(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
		panic(e)
	}
}

// SendMessage : parses and drops message to the appropiate channel as defined in the incoming payload
func SendMessage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	responsePayload := struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
		Err    string `json:"err"`
	}{}

	message := Message{}
	messagePayload := json.NewDecoder(req.Body)
	err := messagePayload.Decode(&message)
	check(err)

	if topic, exists := Queues[message.Topic]; exists {
		topic.Insert(message.Msg)
		responsePayload.Status = "Success"
		responsePayload.Msg = fmt.Sprintf("Recieved Msg %s // ID: %d // Topic: %s", message.Msg, message.ID, message.Topic)

		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	responsePayload.Status = "Fail"
	responsePayload.Err = "Topic doesnt exist"

	json.NewEncoder(res).Encode(responsePayload)
}

//GetMessage : retrieves the top message from the appropriate queue
func GetMessage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	requestPayload := struct {
		TopicName string `json:"topic_name"`
	}{}

	responsePayload := struct {
		Topic string `json:"topic"`
		Msg   Item   `json:"msg"`
		Err   string `json:"err"`
	}{}

	topic := requestPayload
	topicPayload := json.NewDecoder(req.Body)
	err := topicPayload.Decode(&topic)
	check(err)

	if topicName, exists := Queues[topic.TopicName]; exists {
		item := topicName.Get()
		responsePayload.Topic = topicName.TopicName
		responsePayload.Msg = item

		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	responsePayload.Err = "Topic doesnt exist"

	json.NewEncoder(res).Encode(responsePayload)

}

// CreateTopic : Adds a new channel to the message queue
func CreateTopic(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	requestPayload := struct {
		TopicName string `json:"topic_name"`
	}{}
	responsePayload := struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
		Err    string `json:"err"`
	}{}

	topic := requestPayload
	topicPayload := json.NewDecoder(req.Body)
	err := topicPayload.Decode(&topic)
	check(err)

	if len(Queues) >= 25 {
		responsePayload.Status = "Fail"
		responsePayload.Err = "Max amount of topics reached."

		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	if _, exists := Queues[topic.TopicName]; exists {
		responsePayload.Status = "Fail"
		responsePayload.Err = "Topic already exists."

		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	newTopicQueue := queue.New(topic.TopicName)
	Queues[topic.TopicName] = &newTopicQueue

	responsePayload.Status = "Success"
	responsePayload.Msg = fmt.Sprintf("Topic %s Added.", topic.TopicName)

	json.NewEncoder(res).Encode(&responsePayload)
}

// Length : returns the length of a queue
func Length(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	requestPayload := struct {
		TopicName string `json:"topic_name"`
	}{}
	responsePayload := struct {
		Status string `json:"status"`
		Length int    `json:"length"`
		Err    string `json:"err"`
	}{}

	topic := requestPayload
	topicPayload := json.NewDecoder(req.Body)
	err := topicPayload.Decode(&topic)
	check(err)

	if userTopic, exists := Queues[topic.TopicName]; exists {
		responsePayload.Status = "Success"
		responsePayload.Length = len(userTopic.Queue)

		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	responsePayload.Status = "Fail"
	responsePayload.Err = "Topic doesnt exists."

	json.NewEncoder(res).Encode(responsePayload)
}

func loaderIoToken(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("loaderio-3000acba7e633b71b1d2d9439c376dd8"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/topic", CreateTopic).Methods("PUT")
	router.HandleFunc("/topic", SendMessage).Methods("POST")
	router.HandleFunc("/topic", GetMessage).Methods("GET")
	router.HandleFunc("/length", Length).Methods("GET")
	router.HandleFunc("/loaderio-3000acba7e633b71b1d2d9439c376dd8/", loaderIoToken)
	//router.HandleFunc("/jobs", PrintMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
