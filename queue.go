package main

import (
	"time"
)

// Topic the struct for storing message queue topics
type Topic struct {
	TopicName string `json:"topic_name"`
	Queue     []Item `json:"queue"`
}

// Item the payload struct for the queue
type Item struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
	Topic     string `json:"topic"`
}

//New creates and returns a new queue
func (q *Topic) New(topicName string) []Item {
	var queue []Item
	return queue
}

//Insert adds an item to the bottom of the queue
func (q *Topic) Insert(targetQueue Topic, message string) {

	item := Item{}
	item.ID = len(targetQueue.Queue) + 1
	item.TimeStamp = time.Now().Format(time.RFC3339)
	item.Message = message
	item.Topic = targetQueue.TopicName

	targetQueue.Queue = append(targetQueue.Queue, item)
}

//Delete removes the top item from the queue
func (q *Topic) Delete(targetQueue Topic) {

	targetQueue.Queue = targetQueue.Queue[1:len(targetQueue.Queue)]
}
