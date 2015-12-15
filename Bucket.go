package b2

import (
	"encoding/json"
	"errors"
	"net/url"
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

	resp, err := b2.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var bucket Bucket
	for dec.More() {
		if err := dec.Decode(&bucket); err != nil {
			return nil, err
		}
	}

	return &bucket, nil
}

func (b2 *B2) DeleteBucket(bucketId string) (*Bucket, error) {
	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	v.Set("bucketId", bucketId)
	uri := b2.ApiURL + "/b2api/v1/b2_delete_bucket?" + v.Encode()

	resp, err := b2.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var bucket Bucket
	for dec.More() {
		if err := dec.Decode(&bucket); err != nil {
			return nil, err
		}
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
	v.Set("bucketId", bucketId)
	v.Set("bucketType", bucketType)
	uri := b2.ApiURL + "/b2api/v1/b2_update_bucket?" + v.Encode()

	resp, err := b2.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var bucket Bucket
	for dec.More() {
		if err := dec.Decode(&bucket); err != nil {
			return nil, err
		}
	}

	return &bucket, nil
}

func (b2 *B2) ListBuckets() (*Buckets, error) {
	v := url.Values{}
	v.Set("accountId", b2.AccountID)
	uri := b2.ApiURL + "/b2api/v1/b2_list_buckets?" + v.Encode()

	resp, err := b2.client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var buckets Buckets
	for dec.More() {
		if err := dec.Decode(&buckets); err != nil {
			return nil, err
		}
	}

	return &buckets, nil
}
