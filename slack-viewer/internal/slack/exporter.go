package slack

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/slack-viewer/internal/dtos"
	"github.com/slack-viewer/internal/mongo"
)

var (
	mongoMgt   *mongo.MongoManager
	mainFolder = "C:\\Users\\Ismael\\Downloads\\slack_history\\mar_2023"
)

/* func main() {

	//importData()
	//parseFiles(mainFolder)

	//parseAndWriteMessages()
	buildSlackHistoryFromJSONFiles()

} */

func buildSlackHistoryFromJSONFiles() {
	// Initialize SlackHistory struct
	slackHistory := dtos.SlackHistory{}

	// Find user data in JSON file
	userData, err := FindUserDataInJSONFile("erik.nelson", mainFolder)
	if err != nil {
		panic(err)
	}

	// Check if user data is not empty
	if userData.Name != "" {

		// Start the timer to track how long it takes to generate the report
		start := time.Now()

		// Get the base date from the main folder path
		baseDate := filepath.Base(mainFolder)

		// Set the user data in SlackHistory
		slackHistory.User = userData

		// Print some useful information for debugging purposes
		fmt.Println("User data:", userData.Name)
		fmt.Println("Base date:", baseDate)

		// Get DMs data from JSON files
		dms, err := getDMsDataFromJSONFiles(userData, mainFolder)
		if err != nil {
			panic(err)
		}
		// Set the DMs data in SlackHistory
		slackHistory.DMs = dms

		// Print the number of DMs groups found
		fmt.Println("DMs group count:", len(dms))

		// Get MPIMs data from JSON files
		mpims, err := getMpimsDataFromJSONFiles(userData, mainFolder)
		if err != nil {
			panic(err)
		}
		// Set the MPIMs data in SlackHistory
		slackHistory.Mpims = mpims

		// Print the number of MPIMs groups found
		fmt.Println("MPIMs group count:", len(mpims))

		// Get groups data from JSON files
		groups, err := getGroupsDataFromJSONFiles(userData, mainFolder)
		if err != nil {
			panic(err)
		}
		// Set the groups data in SlackHistory
		slackHistory.Groups = groups

		// Print the number of groups found
		fmt.Println("Groups count:", len(groups))

		// Get channels data from JSON files
		channels, err := getChannelsDataFromJSONFiles(userData, mainFolder)
		if err != nil {
			panic(err)
		}
		// Set the channels data in SlackHistory
		slackHistory.Channels = channels

		// Print the number of channels found
		fmt.Println("Channels count:", len(channels))

		// Generate the Slack history report
		GenerateSlackHistoryReport(slackHistory, fmt.Sprintf("%s-%s.txt", slackHistory.User.Name, baseDate))

		// Calculate how long it took to generate the report and print it
		elapsed := time.Since(start)
		fmt.Printf("\nReport generation took %s to complete.\n\n", elapsed)
	}
}

func getChannelsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserChannelsDataInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getGroupsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserGroupsDataInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getMpimsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserMPIMsInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func getDMsDataFromJSONFiles(user dtos.SlackUser, basePath string) (result []dtos.SlackHistoryGroup, err error) {
	result, err = FindUserDMsInJSONFiles(user, basePath)

	if err != nil {
		panic(err)
	}

	return result, nil

}

func importData() {

	start := time.Now()

	mongoMgt, err := mongo.NewMongoManager()
	mongoMgt.DbName += filepath.Base(mainFolder)

	if err != nil {
		fmt.Println("Erro on main: ", err)
	}

	//mongoMgt.DeleteDatabase(mongoMgt.dbName)

	defer mongoMgt.Disconnect()

	fmt.Println(start)

	totalRecords := mongo.ReadSlackFilesParallel(mainFolder, *mongoMgt, 10)
	totalUsers := mongo.PersistSlackUser(fmt.Sprintf("%s/%s", mainFolder, "users.json"), *mongoMgt)
	totalGroups := mongo.PersistSlackGroups(fmt.Sprintf("%s/%s", mainFolder, "groups.json"), *mongoMgt)
	totalChannels := mongo.PersistSlackChannels(fmt.Sprintf("%s/%s", mainFolder, "channels.json"), *mongoMgt)
	totalDms := mongo.PersistSlackDMsJson(fmt.Sprintf("%s/%s", mainFolder, "dms.json"), *mongoMgt)
	totalMPIMs := mongo.PersistSlackMPIMsJson(fmt.Sprintf("%s/%s", mainFolder, "mpims.json"), *mongoMgt)

	elapsed := time.Since(start)
	log.Printf("\n\nImport is done! A total of %d slack users, %d groups, %d channels, %d dms, %d multiparty messages and %d slack messages were imported to MongoDB in %s \n", totalUsers, totalGroups, totalChannels, totalDms, totalMPIMs, totalRecords, elapsed)
}
