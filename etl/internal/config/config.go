package config

import "os"

const BatchSize = 1000

const (
	STG = "STG"
	ODS = "ODS"
	DDS = "DDS"
	DM  = "DM"
)

func GetStage() string {
	var res string
	switch os.Getenv("ETL_STAGE") {
	case STG:
		res = STG
	case ODS:
		res = ODS
	case DDS:
		res = DDS
	case DM:
		res = DM
	default:
		res = "unknown stage"
	}
	return res
}
