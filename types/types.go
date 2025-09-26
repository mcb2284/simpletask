package types

import (
	"fmt"
	"time"
)

type Request struct {
	ID     string `form:"id"`
	Status string `form:"status"`
	Order  string `form:"order"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type Task struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	Status   string    `json:"status"`
	DueDate  time.Time `json:"due_date"`
	IdempKey string    `json:"idemp_key,omitempty" gorm:"uniqueIndex"`
	UserID   string    `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
}

type TaskSummary struct {
	Pending    int `json:"pending"`
	InProgress int `json:"in_progress"`
	Done       int `json:"done"`
}

type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (req *Request) Validate() error {

	validStatuses := map[string]bool{
		"pending":     true,
		"in_progress": true,
		"done":        true,
	}

	validOrders := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if req.Status != "" && !validStatuses[req.Status] {
		return fmt.Errorf("invalid status: %s", req.Status)
	}

	if req.Order != "" && !validOrders[req.Order] {
		return fmt.Errorf("invalid order: %s", req.Order)
	}

	return nil

}

// Set defaults
func (req *Request) SetDefaults() {
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}
	if req.Order == "" {
		req.Order = "desc"
	}
}
