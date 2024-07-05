package internal

import "github.com/av-belyakov/shaper_stix_2.1/datamodels"

// searchCaseId выполняет поиск id кейса
func searchCaseId(tmf datamodels.ChanOutputDecodeJSON) (float64, bool) {
	var (
		cid float64
		ok  bool
	)

	if tmf.FieldBranch == "event.object.caseId" {
		cid, ok = tmf.Value.(float64)
	}

	return cid, ok
}
