package requests

type NetworkDoc struct {
	Project                  int    `json:"Project"`
	WBSElement               int    `json:"WBSElement"`
	Network                  int    `json:"Network"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
