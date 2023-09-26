package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	ProjectDoc 		*[]ProjectDoc	    `json:"ProjectDoc"`
	WBSElementDoc	*[]WBSElementDoc	`json:"WBSElementDoc"`
	NetworkDoc   	*[]NetworkDoc	    `json:"NetworkDoc"`
}

type ProjectDoc struct {
	Project                  int    `json:"Project"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}

type WBSElementDoc struct {
	Project                  int    `json:"Project"`
	WBSElement               int    `json:"WBSElement"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}

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

package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToProjectDoc(rows *sql.Rows) (*[]ProjectDoc, error) {
	defer rows.Close()
	projectDoc := make([]ProjectDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &ProjectDoc{}

		err := rows.Scan(
			&pm.Project,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &projectDoc, err
		}

		data := pm
		projectDoc = append(projectDoc, ProjectDoc{
			Project:                  data.Project,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &projectDoc, nil
	}

	return &projectDoc, nil
}

func ConvertToWBSElementDoc(rows *sql.Rows) (*[]WBSElementDoc, error) {
	defer rows.Close()
	wBSElementDoc := make([]WBSElementDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &WBSElementDoc{}

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &wBSElementDoc, err
		}

		data := pm
		wBSElementDoc = append(wBSElementDoc, WBSElementDoc{
			Project:                  data.Project,
			WBSElement:      		  data.WBSElement,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &wBSElementDoc, nil
	}

	return &wBSElementDoc, nil
}

func ConvertToNetworkDoc(rows *sql.Rows) (*[]NetworkDoc, error) {
	defer rows.Close()
	networkDoc := make([]NetworkDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &NetworkDoc{}

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.Network,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &wBSElementDoc, err
		}

		data := pm
		wBSElementDoc = append(wBSElementDoc, WBSElementDoc{
			Project:                  data.Project,
			WBSElement:      		  data.WBSElement,
			Network:      	    	  data.Network,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &networkDoc, nil
	}

	return &networkDoc, nil
}
