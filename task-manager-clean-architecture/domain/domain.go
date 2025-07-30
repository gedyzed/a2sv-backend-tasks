package domain

import "context"

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}

type User struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserRepository interface {

	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, username string) error

}

type TaskRepository interface {

	GetByID(ctx context.Context, id string) (*Task, error)
	GetTasks(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, task *Task) (*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id string) error
}
