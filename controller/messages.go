package controller

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

// ForumMessage represents messages in /message
type ForumMessage struct {
	Id         uuid.UUID `json:"id"`
	AuthorId   string    `json:"author_id"`
	AuthorName string    `json:"author_name"`
	Text       string    `json:"text"`
	PostedAt   time.Time `json:"posted_at"`
}

// Message represents message instance from DB
type Message struct {
	Id         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	CategoryId uuid.UUID `json:"category_id"`
	PostedAt   time.Time `json:"posted_at"`
	AuthorId   uuid.UUID `json:"author_id"`
}

func (c *Controller) GetMessages(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	messages, err := c.getMessages(id)
	if err != nil {
		log.Printf("Cannot extract messages from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		log.Printf("Error encoding messages to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
//80c79ade-c37e-462a-9c91-e4805fd2b183
func (c *Controller) getMessages(id string) ([]ForumMessage, error) {
	rows, err := c.DB.Query(
		fmt.Sprintf(`
			SELECT m.id, u.id, u.name, m.text, m.posted_at 
			FROM messages m
			INNER JOIN users u ON m.author_id = u.id AND m.category_id = '%v'
			ORDER BY m.posted_at DESC `, id,
		),
	)
	if err != nil {
		return make([]ForumMessage, 0), err
	}

	var messages []ForumMessage
	for rows.Next() {
		msg := ForumMessage{}
		err := rows.Scan(&msg.Id, &msg.AuthorId, &msg.AuthorName, &msg.Text, &msg.PostedAt)
		if err != nil {
			return make([]ForumMessage, 0), err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return make([]ForumMessage, 0), err
	}

	if len(messages) == 0 {
		return make([]ForumMessage, 0), nil
	}
	return messages, nil
}

func (c *Controller) AddMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	message.Text = r.FormValue("text")
	message.CategoryId, _ = uuid.FromString(r.FormValue("category_id"))

	message.Id, _ = uuid.NewV4()

	row := c.DB.QueryRow("SELECT id FROM users ORDER BY id LIMIT 1")
	if err := row.Scan(&message.AuthorId); err != nil {
		log.Printf("Cannot get first user from db, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := c.DB.Exec("INSERT INTO messages VALUES($1, $2, $3, $4, $5)",
		message.Id, message.Text, message.CategoryId, time.Now(), message.AuthorId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Cannot execute message, error: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("Error encoding message to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
