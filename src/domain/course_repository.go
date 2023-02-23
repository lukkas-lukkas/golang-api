package domain

type CourseRepository interface {
	Insert(course Course) error
}
