package data

import (
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
)

type Report struct {
	Id               string `json:"id"`
	Source           string `json:"source"`
	SourceIdentityId string `json:"sourceIdentityIdA"`
	Reference        struct {
		ReferenceId   string `json:"referenceId"`
		ReferenceType string `json:"referenceType"`
	} `json:"reference"`
	State            string `json:"state"`
	Payload          struct {
		Source                string `json:"source"`
		ReportType            string `json:"reportType"`
		Message               string `json:"message"`
		ReportId              string `json:"reportId"`
		ReferenceResourceId   string `json:"referenceResourceId"`
		ReferenceResourceType string `json:"referenceResourceType"`
	} `json:"payload"`
	Created time.Time `json:"created"`
}

type ReportList struct {
	Size     int `json:"size"`
	Offset   string `json:"nextOffset"`
	Elements []Report `json:"elements"`
}

type ReportApi interface{
	GetReports() ReportList
	ResolveReport(reportId string)
	BlockReport(reportId string)
}

type StubReport struct {
	reports ReportList
}

func CreateReportStub() ReportApi {
	dat, err := ioutil.ReadFile("reports.json")
	if err != nil {
		log.Fatal("Can't read reports: %s", err)
	}
	reports := ReportList{}
	json.Unmarshal(dat, &reports)
	return &StubReport {
		reports: reports,
	}
}

func (p* StubReport) setReportState(reportId string, state string) {
	for i, report := range p.reports.Elements {
		if report.Id == reportId {
			log.Printf("Found report %s", reportId)
			p.reports.Elements[i].State = state
		}
	}
}

func (p* StubReport) GetReports() ReportList {
	rl := ReportList{}
	rl.Size = p.reports.Size
	rl.Offset = p.reports.Offset
	for _, report := range p.reports.Elements {
		if report.State != "RESOLVED" {
			rl.Elements = append(rl.Elements, report)
		}
	}
	return rl
}

func (p* StubReport) ResolveReport(reportId string) {
	p.setReportState(reportId, "RESOLVED")
}

func (p* StubReport) BlockReport(reportId string) {
	p.setReportState(reportId, "BLOCKED")
}

