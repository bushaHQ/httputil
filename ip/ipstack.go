// Package ipstack provides info on IP address location
// using the http://api.ipstack.com service.
package ip

import (
	"encoding/json"
	"fmt"
	"github.com/bushaHQ/httputil/request"
	"net/url"
)

var ipstackURI = "http://api.ipstack.com"

var (
	IpKey = "IP_KEY"
)

// IPInfo wraps json response
type IPInfo struct {
	IP            string  `json:"ip,omitempty"`
	Type          string  `json:"type,omitempty"`
	ContinentCode string  `json:"continent_code,omitempty"`
	ContinentName string  `json:"continent_name,omitempty"`
	CountryCode   string  `json:"country_code,omitempty"`
	CountryName   string  `json:"country_name,omitempty"`
	RegionCode    string  `json:"region_code,omitempty"`
	RegionName    string  `json:"region_name,omitempty"`
	City          string  `json:"city,omitempty"`
	Zip           string  `json:"zip,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Location      struct {
		GeonameID float64 `json:"geoname_id,omitempty"`
		Capital   string  `json:"capital,omitempty"`
		Languages []struct {
			Code   string `json:"code,omitempty"`
			Name   string `json:"name,omitempty"`
			Native string `json:"native,omitempty"`
		} `json:"languages,omitempty"`
		CountryFlag             string `json:"country_flag,omitempty"`
		CountryFlagEmoji        string `json:"country_flag_emoji,omitempty"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode,omitempty"`
		CallingCode             string `json:"calling_code,omitempty"`
		IsEu                    bool   `json:"is_eu,omitempty"`
	} `json:"location,omitempty"`
}

//sets the ipstack key for
func Init(ipKey string) {
	IpKey = ipKey
}

// MyIP provides information about the public IP address of the client.
func MyIP() (*IPInfo, error) {
	return getInfo(fmt.Sprintf("%s/json", ipstackURI))
}

// ForeignIP provides information about the given IP address (IPv4 or IPv6)
func ForeignIP(ip string) (*IPInfo, error) {

	if ip == "" {
		return nil, fmt.Errorf("empty ip address")
	}

	return getInfo(fmt.Sprintf("/%s?access_key=%s", ip, IpKey))
}

// Undercover code that makes the real call to the webservice
func getInfo(path string) (*IPInfo, error) {
	u, err := url.Parse(ipstackURI)
	if err != nil {
		return &IPInfo{}, err
	}
	req := request.Req{
		BaseURL: u,
		Path:    path,
		Body:    nil,
		Header:  nil,
		Method:  "GET",
	}

	res, err := request.Do(req)
	if err != nil {
		return &IPInfo{}, err
	}

	var ipInfo IPInfo
	err = json.Unmarshal(res.Body, &ipInfo)
	if err != nil {
		return nil, err
	}

	return &ipInfo, nil
}
