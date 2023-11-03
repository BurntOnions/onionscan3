package deanonymization

import (
	"regexp"
	"strings"

	"github.com/burntonions/onionscan3/onionscan/config"
	"github.com/burntonions/onionscan3/onionscan/report"
)

// EmailScan extracts anything which resembles an email address from the current crawl.
func EmailScan(osreport *report.OnionScanReport, anonreport *report.AnonymityReport, osc *config.OnionScanConfig) {

	mailRegexp := regexp.MustCompile(`(\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,56})`)

	for _, id := range osreport.Crawls {
		crawlRecord, _ := osc.Database.GetCrawlRecord(id)
		if strings.Contains(crawlRecord.Page.Headers.Get("Content-Type"), "text/html") {
			foundEmail := mailRegexp.FindAllString(crawlRecord.Page.Snapshot, -1)
			for _, email := range foundEmail {
				anonreport.EmailAddresses = append(anonreport.EmailAddresses, email)
				osc.Database.InsertRelationship(osreport.HiddenService, "snapshot", "email-address", email)
			}
		}
	}
}
