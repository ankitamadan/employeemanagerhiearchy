package main

//PopulateEmployeeNameID create employee id name relationship
func PopulateEmployeeNameID() map[int]string {
	var employeeIDNameMap = make(map[int]string)

	employeeIDNameMap[100] = "Alan"
	employeeIDNameMap[220] = "Martin"
	employeeIDNameMap[150] = "Jamie"
	employeeIDNameMap[275] = "Alex"
	employeeIDNameMap[400] = "Steve"
	employeeIDNameMap[190] = "David"

	return employeeIDNameMap
}

//PopulateEmployeeManagerRelationship create employee manager relationship
func PopulateEmployeeManagerRelationship() map[int]int {
	var employeeIDManagerIDMap = make(map[int]int)

	employeeIDManagerIDMap[100] = 150
	employeeIDManagerIDMap[220] = 100
	employeeIDManagerIDMap[150] = 0
	employeeIDManagerIDMap[275] = 100
	employeeIDManagerIDMap[400] = 150
	employeeIDManagerIDMap[190] = 400

	return employeeIDManagerIDMap
}
