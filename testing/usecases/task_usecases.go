package usecases

import (
	"context"
	"task-manager-test/domain"
)

type TaskUsecase interface {

	GetTasks(ctx context.Context)([]domain.Task, error)
	GetTaskById(ctx context.Context, id string)(*domain.Task, error)
	AddTask(ctx context.Context, task *domain.Task)(*domain.Task, error)
	DeleteTask(ctx context.Context, id string) error 
	UpdateTask(ctx context.Context, task *domain.Task)(*domain.Task, error)
}

type taskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (t *taskUsecase) GetTasks(ctx context.Context)([]domain.Task, error){
	return t.repo.GetTasks(ctx)
}

func (t *taskUsecase) GetTaskById(ctx context.Context, id string)(*domain.Task, error){
	return t.repo.GetByID(ctx, id)
}


func (t *taskUsecase) AddTask(ctx context.Context, task *domain.Task)(*domain.Task, error){
	return t.repo.Create(ctx, task)
}


func (t *taskUsecase) DeleteTask(ctx context.Context, id string) error {
	return t.repo.Delete(ctx, id)
}


func (t *taskUsecase) UpdateTask(ctx context.Context, task *domain.Task)(*domain.Task, error){
	return t.repo.Update(ctx, task)
}



