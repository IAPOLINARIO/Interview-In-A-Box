package slack

import (
	"time"
)

const (
	defaultUserFileName     = "users.json"
	defaultDMFileName       = "dms.json"
	defaultMPIMFileName     = "mpims.json"
	defaultGroupsFileName   = "groups.json"
	defaultChannelsFileName = "channels.json"
)

func ConvertTimeStamp(ts string) string {
	unixTime, err := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	if err != nil {
		panic(err)
	}
	seconds, err := time.ParseDuration(ts + "s")
	if err != nil {
		panic(err)

	}
	t := unixTime.Add(seconds)

	return t.Format("2006-01-02 15:04:05")

}
