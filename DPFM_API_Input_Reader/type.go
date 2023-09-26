package dpfm_api_input_reader

type SDC struct {
	ConnectionKey     string     `json:"connection_key"`
	Result            bool       `json:"result"`
	RedisKey          string     `json:"redis_key"`
	Filepath          string     `json:"filepath"`
	APIStatusCode     int        `json:"api_status_code"`
	RuntimeSessionID  string     `json:"runtime_session_id"`
	BusinessPartnerID *int       `json:"business_partner"`
	ServiceLabel      string     `json:"service_label"`
	APIType           string     `json:"api_type"`
	ProjectDoc        ProjectDoc `json:"Project"`
	APISchema         string     `json:"api_schema"`
	Accepter          []string   `json:"accepter"`
	Deleted           bool       `json:"deleted"`
}

type ProjectDoc struct {
	Project                  *int          `json:"Project"`
	DocType                  *string       `json:"DocType"`
	DocVersionID             *int          `json:"DocVersionID"`
	DocID                    *string       `json:"DocID"`
	FileExtension            *string       `json:"FileExtension"`
	FileName                 *string       `json:"FileName"`
	FilePath                 *string       `json:"FilePath"`
	DocIssuerBusinessPartner *int          `json:"DocIssuerBusinessPartner"`
	WBSElementDoc            WBSElementDoc `json:"WBSElementDoc"`
}

type WBSElementDoc struct {
	Project                  int           `json:"Project"`
	WBSElement               int           `json:"WBSElement"`
	DocType                  string        `json:"DocType"`
	DocVersionID             int           `json:"DocVersionID"`
	DocID                    string        `json:"DocID"`
	FileExtension            string        `json:"FileExtension"`
	FileName                 string        `json:"FileName"`
	FilePath                 string        `json:"FilePath"`
	DocIssuerBusinessPartner int           `json:"DocIssuerBusinessPartner"`
	WBSElementDoc            WBSElementDoc `json:"WBSElementDoc"`
}

type NetworkDoc struct {
	Project                  int           `json:"Project"`
	WBSElement               int           `json:"WBSElement"`
	Network                  int           `json:"Network"`
	DocType                  string        `json:"DocType"`
	DocVersionID             int           `json:"DocVersionID"`
	DocID                    string        `json:"DocID"`
	FileExtension            string        `json:"FileExtension"`
	FileName                 string        `json:"FileName"`
	FilePath                 string        `json:"FilePath"`
	DocIssuerBusinessPartner int           `json:"DocIssuerBusinessPartner"`
	WBSElementDoc            WBSElementDoc `json:"WBSElementDoc"`
}
