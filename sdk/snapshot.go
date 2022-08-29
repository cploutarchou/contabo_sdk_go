package sdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type Snapshots struct {
	Snapshot Snapshot `json:"snapshots"`
}
type Snapshot struct {
	Pagination struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Page          int `json:"page"`
	} `json:"_pagination"`
	Data []struct {
		TenantId       string    `json:"tenantId"`
		CustomerId     string    `json:"customerId"`
		SnapshotId     string    `json:"snapshotId"`
		Name           string    `json:"name"`
		Description    string    `json:"description"`
		InstanceId     int       `json:"instanceId"`
		CreatedDate    time.Time `json:"createdDate"`
		AutoDeleteDate time.Time `json:"autoDeleteDate"`
		ImageId        string    `json:"imageId"`
		ImageName      string    `json:"imageName"`
	} `json:"data"`
	Links struct {
		First    string `json:"first"`
		Next     string `json:"next"`
		Self     string `json:"self"`
		Previous string `json:"previous"`
		Last     string `json:"last"`
	} `json:"_links"`
}
type CreateSnapShopRes struct {
	Data []struct {
		TenantId   string `json:"tenantId"`
		CustomerId string `json:"customerId"`
		SnapshotId string `json:"snapshotId"`
		ImageId    string `json:"imageId"`
		ImageName  string `json:"imageName"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

func (s *Snapshots) GetInstanceSnapshots(instanceId int) (*Snapshot, error) {

	url := fmt.Sprintf("%s/%d/snapshots", ComputeInstancesUrl, instanceId)
	res, _ := Do(GET, URL(url), nil)
	var snapshots Snapshot
	err := json.Unmarshal(res, &snapshots)
	if err != nil {
		return nil, err
	}
	return &snapshots, nil

}
func (s *Snapshots) GetSnapshot(instanceId int, snapshotId string) (*Snapshot, error) {

	url := fmt.Sprintf("%s/%d/snapshots/%s", ComputeInstancesUrl, instanceId, snapshotId)
	res, _ := Do(GET, URL(url), nil)
	var snapshots Snapshot
	err := json.Unmarshal(res, &snapshots)
	if err != nil {
		return nil, err
	}
	return &snapshots, nil

}

func (s *Snapshots) CreateSnapshot(instanceId int, name, description string) (*CreateSnapShopRes, error) {
	type Payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	data := Payload{
		Name:        name,
		Description: description,
	}
	url := fmt.Sprintf("%s/%d/snapshots", ComputeInstancesUrl, instanceId)
	fmt.Println(url)
	res, _ := Do(POST, URL(url), data)
	var snapshots CreateSnapShopRes
	err := json.Unmarshal(res, &snapshots)
	if err != nil {
		return nil, err
	}
	return &snapshots, nil
}
