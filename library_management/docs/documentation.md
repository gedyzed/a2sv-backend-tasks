# Library Management System

This is a simple **console-based Library Management System** written in Go. It allows users to:

- Add, remove, borrow, and return books
- List all available books
- View books borrowed by a member

The code is structured using Go packages: **controllers**, **services**, and **models**.

---

## Packages

### models

Contains basic data structures:
- **Book**: holds `ID`, `Title`, `Author`
- **Member**: holds `ID`, `Name`, `BorrowedBooks []Book`

### services

Contains the `Library` struct and core logic:
- `AddBook(book Book)`
- `RemoveBook(bookID int)`
- `BorrowBook(bookID, memberID int)`
- `ReturnBook(bookID, memberID int)`

Data members:
- `Books map[int]Book` — available books
- `Members map[int]*Member` — registered members

### controllers

Handles console input and UI display:
- `ProcessInput()` function drives the command-line menu
- Formats and prints books/member info to the terminal

---

## How to Run

```bash
go run main.go
