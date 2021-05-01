package database

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Tersedia = 1
	Kosong   = 0
)
const (
	admin = 1
	user  = 2
)

type Category struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Books []Book
}
type Book struct {
	gorm.Model

	Title       string
	Description string
	Author      string
	Year        int
	Stock       int
	Status      int
	Category    Category
	CategoryID  uint
	Loans       []Loan
}

type User struct {
	gorm.Model

	Name              string
	Address           string
	Photo             string
	Email             string
	Email_verified_at time.Time
	Password          string
	Role              int
	Loans             []Loan
	Registrations     []Registration
}

type Loan struct {
	gorm.Model

	UserID       uint
	BookID       uint
	BorrowedDate time.Time
	DueDate      time.Time
	ReturnDate   time.Time
	FinePayment  FinePayment
}
type FinePayment struct {
	gorm.Model
	Receipt string
	Amount  int64
	LoanID  uint
}

type Registration struct {
	gorm.Model
	RegistrationToken string
	TimeCreated       time.Time
	UserID            uint
}

type Transaction struct {
	IDBook       int
	Title        string
	Amount       int64
	Receipt      string
	BorrowedDate time.Time
	DueDate      time.Time
}

var db *gorm.DB

func StartDB(dataSourceName string) error {
	fmt.Println("datasourceName: ", dataSourceName)
	var err error
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	return err
}

func OpenDB(dataSourceName string) (db *gorm.DB, err error) {
	fmt.Println("datasourceName: ", dataSourceName)
	return gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
}

func InitDBTable() {
	db.AutoMigrate(&Category{}, &Book{}, &User{}, &Loan{}, &FinePayment{}, &Registration{})

	// Create Fresh Category Table
	if (db.Migrator().HasTable(&Category{})) {
		fmt.Println("Category table exist")
		db.Migrator().DropTable(&Category{})
	}
	db.Migrator().CreateTable(&Category{})

	// Create Fresh Book Table
	if (db.Migrator().HasTable(&Book{})) {
		fmt.Println("Book table exist")
		db.Migrator().DropTable(&Book{})
	}
	db.Migrator().CreateTable(&Book{})

	// Create Fresh User Table
	if (db.Migrator().HasTable(&User{})) {
		fmt.Println("User table exist")
		db.Migrator().DropTable(&User{})
	}
	db.Migrator().CreateTable(&User{})

	// Create Fresh Loan Table
	if (db.Migrator().HasTable(&Loan{})) {
		fmt.Println("Loan table exist")
		db.Migrator().DropTable(&Loan{})
	}
	db.Migrator().CreateTable(&Loan{})

	// Create Fresh FinePayment Table
	if (db.Migrator().HasTable(&FinePayment{})) {
		fmt.Println("FinePayment table exist")
		db.Migrator().DropTable(&FinePayment{})
	}
	db.Migrator().CreateTable(&FinePayment{})

	// Create Fresh Registration Table
	if (db.Migrator().HasTable(&Registration{})) {
		fmt.Println("Registration table exist")
		db.Migrator().DropTable(&Registration{})
	}
	db.Migrator().CreateTable(&Registration{})

}

func CategoryDataSeedInit(totalData int) {
	for i := 0; i < totalData; i++ {
		category := Category{Name: "Name" + strconv.Itoa(i)}
		db.Create(&category)
	}
}

func BookDataSeedInit(totalData int) {
	// collect Category table rows Primary Key (ID)
	//SELECT * FROM category;
	var categories []Category
	db.Find(&categories)

	var categoryIds []uint
	for _, category := range categories {
		// fmt.Println("ID:", category.ID)
		categoryIds = append(categoryIds, category.ID)
	}

	// Create Dummy Data for Book Table

	var counter int
	for i := 0; i < totalData; i++ {

		// Distribute Category Table ID evenly as foreign key
		counter = i % int(len(categoryIds))

		book := Book{
			Title:       "Title" + strconv.Itoa(i),
			Description: "Description" + strconv.Itoa(i),
			Author:      "Author" + strconv.Itoa(i),
			Year:        2000 + i,
			Stock:       i,
			Status:      i % 2,
			CategoryID:  categoryIds[counter],
		}
		db.Create(&book)
	}
}

func UserDataSeedInit(totalData int) {
	// Create Dummy Data for User Table
	for i := 0; i < totalData; i++ {

		user := User{
			Name:              "User" + strconv.Itoa(i),
			Address:           "Address" + strconv.Itoa(i),
			Photo:             "PhotoPath" + strconv.Itoa(i),
			Email:             "email" + strconv.Itoa(i) + "@gmail.com",
			Email_verified_at: time.Now(),
			Password:          "Password" + strconv.Itoa(i),
			Role:              (i % 2) + 1,
		}
		db.Create(&user)
	}
}
func LoanAndFinePaymentDataSeedInit(totalData int) {
	//Loan & Fine Payment has 1 on 1 relational mapping so the totalData should be the same
	LoanDataSeedInit(totalData)
	FinePaymentDataSeedInit(totalData)
}

func LoanDataSeedInit(totalData int) {
	// collect Users table rows Primary Key (ID)
	//SELECT * FROM users;
	var users []User
	db.Find(&users)

	var userIds []uint
	for _, user := range users {
		// fmt.Println("ID:", category.ID)
		userIds = append(userIds, user.ID)
	}
	// collect Books table rows Primary Key (ID)
	//SELECT * FROM books;
	var books []Book
	db.Find(&books)

	var bookIds []uint
	for _, book := range books {
		// fmt.Println("ID:", category.ID)
		bookIds = append(bookIds, book.ID)
	}

	// Create Dummy Data for Loan Table
	var counter1 int
	var counter2 int
	for i := 0; i < totalData; i++ {

		// Distribute Category Table ID evenly as foreign key
		counter1 = i % int(len(userIds))
		counter2 = i % int(len(bookIds))

		loan := Loan{
			UserID:       userIds[counter1],
			BookID:       bookIds[counter2],
			BorrowedDate: time.Now(),
			DueDate:      time.Now().AddDate(0, 0, 3),
			ReturnDate:   time.Now().AddDate(0, 0, 2),
		}
		db.Create(&loan)
	}
}

func FinePaymentDataSeedInit(totalData int) {
	// collect Loans table rows Primary Key (ID)
	//SELECT * FROM Loan;
	var loans []Loan
	db.Find(&loans)

	var loanIds []uint
	for _, loan := range loans {
		// fmt.Println("ID:", category.ID)
		loanIds = append(loanIds, loan.ID)
	}

	// Create Dummy Data for Loan Table
	var counter int
	for i := 0; i < totalData; i++ {

		// Distribute Category Table ID evenly as foreign key
		counter = i % int(len(loanIds))
		finePayment := FinePayment{
			Receipt: "Receipt" + strconv.Itoa(i),
			Amount:  int64(10000 + i),
			LoanID:  loanIds[counter],
		}
		db.Create(&finePayment)
	}
}

//CRUD Category

// r.GET("/categories", router.ListAllCategories)
func ListAllCategories() ([]Category, error) {
	var categories []Category
	result := db.Find(&categories)
	if result.Error != nil {
		return categories, result.Error
	} else if result.RowsAffected < 1 {
		return categories, fmt.Errorf("table is empty")
	}
	return categories, nil
}

// r.GET("/categories/:id", router.ListCategory)
func ListCategory(id int) (Category, error) {
	var category Category
	result := db.Where("id=?", id).Find(&category)
	if result.Error != nil {
		return category, result.Error
	} else if result.RowsAffected < 1 {
		return category, fmt.Errorf("can't find the row with ID: %v", id)
	}
	return category, nil
}

// r.POST("/categories", router.AddCategory)
func AddCategory(name string) error {
	var category Category
	category.Name = name
	result := db.Create(&category)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("adding new row is failed")
	}
	return nil

}

// r.PUT("/categories/:id", router.UpdateCategory)
func UpdateCategory(id int, name string) error {
	var category Category
	result := db.Where("id=?", id).Find(&category)
	fmt.Println("UPDATECATEGORY ERROR: ", result.Error)
	fmt.Println("UPDATECATEGORY name: ", name)
	fmt.Println("UPDATECATEGORY id: ", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row with id: %v", id)
	}

	category.Name = name
	fmt.Println("UPDATECATEGORY Category Object: ", category, category.Name)
	result = db.Save(&category)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating row with id: %v is failed", id)
	}
	return nil
}

func UpdateCategoryByObject(category Category) error {
	result := db.Save(category)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't update the row, updating row with id: %v is failed", category.ID)
	}
	return nil

}

// r.DELETE("/categories", router.DeleteAllCategories)
func DeleteAllCategories() error {

	// result := db.Model(&Category{}).Delete(&Category{})
	var categories []Category
	result := db.Find(&categories)
	for _, category := range categories {
		result = db.Delete(&category)
	}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table categories is empty")
	}
	return nil
}

// r.DELETE("/categories/:id", router.DeleteCategory)
func DeleteCategory(id uint) error {
	var category Category
	category.ID = id
	result := db.Delete(&category)
	fmt.Println("Result.Error: ", result.Error)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, deleting category id:%v is failed", id)
	}
	return nil
}

// CRUD Book

// r.GET("/books", router.ListAllBooks)
func ListAllBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return books, result.Error
	} else if result.RowsAffected < 1 {
		return books, fmt.Errorf("table is empty")
	}
	return books, nil
}

// r.GET("/books/:id", router.ListBook)
func ListBook(id int) (Book, error) {
	var book Book
	result := db.Where("id=?", id).Find(&book)
	if result.Error != nil {
		return book, result.Error
	} else if result.RowsAffected < 1 {
		return book, fmt.Errorf("can't find the row, get row with id: %v is failed", id)
	}
	return book, nil

}

// r.POST("/books", router.AddBook)
func AddBook(id_category int, title string, description string, author string, year int, stock int) error {
	var book Book
	book.CategoryID = uint(id_category)
	book.Title = title
	book.Description = description
	book.Author = author
	book.Year = year
	book.Stock = stock
	if book.Stock > 0 {
		book.Status = 1
	} else {
		book.Status = 0
	}
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("adding new book failed")
	}
	//update category
	category, _ := ListCategory(id_category)
	category.Books = append(category.Books, book)
	_ = UpdateCategoryByObject(category)
	return nil
}

// r.PUT("/books/:id", router.UpdateBook)
func UpdateBook(id int, id_category int, title string, description string, author string, year int, stock int) error {
	var book Book
	result := db.Where("id=?", id).Find(&book)
	// fmt.Println("UPDATECATEGORY ERROR: ", result.Error)
	// fmt.Println("UPDATECATEGORY name: ", name)
	// fmt.Println("UPDATECATEGORY id: ", id)
	if result.Error != nil {
		return result.Error
	}
	book.CategoryID = uint(id_category)
	book.Title = title
	book.Description = description
	book.Author = author
	book.Year = year
	book.Stock = stock
	if book.Stock > 0 {
		book.Status = 1
	} else {
		book.Status = 0
	}
	// fmt.Println("UPDATECATEGORY Category Object: ", category, category.Name)
	result = db.Save(&book)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating new book with id:%v failed", id)
	}
	return nil
}

// r.DELETE("/books", router.DeleteAllBooks)
func DeleteAllBooks() error {

	// result := db.Model(&Book{}).Delete(&Book{})
	var books []Book
	result := db.Find(&books)
	for _, book := range books {
		result = db.Delete(&book)
	}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table books is empty")
	}
	return nil
}

// r.DELETE("/books/:id", router.DeleteBook)
func DeleteBook(id uint) error {
	var book Book
	book.ID = id
	result := db.Delete(&book)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, deleting row with id: %v is failed", id)
	}
	return nil
}
func GetAllBook() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	return books, result.Error
}
func UpdateBookByObject(book Book) error {
	result := db.Save(book)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating row with id: %v is failed", book.ID)
	}
	return nil
}

// CRUD User

// r.GET("/users", router.ListAllUsers)
func ListAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	} else if result.RowsAffected < 1 {
		return users, fmt.Errorf("table is empty")
	}
	return users, nil
}

// r.GET("/users/:id", router.ListUser)
func ListUser(id int) (User, error) {
	var user User
	result := db.Where("id=?", id).Find(&user)
	if result.Error != nil {
		return user, result.Error
	} else if result.RowsAffected < 1 {
		return user, fmt.Errorf("can't find the row, get row with id: %v is failed", id)
	}
	return user, nil

}

// r.POST("/users", router.AddUser)
func AddUser(name string, address string, photo string, email string, password string, role int) error {
	var user User
	user.Name = name
	user.Address = address
	user.Photo = photo
	user.Email = email
	user.Password = password
	user.Role = role
	user.Email_verified_at = time.Now()

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("adding new user failed")
	}
	return nil
}

// r.PUT("/users/:id", router.UpdateUser)
func UpdateUser(id int, name string, address string, photo string, email string, password string, role int) error {
	var user User

	result := db.Where("id=?", id).Find(&user)
	// fmt.Println("UPDATECATEGORY ERROR: ", result.Error)
	// fmt.Println("UPDATECATEGORY name: ", name)
	// fmt.Println("UPDATECATEGORY id: ", id)
	if result.Error != nil {
		return result.Error
	}
	user.Name = name
	user.Address = address
	user.Photo = photo
	user.Email = email
	user.Password = password
	user.Role = role

	// fmt.Println("UPDATECATEGORY Category Object: ", category, category.Name)
	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating new user with id:%v failed", id)
	}
	return nil
}

// r.DELETE("/users", router.DeleteAllUsers)
func DeleteAllUsers() error {

	// result := db.Model(&Book{}).Delete(&Book{})
	var users []User
	result := db.Find(&users)
	for _, user := range users {
		result = db.Delete(&user)
	}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table users is empty")
	}
	return nil
}

// r.DELETE("/users/:id", router.DeleteUser)
func DeleteUser(id uint) error {
	var user User
	user.ID = id
	result := db.Delete(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, deleting row with id: %v is failed", id)
	}
	return nil
}

func IsValidUserData(user User) bool {
	return true
}

// func AddUser(user User) error {
// 	result := db.Create(&user)
// 	return result.Error
// }

func GetAllUser() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

// CRUD Loan

// r.GET("/loans", router.ListAllLoans)
func ListAllLoans() ([]Loan, error) {
	var loans []Loan
	result := db.Find(&loans)
	if result.Error != nil {
		return loans, result.Error
	} else if result.RowsAffected < 1 {
		return loans, fmt.Errorf("table is empty")
	}
	return loans, nil
}

// r.GET("/loans/:id", router.ListLoan)
func ListLoan(id int) (Loan, error) {
	var loan Loan
	result := db.Where("id=?", id).Find(&loan)
	if result.Error != nil {
		return loan, result.Error
	} else if result.RowsAffected < 1 {
		return loan, fmt.Errorf("can't find the row, get row with id: %v is failed", id)
	}
	return loan, nil

}

// r.POST("/loans", router.AddLoan)
func AddLoan(idUser uint, idBook uint, borrowedDate time.Time, dueDate time.Time, returnDate time.Time) error {
	var loan Loan
	loan.UserID = idUser
	loan.BookID = idBook
	loan.BorrowedDate = borrowedDate
	loan.DueDate = dueDate
	loan.ReturnDate = returnDate

	result := db.Create(&loan)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("adding new user failed")
	}

	//update user
	user, _ := ListUser(int(idUser))
	user.Loans = append(user.Loans, loan)
	//update book
	book, _ := ListUser(int(idBook))
	book.Loans = append(book.Loans, loan)

	return nil
}

// r.PUT("/loans/:id", router.UpdateLoan)
func UpdateLoan(id int, idUser uint, idBook uint, borrowedDate time.Time, dueDate time.Time, returnDate time.Time) error {
	var loan Loan

	result := db.Where("id=?", id).Find(&loan)

	if result.Error != nil {
		return result.Error
	}
	loan.UserID = idUser
	loan.BookID = idBook
	loan.BorrowedDate = borrowedDate
	loan.DueDate = dueDate
	loan.ReturnDate = returnDate

	result = db.Save(&loan)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating new user with id:%v failed", id)
	}
	return nil
}

// r.DELETE("/loans", router.DeleteAllLoans)
func DeleteAllLoans() error {

	// result := db.Model(&Book{}).Delete(&Book{})
	var loans []Loan
	result := db.Find(&loans)
	for _, loan := range loans {
		result = db.Delete(&loan)
	}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table users is empty")
	}
	return nil
}

// r.DELETE("/loans/:id", router.DeleteLoan)
func DeleteLoan(id uint) error {
	var loan Loan
	loan.ID = id
	result := db.Delete(&loan)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, deleting row with id: %v is failed", id)
	}
	return nil
}

//CRUD FinePayment

// r.GET("/fine-payments", router.ListAllFinePayments)
func ListAllFinePayments() ([]FinePayment, error) {
	var finePayments []FinePayment
	result := db.Find(&finePayments)
	if result.Error != nil {
		return finePayments, result.Error
	} else if result.RowsAffected < 1 {
		return finePayments, fmt.Errorf("table is empty")
	}
	return finePayments, nil
}

// r.GET("/fine-payments/:id", router.ListFinePayment)
func ListFinePayment(id int) (FinePayment, error) {
	var finePayment FinePayment
	result := db.Where("id=?", id).Find(&finePayment)
	if result.Error != nil {
		return finePayment, result.Error
	} else if result.RowsAffected < 1 {
		return finePayment, fmt.Errorf("can't find the row, get row with id: %v is failed", id)
	}
	return finePayment, nil

}

// r.POST("/fine-payments", router.AddFinePayment)
func AddFinePayment(receipt string, amount int64, idLoan uint) error {
	var finePayment FinePayment
	finePayment.Receipt = receipt
	finePayment.Amount = amount
	finePayment.LoanID = idLoan

	result := db.Create(&finePayment)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("adding new user failed")
	}
	return nil
}

// r.PUT("/fine-payments/:id", router.UpdateFinePayment)
func UpdateFinePayment(id int, receipt string, amount int64, idLoan uint) error {
	var finePayment FinePayment

	result := db.Where("id=?", id).Find(&finePayment)

	if result.Error != nil {
		return result.Error
	}
	finePayment.Receipt = receipt
	finePayment.Amount = amount
	finePayment.LoanID = idLoan

	result = db.Save(&finePayment)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, updating new user with id:%v failed", id)
	}
	return nil
}

// r.DELETE("/fine-payments", router.DeleteAllFinePayments)
func DeleteAllFinePayments() error {

	// result := db.Model(&Book{}).Delete(&Book{})
	var finePayments []FinePayment
	result := db.Find(&finePayments)
	for _, finePayment := range finePayments {
		result = db.Delete(&finePayment)
	}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("table users is empty")
	}
	return nil
}

// r.DELETE("/fine-payments/:id", router.DeleteFinePayment)
func DeleteFinePayment(id uint) error {
	var finePayment FinePayment
	finePayment.ID = id
	result := db.Delete(&finePayment)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("can't find the row, deleting row with id: %v is failed", id)
	}
	return nil
}

func GetBookRent(user User) ([]Book, error) {
	var books []Book
	result := db.Where("email=?", user.Email).Find(&books)
	return books, result.Error

}

func GetTransactionList(user User) ([]Transaction, error) {
	var transactions []Transaction
	result := db.Raw("SELECT id_book, title, amount, receipt, borrowed_date, due_date FROM loans, fine_payments, book WHERE id_user = ?", user.ID).Scan(&transactions)

	return transactions, result.Error

}

func GetCategoryBasedBookList(idCategory int) ([]Book, error) {
	var books []Book
	result := db.Where("id_category=?", idCategory).Find(&books)
	return books, result.Error
}
func GetNewestList() ([]Book, error) {
	var books []Book
	result := db.Where("year>?", time.Now().AddDate(-2, 0, 0)).Find(&books)
	return books, result.Error

}
func GetUserByEmail(email string) (User, error) {
	var user User
	result := db.Where("email=?", email).Find(user)
	return user, result.Error
}

func UpdateUserByObject(user User) error {
	result := db.Save(&user)
	return result.Error
}

func AddRegistration(token string) error {
	registration := Registration{
		RegistrationToken: token,
		TimeCreated:       time.Now(),
		UserID:            1,
	}
	db.Create(&registration)
	return nil
}
func GetRegistration(email string, token string) (User, error) {
	var registration Registration
	var user User
	result := db.Where("registration_token=?", token).Find(&registration)
	fmt.Printf("GETREGISTRATION: EMAIL: %v token: %v IdUser:%v\n", email, token, registration.UserID)
	result = db.Where("id=?", registration.UserID).Find(&user)

	return user, result.Error
}
