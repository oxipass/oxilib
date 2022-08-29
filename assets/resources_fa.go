package assets

import (
	"embed"
	"encoding/json"
	"strings"
)

var (
	//go:embed fa/brands/*.svg
	iconsBrandsFS embed.FS
	//go:embed fa/regular/*.svg
	iconsRegularFS embed.FS
	//go:embed fa/solid/*.svg
	iconsSolidFS embed.FS
	//go:embed fa/search.json
	searchResBin []byte

	searchRes map[string]string
)

const cFAprefix = "fa/"
const cRegularPrefix = "regular/"
const cSolidPrefix = "solid/"
const cBrandsPrefix = "brands/"

const cSVGExt = ".svg"
const cDefaultIcon = "solid/file"

func getSVGImageFromFilePath(filePath string) (svgImage string, err error) {
	var svgBytes []byte

	if strings.HasPrefix(filePath, cRegularPrefix) {
		svgBytes, err = iconsRegularFS.ReadFile(cFAprefix + filePath + cSVGExt)
	} else if strings.HasPrefix(filePath, cSolidPrefix) {
		svgBytes, err = iconsSolidFS.ReadFile(cFAprefix + filePath + cSVGExt)
	} else if strings.HasPrefix(filePath, cBrandsPrefix) {
		svgBytes, err = iconsBrandsFS.ReadFile(cFAprefix + filePath + cSVGExt)
	} else {
		svgBytes, err = iconsSolidFS.ReadFile(cFAprefix + cDefaultIcon + cSVGExt)
	}

	if err != nil {
		svgBytes, err = iconsSolidFS.ReadFile(cFAprefix + cDefaultIcon + cSVGExt)
		if err != nil {
			return "", err
		}
	}

	return string(svgBytes), nil
}

func CheckIfExistsInFontAwesome(faCheckValue string) bool {
	var err error
	if strings.HasPrefix(faCheckValue, cRegularPrefix) {
		_, err = iconsRegularFS.ReadFile(cFAprefix + faCheckValue + cSVGExt)
	} else if strings.HasPrefix(faCheckValue, cSolidPrefix) {
		_, err = iconsSolidFS.ReadFile(cFAprefix + faCheckValue + cSVGExt)
	} else if strings.HasPrefix(faCheckValue, cBrandsPrefix) {
		_, err = iconsBrandsFS.ReadFile(cFAprefix + faCheckValue + cSVGExt)
	} else {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func SearchFontAwesomeList(term string) (faValues []string) {
	if searchRes == nil {
		searchRes = make(map[string]string)
		err := json.Unmarshal(searchResBin, &searchRes)
		if err != nil {
			return nil
		}
	}
	lowerTerm := strings.ToLower(term)
	for faKey, faSearchTerms := range searchRes {
		if strings.Contains(faKey, lowerTerm) || strings.Contains(faSearchTerms, lowerTerm) {
			faValues = append(faValues, faKey)
		}
	}
	return faValues
}
