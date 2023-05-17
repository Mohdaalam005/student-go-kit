package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/mohdaalam/005/student/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	CreateStudent(ctx context.Context, student StudentRequest) (StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]Student, error)
	GetStudentById(ctx context.Context, id int) (Student, error)
	UpdateStudent(ctx context.Context, student Student, studentId int) (Student, error)
	DeleteStudentById(ctx context.Context, id int) error
}

type repository struct {
	Db     sql.DB
	Logger logrus.Logger
}

// UpdateStudent implements Repository
func (r *repository) UpdateStudent(ctx context.Context, student Student, studentId int) (Student, error) {
	studentById, err := models.FindStudent(ctx, &r.Db, studentId)
	if err != nil {
		r.Logger.Info("id not present")
		errMsg := "Student not found"
		errJSON, _ := json.Marshal(map[string]string{"error": errMsg})
		return Student{}, errors.New(string(errJSON))
	}
	studentById.Dob = student.Dob
	studentById.Name = student.Name
	studentById.Gender = student.Gender
	_, err = studentById.Update(ctx, &r.Db, boil.Infer())
	if err != nil {
		return Student{}, err
	}

	r.Logger.Info("update student successfully id = ",student.ID)
	return Student{
		ID:     studentId,
		Name:   studentById.Name,
		Gender: student.Gender,
		Dob:    student.Dob,
	}, nil

}

// DeleteStudentById implements Repository
func (r *repository) DeleteStudentById(ctx context.Context, id int) error {
	r.Logger.Info("deleting student by id")
	student, err := models.FindStudent(ctx, &r.Db, id)
	student.Delete(ctx, &r.Db)
	if err != nil {
		r.Logger.Info("id not present")
	}
	student.Delete(ctx, &r.Db)
	return nil
}

// GetStudentById implements Repository
func (r *repository) GetStudentById(ctx context.Context, id int) (Student, error) {
	student, err := models.FindStudent(ctx, &r.Db, id)
	if err != nil {
		r.Logger.Info("id not present")
		errMsg := "Student not found"
		errJSON, _ := json.Marshal(map[string]string{"error": errMsg})
		return Student{}, errors.New(string(errJSON))
	}
	return Student{
		ID:     student.ID,
		Name:   student.Name,
		Gender: student.Gender,
		Dob:    student.Dob,
	}, nil
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
