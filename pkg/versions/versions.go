package versions

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const BaseUrl = "https://developers.redhat.com/content-gateway/rest/mirror/pub/openshift-v4/clients/crc/"

type ReleaseInfoSpec struct {
	CrcVersion       string
	GitSHA           string
	OpenShiftVersion string
}

type ReleaseInfo struct {
	Version ReleaseInfoSpec
}

type VersionNotes struct {
	MissingReleaseInfo bool
}

//go:embed latestVersionInfo.json
var latestVersionInfo string

var KnownVersionInfo map[string]ReleaseInfo

var knownVersions = map[string]VersionNotes{
	"1.0.0":  {},
	"1.1.0":  {},
	"1.10.0": {},
	"1.11.0": {},
	"1.12.0": {},
	"1.13.0": {},
	"1.14.0": {},
	"1.15.0": {},
	"1.16.0": {},
	"1.17.0": {},
	"1.18.0": {},
	"1.19.0": {},
	"1.2.0":  {},
	"1.20.0": {},
	"1.21.0": {},
	"1.22.0": {},
	"1.23.1": {},
	"1.24.0": {MissingReleaseInfo: true},
	"1.25.0": {},
	"1.26.0": {},
	"1.27.0": {},
	"1.28.0": {},
	"1.29.1": {},
	"1.3.0":  {},
	"1.30.1": {},
	"1.31.2": {},
	"1.32.1": {},
	"1.33.1": {},
	"1.34.0": {},
	"1.35.0": {},
	"1.36.0": {},
	"1.37.0": {},
	"1.4.0":  {},
	"1.5.0":  {},
	"1.6.0":  {},
	"1.7.0":  {},
	"1.8.0":  {},
	"1.9.0":  {},
	"latest": {},
}

func init() {
	json.Unmarshal([]byte(latestVersionInfo), &KnownVersionInfo)
}

func GetAllVersionInfo() (versions map[string]ReleaseInfo, err error) {
	versions = make(map[string]ReleaseInfo, len(versions))

	for version, notes := range knownVersions {
		ignoreErrors := false
		if notes.MissingReleaseInfo {
			fmt.Fprintf(os.Stderr, "WARNING: version %v is not known to have release info. Trying optimistically but ignoring errors.\n", version)
			ignoreErrors = true
		}
		var info ReleaseInfo
		info, err = GetInfoForVersion(version)

		if err != nil && !ignoreErrors {
			return
		} else if ignoreErrors {
			continue
		}

		versions[version] = info
	}

	return versions, nil
}

func GetInfoForVersion(version string) (info ReleaseInfo, err error) {
	resp, err := http.Get(BaseUrl + version + "/release-info.json")

	if err != nil {
		err = fmt.Errorf("HTTP error: %w", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("HTTP response was %v when getting info for version %v", resp.StatusCode, version)
		return
	}

	dec := json.NewDecoder(resp.Body)

	if err = dec.Decode(&info); err != nil {
		err = fmt.Errorf("json error: %w", err)
		return
	}

	return
}
