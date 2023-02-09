package basicauth

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"net/http"
	"usr/local/go/server/gen/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ExpectedUserInformation struct {
	MailAddress string
	Password    string
}

// TODO:create as enum. only Exmployee and Administrator
type Executer struct {
	Executer string
}

func BasicAuth(ctx context.Context, db boil.ContextExecutor, executer Executer, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the username and password from the request
		// Authorization header. If no Authentication header is present
		// or the header value is invalid, then the 'ok' return value
		// will be false.
		mailAddress, password, ok := r.BasicAuth()

		if ok {
			// Calculate SHA-256 hashes for the provided and expected passwords.
			passwordHash := sha256.Sum256([]byte(password))

			expectedUserInformation, err := GetExpectedUserInformation(ctx, db, mailAddress, executer)
			if err != nil {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, fmt.Sprintf("InternalServerError:error:%s", err), http.StatusInternalServerError)
				return
			}
			if expectedUserInformation == nil {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized.Specified mailAddress is not found in registered user datas.", http.StatusUnauthorized)
				return
			}

			expectedPasswordHash := sha256.Sum256([]byte(expectedUserInformation.Password))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided password hash equal the expected password hash.
			// ConstantTimeCompare will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate password
			// before checking the return values to avoid leaking information.
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			// If the password are correct, then call the next handler in the chain.
			// Make sure to return afterwards, so that none of the code below is run.
			if passwordMatch {
				next.ServeHTTP(w, r)
				return
			} else {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized.Please check to see if password is correct.", http.StatusUnauthorized)
				return
			}
		}

		// If the Authentication header is not present or invalid,
		// then set a WWW-Authenticate header to inform the client
		// that we expect them to use basic authentication and
		// send a 401 Unauthorized response.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.Please check to see if Authentication header is not present or invalid.", http.StatusUnauthorized)
	})
}

func GetExpectedUserInformation(ctx context.Context, db boil.ContextExecutor, mailAddress string, executer Executer) (*ExpectedUserInformation, error) {
	switch executer.Executer {
	case "Employee":
		employees, err := models.Employees(qm.Where("mail_address=?", mailAddress)).All(ctx, db)
		if err != nil {
			return nil, err
		}
		if employees == nil {
			return nil, nil
		}
		return &ExpectedUserInformation{employees[0].MailAddress, employees[0].Password}, nil
	case "Administrator":
		administrators, err := models.Administrators(qm.Where("mail_address=?", mailAddress)).All(ctx, db)
		if err != nil {
			return nil, err
		}
		if administrators == nil {
			return nil, nil
		}
		return &ExpectedUserInformation{administrators[0].MailAddress, administrators[0].Password}, nil
	}
	return nil, fmt.Errorf("not support specified Executer:%s", executer.Executer)
}
