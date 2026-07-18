package namecheap

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	domainsSRVGetRecords = "namecheap.domains.dns.getsrvrecords"
	domainsSRVSetRecords = "namecheap.domains.dns.setsrvrecords"
)

type DomainSRVGetRecordsResult struct {
	Result []DomainSRVRecord `xml:"Result"`
}

type DomainSRVRecord struct {
	Service  string `xml:"Service"`
	Protocol string `xml:"Protocol"`
	Priority int    `xml:"Priority"`
	Port     int    `xml:"Port"`
	Target   string `xml:"Target"`
	Weight   int    `xml:"Weight"`
}

type DomainSRVSetRecordsResult struct {
	Inserted int `xml:"Result>Inserted"`
	Deleted  int `xml:"Result>Deleted"`
	Updated  int `xml:"Result>Updated"`
}

func (client *Client) DomainsSRVGetRecords(sld, tld string) (*DomainSRVGetRecordsResult, error) {
	requestInfo := &ApiRequest{
		command: domainsSRVGetRecords,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}

	return resp.DomainSRVRecords, nil
}

func (client *Client) DomainSRVSetRecords(
	sld, tld string, records []DomainSRVRecord,
) (*DomainSRVSetRecordsResult, error) {
	requestInfo := &ApiRequest{
		command: domainsSRVSetRecords,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)
	requestInfo.params.Set("SrvCount", strconv.Itoa(len(records)))

	for i, r := range records {
		requestInfo.params.Set(fmt.Sprintf("Service%v", i+1), r.Service)
		requestInfo.params.Set(fmt.Sprintf("Protocol%v", i+1), r.Protocol)
		requestInfo.params.Set(fmt.Sprintf("Priority%v", i+1), strconv.Itoa(r.Priority))
		requestInfo.params.Set(fmt.Sprintf("Port%v", i+1), strconv.Itoa(r.Port))
		requestInfo.params.Set(fmt.Sprintf("Target%v", i+1), r.Target)
		requestInfo.params.Set(fmt.Sprintf("Weight%v", i+1), strconv.Itoa(r.Weight))
	}

	resp, err := client.doSRV(requestInfo)
	if err != nil {
		return nil, err
	}
	return resp.DomainSRVSetRecords, nil
}
