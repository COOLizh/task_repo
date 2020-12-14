/*Package client sends a request to the executor and receives a response to the user's decision*/
package client

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	config "github.com/COOLizh/task_repo/configs"
	"github.com/COOLizh/task_repo/pkg/models"
	"github.com/COOLizh/task_repo/pkg/pb"
	"google.golang.org/grpc"
)

// getUniqueSolutionID generates unique solution id
func getUniqueSolutionID() (string, error) {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}

// SendSolution sends code to executor and gets the ruslt of testing
func SendSolution(userSolution []byte, language string, timeLimit,
	memoryLimit int64, testCases []models.TestCase) (*models.SolutionResult, error) {
	solutionID, err := getUniqueSolutionID()
	if err != nil {
		return nil, err
	}
	reqTestCases := make([]*pb.CodeHandleRequest_TestCase, len(testCases))
	for i := 0; i < len(testCases); i++ {
		reqTestCases[i] = &pb.CodeHandleRequest_TestCase{
			TestData: []byte(testCases[i].TestData),
			Answer:   []byte(testCases[i].Answer),
		}
	}
	req := &pb.CodeHandleRequest{
		SolutionID:  solutionID,
		Solution:    userSolution,
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   reqTestCases,
	}
	resp, err := getSolutionResults(req)
	if err != nil {
		return nil, err
	}
	results := make([]*models.TestResult, resp.TestsData.PassedTestsCount)
	for i := int64(0); i < resp.TestsData.PassedTestsCount; i++ {
		results[i] = &models.TestResult{
			Status: resp.TestsData.TestResults[i].Result,
			Time:   resp.TestsData.TestResults[i].TimeSpent,
		}
	}
	if resp.TestsData.PassedTestsCount != int64(len(testCases)) {
		results = append(results, &models.TestResult{
			Status: resp.TestsData.TestResults[resp.TestsData.PassedTestsCount].Result,
			Time:   0,
		})
	}

	return &models.SolutionResult{
		ID:               solutionID,
		PassedTestsCount: resp.TestsData.PassedTestsCount,
		TestsCount:       int64(len(testCases)),
		Results:          results,
	}, nil
}

// getClientConnection returns client connection
func getClientConnection(target string) (*grpc.ClientConn, error) {
	return grpc.Dial(target, grpc.WithInsecure())
}

/* getCodeHandleResponse sends a request to the executor to test the task
The response is the task ID and whether the job was created for the task */
func getCodeHandleResponse(conn *grpc.ClientConn, req *pb.CodeHandleRequest) (*pb.CodeHandleResponse, error) {
	codeHandleClient := pb.NewCodeHandlerClient(conn)
	return codeHandleClient.CodeHandle(context.Background(), req)
}

/* getStatusCheckResponse every 2 seconds sends a request to the executor to check the complete testing of the task.
If the task is not tested, it sends a new request after 2 seconds. If it is tested, it returns the test results */
func getStatusCheckResponse(conn *grpc.ClientConn, codeHandleResult *pb.CodeHandleResponse) (*pb.StatusHandleResponse, error) {
	statusCheckClient := pb.NewStatusHandlerClient(conn)
	statusRequest := &pb.StatusHandleRequest{
		ID: codeHandleResult.ID,
	}
	var statusResponse *pb.StatusHandleResponse
	var err error
	var isReadyTask bool
	for !isReadyTask {
		time.Sleep(2 * time.Second)
		statusResponse, err = statusCheckClient.StatusCheck(context.Background(), statusRequest)
		if err != nil {
			return nil, err
		}
		isReadyTask = statusResponse.Ready
	}
	return statusResponse, nil
}

// getSolutionResults returns the result of the user's solution
func getSolutionResults(req *pb.CodeHandleRequest) (*pb.StatusHandleResponse, error) {
	conf := config.New()
	conn, err := getClientConnection(conf.ExecutionerPort)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	codeHandleResult, err := getCodeHandleResponse(conn, req)
	if err != nil {
		return nil, err
	}
	return getStatusCheckResponse(conn, codeHandleResult)
}
