package user

// IPStack details for particular ip address
type IPStack struct {
	IP            string   `json:"ip"`
	Type          string   `json:"type"`
	ContinentCode string   `json:"continent_code"`
	ContinentName string   `json:"continent_name"`
	CountryCode   string   `json:"country_code"`
	CountryName   string   `json:"country_name"`
	RegionCode    string   `json:"region_code"`
	RegionName    string   `json:"region_name"`
	City          string   `json:"city"`
	Zip           string   `json:"zip"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Location      Location `json:"location"`
}

// Location detail of particular ip address
type Location struct {
	GeonameID               int         `json:"geoname_id"`
	Capital                 string      `json:"capital"`
	Languages               []Languages `json:"languages"`
	CountryFlag             string      `json:"country_flag"`
	CountryFlagEmoji        string      `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string      `json:"country_flag_emoji_unicode"`
	CallingCode             string      `json:"calling_code"`
	IsEu                    bool        `json:"is_eu"`
}

// Languages detail of particular ip address
type Languages struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

type IPAddr struct {
	Host string `json:"ipAddr"`
}
