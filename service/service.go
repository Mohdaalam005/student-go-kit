package service

import (
	"context"

	"github.com/mohdaalam/005/student/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateStudent(ctx context.Context, student repository.StudentRequest) (repository.StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]repository.Student, error)
	GetStudentById(ctx context.Context, id int) (repository.Student, error)
	UpdateStudent(ctx context.Context, student repository.Student, studentId int) (repository.Student, error)
	DeleteStudentById(ctx context.Context, id int) error
}

type serivce struct {
	repository repository.Repository
	log        logrus.Logger
}

// UpdateStudent implements Service
func ( s *serivce) UpdateStudent(ctx context.Context, student repository.Student, studentId int) (repository.Student, error) {
	updteStudent ,err := s.repository.UpdateStudent(ctx,student, studentId);
	if err != nil {
		return  repository.Student{}, err
	}
	return updteStudent ,nil
}


// DeleteStudentById implements Service
func (s *serivce) DeleteStudentById(ctx context.Context, id int) error {
	s.repository.DeleteStudentById(ctx, id)
	return nil
}

// GetStudentById implements Service
func (s *serivce) GetStudentById(ctx context.Context, id int) (repository.Student, error) {
	student, err := s.repository.GetStudentById(ctx, id)
	if err != nil {
		s.log.Info("error getting students", err)

	}
	return repository.Student{
		ID:     student.ID,
		Name:   student.Name,
		Gender: student.Gender,
		Dob:    student.Dob,
	}, nil
}

// GetAllStudent implements Service
func (s *serivce) GetAllStudent(ctx context.Context) ([]repository.Student, error) {
	s.log.Info("getting all students")
	students, err := s.repository.GetAllStudent(ctx)
	if err != nil {
		s.log.Error("error getting students", err)
		return nil, err
	}
	return students, nil
}

// GetAllStudent implements Service

// CreateStudent implements Service
func (s serivce) CreateStudent(ctx context.Context, student repository.StudentRequest) (repository.StudentResponse, error) {
	data, err := s.repository.CreateStudent(ctx, student)
	s.log.Info("created new student !!")
	if err != nil {
		s.log.Log(logrus.ErrorLevel, err)
	}
	return repository.StudentResponse{
		ID: data.ID,
	}, nil

}

//	func NewService(repo repository.Repository, logg logrus.Logger) Service {
//		return &serivce{
//			repository: repo,
//			log:        logg,
//		}
//	}
func NewService(repo repository.Repository, log logrus.Logger) Service {
	return &serivce{
		repository: repo,
		log:        log,
	}
}
