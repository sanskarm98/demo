package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"demo/models"
	"demo/store"

	"github.com/gorilla/mux"
)

type Server struct {
	Store *store.EmployeeStore
}

func NewServer(store *store.EmployeeStore) *Server {
	return &Server{Store: store}
}

func (s *Server) HandleCreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdEmployee := s.Store.CreateEmployee(emp.Name, emp.Position, emp.Salary)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEmployee)
}

func (s *Server) HandleGetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := s.Store.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(employee)
}

func (s *Server) HandleUpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedEmployee, err := s.Store.UpdateEmployee(emp.ID, emp.Name, emp.Position, emp.Salary)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedEmployee)
}

func (s *Server) HandleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = s.Store.DeleteEmployee(id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) HandleListEmployees(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 0 // Set to 0 to fetch all employees
	}

	employees := s.Store.ListEmployees(page, pageSize)
	json.NewEncoder(w).Encode(employees)
}
