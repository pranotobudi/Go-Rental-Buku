package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pranotobudi/Go-Rental-Buku/config"
	"github.com/pranotobudi/Go-Rental-Buku/database"
	"github.com/pranotobudi/Go-Rental-Buku/router"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "admin"
// 	dbname   = "rental_buku"
// )

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// fmt.Println("bismillah")
	conf := config.New()

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", conf.Postgresql.Host, conf.Postgresql.User, conf.Postgresql.Password, conf.Postgresql.DBName, conf.Postgresql.Port)

	err := database.StartDB(dataSourceName)
	// db, err := database.OpenDB(dataSourceName)
	if err != nil {
		fmt.Println("database connection FAILED: ", err)
	} else {
		fmt.Println("database connection SUCCESS.. ")
	}
	// seeders.DBSeed(db)
	database.InitDBTable()
	database.CategoryDataSeedInit(10)
	database.BookDataSeedInit(20)
	database.UserDataSeedInit(5)
	database.LoanAndFinePaymentDataSeedInit(30)

	r := gin.Default()

	//CRUD Category

	r.GET("/categories", router.TokenAuthMiddleware(), router.ListAllCategories)
	r.GET("/categories/:id", router.ListCategory)
	r.POST("/categories", router.TokenAuthMiddleware(), router.AddCategory)
	r.PUT("/categories/:id", router.TokenAuthMiddleware(), router.UpdateCategory)
	r.DELETE("/categories", router.TokenAuthMiddleware(), router.DeleteAllCategories)
	r.DELETE("/categories/:id", router.TokenAuthMiddleware(), router.DeleteCategory)

	// CRUD Book

	r.GET("/books", router.TokenAuthMiddleware(), router.ListAllBooks)
	r.GET("/books/:id", router.TokenAuthMiddleware(), router.ListBook)
	r.POST("/books", router.TokenAuthMiddleware(), router.AddBook)
	r.PUT("/books/:id", router.TokenAuthMiddleware(), router.UpdateBook)
	r.DELETE("/books", router.TokenAuthMiddleware(), router.DeleteAllBooks)
	r.DELETE("/books/:id", router.TokenAuthMiddleware(), router.DeleteBook)

	//CRUD User

	r.GET("/users", router.TokenAuthMiddleware(), router.ListAllUsers)
	r.GET("/users/:id", router.TokenAuthMiddleware(), router.ListUser)
	r.POST("/users", router.TokenAuthMiddleware(), router.AddUser)
	r.PUT("/users/:id", router.TokenAuthMiddleware(), router.UpdateUser)
	r.DELETE("/users", router.TokenAuthMiddleware(), router.DeleteAllUsers)
	r.DELETE("/users/:id", router.TokenAuthMiddleware(), router.DeleteUser)

	//CRUD Loan

	r.GET("/loans", router.TokenAuthMiddleware(), router.ListAllLoans)
	r.GET("/loans/:id", router.TokenAuthMiddleware(), router.ListLoan)
	r.POST("/loans", router.TokenAuthMiddleware(), router.AddLoan)
	r.PUT("/loans/:id", router.TokenAuthMiddleware(), router.UpdateLoan)
	r.DELETE("/loans", router.TokenAuthMiddleware(), router.DeleteAllLoans)
	r.DELETE("/loans/:id", router.TokenAuthMiddleware(), router.DeleteLoan)

	//CRUD FinePayment

	r.GET("/fine-payments", router.TokenAuthMiddleware(), router.ListAllFinePayments)
	r.GET("/fine-payments/:id", router.TokenAuthMiddleware(), router.ListFinePayment)
	r.POST("/fine-payments", router.TokenAuthMiddleware(), router.AddFinePayment)
	r.PUT("/fine-payments/:id", router.TokenAuthMiddleware(), router.UpdateFinePayment)
	r.DELETE("/fine-payments", router.TokenAuthMiddleware(), router.DeleteAllFinePayments)
	r.DELETE("/fine-payments/:id", router.TokenAuthMiddleware(), router.DeleteFinePayment)

	r.GET("/users/list", router.UserList)
	r.POST("/register", router.UserRegister)
	r.GET("/register/confirmation", router.EmailConfirmation)
	r.POST("/login", router.UserLogin)
	// r.POST("/logout", router.UserLogout)

	r.POST("/resetpassword", router.ResetPassword)
	r.GET("/resetpassword/submitpage", router.ResetPasswordSubmitGet)
	r.POST("/resetpassword/submitpage", router.ResetPasswordSubmitPost)

	r.GET("/mybooks", router.TokenAuthMiddleware(), router.BookRentList)
	r.GET("/mytransactions", router.TokenAuthMiddleware(), router.TransactionList)
	r.GET("/catalog", router.TokenAuthMiddleware(), router.GetCatalog)
	r.GET("/category", router.TokenAuthMiddleware(), router.GetCategoryBasedBook)
	r.GET("/newest", router.TokenAuthMiddleware(), router.GetNewest)

	//Pinjam Buku
	r.POST("/books/:id/borrow", router.BorrowBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
