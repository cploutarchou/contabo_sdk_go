package sdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type Instances struct {
	ListInstances ListInstances
}

type ListInstances struct {
	Pagination struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Page          int `json:"page"`
	} `json:"_pagination"`
	Data []struct {
		TenantId      string `json:"tenantId"`
		CustomerId    string `json:"customerId"`
		AdditionalIps []struct {
			V4 struct {
				Ip          string `json:"ip"`
				NetmaskCidr int    `json:"netmaskCidr"`
				Gateway     string `json:"gateway"`
			} `json:"v4"`
		} `json:"additionalIps"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		InstanceId  int    `json:"instanceId"`
		Region      string `json:"region"`
		ProductId   string `json:"productId"`
		ImageId     string `json:"imageId"`
		IpConfig    struct {
			V4 struct {
				Ip          string `json:"ip"`
				NetmaskCidr int    `json:"netmaskCidr"`
				Gateway     string `json:"gateway"`
			} `json:"v4"`
			V6 struct {
				Ip          string `json:"ip"`
				NetmaskCidr int    `json:"netmaskCidr"`
				Gateway     string `json:"gateway"`
			} `json:"v6"`
		} `json:"ipConfig"`
		MacAddress  string    `json:"macAddress"`
		RamMb       int       `json:"ramMb"`
		CpuCores    int       `json:"cpuCores"`
		OsType      string    `json:"osType"`
		DiskMb      int       `json:"diskMb"`
		SshKeys     []int     `json:"sshKeys"`
		CreatedDate time.Time `json:"createdDate"`
		CancelDate  string    `json:"cancelDate"`
		Status      string    `json:"status"`
		VHostId     int       `json:"vHostId"`
		AddOns      []struct {
			Id       int `json:"id"`
			Quantity int `json:"quantity"`
		} `json:"addOns"`
		ErrorMessage string `json:"errorMessage"`
		ProductType  string `json:"productType"`
		DefaultUser  string `json:"defaultUser"`
	} `json:"data"`
	Links struct {
		First    string `json:"first"`
		Previous string `json:"previous"`
		Self     string `json:"self"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"_links"`
}

// Get returns the list of available resources
func (l *ListInstances) Get(page, size *int) *ListInstances {
	url := getPage(page, size)
	res, _ := Do(GET, URL(url), nil)
	var re ListInstances
	err := json.Unmarshal(res, &re)
	if err != nil {
		return nil
	}
	return &re
}

func getPage(page *int, size *int) string {
	var url string
	if page != nil && size != nil {
		url = fmt.Sprintf("%s?page=%d&size=%d", string(ComputeInstancesUrl), page, size)
	} else {
		url = string(ComputeInstancesUrl) + "?page=1&size=10"
	}
	return url
}
