## 📚 **Task Manager API Documentation**


### 🔹 `GET /tasks`

* **Description:** Get a list of all tasks.
* **Request Params:** None
* **Response:**

```json
[
  {
    "id": 1,
    "title": "Implement login",
    "description": "Create login route and logic",
    "due_date": "2025-07-20",
    "status": "In Progress"
  },
  ...
]
```

* **Status Code:** `200 OK`

---

### 🔹 `GET /tasks/:id`

* **Description:** Get a single task by its ID.
* **Path Parameter:**

  * `id` (integer) — ID of the task to retrieve
* **Example Request:**
  `GET /tasks/2`
* **Response:**

```json
{
  "id": 2,
  "title": "Set up database",
  "description": "Configure PostgreSQL for dev environment",
  "due_date": "2025-07-18",
  "status": "Completed"
}
```

* **Status Codes:**

  * `200 OK`
  * `404 Not Found` if task is missing

---

### 🔹 `POST /tasks`

* **Description:** Create a new task.
* **Request Body (JSON):**

```json
{
  "title": "New Task",
  "description": "Task details go here",
  "due_date": "2025-07-30",
  "status": "Pending"
}
```

* **Response:**

```json
{
  "message": "Task added successfully",
  "task": {
    "id": 4,
    "title": "New Task",
    "description": "Task details go here",
    "due_date": "2025-07-30",
    "status": "Pending"
  }
}
```

* **Status Code:** `201 Created`

---

### 🔹 `DELETE /tasks/:id`

* **Description:** Delete a task by ID.
* **Path Parameter:**

  * `id` (integer) — ID of the task to delete
* **Example Request:**
  `DELETE /tasks/1`
* **Response:**

```json
{
  "message": "Task deleted successfully"
}
```

* **Status Codes:**

  * `200 OK`
  * `404 Not Found` if task is not found

---

### 🔹 `PUT /tasks/:id`

* **Description:** Update a task by ID.
* **Path Parameter:**

  * `id` (integer) — ID of the task to update
* **Request Body (JSON):**

```json
{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2025-08-01",
  "status": "Completed"
}
```

* **Response:**

```json
{
  "message": "Task updated successfully",
  "task": {
    "id": 1,
    "title": "Updated Task Title",
    "description": "Updated description",
    "due_date": "2025-08-01",
    "status": "Completed"
  }
}
```

* **Status Codes:**

  * `200 OK`
  * `404 Not Found` if task not found

