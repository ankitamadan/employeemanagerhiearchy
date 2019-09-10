package main

import (
	"fmt"
	"log"
	"sort"
)

var result = make(map[string]int)
var errorList []string
var resultManagerEmployeeMap = make(map[string][]string)

func createEmployeeManagerMap() map[string]string {
	var managerEmployeeMap = make(map[string]string)
	var employeeIDNameMap = PopulateEmployeeNameID()
	var employeeIDManagerIDMap = PopulateEmployeeManagerRelationship()

	for empID, empName := range employeeIDNameMap {
		for employeeID, managerID := range employeeIDManagerIDMap {
			if empID == employeeID {
				if employeeIDNameMap[managerID] == "" {
					managerEmployeeMap[empName] = empName
				} else {
					managerEmployeeMap[empName] = employeeIDNameMap[managerID]
				}
			}
		}
	}

	validateIfEmployeeHasAManager(managerEmployeeMap)
	return managerEmployeeMap
}

func managerDirectReporteeMap() {

	var employeeManagerMap = createEmployeeManagerMap()
	var managerEmployeeMap = make(map[string][]string)

	var directReportList []string

	for key, value := range employeeManagerMap {
		if key != value {
			directReportList = managerEmployeeMap[value]
			if len(directReportList) == 0 {
				directReportList := []string{}
				managerEmployeeMap[value] = directReportList
			}
			directReportList = append(directReportList, key)
			managerEmployeeMap[value] = directReportList
		}
	}

	for key := range employeeManagerMap {
		createManagerEmployeeCountMap(key, managerEmployeeMap)
	}
}

func createManagerEmployeeCountMap(manager string, managerEmployeeMap map[string][]string) int {
	var count = 0

	if _, exist := managerEmployeeMap[manager]; !exist {
		result[manager] = 0
		return 0
	} else if _, exist = result[manager]; exist {
		count = result[manager]
	} else {
		var directReportList []string
		directReportList = managerEmployeeMap[manager]
		resultManagerEmployeeMap[manager] = directReportList
		count = len(directReportList)
		for _, directReportEmp := range directReportList {
			count = count + createManagerEmployeeCountMap(directReportEmp, managerEmployeeMap)
		}
		result[manager] = count
	}

	return count
}

func sortManagerByNumberOfReporteesAndReturnCEO() string {

	sortedByHighestHierarchyManager := make([]string, 0, len(result))

	for key := range result {
		sortedByHighestHierarchyManager = append(sortedByHighestHierarchyManager, key)
	}
	sort.Slice(sortedByHighestHierarchyManager, func(i, j int) bool {
		return result[sortedByHighestHierarchyManager[i]] > result[sortedByHighestHierarchyManager[j]]
	})

	for _, mgrName := range sortedByHighestHierarchyManager {
		for empName, _ := range resultManagerEmployeeMap {
			if empName == mgrName {
				log.Printf("%v reports directly to %s\n", resultManagerEmployeeMap[mgrName], mgrName)
			} else {
				errInvalidManager := IsEmployeeManager(mgrName, resultManagerEmployeeMap[mgrName])
				if errInvalidManager != nil {
					errorList = append(errorList, errInvalidManager.Error())
				}
			}
		}
	}

	displayErrors(errorList)
	return sortedByHighestHierarchyManager[0]
}

//PrettyPrintEmployeeManagerHierarchy pretty print employee manager hierarchy
func PrettyPrintEmployeeManagerHierarchy() {
	managerDirectReporteeMap()
	ceo := sortManagerByNumberOfReporteesAndReturnCEO()

	fmt.Printf("\n\n")
	fmt.Printf("**********EMPLOYEE HIERARCHY**********\n")
	fmt.Printf("---------------%s(CEO)-------------\n", ceo)
	fmt.Printf("--------------%v------------\n", resultManagerEmployeeMap[ceo])

	employeeManagerMap := resultManagerEmployeeMap[ceo]
	for _, value := range employeeManagerMap {
		fmt.Printf("-----%v----", resultManagerEmployeeMap[value])
	}
}

func displayErrors(errorList []string) {
	uniqueErrors := UniqueValuesInList(errorList)
	for _, value := range uniqueErrors {
		log.Printf("%s", value)
	}
}

func validateIfEmployeeHasAManager(managerEmployeeMap map[string]string) {
	if err := DoesEmployeeHaveManager(managerEmployeeMap); err != nil {
		errorList = append(errorList, err.Error())
	}
}
