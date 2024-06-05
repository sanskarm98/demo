package tests

import (
	"testing"

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

	page1 := empStore.ListEmployees(1, 2)
	assert.Len(t, page1, 2)
	assert.Contains(t, page1, emp1)
	assert.Contains(t, page1, emp2)

	page2 := empStore.ListEmployees(2, 2)
	assert.Len(t, page2, 1)
	assert.Contains(t, page2, emp3)

	page3 := empStore.ListEmployees(3, 2)
	assert.Len(t, page3, 0) // No more employees on page 3
}
