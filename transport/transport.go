package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mohdaalam/005/student/endpoints"
	"github.com/mohdaalam/005/student/models"
	"github.com/mohdaalam/005/student/repository"
)

func NewHTTPServer(ctx context.Context, endpoint endpoints.Endpoints) http.Handler {
	route := mux.NewRouter()
	route.Use(middleWare)
	route.Methods("POST").Path("/students").Handler(httptransport.NewServer(
		endpoint.CreateStudent,
		decodeCreateRequest,
		encodeResponse,
	))

	route.Methods("GET").Path("/students").Handler(httptransport.NewServer(
		endpoint.GetAllStudent,
		decodeGetAllStudent,
		encodeResponse,
	))
	route.Methods("GET").Path("/students/{id}").Handler(httptransport.NewServer(
		endpoint.GetStudentById,
		decodeGetStudentById,
		encodeResponse,
	))
	route.Methods("DELETE").Path("/students/{id}").Handler(httptransport.NewServer(
		endpoint.DeleteStudentById,
		decodeDeleteStudentById,
		encodeResponse,
	))
	route.Methods("PUT").Path("/students/{id}").Handler(httptransport.NewServer(
		endpoint.UpdateStudent,
		decodeUpdateStudent,
		encodeResponse,
	))

	return route

}

func decodeUpdateStudent(ctx context.Context, r *http.Request) (interface{}, error) {
	var student models.Student
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	json.NewDecoder(r.Body).Decode(&student)
	student.ID = id
	return student, nil

}

func decodeDeleteStudentById(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}
	student := &models.Student{
		ID: id,
	}

	return student, nil
}

func decodeGetStudentById(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}
	student := &models.Student{
		ID: id,
	}

	return student, nil
}

func decodeGetAllStudent(ctx context.Context, r *http.Request) (interface{}, error) {
	var students []models.Student
	// json.NewDecoder(r.Body).Decode(&students)
	return students, nil
}

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}

func decodeCreateRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var student repository.StudentRequest
	json.NewDecoder(request.Body).Decode(&student)
	return student, nil
}

func middleWare(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/josn")
		handle.ServeHTTP(w, r)
	})
}
