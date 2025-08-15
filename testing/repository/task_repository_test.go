package repository_test

import (
	"context"
	"errors"
	"task-manager-test/domain"
	"task-manager-test/repository"
	"task-manager-test/repository/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepoTestSuite struct {
	suite.Suite
	repo     domain.TaskRepository
	mockColl *mocks.MongoCollection
}

func (s *TaskRepoTestSuite) SetupTest() {
	s.mockColl = new(mocks.MongoCollection)
	s.repo = repository.NewTaskMongoRepo(s.mockColl)
}

func NewFakeSingleResult(doc interface{}, err error) *mongo.SingleResult {
	return mongo.NewSingleResultFromDocument(doc, err, nil)
}

func (s *TaskRepoTestSuite) TestGetByID_Success() {
	expectedTask := domain.Task{Id: "123", Title: "Test Task"}
	s.mockColl.On("FindOne", mock.Anything, bson.M{"id": "123"}).
		Return(NewFakeSingleResult(expectedTask, nil))

	task, err := s.repo.GetByID(context.Background(), "123")

	s.NoError(err)
	s.Equal(&expectedTask, task)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestGetByID_NotFound() {
	s.mockColl.On("FindOne", mock.Anything, bson.M{"id": "123"}).
		Return(NewFakeSingleResult(nil, errors.New("no documents")))

	task, err := s.repo.GetByID(context.Background(), "123")

	s.Nil(task)
	s.ErrorIs(err, domain.ErrTaskNotFound)
	s.mockColl.AssertExpectations(s.T())
}

func newMockCursor(tasks []domain.Task) *mocks.MongoCursor {
	cursor := new(mocks.MongoCursor)

	for idx := range tasks {
		cursor.On("Next", mock.Anything).Return(true).Once()
		cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			t := args.Get(0).(*domain.Task)
			*t = tasks[idx]
		}).Return(nil).Once()
	}

	cursor.On("Next", mock.Anything).Return(false)
	cursor.On("Close", mock.Anything).Return(nil)

	return cursor
}

func (s *TaskRepoTestSuite) TestGetTasks_FindError() {
	
	s.mockColl.On("Find", mock.Anything, bson.M{}).Return(nil, errors.New("find failed"))

	result, err := s.repo.GetTasks(context.Background())

	s.Nil(result)
	s.ErrorIs(err, domain.ErrWhileReadingData)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestCreate_Success() {
	task := &domain.Task{Id: "123", Title: "New Task"}

	
	s.mockColl.On("FindOne", mock.Anything, bson.M{"id": "123"}).
		Return(NewFakeSingleResult(nil, errors.New("no documents")))

	s.mockColl.On("InsertOne", mock.Anything, task).Return(&mongo.InsertOneResult{}, nil)

	result, err := s.repo.Create(context.Background(), task)

	s.NoError(err)
	s.Equal(task, result)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestCreate_AlreadyExists() {
	task := &domain.Task{Id: "123", Title: "New Task"}


	s.mockColl.On("FindOne", mock.Anything, bson.M{"id": "123"}).
		Return(NewFakeSingleResult(task, nil))

	result, err := s.repo.Create(context.Background(), task)

	s.Nil(result)
	s.ErrorIs(err, domain.ErrWhileReadingData)
	s.mockColl.AssertExpectations(s.T())
}


func (s *TaskRepoTestSuite) TestUpdate_Success() {
	task := &domain.Task{Id: "123", Title: "Updated Task"}

	s.mockColl.On("UpdateOne", mock.Anything, bson.M{"id": "123"}, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	result, err := s.repo.Update(context.Background(), task)

	s.NoError(err)
	s.Equal(task, result)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestUpdate_NotFound() {
	task := &domain.Task{Id: "123", Title: "Updated Task"}

	s.mockColl.On("UpdateOne", mock.Anything, bson.M{"id": "123"}, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0}, nil)

	result, err := s.repo.Update(context.Background(), task)

	s.Nil(result)
	s.ErrorIs(err, domain.ErrTaskNotFound)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestDelete_Success() {
	s.mockColl.On("DeleteOne", mock.Anything, bson.M{"id": "123"}).
		Return(&mongo.DeleteResult{DeletedCount: 1}, nil)

	err := s.repo.Delete(context.Background(), "123")

	s.NoError(err)
	s.mockColl.AssertExpectations(s.T())
}

func (s *TaskRepoTestSuite) TestDelete_NotFound() {
	s.mockColl.On("DeleteOne", mock.Anything, bson.M{"id": "123"}).
		Return(&mongo.DeleteResult{DeletedCount: 0}, nil)

	err := s.repo.Delete(context.Background(), "123")

	s.ErrorIs(err, domain.ErrTaskNotFound)
	s.mockColl.AssertExpectations(s.T())
}



func TestTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
