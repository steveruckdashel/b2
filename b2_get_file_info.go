package b2

/*
b2_get_file_info
Gets information about one file stored in B2.

Request
Request HTTP Headers

Authorization
required

The account authorization token returned by b2_authorize_account.

Request HTTP Message Body Parameters

fileId
required

The ID of the file, as returned by b2_upload_file, b2_list_file_names, or b2_list_file_versions.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and no file info will be returned. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
fileId
The unique identifier for this version of this file. Used with b2_get_file_info, b2_download_file_by_id, and b2_delete_file_version.

fileName
The name of this file, which can be used with b2_download_file_by_name.

accountId
Your account ID.

contentSha1
The Sha1 hash of the contents of the file stored in B2.

bucketId
The bucket that the file is in.

contentLength
The number of bytes stored in the file.

contentType
The MIME type of the file.

fileInfo
The custom information that was uploaded with the file. This is a JSON object, holding the name/value pairs that were uploaded with the file.

Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

api_url = "" # Provided by b2_authorize_account
file_id = "" # The ID of the file you want to get info on
account_authorization_token = "" # Provided by b2_authorize_account
request = urllib2.Request(
    '%s/b2api/v1/b2_get_file_info' % api_url,
    json.dumps({ 'fileId' : file_id }),
    headers = { 'Authorization': account_authorization_token }
    )
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
    "accountId": "7eecc42b9675",
    "bucketId": "e73ede9c9c8412db49f60715",
    "contentLength": 122573,
    "contentSha1": "a01a21253a07fb08a354acd30f3a6f32abb76821",
    "contentType": "image/jpeg",
    "fileId": "4_ze73ede9c9c8412db49f60715_f100b4e93fbae6252_d20150824_m224353_c900_v8881000_t0001",
    "fileInfo": {},
    "fileName": "akitty.jpg"
}
*/
