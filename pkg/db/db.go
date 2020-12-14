// Package db implements mock database functions
package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/COOLizh/task_repo/pkg/models"
)

// dbC struct for connecting to the database
type dbC struct {
	pool *pgxpool.Pool
	conn *pgxpool.Conn
}

var db = dbC{}

// Connect function gets a database connection
func Connect(connectionString string) (*pgxpool.Conn, *pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, nil, err
	}
	log.Println("Connected")

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, nil, err
	}
	log.Println("Get connection")

	db.pool = pool
	db.conn = conn
	return conn, pool, nil
}

// AddUser function adds into db
func AddUser(user *models.User) (int, error) {
	err := db.conn.QueryRow(context.Background(), "insert into users(login, password) values($1, $2) returning id", user.Username, user.Password).Scan(&user.ID)
	return user.ID, err
}

// GetUserByUsername returns User if it was found in db, and error if not
func GetUserByUsername(username string) (*models.User, error) {
	row := db.conn.QueryRow(context.Background(), "select * from users where login = $1", username)
	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//  "TODO "AddTask""
//func AddTask(task *models.Task) (int, error) {
//	err := db.conn.QueryRow(context.Background(), "insert into tasks(name, description, time_limit, memory) values($1, $2, $3, $4) returning id", task.Name, task.Description, task.TimeLimit, task.Memory).Scan(&task.ID)
//	return task.ID, err
//}

// GetAllTasks returns all tasks list from database
func GetAllTasks() ([]models.Task, error) {
	rows, err := db.conn.Query(context.Background(), "select * from tasks")
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	for rows.Next() {
		t := models.Task{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.TimeLimit, &t.Memory)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, err
}

// GetTaskByID gets task by ID
func GetTaskByID(id int) (*models.Task, error) {
	row := db.conn.QueryRow(context.Background(), "select * from tasks where id = $1", id)
	task := models.Task{}
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.TimeLimit, &task.Memory)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// "TODO "AddTestCase""
//func AddTestCase(testCase *models.TestCase) (int, error) {
//	err := db.conn.QueryRow(context.Background(), "insert into test_cases(task_id, test_data, answer) values($1, $2, $3) returning id", testCase.TaskID, testCase.TestData, testCase.Answer).Scan(&testCase.ID)
//	return testCase.ID, err
//}

// GetTestCasesByTaskID gets test cases by task id
func GetTestCasesByTaskID(taskID int) ([]models.TestCase, error) {
	var err error
	rows, err := db.conn.Query(context.Background(), "select * from test_cases where task_id = $1", taskID)
	if err != nil {
		return nil, err
	}
	var testCases []models.TestCase
	for rows.Next() {
		testCase := models.TestCase{}
		err = rows.Scan(&testCase.ID, &testCase.TaskID, &testCase.TestData, &testCase.Answer)
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, testCase)
	}

	return testCases, err
}
