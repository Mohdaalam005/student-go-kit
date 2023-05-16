package transport

import (
	"context"
	"encoding/json"
	"net/http"

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
		endpoint.CreateStudent ,
		decodeCreateRequest,
		encodeResponse ,
	))

	route.Methods("GET").Path("/students").Handler(httptransport.NewServer(
		endpoint.GetAllStudent,
		decodeGetAllStudent,
		encodeResponse,
	))

	return route

}

func decodeGetAllStudent(ctx context.Context , r *http.Request) (interface{}, error) {
	var students []models.Student
	// json.NewDecoder(r.Body).Decode(&students)
	return students, nil
}

func encodeResponse(ctx context.Context, writer http.ResponseWriter , response interface{})error{
	return json.NewEncoder(writer).Encode(response) ;
}

func decodeCreateRequest(ctx context.Context ,request *http.Request)(interface{},error) {
	var student repository.StudentRequest
	json.NewDecoder(request.Body).Decode(&student)
	return student , nil 
}


func middleWare(handle http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/josn")
		handle.ServeHTTP(w,r)
	})
}