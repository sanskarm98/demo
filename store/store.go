package store

import (
	"errors"
	"sync"

	"demo/models"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type EmployeeStore struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
}

func (s *EmployeeStore) CreateEmployee(name, position string, salary float64) models.Employee {
	s.mu.Lock()
	defer s.mu.Unlock()

	employee := models.Employee{
		ID:       s.nextID,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	s.employees[s.nextID] = employee
	s.nextID++
	return employee
}

func (s *EmployeeStore) GetEmployeeByID(id int) (models.Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	employee, exists := s.employees[id]
	if !exists {
		return models.Employee{}, ErrEmployeeNotFound
	}
	return employee, nil
}

func (s *EmployeeStore) UpdateEmployee(id int, name, position string, salary float64) (models.Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	employee, exists := s.employees[id]
	if !exists {
		return models.Employee{}, ErrEmployeeNotFound
	}

	employee.Name = name
	employee.Position = position
	employee.Salary = salary
	s.employees[id] = employee
	return employee, nil
}

func (s *EmployeeStore) DeleteEmployee(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.employees[id]
	if !exists {
		return ErrEmployeeNotFound
	}

	delete(s.employees, id)
	return nil
}

// ListEmployees returns a paginated list of employees.
// If page and pageSize are 0, it returns all employees.
func (s *EmployeeStore) ListEmployees(page, pageSize int) []models.Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()

	employees := make([]models.Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		employees = append(employees, emp)
	}

	// If page or pageSize is not provided, return all employees.
	if page == 0 || pageSize == 0 {
		return employees
	}

	start := (page - 1) * pageSize
	if start >= len(employees) {
		return []models.Employee{}
	}

	end := start + pageSize
	if end > len(employees) {
		end = len(employees)
	}

	return employees[start:end]
}
