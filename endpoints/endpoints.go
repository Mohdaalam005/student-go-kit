package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mohdaalam/005/student/models"
	"github.com/mohdaalam/005/student/repository"
	"github.com/mohdaalam/005/student/service"
)

type Endpoints struct {
	CreateStudent endpoint.Endpoint
	GetAllStudent endpoint.Endpoint
	GetStudentById endpoint.Endpoint
}

func NewEnpoints(service service.Service) Endpoints{
	return Endpoints{
		CreateStudent: makeCreateEndpoint(service),
		GetAllStudent: makeGetAllEndpoints(service),
		GetStudentById : makeGetStudentByIdEndpoints(service), 
	}
}

func makeGetStudentByIdEndpoints(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.Student)
		res,_ := service.GetStudentById(ctx ,req.ID)
		return repository.Student{
			ID: res.ID,
			Name: res.Name,
			Gender: res.Gender,
			Dob: res.Dob,
		},nil
	}
}

func makeGetAllEndpoints(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res , _ := service.GetAllStudent(ctx)
		return res,nil
	}
}

func makeCreateEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		student := request.(repository.StudentRequest)
		res , err:= 	service.CreateStudent(ctx,student)
		if err != nil {
			return err ,err
		}
		return repository.StudentResponse{
			ID: res.ID,
		},nil
	}
}