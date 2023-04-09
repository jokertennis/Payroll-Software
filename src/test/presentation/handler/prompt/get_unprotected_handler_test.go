package prompt

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnprotectedHandler(t *testing.T) {
	cases := map[string]struct {
		email              string
		password           string
		expectedStatusCode int
		expectedResponse   string
	}{
		"execute by registered employee": {
			email:              "potter@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "{\"message\":\"This is the unprotected handler\"}\n",
		},
		"execute by registered administrator": {
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "{\"message\":\"This is the unprotected handler\"}\n",
		},
		"Authentication header is not present": {
			email:              "",
			password:           "",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "{\"message\":\"This is the unprotected handler\"}\n",
		},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/unprotected", nil)
	assert.Nil(t, err)

	for _, value := range cases {
		if value.email == "" && value.password == "" {
			req.Header = http.Header{}
		} else {
			req.SetBasicAuth(value.email, value.password)
		}
		res, err := client.Do(req)
		assert.Nil(t, err)
		defer res.Body.Close()
		assert.Equal(t, value.expectedStatusCode, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, value.expectedResponse, string(data))
	}
}