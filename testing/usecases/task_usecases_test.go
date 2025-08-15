package usecases_test

import (
	"context"
	"errors"
	domain "task-manager-test/domain"
	"task-manager-test/domain/mocks"
	"task-manager-test/usecases"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskUsecaseSuite struct {
	suite.Suite
	repo    *mocks.TaskRepository
	usecase usecases.TaskUsecase
}

func (s *TaskUsecaseSuite) SetupTest() {
	s.repo = new(mocks.TaskRepository)
	s.usecase = usecases.NewTaskUsecase(s.repo)
}


func (s *TaskUsecaseSuite) TestGetTasks_Success() {
	ctx := context.Background()
	tasks := []domain.Task{
		{Id: "1", Title: "test1"},
		{Id: "2", Title: "test2"},
	}

	s.repo.On("GetTasks", ctx).Return(tasks, nil)
	result, err := s.usecase.GetTasks(ctx)

	s.NoError(err)
	s.Equal(tasks, result)
	s.repo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestGetTasks_Error() {
	ctx := context.Background()
	mockErr := errors.New("db: error")

	s.repo.On("GetTasks", ctx).Return(nil, mockErr)
	result, err := s.usecase.GetTasks(ctx)

	s.Error(err)
	s.Equal(mockErr, err)
	s.Nil(result)
	s.repo.AssertExpectations(s.T())
}

// -------------------- GetTaskById --------------------

func (s *TaskUsecaseSuite) TestGetTaskById_Success() {
	ctx := context.Background()
	task := &domain.Task{Id: "1", Title: "test task"}

	s.repo.On("GetByID", ctx, "1").Return(task, nil)
	result, err := s.usecase.GetTaskById(ctx, "1")

	s.NoError(err)
	s.Equal(task, result)
	s.repo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestGetTaskById_Error() {
	ctx := context.Background()
	mockErr := errors.New("not found")

	s.repo.On("GetByID", ctx, "1").Return(nil, mockErr)
	result, err := s.usecase.GetTaskById(ctx, "1")

	s.Error(err)
	s.Equal(mockErr, err)
	s.Nil(result)
	s.repo.AssertExpectations(s.T())
}


func (s *TaskUsecaseSuite) TestAddTask_Success() {
	ctx := context.Background()
	task := &domain.Task{Id: "1", Title: "new task"}

	s.repo.On("Create", ctx, task).Return(task, nil)
	result, err := s.usecase.AddTask(ctx, task)

	s.NoError(err)
	s.Equal(task, result)
	s.repo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestAddTask_Error() {
	ctx := context.Background()
	task := &domain.Task{Id: "1", Title: "new task"}
	mockErr := errors.New("insert failed")

	s.repo.On("Create", ctx, task).Return(nil, mockErr)
	result, err := s.usecase.AddTask(ctx, task)

	s.Error(err)
	s.Equal(mockErr, err)
	s.Nil(result)
	s.repo.AssertExpectations(s.T())
}


func (s *TaskUsecaseSuite) TestDeleteTask_Success() {
	ctx := context.Background()

	s.repo.On("Delete", ctx, "1").Return(nil)
	err := s.usecase.DeleteTask(ctx, "1")

	s.NoError(err)
	s.repo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestDeleteTask_Error() {
	ctx := context.Background()
	mockErr := errors.New("delete failed")

	s.repo.On("Delete", ctx, "1").Return(mockErr)
	err := s.usecase.DeleteTask(ctx, "1")

	s.Error(err)
	s.Equal(mockErr, err)
	s.repo.AssertExpectations(s.T())
}


func (s *TaskUsecaseSuite) TestUpdateTask_Success() {
	ctx := context.Background()
	task := &domain.Task{Id: "1", Title: "updated task"}

	s.repo.On("Update", ctx, task).Return(task, nil)
	result, err := s.usecase.UpdateTask(ctx, task)

	s.NoError(err)
	s.Equal(task, result)
	s.repo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestUpdateTask_Error() {
	ctx := context.Background()
	task := &domain.Task{Id: "1", Title: "updated task"}
	mockErr := errors.New("update failed")

	s.repo.On("Update", ctx, task).Return(nil, mockErr)
	result, err := s.usecase.UpdateTask(ctx, task)

	s.Error(err)
	s.Equal(mockErr, err)
	s.Nil(result)
	s.repo.AssertExpectations(s.T())
}


func TestTaskUsecase(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
