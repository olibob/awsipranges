package awsipranges

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// AwsIPRangesJSON All IP ranges for amazon services
	AwsIPRangesJSON = "https://ip-ranges.amazonaws.com/ip-ranges.json"
)

// AwsIPPrefix blaÒ
type AwsIPPrefix struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

// AwsIPs bla
type AwsIPs struct {
	SyncToken  string        `json:"syncToken"`
	CreateDate string        `json:"createDate"`
	Prefixes   []AwsIPPrefix `json:"prefixes"`
}

// AwsIPRanges Get a JSON file from AWS containing
// all IP ranges for amazon services
func AwsIPRanges() (AwsIPs, error) {
	ips := AwsIPs{}

	resp, err := http.Get(AwsIPRangesJSON)
	if err != nil {
		return ips, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ips, err
	}

	err = json.Unmarshal(body, &ips)
	if err != nil {
		return ips, err
	}

	return ips, err
}
