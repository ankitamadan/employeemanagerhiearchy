package main

import (
	"fmt"

	"errors"
)

//IsEmployeeManager validates if employee is a manager
func IsEmployeeManager(managerName string, directReporteeList []string) error {
	if len(directReporteeList) == 0 {
		var employeeNotManager = fmt.Sprintf("Employee: %s is not a manager", managerName)
		return errors.New(employeeNotManager)
	}

	return nil
}

//DoesEmployeeHaveManager validates if employee has a manager
func DoesEmployeeHaveManager(managerEmployeeMap map[string]string) error {
	if key := IsKeyValueEqualInMap(managerEmployeeMap); !(key == "") {
		var employeeDoesNotHaveManager = fmt.Sprintf("Employee: %s does not have a manager with error code", key)
		return errors.New(employeeDoesNotHaveManager)
	}

	return nil
}
