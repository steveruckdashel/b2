package b2

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type fileUpload_t struct {
	BucketID           string `json:"bucketId"`
	UploadURL          string `json:"uploadUrl"`
	AuthorizationToken string `json:"authorizationToken"`
}

type FileData struct {
	FileID        string            `json:"fileId"`
	FileName      string            `json:"fileName"`
	AccountID     string            `json:"accountId"`
	BucketID      string            `json:"bucketId"`
	ContentLength int               `json:"contentLength"`
	ContentSHA1   string            `json:"contentSha1"`
	ContentType   string            `json:"contentType"`
	FileInfo      map[string]string `json:"fileInfo"`
}

type SimpleFileData struct {
	FileID          string `json:"fileId"`
	FileName        string `json:"fileName"`
	Action          string `json:"action,omitempty"`
	Size            int    `json:"size,omitempty"`
	UploadTimestamp int    `json:"uploadTimestamp,omitempty"`
}

type SimpleFileDataArray struct {
	Files        SimpleFileData `json:"files"`
	NextFileName string         `json:"nextFileName,omitempty"`
	NextFileID   string         `json:"nextFileId,omitempty"`
}

func (b2 *B2) UploadFile(bucket_id, file_name, content_type string, file_reader io.Reader) (*FileData, error) {
	v := url.Values{}
	v.Set("bucketId", bucket_id)
	uri := b2.ApiURL + "/b2api/v1/b2_get_upload_url?" + v.Encode()
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
	dec := json.NewDecoder(resp.Body)

	var fileUpload fileUpload_t
	for dec.More() {
		if err := dec.Decode(&fileUpload); err != nil {
			return nil, err
		}
	}

	b, err := ioutil.ReadAll(file_reader)
	if err != nil {
		return nil, err
	}
	hsh := make(chan string, 1)
	go func() {
		hsh <- fmt.Sprintf("%x", sha1.Sum(b))
	}()

	up_req, err := http.NewRequest("POST", fileUpload.UploadURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	up_req.Header.Set("Authorization", fileUpload.AuthorizationToken)
	up_req.Header.Add("X-Bz-File-Name", file_name)
	up_req.Header.Add("Content-Type", content_type)
	up_req.Header.Add("X-Bz-Content-Sha1", <-hsh)

	up_resp, err := client.Do(up_req)
	if err != nil {
		return nil, err
	}
	defer up_resp.Body.Close()
	up_dec := json.NewDecoder(up_resp.Body)

	var fileData FileData
	for up_dec.More() {
		if err := up_dec.Decode(&fileData); err != nil {
			return nil, err
		}
	}

	return &fileData, nil
}

func (b2 *B2) GetFileInfo(file_id string) (*FileData, error) {
	v := url.Values{}
	v.Set("fileId", file_id)
	uri := b2.ApiURL + "/b2api/v1/b2_get_file_info?" + v.Encode()
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
	dec := json.NewDecoder(resp.Body)

	var fileData FileData
	for dec.More() {
		if err := dec.Decode(&fileData); err != nil {
			return nil, err
		}
	}

	return &fileData, nil
}

func (b2 *B2) ListFileNames(bucket_id, fileName string, maxFileCount int) ([]SimpleFileData, error) {
	uri_base := b2.ApiURL + "/b2api/v1/b2_list_file_names?"
	arr := make([]SimpleFileData, 0)

	for {
		v := url.Values{}
		v.Set("bucketId", bucket_id)
		if len(fileName) > 0 {
			v.Set("startFileName", fileName)
		}
		if maxFileCount > 0 {
			v.Set("maxFileCount", strconv.Itoa(maxFileCount))
		}
		uri := uri_base + v.Encode()
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
		dec := json.NewDecoder(resp.Body)

		var fileData SimpleFileDataArray
		for dec.More() {
			if err := dec.Decode(&fileData); err != nil {
				return arr, err
			}
		}

		arr = append(arr, fileData.Files)
		if len(fileData.NextFileName) <= 0 {
			break
		} else {
			fileName = fileData.NextFileName
		}
	}

	return arr, nil
}

func (b2 *B2) ListFileVersions(bucket_id, fileName, fileId string, maxFileCount int) ([]SimpleFileData, error) {
	uri_base := b2.ApiURL + "/b2api/v1/b2_list_file_versions?"
	arr := make([]SimpleFileData, 0)

	for {
		v := url.Values{}
		v.Set("bucketId", bucket_id)
		if len(fileName) > 0 {
			v.Set("startFileName", fileName)
		}
		if len(fileId) > 0 {
			v.Set("startFileId", fileId)
		}
		if maxFileCount > 0 {
			v.Set("maxFileCount", strconv.Itoa(maxFileCount))
		}
		uri := uri_base + v.Encode()
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
		dec := json.NewDecoder(resp.Body)

		var fileData SimpleFileDataArray
		for dec.More() {
			if err := dec.Decode(&fileData); err != nil {
				return arr, err
			}
		}

		arr = append(arr, fileData.Files)
		if len(fileData.NextFileName) <= 0 || len(fileData.NextFileID) <= 0 {
			break
		} else {
			fileName = fileData.NextFileName
			fileId = fileData.NextFileID
		}
	}

	return arr, nil
}

func (b2 *B2) HideFile(bucket_id, file_name string) (*SimpleFileData, error) {
	v := url.Values{}
	v.Set("bucketId", bucket_id)
	v.Set("fileName", file_name)
	uri := b2.ApiURL + "/b2api/v1/b2_hide_file?" + v.Encode()
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
	dec := json.NewDecoder(resp.Body)

	var fileData SimpleFileData
	for dec.More() {
		if err := dec.Decode(&fileData); err != nil {
			return nil, err
		}
	}

	return &fileData, nil
}

func (b2 *B2) DeleteFileVersion(file_name, file_id string) (*SimpleFileData, error) {
	v := url.Values{}
	v.Set("fileId", file_id)
	v.Set("fileName", file_name)
	uri := b2.ApiURL + "/b2api/v1/b2_delete_file_version?" + v.Encode()
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
	dec := json.NewDecoder(resp.Body)

	var fileData SimpleFileData
	for dec.More() {
		if err := dec.Decode(&fileData); err != nil {
			return nil, err
		}
	}

	return &fileData, nil
}

func (b2 *B2) DownloadFileByID(file_id string) (*FileData, io.Reader, error) {
	v := url.Values{}
	v.Set("fileId", file_id)
	uri := b2.DownloadURL + "/b2api/v1/b2_download_file_by_id?" + v.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	fileData := &FileData{
		FileID:      strings.Join(resp.Header["X-Bz-File-Id"], ","),
		FileName:    strings.Join(resp.Header["X-Bz-File-Name"], ","),
		ContentSHA1: strings.Join(resp.Header["X-Bz-Content-Sha1"], ","),
		ContentType: strings.Join(resp.Header["Content-Type"], ","),
		FileInfo:    make(map[string]string),
	}
	if len(resp.Header["Content-Length"]) > 1 {
		return nil, nil, errors.New("Unexpected Content-Length header value")
	} else if len(resp.Header["content-Length"]) == 1 {
		length, err := strconv.Atoi(resp.Header["Content-Length"][0])
		if err != nil {
			return nil, nil, err
		}
		fileData.ContentLength = length
	}

	info_key := "X-Bz-Info-"
	for key, values := range resp.Header {
		if info_key == key[:len(info_key)] {
			fileData.FileInfo[key] = strings.Join(values, ",")
		}
	}

	return fileData, resp.Body, nil
}

func (b2 *B2) DownloadFileByName(bucket_name, file_name string) (*FileData, io.Reader, error) {
	uri := b2.DownloadURL + fmt.Sprintf("/file/%s/%s", bucket_name, file_name)
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Authorization", b2.AuthorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	fileData := &FileData{
		FileID:      strings.Join(resp.Header["X-Bz-File-Id"], ","),
		FileName:    strings.Join(resp.Header["X-Bz-File-Name"], ","),
		ContentSHA1: strings.Join(resp.Header["X-Bz-Content-Sha1"], ","),
		ContentType: strings.Join(resp.Header["Content-Type"], ","),
		FileInfo:    make(map[string]string),
	}
	if len(resp.Header["Content-Length"]) > 1 {
		return nil, nil, errors.New("Unexpected Content-Length header value")
	} else if len(resp.Header["content-Length"]) == 1 {
		length, err := strconv.Atoi(resp.Header["Content-Length"][0])
		if err != nil {
			return nil, nil, err
		}
		fileData.ContentLength = length
	}

	info_key := "X-Bz-Info-"
	for key, values := range resp.Header {
		if info_key == key[:len(info_key)] {
			fileData.FileInfo[key] = strings.Join(values, ",")
		}
	}

	return fileData, resp.Body, nil
}
