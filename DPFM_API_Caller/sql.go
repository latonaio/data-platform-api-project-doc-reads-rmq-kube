package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-project-doc-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-project-doc-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var projectDoc *[]dpfm_api_output_formatter.ProjectDoc
	var wBSElementDoc *[]dpfm_api_output_formatter.WBSElementDoc
	var networkDoc *[]dpfm_api_output_formatter.NetworkDoc

	for _, fn := range accepter {
		switch fn {
		case "ProjectDoc":
			func() {
				projectDoc = c.ProjectDoc(input, output, errs, log)
			}()
		case "WBSElementDoc":
			func() {
				wBSElementDoc = c.WBSElementDoc(input, output, errs, log)
			}()
		case "NetworkDoc":
			func() {
				networkDoc = c.NetworkDoc(input, output, errs, log)
			}()

		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProjectDoc:    projectDoc,
		WBSElementDoc: wBSElementDoc,
		NetworkDoc:    networkDoc,
	}

	return data
}

func (c *DPFMAPICaller) ProjectDoc(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProjectDoc {
	where := "WHERE 1 = 1"

	if input.ProjectDoc.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.ProjectDoc.Project)
	}
	if input.ProjectDoc.DocType != nil && len(*input.ProjectDoc.DocType) != 0 {
		where = fmt.Sprintf("%s\nAND DocType = '%v'", where, *input.ProjectDoc.DocType)
	}
	if input.ProjectDoc.DocIssuerBusinessPartner != nil && *input.ProjectDoc.DocIssuerBusinessPartner != 0 {
		where = fmt.Sprintf("%s\nAND DocIssuerBusinessPartner = %v", where, *input.ProjectDoc.DocIssuerBusinessPartner)
	}
	groupBy := "\nGROUP BY Project, DocType, DocIssuerBusinessPartner "

	rows, err := c.db.Query(
		`SELECT
		Project, DocType, MAX(DocVersionID), DocID, FileExtension, FileName, FilePath, DocIssuerBusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_project_doc_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProjectDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) WBSElementDoc(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WBSElementDoc {
	where := "WHERE 1 = 1"

	if input.ProjectDoc.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.ProjectDoc.Project)
	}
	if input.ProjectDoc.WBSElementDoc.WBSElement != nil {
		where = fmt.Sprintf("%s\nAND WBSElement = %d", where, *input.ProjectDoc.WBSElementDoc.WBSElement)
	}
	if input.ProjectDoc.WBSElementDoc.DocType != nil {
		where = fmt.Sprintf("%s\nAND DocType = '%v'", where, *input.HeaderDoc.WBSElementDoc.DocType)
	}
	if input.ProjectDoc.WBSElementDoc.DocIssuerBusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND DocIssuerBusinessPartner = %v", where, *input.HeaderDoc.WBSElementDoc.DocIssuerBusinessPartner)
	}
	groupBy := "\nGROUP BY Project, WBSElement, DocType, DocIssuerBusinessPartner "

	rows, err := c.db.Query(
		`SELECT
		Project, WBSElement, DocType, MAX(DocVersionID), DocID, FileExtension, FileName, FilePath, DocIssuerBusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_wbs_element_doc_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToWBSElementDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) NetworkDoc(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.NetworkDoc {
	where := "WHERE 1 = 1"

	if input.ProjectDoc.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.ProjectDoc.Project)
	}
	if input.ProjectDoc.NetworkDoc.WBSElement != nil {
		where = fmt.Sprintf("%s\nAND WBSElement = %d", where, *input.ProjectDoc.NetworkDoc.WBSElement)
	}
	if input.ProjectDoc.NetworkDoc.Network != nil {
		where = fmt.Sprintf("%s\nAND Network = %d", where, *input.ProjectDoc.NetworkDoc.Network)
	}
	if input.ProjectDoc.NetworkDoc.DocType != nil {
		where = fmt.Sprintf("%s\nAND DocType = '%v'", where, *input.HeaderDoc.NetworkDoc.DocType)
	}
	if input.ProjectDoc.NetworkDoc.DocIssuerBusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND DocIssuerBusinessPartner = %v", where, *input.HeaderDoc.NetworkDoc.DocIssuerBusinessPartner)
	}
	groupBy := "\nGROUP BY Project, WBSElement, network, DocType, DocIssuerBusinessPartner "

	rows, err := c.db.Query(
		`SELECT
		Project, WBSElement, network, DocType, MAX(DocVersionID), DocID, FileExtension, FileName, FilePath, DocIssuerBusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_network_doc_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToNetworkDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
