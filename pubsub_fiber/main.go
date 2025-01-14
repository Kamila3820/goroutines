package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type PubSub struct {
	subs []chan Message
	mu   sync.Mutex
}

// Adds a new subscriber channel
func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	chann := make(chan Message, 1)
	ps.subs = append(ps.subs, chann)

	return chann
}

// Sends a message to all subscriber channels
func (ps *PubSub) Publish(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, sub := range ps.subs {
		sub <- *msg
	}
}

// Removes a subscriber channel
func (ps *PubSub) UnSub(chann chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i, sub := range ps.subs {
		if sub == chann {
			ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			close(chann)
			break
		}
	}
}

func main() {
	app := fiber.New()

	pubsub := &PubSub{}

	app.Post("/publisher", func(c *fiber.Ctx) error {
		message := new(Message)
		if err := c.BodyParser(message); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		pubsub.Publish(message)
		return c.JSON(&fiber.Map{
			"message": "Add to subscriber",
		})
	})

	sub1 := pubsub.Subscribe()
	go func() {
		for msg := range sub1 {
			fmt.Printf("Sub1: receive %v from a publisher\n", msg)
		}
	}()

	sub2 := pubsub.Subscribe()
	go func() {
		for msg := range sub2 {
			fmt.Printf("Sub2: receive %v from a publisher\n", msg)
		}
	}()

	app.Listen(":8888")
}
