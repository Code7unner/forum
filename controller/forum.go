package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Category struct {
	Id       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ParentId sql.NullString `json:"parent_id,omitempty"`
}

// @route  GET api/forum/categories
// @desc   Get all categories fields from db
// @access Private
func (c *Controller) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := c.getCategories()

	if err != nil {
		log.Printf("Cannot extract categories from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Printf("Error encoding categories to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getCategories() ([]Category, error) {
	rows, err := c.DB.Query("SELECT * FROM categories ORDER BY categories.name LIMIT 10")
	if err != nil {
		return make([]Category, 0), err
	}

	var categories []Category
	for rows.Next() {
		ct := Category{}
		err := rows.Scan(&ct.Id, &ct.Name, &ct.ParentId)
		if err != nil {
			return make([]Category, 0), err
		}
		categories = append(categories, ct)
	}

	if err = rows.Err(); err != nil {
		return make([]Category, 0), err
	}

	if len(categories) == 0 {
		return make([]Category, 0), nil
	}
	return categories, nil
}