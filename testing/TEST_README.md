# Go Tests for Clean Architecture Project

This directory contains comprehensive unit tests for all components of the clean architecture project. The tests are organized by layer and follow Go testing best practices.

## 📁 Test Structure

```
├── domain/
│   ├── domain_test.go          # Domain entities and interfaces tests
├── usecases/
│   ├── user_usecases_test.go   # User business logic tests
│   └── task_usecases_test.go   # Task business logic tests
├── repository/
│   ├── user_repository_test.go # User data access tests
│   └── task_repository_test.go # Task data access tests
├── infrastructure/
│   ├── password_service_test.go # Password hashing tests
│   └── jwt_service_test.go     # JWT token tests
├── delivery/
│   ├── controllers/
│   │   └── controller_test.go  # HTTP controller tests
│   └── routers/
│       └── route_test.go       # Router configuration tests
├── test_runner.go              # Test runner script
└── TEST_README.md              # This file
```

## 🚀 Running Tests

### Quick Start
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Using the Test Runner
```bash
# Run the custom test runner
go run test_runner.go
```

### Running Specific Package Tests
```bash
# Test domain layer
go test ./domain

# Test use cases
go test ./usecases

# Test repositories
go test ./repository

# Test infrastructure
go test ./infrastructure

# Test controllers
go test ./delivery/controllers

# Test routers
go test ./delivery/routers
```

## 🧪 Test Coverage

### Domain Layer (`domain/domain_test.go`)
- **Entity Tests**: Validates User and Task structs
- **Interface Tests**: Tests repository interfaces with mock implementations
- **JSON Tag Tests**: Ensures proper JSON serialization
- **Mock Repository Tests**: Tests mock implementations of UserRepository and TaskRepository

**Key Test Cases:**
- Task and User struct field validation
- Mock repository CRUD operations
- Error handling for not found scenarios
- Data consistency checks

### Use Cases Layer (`usecases/`)
#### User Use Cases (`user_usecases_test.go`)
- **Registration Tests**: User creation with password hashing
- **Login Tests**: Authentication with password verification
- **Admin Promotion Tests**: Role elevation functionality
- **Error Handling**: Duplicate users, invalid credentials, service errors

**Key Test Cases:**
- Successful user registration
- Duplicate username handling
- Password hashing integration
- Login with correct/incorrect credentials
- Admin promotion workflow
- Service layer error propagation

#### Task Use Cases (`task_usecases_test.go`)
- **CRUD Operations**: Create, Read, Update, Delete tasks
- **Error Handling**: Task not found, validation errors
- **Integration Tests**: Full workflow testing

**Key Test Cases:**
- Task creation and retrieval
- Task updates and status changes
- Task deletion and verification
- Empty task list handling
- Error scenarios (not found, invalid data)

### Repository Layer (`repository/`)
#### User Repository (`user_repository_test.go`)
- **Database Operations**: Mock database interactions
- **Data Persistence**: User storage and retrieval
- **Error Scenarios**: Database errors, not found cases

**Key Test Cases:**
- User creation and storage
- User retrieval by username
- User role updates
- Duplicate user handling
- Database error simulation

#### Task Repository (`task_repository_test.go`)
- **Task CRUD**: Complete task lifecycle management
- **Bulk Operations**: Multiple task handling
- **Data Integrity**: Task data consistency

**Key Test Cases:**
- Task creation and storage
- Task retrieval by ID
- Task updates and modifications
- Task deletion and cleanup
- Multiple task operations

### Infrastructure Layer (`infrastructure/`)
#### Password Service (`password_service_test.go`)
- **Password Hashing**: bcrypt implementation testing
- **Password Verification**: Hash comparison testing
- **Security Tests**: Salt generation, hash uniqueness

**Key Test Cases:**
- Password hashing functionality
- Password verification (correct/incorrect)
- Hash uniqueness (same password, different hashes)
- Empty password handling
- Special character passwords

#### JWT Service (`jwt_service_test.go`)
- **Token Generation**: JWT creation and validation
- **User Data**: Token payload verification
- **Error Handling**: Invalid user data

**Key Test Cases:**
- Token generation for valid users
- Token uniqueness for different users
- Token consistency for same user
- Special character handling in usernames
- Empty user data handling

### Delivery Layer (`delivery/`)
#### Controllers (`controllers/controller_test.go`)
- **HTTP Endpoints**: All API endpoint testing
- **Request/Response**: JSON handling and status codes
- **Error Handling**: HTTP error responses
- **Middleware**: Authentication and validation

**Key Test Cases:**
- User registration endpoint
- User login endpoint
- Admin promotion endpoint
- Task CRUD endpoints
- Invalid input handling
- HTTP status code verification
- JSON response validation

#### Routers (`routers/route_test.go`)
- **Route Configuration**: Endpoint registration
- **HTTP Methods**: GET, POST, PUT, DELETE handling
- **Middleware**: CORS, authentication middleware
- **Error Handling**: 404, 405 status codes

**Key Test Cases:**
- Route existence verification
- HTTP method validation
- API prefix validation
- Route parameter handling
- CORS middleware testing
- Non-existent route handling

## 🛠️ Mock Implementations

The test suite includes comprehensive mock implementations:

### Mock Database
- **MockDB**: In-memory database for testing
- **MockTaskDB**: Task-specific database operations
- **Data Persistence**: Simulates real database behavior

### Mock Use Cases
- **MockUserUsecase**: User business logic simulation
- **MockTaskUsecase**: Task business logic simulation
- **Error Simulation**: Controlled error scenarios

### Mock Services
- **MockOtherServices**: Password and JWT service simulation
- **Configurable Behavior**: Customizable mock responses

## 📊 Test Metrics

### Coverage Goals
- **Line Coverage**: >90%
- **Function Coverage**: >95%
- **Branch Coverage**: >85%

### Test Categories
- **Unit Tests**: Individual component testing
- **Integration Tests**: Component interaction testing
- **Error Tests**: Error scenario validation
- **Edge Case Tests**: Boundary condition testing

## 🔧 Test Configuration

### Environment Variables
```bash
# Set test mode
export GIN_MODE=test

# Database configuration (if needed)
export TEST_DB_URL="test_database_url"
```

### Test Flags
```bash
# Run tests with race detection
go test -race ./...

# Run tests with timeout
go test -timeout 30s ./...

# Run tests with CPU profiling
go test -cpuprofile=cpu.prof ./...

# Run tests with memory profiling
go test -memprofile=mem.prof ./...
```

## 🐛 Debugging Tests

### Verbose Output
```bash
# Run specific test with verbose output
go test -v -run TestUserUsecase_Register_Success ./usecases
```

### Test Failures
```bash
# Run tests and stop on first failure
go test -failfast ./...

# Run tests with detailed output
go test -v -count=1 ./...
```

### Coverage Analysis
```bash
# Generate detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

## 📝 Best Practices

### Test Naming
- Use descriptive test names: `TestComponent_Method_Scenario`
- Include expected outcome in test name
- Use table-driven tests for multiple scenarios

### Test Organization
- Group related tests together
- Use subtests for complex scenarios
- Separate setup, execution, and verification

### Mock Usage
- Use mocks for external dependencies
- Keep mocks simple and focused
- Test mock behavior independently

### Error Testing
- Test both success and failure scenarios
- Verify error messages and types
- Test edge cases and boundary conditions

## 🚨 Common Issues

### Import Errors
```bash
# Ensure all dependencies are available
go mod tidy
go mod download
```

### Test Timeouts
```bash
# Increase timeout for slow tests
go test -timeout 60s ./...
```

### Coverage Issues
```bash
# Check for untested code
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | grep -E "(0\.0%|100\.0%)"
```

## 📚 Additional Resources

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Go Test Coverage](https://blog.golang.org/cover)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go Testing Best Practices](https://github.com/golang/go/wiki/CodeReviewComments#tests)

## 🤝 Contributing

When adding new tests:
1. Follow the existing naming conventions
2. Include both success and failure scenarios
3. Add appropriate mock implementations
4. Update this README if adding new test categories
5. Ensure tests pass before submitting

## 📞 Support

For test-related issues:
1. Check the test output for specific error messages
2. Verify all dependencies are properly installed
3. Ensure the test environment is correctly configured
4. Review the test documentation for common solutions
