package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/anmol1377/student-api/internal/config"
	"github.com/anmol1377/student-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students 
(id INTEGER PRIMARY KEY AUTOINCREMENT, 
name TEXT, 
age INTEGER, 
email TEXT,
 created_at TEXT, 
 updated_at TEXT)`)

	if err != nil {
		return nil, err
	}

	return &sqlite{
		DB: db,
	}, nil

}

func (s sqlite) CreateStudent(name string, email string, age int, createdAt string, updatedAt string) (int64, error) {

	statement, err := s.DB.Prepare("INSERT INTO students (name, email, age,createdAt,updatedAt ) VALUES (?, ?, ? ,? ,?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(name, email, age, createdAt, updatedAt)

	if err != nil {
		return 0, err
	}

	LastId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return LastId, nil

}

func (s sqlite) GetStudent(id int64) (types.Student, error) {

	stmt, err := s.DB.Prepare("SELECT * FROM students WHERE ID = ?")

	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(student.Id, student.Name, student.Age, student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no rows found %w", err)
		}
		return types.Student{}, fmt.Errorf("query error %w", err)
	}
	return student, nil

}
