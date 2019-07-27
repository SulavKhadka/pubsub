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
var Queues = make(map[string]Topic)

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
		queue.Insert(topic, message.Msg)
		responsePayload.Status = "Success"
		responsePayload.Msg = fmt.Sprintf("Recieved Msg %s // ID: %d // Topic: %s", message.Msg, message.ID, message.Topic)

		fmt.Println(responsePayload)
		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	responsePayload.Status = "Fail"
	responsePayload.Err = "Topic doesnt exist"

	fmt.Println(responsePayload)
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

		fmt.Println(responsePayload)
		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	if _, exists := Queues[topic.TopicName]; exists {
		responsePayload.Status = "Fail"
		responsePayload.Err = "Topic already exists."

		fmt.Println(responsePayload)
		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	newTopicQueue := queue.New(topic.TopicName)
	Queues[topic.TopicName] = Topic{topic.TopicName, newTopicQueue}

	responsePayload.Status = "Success"
	responsePayload.Msg = fmt.Sprintf("Topic %s Added.", topic.TopicName)

	fmt.Println(responsePayload)
	json.NewEncoder(res).Encode(&responsePayload)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/newTopic", CreateTopic).Methods("POST")
	router.HandleFunc("/sendMessage", SendMessage).Methods("POST")
	//router.HandleFunc("/jobs", PrintMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
