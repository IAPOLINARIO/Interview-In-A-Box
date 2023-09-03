package dtos

type SlackHistoryStats struct {
	UserDisplayName string `json:"user_display_name"`
	TotalDMs        int    `json:"total_dms"`
	TotalMpims      int    `json:"total_mpims"`
	TotalGroups     int    `json:"total_groups"`
	TotalChannels   int    `json:"total_channels"`
}
type SlackHistory struct {
	Month    string
	User     SlackUser
	DMs      []SlackHistoryGroup
	Mpims    []SlackHistoryGroup
	Groups   []SlackHistoryGroup
	Channels []SlackHistoryGroup
}

type SlackHistoryGroup struct {
	GroupName string
	Groups    []SlackMessagesGroup
}

type SlackMessagesGroup struct {
	Context      string
	MessageCount int
	Messages     []SlackMessage
}

type SlackMessage struct {
	Attachments []Attachment `json:"attachments"`
	Blocks      []Block      `json:"blocks"`
	SourceTeam  string       `json:"source_team"`
	Team        string       `json:"team"`
	Text        string       `json:"text"`
	Ts          string       `json:"ts"`
	Type        string       `json:"type"`
	User        string       `json:"user"`
	UserProfile UserProfile  `json:"user_profile"`
	UserTeam    string       `json:"user_team"`
}

type UserProfile struct {
	AvatarHash        string `json:"avatar_hash"`
	DisplayName       string `json:"display_name"`
	FirstName         string `json:"first_name"`
	Image72           string `json:"image_72"`
	IsRestricted      bool   `json:"is_restricted"`
	IsUltraRestricted bool   `json:"is_ultra_restricted"`
	Name              string `json:"name"`
	RealName          string `json:"real_name"`
	Team              string `json:"team"`
}

type Block struct {
	BlockID  string     `json:"block_id"`
	Elements []Elements `json:"elements"`
	Type     string     `json:"type"`
}

type Elements struct {
	Element []Element `json:"elements"`
	Type    string    `json:"type"`
}

type Element struct {
	Text string `json:"text"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Attachment struct {
	Fallback    string `json:"fallback"`
	FromURL     string `json:"from_url"`
	ID          int64  `json:"id"`
	ImageBytes  int64  `json:"image_bytes"`
	ImageHeight int64  `json:"image_height"`
	ImageURL    string `json:"image_url"`
	ImageWidth  int64  `json:"image_width"`
	ServiceName string `json:"service_name"`
	Text        string `json:"text"`
	Title       string `json:"title"`
	TitleLink   string `json:"title_link"`
}

type SlackUser struct {
	Deleted   bool    `json:"deleted"`
	ID        string  `json:"id"`
	IsAppUser bool    `json:"is_app_user"`
	IsBot     bool    `json:"is_bot"`
	Name      string  `json:"name"`
	Profile   Profile `json:"profile"`
	TeamID    string  `json:"team_id"`
	Updated   int64   `json:"updated"`
}

type Profile struct {
	AvatarHash             string        `json:"avatar_hash"`
	DisplayName            string        `json:"display_name"`
	DisplayNameNormalized  string        `json:"display_name_normalized"`
	Email                  string        `json:"email"`
	Fields                 struct{}      `json:"fields"`
	FirstName              string        `json:"first_name"`
	Image1024              string        `json:"image_1024"`
	Image192               string        `json:"image_192"`
	Image24                string        `json:"image_24"`
	Image32                string        `json:"image_32"`
	Image48                string        `json:"image_48"`
	Image512               string        `json:"image_512"`
	Image72                string        `json:"image_72"`
	ImageOriginal          string        `json:"image_original"`
	IsCustomImage          bool          `json:"is_custom_image"`
	LastName               string        `json:"last_name"`
	Phone                  string        `json:"phone"`
	RealName               string        `json:"real_name"`
	RealNameNormalized     string        `json:"real_name_normalized"`
	Skype                  string        `json:"skype"`
	StatusEmoji            string        `json:"status_emoji"`
	StatusEmojiDisplayInfo []interface{} `json:"status_emoji_display_info"`
	StatusExpiration       int64         `json:"status_expiration"`
	StatusText             string        `json:"status_text"`
	StatusTextCanonical    string        `json:"status_text_canonical"`
	Team                   string        `json:"team"`
	Title                  string        `json:"title"`
}

type SlackGroup struct {
	Created    int64    `json:"created"`
	Creator    string   `json:"creator"`
	ID         string   `json:"id"`
	IsArchived bool     `json:"is_archived"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	Purpose    Purpose  `json:"purpose"`
	Topic      Topic    `json:"topic"`
}

type Purpose struct {
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
	Value   string `json:"value"`
}

type Topic struct {
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
	Value   string `json:"value"`
}

type SlackChannel struct {
	Created    int64    `json:"created"`
	Creator    string   `json:"creator"`
	ID         string   `json:"id"`
	IsArchived bool     `json:"is_archived"`
	IsGeneral  bool     `json:"is_general"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	Purpose    Purpose  `json:"purpose"`
	Topic      Topic    `json:"topic"`
}

type SlackDM struct {
	Created int64    `json:"created"`
	ID      string   `json:"id"`
	Members []string `json:"members"`
}

type MPIM struct {
	Created    int64    `json:"created"`
	Creator    string   `json:"creator"`
	ID         string   `json:"id"`
	IsArchived bool     `json:"is_archived"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	Purpose    Purpose  `json:"purpose"`
	Topic      Topic    `json:"topic"`
}
