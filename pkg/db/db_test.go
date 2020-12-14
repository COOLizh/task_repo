package db

// "TODO "TestCases""
//
//import (
//	"context"
//	"testing"
//
//	"gitlab.com/greenteam1/task_repo/pkg/models"
//
//	"github.com/jackc/pgconn"
//
//	"github.com/jackc/pgx/v4"
//)
//
//var user = models.User{ID: 1, Username: "ann", Password: "ann"}
//
//func TestInsertGet(t *testing.T) {
//	t.Parallel()
//
//	conn := connect(t)
//	defer closeConn(t, conn)
//
//	mustExec(t, conn, `create temporary table users (
//	id int primary key,
//	username varchar not null,
//	login varchar
//	);`)
//
//	addUser, err := conn.Exec(context.Background(), `insert into users (id, username, login) values ($1, $2, $3)`, user.ID, user.Username, user.Password)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if string(addUser) != "INSERT 0 1" {
//		t.Fatalf("want %s, got %s", "INSERT 0 1", addUser)
//	}
//
//	row := conn.QueryRow(context.Background(), "select * from users")
//	user := models.User{}
//	err = row.Scan(&user.ID, &user.Username, &user.Password)
//	if err != nil {
//		t.Fatalf("conn.Query failed: %v", err)
//	}
//
//	if user.ID != 1 {
//		t.Error("Select called onDataRow wrong number of times")
//	}
//	if user.Username != "ann" {
//		t.Error("Select called onDataRow wrong number of times")
//	}
//	if user.Password != "ann" {
//		t.Error("Select called onDataRow wrong number of times")
//	}
//}
//func connect(t testing.TB) *pgx.Conn {
//	conn, err := pgx.Connect(context.Background(), "postgres://postgres:secret@localhost:5432/task_repo?sslmode=disable")
//	if err != nil {
//		t.Fatalf("Unable to get connection: %v", err)
//	}
//	return conn
//}
//func closeConn(t testing.TB, conn *pgx.Conn) {
//	err := conn.Close(context.Background())
//	if err != nil {
//		t.Fatalf("conn.Close unexpectedly failed: %v", err)
//	}
//}
//func mustExec(t testing.TB, conn *pgx.Conn, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag) {
//	var err error
//	if commandTag, err = conn.Exec(context.Background(), sql, arguments...); err != nil {
//		t.Fatalf("Exec unexpectedly failed with %v: %v", sql, err)
//	}
//	return
//}
