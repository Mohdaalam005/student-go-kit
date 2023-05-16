package repository

import (
	"context"
	"database/sql"

	"github.com/mohdaalam/005/student/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Repository interface {
	CreateStudent(ctx context.Context, student StudentRequest) (StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]Student, error)
	GetStudentById(ctx context.Context, id int) (Student, error)
}

type repository struct {
	Db     sql.DB
	Logger logrus.Logger
}

// GetStudentById implements Repository
func (r *repository) GetStudentById(ctx context.Context, id int) (Student, error) {
	 student, err :=models.Students(qm.Where("id=?",id)).One(ctx, &r.Db )
	 if err !=nil{
		r.Logger.Info("id not present")
	 }
	 return Student{
		ID: student.ID,
		Name: student.Name,
		Gender: student.Gender,
		Dob: student.Dob,
	 } , nil
}

// GetAllStudent implements Repository
func (r *repository) GetAllStudent(ctx context.Context) ([]Student, error) {
	student, err := models.Students().All(ctx, &r.Db)
	if err != nil {
		r.Logger.Error("GetAllStudent", err)
	}
	var result []Student
	for _, student := range student {
		result = append(result, Student{
			ID:     student.ID,
			Name:   student.Name,
			Gender: student.Gender,
			Dob:    student.Dob,
		})
	}
	return result, nil

}

// CreateStudent implements Repository
func (r *repository) CreateStudent(ctx context.Context, student StudentRequest) (StudentResponse, error) {
	create := models.Student{
		Name:   student.Name,
		Gender: student.Gender,
		Dob:    student.Dob,
	}
	err := create.Insert(ctx, &r.Db, boil.Infer())

	if err != nil {
		panic(err)
	}
	return StudentResponse{
		ID: create.ID,
	}, nil

}

func NewRespostiry(db sql.DB, log logrus.Logger) Repository {
	return &repository{
		Db:     db,
		Logger: log,
	}
}
