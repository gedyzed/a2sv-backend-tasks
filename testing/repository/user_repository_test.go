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


type UserRepoTestSuite struct {
	suite.Suite
	repo     domain.UserRepository
	mockColl *mocks.MongoCollection
}

func (s *UserRepoTestSuite) SetupTest() {
	s.mockColl = new(mocks.MongoCollection)
	s.repo = repository.NewUserMongoRepo(s.mockColl)
}

// Create Tests 

func (s *UserRepoTestSuite) TestCreate_FirstUser_AdminRole() {
	
	user := &domain.User{Username: "admin"}

	mockCursor := new(mocks.MongoCursor) // empty DB
	mockCursor.On("Next", mock.Anything).Return(false)
	mockCursor.On("Close", mock.Anything).Return(nil)

	s.mockColl.On("Find", mock.Anything, bson.M{}).Return(mockCursor, nil)
	s.mockColl.On("InsertOne", mock.Anything, user).
		Return(&mongo.InsertOneResult{}, nil)

	created, err := s.repo.Create(context.Background(), user)
	s.NoError(err)
	s.Equal("admin", created.Role)
	s.mockColl.AssertExpectations(s.T())
	mockCursor.AssertExpectations(s.T())
}

func (s *UserRepoTestSuite) TestCreate_NotFirstUser_RegularRole() {

    user := &domain.User{Username: "user1"}

    mockCursor := new(mocks.MongoCursor)
    mockCursor.On("Next", mock.Anything).Return(true).Once() 
    mockCursor.On("Close", mock.Anything).Return(nil)

    s.mockColl.On("Find", mock.Anything, bson.M{}).Return(mockCursor, nil)
    s.mockColl.On("InsertOne", mock.Anything, user).
        Return(&mongo.InsertOneResult{}, nil)

    created, err := s.repo.Create(context.Background(), user)
    s.NoError(err)
    s.Equal("regular", created.Role) 
    s.mockColl.AssertExpectations(s.T())
    mockCursor.AssertExpectations(s.T())
}

// --- Update Tests ---

func (s *UserRepoTestSuite) TestUpdate_Success() {
	s.mockColl.On("UpdateOne", mock.Anything, bson.M{"username": "user1"}, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	err := s.repo.Update(context.Background(), "user1")
	s.NoError(err)
	s.mockColl.AssertExpectations(s.T())
}

func (s *UserRepoTestSuite) TestUpdate_NotFound() {
	s.mockColl.On("UpdateOne", mock.Anything, bson.M{"username": "user1"}, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0}, nil)

	err := s.repo.Update(context.Background(), "user1")
	s.Error(err) 
	s.mockColl.AssertExpectations(s.T())
}

// --- GetByUsername Tests ---

func (s *UserRepoTestSuite) TestGetByUsername_Success() {

	expectedUser := domain.User{Username: "user1", Role: "regular"}
	s.mockColl.On("FindOne", mock.Anything, bson.M{"username": "user1"}).
		Return(mongo.NewSingleResultFromDocument(expectedUser, nil, nil))

	user, err := s.repo.GetByUsername(context.Background(), "user1")
	s.NoError(err)
	s.Equal(&expectedUser, user)
	s.mockColl.AssertExpectations(s.T())
}

func (s *UserRepoTestSuite) TestGetByUsername_NotFound() {
	
	s.mockColl.On("FindOne", mock.Anything, bson.M{"username": "user1"}).
		Return(mongo.NewSingleResultFromDocument(nil, mongo.ErrNoDocuments, nil))

	user, err := s.repo.GetByUsername(context.Background(), "user1")
	s.Nil(user)
	s.ErrorIs(err, domain.ErrWhileReadingData)
	s.mockColl.AssertExpectations(s.T())
}



func (s *UserRepoTestSuite) TestGetByUsername_OtherError() {
	s.mockColl.On("FindOne", mock.Anything, bson.M{"username": "user1"}).
		Return(mongo.NewSingleResultFromDocument(nil, errors.New("db error"), nil))

	user, err := s.repo.GetByUsername(context.Background(), "user1")
	s.Nil(user)
	s.ErrorIs(err, domain.ErrWhileReadingData)
	s.mockColl.AssertExpectations(s.T())
}


func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
