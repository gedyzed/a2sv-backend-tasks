package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/services"
)

type Library = services.Library
type Book = services.Book
type Member = services.Member

var library Library

func ProcessInput() {

	library = Library{
		Books:   make(map[int]Book),
		Members: make(map[int]Member),
	}

	// Dummy members
	library.Members[1] = Member{ID: 1, Name: "Alice"}
	library.Members[2] = Member{ID: 2, Name: "Bob"}

	// Dummy books
	library.Books[102] = Book{ID: 102, Title: "Clean Code", Author: "Robert C. Martin"}
	library.Books[103] = Book{ID: 103, Title: "Design Patterns", Author: "Erich Gamma"}
	
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")

		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Enter Book Title: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Enter Book Author: ")
			author, _ := reader.ReadString('\n')

			book := Book{
				ID:     id,
				Title:  strings.TrimSpace(title),
				Author: strings.TrimSpace(author),
			}

			library.AddBook(book)

		case "2":
			fmt.Print("Enter Book ID to remove: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			library.RemoveBook(id)

		case "3":
			fmt.Print("Enter Book ID to borrow: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			fmt.Print("Enter Book ID to return: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "5":
			books := library.Books
			if len(books) == 0 {
				fmt.Println("\n No available books in the library.")
				break
			}

			fmt.Println("\n Available Books:")
			fmt.Println("----------------------------------------------------")
			fmt.Printf("%-6s | %-30s | %-20s\n", "ID", "Title", "Author")
			fmt.Println("----------------------------------------------------")

			for _, book := range books {
				fmt.Printf("%-6d | %-30s | %-20s\n", book.ID, book.Title, book.Author)
			}

			

		case "6":
			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

			member, ok := library.Members[memberID]
			if !ok {
				fmt.Println("Member not found.")
				break
			}

			if len(member.BorrowedBooks) == 0 {
				fmt.Printf("%s has not borrowed any books.\n", member.Name)
				break
			}

			fmt.Printf("\nBooks borrowed by %s:\n", member.Name)
			fmt.Println("----------------------------------------------------")
			fmt.Printf("%-6s | %-30s | %-20s\n", "ID", "Title", "Author")
			fmt.Println("----------------------------------------------------")

			for _, book := range member.BorrowedBooks {
				fmt.Printf("%-6d | %-30s | %-20s\n", book.ID, book.Title, book.Author)
			}


		case "7":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
