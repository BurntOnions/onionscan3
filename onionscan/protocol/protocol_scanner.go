package protocol

import (
	"github.com/burntonions/onionscan3/onionscan/config"
	"github.com/burntonions/onionscan3/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
