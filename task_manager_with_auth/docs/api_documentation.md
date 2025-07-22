```markdown
## **User Management API Documentation

## 1. Register User
**Endpoint:** `POST /register`  
Registers a new user. The first registered user becomes an admin; subsequent users are regular users.

### Request
**Body (application/json):**
```json
{
  "username": "string",
  "password": "string"
}
```

### Responses
- **200 OK** (Success):
  ```json
  {
    "message": "User Inserted Successfully",
    "user": {
      "username": "gedion",
      "role": "admin",
      "password": "<hashed_value>"
    }
  }
  ```
- **400 Bad Request** (Invalid input format):
  ```json
  {"error": "Invalid Input Format"}
  ```
- **409 Conflict** (Username exists):
  ```json
  {"error": "username already exists"}
  ```
- **500 Internal Server Error** (Server error):
  ```json
  {"error": "failed to hash password" | "failed to insert user"}
  ```

---

## 2. Login User
**Endpoint:** `POST /login`  
Authenticates a user and returns a JWT token for authorization.

### Request
**Body (application/json):**
```json
{
  "username": "string",
  "password": "string"
}
```

### Responses
- **200 OK** (Success):
  ```json
  {
    "message": "User Logged Successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```
- **401 Unauthorized** (Invalid credentials/input):
  ```json
  {"error": "invalid username or password" | "Invalid input format"}
  ```
- **500 Internal Server Error** (Token generation failed):
  ```json
  {"error": "failed to generate token"}
  ```

---

## 3. Promote User to Admin
**Endpoint:** `POST /promote-admin`  
*Requires admin privileges*  
Promotes a regular user to admin role.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Body (application/json):**
```json
{
  "username": "string"
}
```

### Responses
- **200 OK** (Success):
  ```json
  {"message": "User has been Promoted to admin"}
  ```
- **200 OK** (User already admin):
  ```json
  {"message": "user already has admin role"}
  ```
- **400 Bad Request** (User not found):
  ```json
  {"message": "user not found"}
  ```
- **401 Unauthorized** (Missing/invalid token or non-admin):
  ```json
  {"error": "Unauthorized"}
  ```
- **500 Internal Server Error** (Update failure):
  ```json
  {"message": "error while updating data"}
  ```

---

## Task Management API Documentation 

## 1. Get All Tasks
**Endpoint:** `GET /tasks`  
*Requires authentication*  
Retrieves all tasks from the system.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

### Responses
- **200 OK** (Success):
  ```json
  {
    "tasks": [
      {
        "id": "task123",
        "description": "Complete API documentation",
        "status": "in-progress"
      },
      {
        "id": "task456",
        "description": "Test endpoints",
        "status": "pending"
      }
    ]
  }
  ```
- **401 Unauthorized** (Missing/invalid token):
  ```json
  {"error": "Unauthorized"}
  ```
- **500 Internal Server Error** (Database error):
  ```json
  {"error": "Error while reading data"}
  ```

---

## 2. Get Task by ID
**Endpoint:** `GET /tasks/:id`  
*Requires authentication*  
Retrieves a single task by its ID.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Path Parameters:**
- `id` (string): ID of the task to retrieve

### Responses
- **200 OK** (Success):
  ```json
  {
    "task": {
      "id": "task123",
      "description": "Complete API documentation",
      "status": "in-progress"
    }
  }
  ```
- **401 Unauthorized** (Missing/invalid token):
  ```json
  {"error": "Unauthorized"}
  ```
- **404 Not Found** (Task not found):
  ```json
  {"message": "Task Not Found"}
  ```
- **500 Internal Server Error** (Database error):
  ```json
  {"message": "Error while file reading"}
  ```

---

## 3. Create Task
**Endpoint:** `POST /tasks/create`  
*Requires admin privileges*  
Creates a new task in the system.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Body (application/json):**
```json
{
  "id": "task789",
  "description": "New task description",
  "status": "pending"
}
```

### Responses
- **201 Created** (Success):
  ```json
  {
    "message": "successfully created",
    "task": {
      "id": "task789",
      "description": "New task description",
      "status": "pending"
    }
  }
  ```
- **400 Bad Request** (Invalid input or duplicate ID):
  ```json
  {"message": "Invalid Input Format" | "Task Id already exists"}
  ```
- **401 Unauthorized** (Missing/invalid token or non-admin):
  ```json
  {"error": "Unauthorized"}
  ```
- **500 Internal Server Error** (Database error):
  ```json
  {"message": "Cannot Insert a Record"}
  ```

---

## 4. Update Task
**Endpoint:** `PUT /tasks/edit/:id`  
*Requires admin privileges*  
Updates an existing task. Only non-empty fields will be updated.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Path Parameters:**
- `id` (string): ID of the task to update

**Body (application/json):**
```json
{
  "description": "Updated description",
  "status": "completed"
}
```

### Responses
- **200 OK** (Success):
  ```json
  {
    "message": "Updated Successfully",
    "updated_fields": {
      "description": "Updated description",
      "status": "completed"
    }
  }
  ```
- **400 Bad Request** (Invalid input):
  ```json
  {"error": "Invalid Input Format"}
  ```
- **401 Unauthorized** (Missing/invalid token or non-admin):
  ```json
  {"error": "Unauthorized"}
  ```
- **404 Not Found** (Task not found):
  ```json
  {"message": "Task Not Found"}
  ```
- **500 Internal Server Error** (Database error):
  ```json
  {"error": "Cannot Update Task"}
  ```

---

## 5. Delete Task
**Endpoint:** `DELETE /tasks/delete/:id`  
*Requires admin privileges*  
Deletes a task from the system.

### Request
**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Path Parameters:**
- `id` (string): ID of the task to delete

### Responses
- **200 OK** (Success):
  ```json
  {"message": "Task delete successfully"}
  ```
- **401 Unauthorized** (Missing/invalid token or non-admin):
  ```json
  {"error": "Unauthorized"}
  ```
- **404 Not Found** (Task not found):
  ```json
  {"message": "Task Not Found"}
  ```
- **500 Internal Server Error** (Database error):
  ```json
  {"message": "Cannot delete a record"}
  ```

---

## Authorization Notes
1. All endpoints require JWT authentication in the `Authorization` header
2. Task creation, update, and deletion require admin privileges
3. Users receive JWT token after successful login
4. First registered user becomes admin automatically
```