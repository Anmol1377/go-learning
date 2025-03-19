package storage

import "github.com/anmol1377/student-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int , createdAt string , updatedAt string) (int64, error)
	GetStudent(id int64) (types.Student, error)
}
