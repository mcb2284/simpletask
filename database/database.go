package database

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/simpletask/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Database() {
	var err error
	db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("could not open db")
	}

	db.AutoMigrate(&types.User{}, &types.Task{})
}

func CreateUser(newUser types.User) (*types.User, error) {

	var existingUser types.User
	result := db.Where("email = ?", newUser.Email).First(&existingUser)

	if result.Error == nil {
		return &existingUser, nil
	}

	if result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	newUser.ID = uuid.New().String()
	fmt.Printf("Created new user %v\n", newUser)
	db.Create(&newUser)

	user, _ := GetUser(newUser.ID)

	return user, nil
}

func GetUser(user_id string) (*types.User, error) {

	_, err := uuid.Parse(user_id)
	if err != nil {
		return nil, errors.New("bad uuid")
	}

	var user types.User
	db.Where("id = ?", user_id).First(&user)

	return &user, nil
}

func UpdateUser(updatedUser types.User) (*types.User, error) {
	var existingUser types.User
	result := db.Where("id = ?", updatedUser.ID).First(&existingUser)

	if result.Error == nil {

		db.Model(&existingUser).Updates(updatedUser)
		user, _ := GetUser(updatedUser.ID)
		return user, nil
	} else {
		return nil, result.Error
	}
}

func DeleteUser(user_id string) error {

	result := db.Where("id = ?", user_id).Delete(&types.User{})
	if result.Error != nil {
		return errors.New("missing or invalid user_id")
	}
	return nil
}

func GetTask(task_id string) (*types.Task, error) {
	_, err := uuid.Parse(task_id)
	if err != nil {
		return nil, errors.New("bad uuid")
	}

	var task types.Task
	db.Preload("User").Where("id = ?", task_id).First(&task)

	return &task, nil
}

func GetTasks(req types.Request) []types.Task {
	var tasks []types.Task

	query := db.Model(&tasks)

	if req.ID != "" {
		query = query.Where("user_id = ?", req.ID)
	}

	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	orderClause := fmt.Sprintf("%s %s", "due_date", strings.ToUpper(req.Order))
	query = query.Order(orderClause)

	query.Limit(req.Limit).Offset(req.Offset)

	query.Find(&tasks)

	return tasks

}

func CreateTask(newTask types.Task) (*types.Task, error) {
	var existingTask types.Task

	if newTask.UserID == "" {
		return nil, errors.New("missing userID")
	}

	if newTask.IdempKey != "" {
		result := db.Where("idemp_key = ?", newTask.IdempKey).Preload("User").First(&existingTask)
		if result.Error == nil {
			return &existingTask, nil
		}
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	if newTask.Status == "" {
		newTask.Status = "pending"
	}

	newTask.ID = uuid.New().String()
	fmt.Printf("Created new user %v\n", newTask)
	db.Create(&newTask)

	task, _ := GetTask(newTask.ID)

	return task, nil

}

func UpdateTask(updatedTask types.Task) (*types.Task, error) {
	var existingTask types.Task
	result := db.Where("id = ?", updatedTask.ID).First(&existingTask)

	if result.Error == nil {
		db.Model(&existingTask).Updates(updatedTask)
		task, _ := GetTask(updatedTask.ID)
		return task, nil
	} else {
		return nil, result.Error
	}
}

func DeleteTask(task_id string) error {

	result := db.Where("id = ?", task_id).Delete(&types.Task{})
	if result.Error != nil {
		return errors.New("missing or invalid task_id")
	}
	return nil
}
