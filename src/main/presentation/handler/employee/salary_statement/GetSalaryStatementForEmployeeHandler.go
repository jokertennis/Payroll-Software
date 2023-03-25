package salarystatement

import (
	"usr/local/go/swagger/employee/salary_statement/restapi/operations/get_salary_statement_for_employee"

	"github.com/go-openapi/runtime/middleware"
)

// Embedded GetEmployeeSalaryStatementHandler interface 
// because we cannot define new methods on non-local type.
type GetSalaryStatementForEmployeeHandler struct {
	GetSalaryStatementForEmployeeGetEmployeeSalaryStatementHandler get_salary_statement_for_employee.GetEmployeeSalaryStatementHandler
}

// Handle executing the request and returning a response
func Handle(params get_salary_statement_for_employee.GetEmployeeSalaryStatementParams) middleware.Responder {
	return get_salary_statement_for_employee.NewGetEmployeeSalaryStatementOK()
}