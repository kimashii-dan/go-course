package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Email   string  `db:"email"`
	Balance float64 `db:"balance"`
}

func main() {
	// load variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// build connection string for DB
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	//connect DB with sqlx
	db, err := sqlx.Connect("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// create myself
	daniyar := User{
		Name:    "Daniyar",
		Email:   "daniyarmunsyzbaev@gmail.com",
		Balance: 500,
	}
	err = InserUser(db, daniyar)
	if err != nil {
		log.Fatal(err)
	}

	// create some clown
	clown := User{
		Name:    "Olzhas",
		Email:   "olzhas@gmail.com",
		Balance: 300,
	}
	err = InserUser(db, clown)
	if err != nil {
		log.Fatal(err)
	}

	// get users one by one
	daniyarDB, err := GetUserByID(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got myself by ID: %+v\n", daniyarDB)

	clownDB, err := GetUserByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got clown by ID: %+v\n", clownDB)

	// get all users before transfer
	usersBeforeTransfer, err := GetAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users before transfer")
	for _, user := range usersBeforeTransfer {
		fmt.Println(user)
	}

	// perform transfer
	err = TransferBalance(db, clownDB.ID, daniyarDB.ID, 100)
	if err != nil {
		log.Fatal(err)
	}

	// get all users after transfer
	usersAfterTransfer, err := GetAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users after transfer")
	for _, user := range usersAfterTransfer {
		fmt.Println(user)
	}

}

// insert user into DB
func InserUser(db *sqlx.DB, user User) error {
	_, err := db.NamedExec(`INSERT INTO users (name, email, balance) 
		VALUES (:name, :email, :balance)`, user)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// get all users
func GetAllUsers(db *sqlx.DB) ([]User, error) {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

// get user by id
func GetUserByID(db *sqlx.DB, id int) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

// transfer amount of (?) from one user to another user
func TransferBalance(db *sqlx.DB, fromID int, toID int, amount float64) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// rollback happens only if the transaction is still not commited
	defer tx.Rollback()

	var fromBalance float64
	err = tx.Get(&fromBalance, "SELECT balance FROM users WHERE id = $1", fromID)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return fmt.Errorf("insufficient funds")
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromID)
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
