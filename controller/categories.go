package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Category struct {
	Id       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ParentId sql.NullString `json:"parent_id,omitempty"`
}

func (c *Controller) GetCategory(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	var row *sql.Row
	if len(id) == 0 || id == "" {
		row = c.DB.QueryRow("SELECT * FROM categories AS c WHERE c.name = 'Forum'")
	} else {
		row = c.DB.QueryRow("SELECT * FROM categories AS c WHERE c.id = $1", id)
	}

	ct := Category{}
	if err := row.Scan(&ct.Id, &ct.Name, &ct.ParentId); err != nil {
		log.Printf("Cannot scan root category, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(ct)
	if err != nil {
		log.Printf("Error encoding root category to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	categories, err := c.getSubcategories(id)
	if err != nil {
		log.Printf("Cannot extract subcategories from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Printf("Error encoding subcategories to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getSubcategories(id string) ([]Category, error) {
	rows, err := c.DB.Query(
		fmt.Sprintf(`
			SELECT * FROM categories AS c 
			WHERE c.parent_id = '%v' 
			ORDER BY c.name`, id,
		),
	)
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

func (c *Controller) GetCategories(w http.ResponseWriter, r *http.Request) {
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
	rows, err := c.DB.Query(
		"SELECT * FROM categories ORDER BY categories.name",
	)
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