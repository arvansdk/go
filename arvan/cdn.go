package arvan

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
Domain Info
*/
type DomainInfo struct {
	ID			string		`json:"id"`
	UserID		string		`json:"user_id"`
	Name		string		`json:"name"`
	PlanLevel	int			`json:"plan_level"`
	NSKeys		[2]string	`json:"ns_keys"`
	CurrentNS	[2]string	`json:"current_ns"`
	TargetCName	string		`json:"target_cname"`
	CustomCName	string		`json:"custom_cname"`
	Type		string		`json:"type"`
	Status		string		`json:"status"`
	DNSCloud	bool		`json:"dns_cloud"`
	Restriction	[]string	`json:"restriction"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type DomainInfoResponse struct {
	Data	DomainInfo	`json:"data"`
	Message	string		`json:"message"`
}

func (client *APIClient) GetDomainInfo(domain string) DomainInfo {
	apiResponse := &DomainInfoResponse{}
	res := client.CurlGet(domain, map[string]string{})
	json.Unmarshal(res, &apiResponse)
	return apiResponse.Data
}
/**/

/* DomainsList */
type DomainListResponse struct {
	Data	[]DomainInfo
	Message string
}
/**/

/* Create Domain */
type CreateDomainPayload struct {
	Domain	string	`json:"domain"`
}

func (client *APIClient) CreateDomain(domain string) []byte {
	return client.CurlPost(
		"domains/dns-service",
		CreateDomainPayload{
			Domain: domain,
		},
	)
}
/**/

/* Add DNS Record - Not Completed */
type AddDNSRecordPayload struct {
	// ID				string	`json:"id"`
	Type			string				`json:"type"`
	Name			string				`json:"name"`
	Cloud			bool				`json:"cloud"`
	UpStreamHTTPS	string				`json:"upstream_https"`
	Values			[]RecordValue		`json:"value"`
	IPFilterMode	RecordFilterMode	`json:"ip_filter_mode"`
	TTL				int					`json:"ttl"`
	// IsProtected		bool				`json:"is_protected"`
}

type RecordFilterMode struct {
	Count		string	`json:"count"`
	GeoFilter	string	`json:"geo_filter"`
	Order		string	`json:"order"`
}

type RecordValue struct {
	Country	string	`json:"country"`
	IP		string	`json:"ip"`
	Port	any		`json:"port"`
	Weight	any		`json:"weight"`
}

func (client *APIClient) AddDNSRecord(domain, name, ip string) []byte {
	return client.CurlPost(
		fmt.Sprintf("%s/%s/%s", "domains", domain, "dns-records"),
		AddDNSRecordPayload{
			Type: "A",
			Name: name,
			Cloud: false,
			Values: []RecordValue{{
				Country: "",
				IP: ip,
				Port: nil,
				Weight: nil,
			}},
			TTL: 120,
		},
	)
}
/**/

/* HTTPS Settings */
type SSLConfigPayload struct {
	Certificate		string	`json:"certificate"`
	TLSVersion		string	`json:"tls_version"`
	HTTPSRedirect	bool	`json:"https_redirect"`
	SSLStatus		bool	`json:"ssl_status"`
}

func (client *APIClient) UpdateSSLConfig(domain string) {
	client.CurlPatch(
		fmt.Sprintf("%s/%s/%s", "domains", domain, "ssl"),
		SSLConfigPayload{
			Certificate: "managed",
			TLSVersion: "TLSv1.2",
			HTTPSRedirect: true,
			SSLStatus: true,
		},
	)
}