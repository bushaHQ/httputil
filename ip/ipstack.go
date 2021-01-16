// Package ipstack provides info on IP address location
// using the http://api.ipstack.com service.
package ip

import (
	"encoding/json"
	"fmt"
	"github.com/bushaHQ/httputil/request"
	"github.com/mitchellh/mapstructure"
	"net/url"
)

var ipstackURI = "http://api.ipstack.com"

var (
	IpKey = "IP_KEY"
)

// IPInfo wraps json response
type IPInfo struct {
	IP            string  `mapstructure:"ip,omitempty"`
	Type          string  `mapstructure:"type,omitempty"`
	ContinentCode string  `mapstructure:"continent_code,omitempty"`
	ContinentName string  `mapstructure:"continent_name,omitempty"`
	CountryCode   string  `mapstructure:"country_code,omitempty"`
	CountryName   string  `mapstructure:"country_name,omitempty"`
	RegionCode    string  `mapstructure:"region_code,omitempty"`
	RegionName    string  `mapstructure:"region_name,omitempty"`
	City          string  `mapstructure:"city,omitempty"`
	Zip           string  `mapstructure:"zip,omitempty"`
	Latitude      float64 `mapstructure:"latitude,omitempty"`
	Longitude     float64 `mapstructure:"longitude,omitempty"`
	Location      struct {
		GeonameID float64 `mapstructure:"geoname_id,omitempty"`
		Capital   string  `mapstructure:"capital,omitempty"`
		Languages []struct {
			Code   string `mapstructure:"code,omitempty"`
			Name   string `mapstructure:"name,omitempty"`
			Native string `mapstructure:"native,omitempty"`
		} `mapstructure:"languages,omitempty"`
		CountryFlag             string `mapstructure:"country_flag,omitempty"`
		CountryFlagEmoji        string `mapstructure:"country_flag_emoji,omitempty"`
		CountryFlagEmojiUnicode string `mapstructure:"country_flag_emoji_unicode,omitempty"`
		CallingCode             string `mapstructure:"calling_code,omitempty"`
		IsEu                    bool   `mapstructure:"is_eu,omitempty"`
	} `mapstructure:"location,omitempty"`
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
		return nil, err
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
		return nil, err
	}
	//log.Println(string(res.Body))
	m := map[string]interface{}{}

	var ipInfo IPInfo
	err = json.Unmarshal(res.Body, &m)
	if err != nil {
		return nil, err
	}

	if s, ok := m["success"]; ok {
		v, ok := s.(bool)
		if !ok || !v {
			return nil, fmt.Errorf("%s", m["error"].(map[string]interface{})["info"])
		}

	}

	err = mapstructure.Decode(m, &ipInfo)

	return &ipInfo, err
}
