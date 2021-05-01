package router

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

var MAX_BOOK_LIMIT_PER_DAY = 3

func BorrowBook(c *gin.Context) {
	idStr := c.Param("id")
	id_book, _ := strconv.Atoi(idStr)
	id_user, _ := strconv.Atoi(c.PostForm("id_user"))

	if IsBookAvailable(id_book) {
		if TotalBorrowedBookToday(id_user) < MAX_BOOK_LIMIT_PER_DAY {
			err := UpdateBorrowedBook(id_user, id_book)
			//return JSON
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": "borrow book failed"})
			}
			c.JSON(http.StatusOK, gin.H{"message": "borrow book success"})

		} else {
			c.JSON(http.StatusOK, gin.H{"message": "maximum book limit per day is 3"})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Book is not available"})
		// c.Abort()
		return
	}
}

func IsBookAvailable(id_book int) bool {
	book, _ := database.ListBook(id_book)
	return book.Stock > 0
}

func TotalBorrowedBookToday(id_user int) int {
	fmt.Println("TOTALBORROWEDBOOKTODAY..........")
	fmt.Println("id_user:..........", id_user)
	user, _ := database.ListUser(id_user)
	fmt.Println("user object:..........", user)
	loans := user.Loans
	fmt.Println("loans length:..........", len(loans))
	yearNow, monthNow, dayNow := time.Now().Date()
	counter := 0
	fmt.Printf("yearNow:%v, monthNow:%v, dayNow:%v \n", yearNow, monthNow, dayNow)
	for _, loan := range loans {
		loanYear, loanMonth, loanDay := loan.BorrowedDate.Date()
		fmt.Printf("loanYear:%v, loanMonth:%v, loanDay:%v \n", loanYear, loanMonth, loanDay)
		if (yearNow == loanYear) && (monthNow == loanMonth) && (dayNow == loanDay) {
			counter++
		}
	}
	fmt.Println("COUNTER: ", counter)
	return counter
}

// UpdateBorrowedBook should be atomic database update transaction, success all or failed all
// need code improvement
func UpdateBorrowedBook(id_user int, id_book int) error {
	fmt.Println("UPDATEBORROWEDBOOK..........")
	// update user: add book which is borrowed by user
	user, _ := database.ListUser(id_user)
	loan := database.Loan{
		UserID: uint(id_user),
		BookID: uint(id_book),
	}
	user.Loans = append(user.Loans, loan)
	err := database.UpdateUserByObject(user)
	if err != nil {
		return err
	}
	//update book: substract book's stock
	book, _ := database.ListBook(id_book)
	book.Stock -= 1
	err = database.UpdateBookByObject(book)
	if err != nil {
		return err
	}

	//update loan: add userID and bookID
	err = database.AddLoan(uint(id_user), uint(id_book), time.Now(), time.Now().AddDate(0, 0, 3), time.Now())
	if err != nil {
		return err
	}

	return nil

}
