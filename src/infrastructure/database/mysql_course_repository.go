package infrastructure_database

import (
	"database/sql"
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
)

type MysqlCourseRepository struct {
	Db *sql.DB
}

func (m *MysqlCourseRepository) Insert(course domain.Course) error {
	stmt, err := m.Db.Prepare(`Insert into courses(id, name, description, status) values (?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}