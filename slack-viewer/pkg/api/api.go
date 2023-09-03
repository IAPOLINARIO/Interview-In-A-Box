package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/slack-viewer/pkg/aws"
	"github.com/slack-viewer/pkg/cache"
	"github.com/slack-viewer/pkg/config"
	"github.com/slack-viewer/pkg/dtos"
	"github.com/slack-viewer/pkg/parser"
)

const (
	region = "us-east-2"
)

type InputUserParameters struct {
	userName     string
	sourceFolder string
	acessKey     string
	secretKey    string
	bucketName   string
	startDate    string
	endDate      string
}

type monthResult struct {
	month   string
	history dtos.SlackHistory
}

func init() {
	cache.CurrentCache().GetStats()
}

func init() {
	cache.CurrentCache().GetStats()
}

// StartServer starts the HTTP server and listens for incoming requests on the
// specified address. It initializes the router and sets up the endpoints for
// handling incoming requests. It also adds a Logger middleware to the router
// to log incoming requests and responses. It returns an error if the server
// fails to start or encounters an error while serving requests.
func StartServer(addr string) error {
	// Initialize the router and endpoints
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/slack/user/report", userReportHandler).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Add the Logger middleware to the router
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	handler := corsHandler.Handler(loggedRouter)

	// Start the server
	log.Printf("Server listening on %s", addr)
	err := http.ListenAndServe(addr, handler)

	if err != nil {
		return err
	}
	return err
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func checkUserQueryParameters(r *http.Request) (result InputUserParameters, err error) {
	// Retrieve query parameters
	username := r.URL.Query().Get("username")
	startDate := r.URL.Query().Get("startdate") //012023
	endDate := r.URL.Query().Get("enddate")     //042023

	// Check if required parameters are present
	if username == "" || startDate == "" || endDate == "" {
		err = fmt.Errorf("Missing required query parameters")
		return result, err
	}

	cfg, err := config.GetConfig()

	if err != nil {
		err = fmt.Errorf("Missing required configuration: %s", err)
		return result, err
	}

	result.userName = username
	result.acessKey = cfg.Server.AccessKey
	result.secretKey = cfg.Server.SecretKey
	result.bucketName = cfg.Server.Bucket
	result.startDate = startDate
	result.endDate = endDate

	return result, err
}

func generateReportForMonth(queryParameters InputUserParameters, month string, resultsCh chan<- monthResult, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Generating report for %s\n", month)

	queryParameters.sourceFolder = month
	monthResults, err := BuildSlackHistory(queryParameters)

	if err != nil {
		errCh <- err
		return
	}
	resultsCh <- monthResult{month: month, history: monthResults}
}

func processResultsAndSendResponse(w http.ResponseWriter, resultsCh <-chan monthResult, errCh <-chan error, doneCh <-chan struct{}) {
	result := []dtos.SlackHistory{}

	select {
	case <-doneCh:
		for monthRes := range resultsCh {
			result = append(result, monthRes.history)
		}
		fmt.Printf("Final result for the %d month(s) requested \n", len(result))

		// Encode the array of Slack history object as JSON and return in response
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
	case err := <-errCh:
		http.Error(w, fmt.Sprintf("Error generating the report. Err: %s", err.Error()), http.StatusOK)
		return
	}
}

func userReportHandler(w http.ResponseWriter, r *http.Request) {

	queryParameters, err := checkUserQueryParameters(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	monthList, err := parser.GenerateMonthList(queryParameters.startDate, queryParameters.endDate)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing date. Err: %s", err.Error()), http.StatusOK)
		return
	}

	resultsCh := make(chan monthResult, len(monthList))
	errCh := make(chan error)
	doneCh := make(chan struct{})

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 6) // Limit concurrent goroutines to 6

	for _, month := range monthList {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(m string) {
			defer func() { <-semaphore }()
			generateReportForMonth(queryParameters, m, resultsCh, errCh, &wg)
		}(month)
	}

	go func() {
		wg.Wait()
		close(doneCh)
		close(resultsCh)
		close(errCh)
	}()

	processResultsAndSendResponse(w, resultsCh, errCh, doneCh)

}

func BuildSlackHistory(params InputUserParameters) (dtos.SlackHistory, error) {
	//Create the cache Key
	cacheKey := fmt.Sprintf("%s_%s", params.userName, params.sourceFolder)

	// Initialize SlackHistory struct
	slackHistory := &dtos.SlackHistory{Month: params.sourceFolder}

	fmt.Printf("Retrieving cache for %s \n", cacheKey)
	//Check if the data is already stored in memory
	cachedInMemory, err := cache.CurrentCache().Get(cacheKey)

	if err != nil {
		return *slackHistory, fmt.Errorf("Error retrieving cache. Err: %s", err.Error())
	}

	//If is in memory, just parse and return the data
	if cachedInMemory.User.ID != "" {
		fmt.Printf("Results are cached in memory for %s\n", params.userName)
		slackHistory = cachedInMemory

	} else {

		//Check if the Data is cached
		cachePath := fmt.Sprintf("%s/%s/%s", params.sourceFolder, aws.CacheFolderName, params.userName)
		fmt.Println("Cache path:", cachePath)
		cacheFileExists, err := aws.CheckIfFileExists(cachePath, aws.FullReportName)

		if err != nil {
			return *slackHistory, fmt.Errorf("Error retrieving cache. Err: %s", err.Error())
		}

		if cacheFileExists {
			fmt.Printf("Cache found in S3 bucket...")
			fullFilePath := fmt.Sprintf("%s/%s", cachePath, aws.FullReportName)
			cachedResult, err := aws.ReadS3JSONFile(fullFilePath)

			if err != nil {
				return *slackHistory, fmt.Errorf("Error retrieving cache. Err: %s", err.Error())
			}

			// Parse the JSON file into a SlackHistory object
			err = aws.UnmarshalData(cachedResult, &slackHistory)

			if err != nil {
				return *slackHistory, fmt.Errorf("Error unmarshaling the data. Err: %s", err.Error())
			}

		} else {
			fmt.Printf("Cache not found in memory or in S3 bucket. Generating report file from scratch...\n")

			// Build Slack history object from scratch
			slackHistory, err = aws.BuildSlackHistoryFromS3Bucket(params.userName, params.sourceFolder)

			if err != nil {
				return *slackHistory, fmt.Errorf("Error building from S3 bucket. Err: %s", err.Error())
			}
		}
	}

	err = cache.CurrentCache().Set(cacheKey, slackHistory)

	if err != nil {
		return *slackHistory, fmt.Errorf("Error saving cache in memory. Err: %s", err.Error())
	}

	return *slackHistory, nil
}
