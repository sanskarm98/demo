package tests

import (
	"testing"

	"demo/models"
	"demo/store"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()

	emp := empStore.CreateEmployee("sanskar", "Developer", 60000)

	assert.Equal(t, "sanskar", emp.Name)
	assert.Equal(t, "Developer", emp.Position)
	assert.Equal(t, float64(60000), emp.Salary)
}

func TestGetEmployeeByID(t *testing.T) {
	empStore := store.NewEmployeeStore()

	emp1 := empStore.CreateEmployee("sanskar", "Developer", 60000)
	emp2 := empStore.CreateEmployee("vashu", "Manager", 80000)

	foundEmp1, err1 := empStore.GetEmployeeByID(emp1.ID)
	assert.NoError(t, err1)
	assert.Equal(t, emp1, foundEmp1)

	foundEmp2, err2 := empStore.GetEmployeeByID(emp2.ID)
	assert.NoError(t, err2)
	assert.Equal(t, emp2, foundEmp2)

	_, err3 := empStore.GetEmployeeByID(1000) // Non-existing ID
	assert.Error(t, err3)
}

func TestUpdateEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()

	emp := empStore.CreateEmployee("sanskar", "Developer", 60000)

	updatedEmp, err := empStore.UpdateEmployee(emp.ID, "vashu", "Manager", 80000)
	assert.NoError(t, err)
	assert.Equal(t, "vashu", updatedEmp.Name)
	assert.Equal(t, "Manager", updatedEmp.Position)
	assert.Equal(t, float64(80000), updatedEmp.Salary)

	_, err = empStore.UpdateEmployee(1000, "jack", "Tester", 50000) // Non-existing ID
	assert.Error(t, err)
}

func TestDeleteEmployee(t *testing.T) {
	empStore := store.NewEmployeeStore()

	emp := empStore.CreateEmployee("sanskar", "Developer", 60000)

	err := empStore.DeleteEmployee(emp.ID)
	assert.NoError(t, err)

	_, err = empStore.GetEmployeeByID(emp.ID)
	assert.Error(t, err) // Employee should not exist after deletion

	err = empStore.DeleteEmployee(1000) // Non-existing ID
	assert.Error(t, err)
}

func TestListEmployees(t *testing.T) {
	empStore := store.NewEmployeeStore()

	emp1 := empStore.CreateEmployee("sanskar", "Developer", 60000)
	emp2 := empStore.CreateEmployee("vashu", "Manager", 80000)
	emp3 := empStore.CreateEmployee("jack", "Tester", 50000)

	allEmployees := empStore.ListEmployees(0, 0)
	assert.Len(t, allEmployees, 3)

	// Switch statement to handle different test cases
	page1 := empStore.ListEmployees(1, 2)
	switch {
	case len(page1) != 2:
		t.Errorf("Expected 2 employees on page 1, got %d", len(page1))
	case !containsEmployee(page1, emp1):
		t.Errorf("Employee 1 not found on page 1")
	case !containsEmployee(page1, emp2):
		t.Errorf("Employee 2 not found on page 1")
	}

	page2 := empStore.ListEmployees(2, 2)
	switch {
	case len(page2) != 1:
		t.Errorf("Expected 1 employee on page 2, got %d", len(page2))
	case !containsEmployee(page2, emp3):
		t.Errorf("Employee 3 not found on page 2")
	}

	page3 := empStore.ListEmployees(3, 2)
	switch {
	case len(page3) != 0:
		t.Errorf("Expected 0 employees on page 3, got %d", len(page3))
	}
}

// Helper function to check if an employee is present in a slice of employees
func containsEmployee(employees []models.Employee, emp models.Employee) bool {
	for _, e := range employees {
		if e.ID == emp.ID && e.Name == emp.Name && e.Position == emp.Position && e.Salary == emp.Salary {
			return true
		}
	}
	return false
}
