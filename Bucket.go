package b2

import (
	"encoding/json"
	"errors"
	"net/url"
	"io/ioutil"
	"net/http"
)

var bucketTypes []string = []string{"allPrivate", "allPublic"}

type Bucket struct {
	BucketID   string `json:"bucketId"`
	AccountID  string `json:"accountId"`
	BucketName string `json:"bucketName"`
	BucketType string `json:"bucketType"`
}

type Buckets struct {
	Buckets []Bucket `json:"buckets"`
}

func (b2 *B2) CreateBucket(bucketName, bucketType string) (*Bucket, error) {
	validBucketType := false
	for _, v := range bucketTypes {
		validBucketType = validBucketType || (v == bucketType)
	}
	if !validBucketType {
		return nil, errors.New("invalid bucket type")
	}

	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	v.Set("bucketName", bucketName)
	v.Set("bucketType", bucketType)
	uri := b2.ApiURL + "/b2api/v1/b2_create_bucket?" + v.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errMsg ErrorJSON
		if err := json.Unmarshal(data, &errMsg); err != nil {
			return nil, err
		}
		return nil, backblaze_error(errMsg)
	}

	var bucket Bucket
	if err := json.Unmarshal(data, &bucket); err != nil {
		return nil, err
	}
	
	return &bucket, nil
}

func (b2 *B2) DeleteBucket(bucketId string) (*Bucket, error) {
	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	v.Set("bucketId", bucketId)
	uri := b2.ApiURL + "/b2api/v1/b2_delete_bucket?" + v.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errMsg ErrorJSON
		if err := json.Unmarshal(data, &errMsg); err != nil {
			return nil, err
		}
		return nil, backblaze_error(errMsg)
	}

	var bucket Bucket
	if err := json.Unmarshal(data, &bucket); err != nil {
		return nil, err
	}

	return &bucket, nil
}

func (b2 *B2) UpdateBucket(bucketId, bucketType string) (*Bucket, error) {
	validBucketType := false
	for _, v := range bucketTypes {
		validBucketType = validBucketType || (v == bucketType)
	}
	if !validBucketType {
		return nil, errors.New("invalid bucket type")
	}

	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	v.Set("bucketId", bucketId)
	v.Set("bucketType", bucketType)
	uri := b2.ApiURL + "/b2api/v1/b2_update_bucket?" + v.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errMsg ErrorJSON
		if err := json.Unmarshal(data, &errMsg); err != nil {
			return nil, err
		}
		return nil, backblaze_error(errMsg)
	}

	var bucket Bucket
	if err := json.Unmarshal(data, &bucket); err != nil {
		return nil, err
	}

	return &bucket, nil
}

func (b2 *B2) ListBuckets() (*Buckets, error) {
	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	uri := b2.ApiURL + "/b2api/v1/b2_list_buckets?" + v.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errMsg ErrorJSON
		if err := json.Unmarshal(data, &errMsg); err != nil {
			return nil, err
		}
		return nil, backblaze_error(errMsg)
	}

	var buckets Buckets
	if err := json.Unmarshal(data, &buckets); err != nil {
		return nil, err
	}

	return &buckets, nil
}
