package auth

var TestCasesRegisterUser = []struct {
	username string
	password string
	httpCode int
}{
	{"test", "test", 422},
	//{"", "", 422},
}

// func TestRegisterUser(t *testing.T) {
// 	rec := httptest.NewRecorder()
// 	handler := http.HandlerFunc(RegisterUser)
// 	for _, v := range testCasesRegisterUser {
// 		payload := fmt.Sprintf(`{
// 			"username": "%s",
// 			"password": "%s"
// 			}`, v.username, v.password)

// 		r, _ := http.NewRequest(http.MethodPost, "/registration", strings.NewReader(payload))
// 		handler.ServeHTTP(rec, r)

// 		assert.Equal(t, rec.Code, v.httpCode)
// 	}
// }

// func TestRegisterUser(t *testing.T) {

// }
