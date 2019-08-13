package main

import (
	"fmt"
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
func (q *Topic) New(topicName string) Topic {
	queue := make([]Item, 0, 10)
	newTopic := Topic{topicName, queue}
	return newTopic
}

//Insert adds an item to the bottom of the queue
func (q *Topic) Insert(message string) {

	item := Item{}
	item.ID = len(q.Queue) + 1
	item.TimeStamp = time.Now().Format(time.RFC3339)
	item.Message = message
	item.Topic = q.TopicName

	q.Queue = append(q.Queue, item)
	fmt.Println(len(q.Queue))
}

// Get retrieves the first item in the queue
func (q *Topic) Get() Item {

	if len(q.Queue) > 0 {
		item := q.Queue[0]
		q.Delete()
		return item
	}

	return Item{}
}

//Delete removes the top item from the queue
func (q *Topic) Delete() {
	fmt.Println(len(q.Queue))
	q.Queue = q.Queue[1:len(q.Queue)]
}
