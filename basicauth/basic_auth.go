package basicauth

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"net/http"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/administrator_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
)

type ExpectedUserInformation struct {
	MailAddress string
	Password    string
}

// TODO:create as enum. only Exmployee and Administrator
type Executer struct {
	Executer string
}

func BasicAuth(employeeRepository employee_repository.EmployeeRepository, administratorRepository administrator_repository.AdministratorRepository, executer Executer, httpRequest *http.Request) (mailAddress string, statusCode int, errorMessage error) {
	// Extract the mailAddress and password from the request
	// Authorization header. If no Authentication header is present
	// or the header value is invalid, then the 'ok' return value
	// will be false.
	mailAddress, password, ok := httpRequest.BasicAuth()

	if ok {
		// Calculate SHA-256 hashes for the provided and expected password.
		passwordHash := sha256.Sum256([]byte(password))

		expectedUserInformation, err := GetExpectedUserInformation(employeeRepository, administratorRepository, mailAddress, executer)
		if err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
		}
		if expectedUserInformation == nil {
			return "", http.StatusUnauthorized, fmt.Errorf("Unauthorized. Specified mailAddress is not found in registered user datas.")
		}

		expectedPasswordHash := sha256.Sum256([]byte(expectedUserInformation.Password))

		// Use the subtle.ConstantTimeCompare() function to check if
		// the provided password hash equal the expected password hash.
		// ConstantTimeCompare will return 1 if the values are equal, or 0 otherwise.
		// Importantly, we should to do the work to evaluate password
		// before checking the return values to avoid leaking information.
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		// Make sure that the password is correct.
		if passwordMatch {
			return mailAddress, 0, nil
		} else {
			return "", http.StatusUnauthorized, fmt.Errorf("Unauthorized. Please check to see if password is correct.")
		}
	}

	// If the Authentication header is not present or invalid,
	// then set a WWW-Authenticate header to inform the client
	// that we expect them to use basic authentication and
	// send a 401 Unauthorized response.
	return "", http.StatusUnauthorized, fmt.Errorf("Unauthorized. Please check to see if Authentication header is not present or invalid.")
}

func GetExpectedUserInformation(employeeRepository employee_repository.EmployeeRepository, administratorRepository administrator_repository.AdministratorRepository, mailAddress string, executer Executer) (*ExpectedUserInformation, error) {
	switch executer.Executer {
	case "Employee":
		employee, err := employeeRepository.GetEmployeeByMailAddress(mailAddress)
		if err != nil {
			return nil, err
		}
		if employee == nil {
			return nil, nil
		}
		return &ExpectedUserInformation{employee.MailAddress, employee.Password}, nil
	case "Administrator":
		administrator, err := administratorRepository.GetAdministratorByMailAddress(mailAddress)
		if err != nil {
			return nil, err
		}
		if administrator == nil {
			return nil, nil
		}
		return &ExpectedUserInformation{administrator.MailAddress, administrator.Password}, nil
	}
	return nil, fmt.Errorf("not support specified Executer:%s", executer.Executer)
}
