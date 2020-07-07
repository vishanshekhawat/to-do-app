package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vishn007/go-to-do-app/server/models"
)

// collection object/instance
var dbconnection *sql.DB
var globalVar int = 1

// init() create connection with mongo db
func init() {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/todo"
	globalVar = 2
	// Open the database.
	newConnection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	//defer dbconnection.Close()

	// Ping the database to verify that the connection is valid.
	if err := newConnection.Ping(); err != nil {
		log.Fatalln(err)
	}
	dbconnection = newConnection
	/*
		dbDriver := "mysql"
		dbUser := "root"
		dbPass := "12345678"
		dbName := "todo"
		dbconnection, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
		if err != nil {
			panic(err.Error())
		}
	*/
	fmt.Println(dbconnection)
	fmt.Println("Connected to mysql!")

}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// TaskComplete update task route
func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []models.ToDoList {
	fmt.Println(dbconnection)
	/*
			dsn := "root:12345678@tcp(127.0.0.1:3306)/todo"
			dbconnection, err := sql.Open("mysql", dsn)
			if err != nil {
				log.Fatalln(err)
			}
			//defer dbconnection.Close()

			// Ping the database to verify that the connection is valid.
			if err := dbconnection.Ping(); err != nil {
				log.Fatalln(err)
			}
				// Ping the database to verify that the connection is valid.
		if err := dbconnection.Ping(); err != nil {
			log.Fatalln(err)
		}
	*/
	res := []models.ToDoList{}

	// Use Query to retrieve all users from the database;
	rows, err := dbconnection.Query("select * from todo")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	task := models.ToDoList{}
	//res := []models.ToDoList{}
	// Iterate over all of the rows returned from the database.
	for rows.Next() {
		// Create a new user object.
		var id, status int
		var name string
		err = rows.Scan(&id, &name, &status)
		if err != nil {
			panic(err.Error())
		}
		task.ID = id
		task.Task = name
		task.Status = status
		res = append(res, task)
	}

	return res

}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {

}

// task complete method, update task's status to true
func taskComplete(task string) {

}

// task undo method, update task's status to false
func undoTask(task string) {

}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {

}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	return 1
}
