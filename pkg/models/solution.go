package models

// SolutionResult contains needful information about user solution result
type SolutionResult struct {
	ID               string        `json:"id"`
	PassedTestsCount int64         `json:"passed_tests_count"`
	TestsCount       int64         `json:"tests_count"`
	Results          []*TestResult `json:"results"`
}

// TestResult contains information about test result
type TestResult struct {
	Status string `json:"status"`
	Time   int64  `json:"time"`
}

// SolutionRequest contains fields that server gets from POST-Solution request
type SolutionRequest struct {
	Code     []byte `json:"solution"`
	Language string `json:"language"`
}
