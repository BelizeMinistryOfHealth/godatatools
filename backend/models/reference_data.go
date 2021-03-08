package models

import "fmt"

type GoDataLabTest string

const (
	PCR                     GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_PCR"
	PcrRt                   GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_RT_PCR"
	RapidBiosensorAntigen   GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_SARCOV_2_SD_BIOSENSOR_ANTIGEN_RAPID"
	WholeGenomeSequencing   GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_WHOLE_GENOME_SEQUENCING"
	PartialGenomeSequencing GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_PARTIAL_GENOME_SEQUENCING"
	PANBIO                  GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_PANBIO"
	OTHER                   GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_OTHER_SPECIFY_IN_NOTES_FIELD"
	IgcOrIgm                GoDataLabTest = "LNG_REFERENCE_DATA_CATEGORY_TYPE_OF_LAB_TEST_IG_C_OR_IG_M"
)

type GoDataLabTestStatus string

const (
	LabTestCompleted  GoDataLabTestStatus = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_STATUS_COMPLETED"
	LabTestInProgress GoDataLabTestStatus = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_STATUS_IN_PROGRESS"
)

type GoDataLabTestResult string

const (
	LabTestResultInconclusive              GoDataLabTestResult = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_INCONCLUSIVE"
	LabTestResultNegative                  GoDataLabTestResult = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_NEGATIVE"
	LabTestResultPositive                  GoDataLabTestResult = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_POSITIVE"
	LabTestResultPositiveForOtherPathogens GoDataLabTestResult = "LNG_REFERENCE_DATA_CATEGORY_LAB_TEST_RESULT_POSITIVE_FOR_OTHER_PATHOGENS_SPECIFY_IN_NOTES_FIELD"
)

func ParseLabTestType(s string) (string, error) {
	switch s {
	case string(PCR):
		return "PCR", nil
	case string(PcrRt):
		return "PCR RT", nil
	case string(RapidBiosensorAntigen):
		return "RAPID BIOSENSOR ANTIGEN", nil
	case string(WholeGenomeSequencing):
		return "WHOLE GENOME SEQUENCING", nil
	case string(PartialGenomeSequencing):
		return "PARTIAL GENOME SEQUENCING", nil
	case string(PANBIO):
		return "PANBIO", nil
	case string(IgcOrIgm):
		return "IgC/IgM", nil
	case string(OTHER):
		return "OTHER UNSPECIFIED TEST", nil
	default:
		return "", fmt.Errorf("invalid lab test type: %s", s)
	}
}

func ParseLabTestStatus(s string) (string, error) {
	switch s {
	case string(LabTestCompleted):
		return "COMPLETED", nil
	case string(LabTestInProgress):
		return "IN PROGRESS", nil
	default:
		return "", fmt.Errorf("invalid lab tests status: %s", s)
	}
}

func ParseLabTestResult(s string) (string, error) {
	switch s {
	case string(LabTestResultInconclusive):
		return "INCONCLUSIVE", nil
	case string(LabTestResultNegative):
		return "NEGATIVE", nil
	case string(LabTestResultPositive):
		return "POSITIVE", nil
	case string(LabTestResultPositiveForOtherPathogens):
		return "POSITIVE_FOR_OTHER_PATHOGENS", nil
	default:
		return "", fmt.Errorf("invalid lab test result: %s", s)
	}
}
