package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Message :Struct for storing JSON message payloads
type Message struct {
	Msg   string `json:"msg"`
	ID    int    `json:"id"`
	Topic string `json:"topic"`
}

// Topic : Struct for storing message queue topics
type Topic struct {
	TopicName string      `json:"topic_name"`
	Queue     chan string `json:"queue"`
}

// TopicQueues : Holds all the queues for a given instance for the message queue
var TopicQueues = make(map[string]Topic)

func check(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
		panic(e)
	}
}

// SendMessage : parses and drops message to the appropiate channel as defined in the incoming payload
func SendMessage(msg Message) string {

	// msgStruct := Message{}
	// err := json.Unmarshal([]byte(msg), &msgStruct)
	// check(err)

	if topic, exists := TopicQueues[msg.Topic]; exists {
		topic.Queue <- msg.Msg
	}
	return fmt.Sprintf("Recieved Msg %s.\nID: %d\nTopic: %s\n", msg.Msg, msg.ID, msg.Topic)
}

// PrintMessage : prints incoming message in the specified channel. >> FOR TESTING <<
func PrintMessage() string {
	msg := <-TopicQueues["Sulav"].Queue
	fmt.Println(msg)
	return msg
}

// CreateTopic : Adds a new channel to the message queue
func CreateTopic(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	requestPayload := struct {
		TopicName string `json:"topic_name"`
	}{}
	responsePayload := struct {
		Status string `json:"status"`
		Err    string `json:"err"`
	}{}

	topic := requestPayload
	topicPayload := json.NewDecoder(req.Body)
	err := topicPayload.Decode(&topic)
	check(err)

	if len(TopicQueues) >= 25 {
		responsePayload.Status = "Failed"
		responsePayload.Err = "Max amount of topics reached."

		fmt.Println(responsePayload)
		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	if _, exists := TopicQueues[topic.TopicName]; exists {
		responsePayload.Status = "Failed"
		responsePayload.Err = "Topic already exists."

		fmt.Println(responsePayload)
		json.NewEncoder(res).Encode(responsePayload)
		return
	}

	newTopicQueue := make(chan string)
	TopicQueues[topic.TopicName] = Topic{topic.TopicName, newTopicQueue}

	responsePayload.Status = "Success"
	responsePayload.Err = fmt.Sprintf("Topic %s Added.", topic.TopicName)

	fmt.Println(responsePayload)
	json.NewEncoder(res).Encode(&responsePayload)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/newTopic", CreateTopic).Methods("POST")
	//router.HandleFunc("/sendMessage", SendMessage).Methods("POST")
	//router.HandleFunc("/jobs", PrintMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
