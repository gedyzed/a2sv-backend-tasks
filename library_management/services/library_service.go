package services

import (
	"errors"
	"library_management/models"
)

// alais
type Book = models.Book
type Member = models.Member

type LibraryManager interface {

	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

type Library struct {
	Books map[int]Book
	Members map[int]Member
}


func (l *Library) AddBook (book Book)  {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int)  {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {

	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("Book Not Found")
	}

	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("Member Not Found")
	}

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	l.RemoveBook(bookID)

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {

	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("Member Not Found")
	}


	found := false
	for i, book := range member.BorrowedBooks {
		if book.ID == bookID {
			l.AddBook(book)
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i + 1 :]...)
			l.Members[memberID] = member
			found = true
			break
		}
	}

	if !found {
		return errors.New("Book Not Found")
	}

	return nil
}

func (l *Library) ListAvailableBooks() []Book {

	availabeBooks := []Book{}
	for _, book := range l.Books {
		availabeBooks = append(availabeBooks, book)
	}

	return availabeBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []Book {

	borrowedBooks := []Book{}
	for _, member := range l.Members {
		borrowedBooks = append(borrowedBooks, member.BorrowedBooks...)

	}

	return borrowedBooks
}






