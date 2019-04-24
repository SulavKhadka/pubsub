package msgqueue

import (
	"encoding/json"
	"fmt"
	"log"
)

// Message :Struct for storing JSON message payloads
type Message struct {
	Msg   string `json:"msg"`
	ID    int    `json:"id"`
	Topic string `json:"topic"`
}

// Topic : Struct for storing message queue topics
type Topic struct {
	TopicName   string   `json:"topic_name"`
	ID          int      `json:"id"`
	Subscribers []string `json:"subscribers"`
}

// TopicArray : Holds all the topics for a given instance of the message queue
var TopicArray []string

// SendMessage : parses and drops message to the appropiate channel as defined in the incoming payload
func SendMessage(msg1 string) string {
	msgStruct := Message{}
	err := json.Unmarshal([]byte(msg1), &msgStruct)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	for _, topic := range TopicArray {
		if topic == msgStruct.Topic {
			return fmt.Sprintf("Recieved Msg %s.\nID: %d\nTopic: %s\n", msgStruct.Msg, msgStruct.ID, msgStruct.Topic)
		}
	}

	return "Topic not found."
}

func subscribeToTopic() {

}

// CreateTopic : Adds a new channel to the message queue
func CreateTopic(newTopic string) string {
	if len(TopicArray) >= 25 {
		return "Max amount of topics reached."
	}

	for _, topic := range TopicArray {
		if topic == newTopic {
			return "Topic already exists."
		}
	}
	TopicArray = append(TopicArray, newTopic)
	return fmt.Sprintf("Topic %s Added.", newTopic)
}
