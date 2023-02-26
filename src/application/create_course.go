package application

import (
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	uuid "github.com/satori/go.uuid"
)

type CreateCourse struct {
	Repository domain.CourseRepository
}

func (c *CreateCourse) Execute(dto CourseDto) (domain.Course, error) {
	course := domain.NewCourse(
		uuid.NewV4().String(),
		dto.Name,
		dto.Description,
		"active",
	)

	err := c.Repository.Insert(course)
	if err != nil {
		return domain.Course{}, err
	}

	return course, nil;
}
