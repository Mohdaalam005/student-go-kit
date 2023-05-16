package service

import (
	"context"

	"github.com/mohdaalam/005/student/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateStudent(ctx context.Context, student repository.StudentRequest) (repository.StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]repository.Student, error)
}

type serivce struct {
	repository repository.Repository
	log        logrus.Logger
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
