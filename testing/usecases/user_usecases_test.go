package usecases_test

import (
	"context"
	"errors"
	"testing"
	"task-manager-test/domain"
	domainMocks "task-manager-test/domain/mocks"
	usecaseMocks"task-manager-test/usecases/mocks"
	"task-manager-test/usecases"

	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	repo     *domainMocks.UserRepository
	services *usecaseMocks.OtherServices
	usecase  usecases.UserUsecase
	ctx      context.Context
}


func (s *UserUsecaseSuite) SetupTest() {
	s.repo = new(domainMocks.UserRepository)
	s.services = new(usecaseMocks.OtherServices)
	s.usecase = usecases.NewUserUsecase(s.repo, s.services)
	s.ctx = context.Background()
}


func (s *UserUsecaseSuite) TestRegister_Success() {

	user := &domain.User{Username: "test", Password: "123"}

	s.repo.On("GetByUsername", s.ctx, "test").Return(nil, nil)
	s.services.On("HashPassword", "123").Return("hashed", nil)
	s.repo.On("Create", s.ctx, &domain.User{Username: "test", Password: "hashed"}).
		Return(&domain.User{UserID: "1", Username: "test", Password: "hashed"}, nil)

	result, err := s.usecase.Register(s.ctx, user)
	s.NoError(err)
	s.Equal("1", result.UserID)
	s.Equal("hashed", result.Password)
	s.repo.AssertExpectations(s.T())
	s.services.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestRegister_UsernameTaken() {

	existing := &domain.User{UserID: "1", Username: "test"}
	s.repo.On("GetByUsername", s.ctx, "test").Return(existing, nil)

	result, err := s.usecase.Register(s.ctx, &domain.User{Username: "test"})
	s.ErrorIs(err, domain.ErrUserAlreadyExists)
	s.Nil(result)
	s.repo.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestRegister_HashError() {

	s.repo.On("GetByUsername", s.ctx, "test").Return(nil, nil)
	s.services.On("HashPassword", "123").Return("", errors.New("hash fail"))

	result, err := s.usecase.Register(s.ctx, &domain.User{Username: "test", Password: "123"})
	s.Error(err)
	s.Nil(result)
	s.services.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestLogin_Success() {

	user := &domain.User{Username: "test", Password: "123"}
	existing := &domain.User{Username: "test", Password: "hashed"}

	s.repo.On("GetByUsername", s.ctx, "test").Return(existing, nil)
	s.services.On("CompareHashAndPassword", "hashed", "123").Return(nil)
	s.services.On("GenerateToken", user).Return("token123", nil)

	token, err := s.usecase.Login(s.ctx, user)
	s.NoError(err)
	s.Equal("token123", token)
	s.repo.AssertExpectations(s.T())
	s.services.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestLogin_InvalidUsername() {

	s.repo.On("GetByUsername", s.ctx, "test").Return(nil, errors.New("not found"))

	token, err := s.usecase.Login(s.ctx, &domain.User{Username: "test", Password: "123"})
	s.ErrorIs(err, domain.ErrUsernameOrPassword)
	s.Empty(token)
}

func (s *UserUsecaseSuite) TestLogin_InvalidPassword() {

	user := &domain.User{Username: "test", Password: "wrong"}
	existing := &domain.User{Username: "test", Password: "hashed"}

	s.repo.On("GetByUsername", s.ctx, "test").Return(existing, nil)
	s.services.On("CompareHashAndPassword", "hashed", "wrong").Return(errors.New("mismatch"))

	token, err := s.usecase.Login(s.ctx, user)
	s.ErrorIs(err, domain.ErrUsernameOrPassword)
	s.Empty(token)
}

func (s *UserUsecaseSuite) TestPromoteAdmin_Success() {

	s.repo.On("Update", s.ctx, "test").Return(nil)
	err := s.usecase.PromoteAdmin(s.ctx, "test")
	s.NoError(err)
	s.repo.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestPromoteAdmin_Error() {
	
	s.repo.On("Update", s.ctx, "test").Return(errors.New("db fail"))
	err := s.usecase.PromoteAdmin(s.ctx, "test")
	s.Error(err)
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
