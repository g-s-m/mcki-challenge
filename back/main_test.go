package main

import (
	"back/data"
	"testing"
)

func TestChangeReport(t *testing.T) {
	reportApi := data.CreateReportStub()
	reports := reportApi.GetReports()

	reportApi.BlockReport(reports.Elements[0].Id)
	newReports := reportApi.GetReports()
	if newReports.Elements[0].State != "BLOCKED" {
		t.Error("BlockReport failed, actual: ", newReports.Elements[0].State, "expected: BLOCKED")
	}
	reportApi.ResolveReport(newReports.Elements[0].Id)
	newReports2 := reportApi.GetReports()
	if len(newReports2.Elements) != len(newReports.Elements) - 1 {
		t.Error("ResolvedReport failed")
	}


}

