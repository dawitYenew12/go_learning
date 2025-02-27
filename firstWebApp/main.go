package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

var database *sql.DB

func init() {
	var err error
	database, err = sql.Open("postgres", "postgres://postgres:Dawit0646@localhost:5432/tasks?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
}

// Handler functions
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to mark a task as complete
	fmt.Fprintf(w, "Complete Task")
}

func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to permanently delete a task
	fmt.Fprintf(w, "Delete Task")
}

func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to show deleted tasks
	fmt.Fprintf(w, "Show Trash Tasks")
}

func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to move a task to the trash/recycle bin
	fmt.Fprintf(w, "Trash Task")
}

func EditTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to edit a task
	fmt.Fprintf(w, "Edit Task")
}

func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to show completed tasks
	fmt.Fprintf(w, "Show Completed Tasks")
}

func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to restore a task from the trash
	fmt.Fprintf(w, "Restore Task")
}

func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to add a new task
	fmt.Fprintf(w, "Add Task")
}

func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to update a task
	fmt.Fprintf(w, "Update Task")
}

func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {
	// Logic to search for tasks
	fmt.Fprintf(w, "Search Task")
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	// Logic to show the login page
	fmt.Fprintf(w, "Login Page")
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	// Logic to handle user registration
	fmt.Fprintf(w, "Register User")
}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	// Logic to handle admin actions
	fmt.Fprintf(w, "Admin Page")
}

func PostAddUser(w http.ResponseWriter, r *http.Request) {
	// Logic to add a new user
	fmt.Fprintf(w, "Add User")
}

func PostChange(w http.ResponseWriter, r *http.Request) {
	// Logic to change user password
	fmt.Fprintf(w, "Change Password")
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Logic to handle user logout
	fmt.Fprintf(w, "Logout")
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	getTaskSQL := "SELECT id, title, content, created_date FROM task"
	rows, err := database.Query(getTaskSQL)

	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		log.Println("Error fetching tasks:", err)
		return
	}

	defer rows.Close()

	var tasks []map[string]interface{}

	for rows.Next() {
		var TaskID int
		var TaskTitle, TaskContent, TaskCreated string

		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			http.Error(w, "Error scanning task data", http.StatusInternalServerError)
			log.Println("Error scanning task data:", err)	
			return
		}

		// Replace new lines in content with HTML <br> tags
		TaskContent = strings.Replace(TaskContent, "\n", "<br>", -1)

		// Create a map for each task
		task := map[string]interface{}{
			"id":         TaskID,
			"title":      TaskTitle,
			"content":    TaskContent,
			"created_at": TaskCreated,
		}

		// Add the task to the tasks slice
		tasks = append(tasks, task)
	}

	// Check for row iteration errors
	err = rows.Err()
	if err != nil {
		http.Error(w, "Error processing tasks", http.StatusInternalServerError)
		log.Println("Error processing tasks:", err)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create the JSON response
	jsonResponse, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		log.Println("Error generating JSON response:", err)
		return
	}

	// Log the JSON response
	log.Println(string(jsonResponse))

	// Send the JSON response
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	http.HandleFunc("/trash/", TrashTaskFunc)
	http.HandleFunc("/edit/", EditTaskFunc)
	http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	http.HandleFunc("/restore/", RestoreTaskFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/update/", UpdateTaskFunc)
	http.HandleFunc("/search/", SearchTaskFunc)
	http.HandleFunc("/login", GetLogin)
	http.HandleFunc("/register", PostRegister)
	http.HandleFunc("/admin", HandleAdmin)
	http.HandleFunc("/add_user", PostAddUser)
	http.HandleFunc("/change", PostChange)
	http.HandleFunc("/logout", HandleLogout)
	http.HandleFunc("/", ShowAllTasksFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
