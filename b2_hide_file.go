package b2

//b2_hide_file
//Hides a file so that downloading by name will not find the file, but previous versions of the file are still stored. See File Versions about what it means to hide a file.
//
//Request
//Request HTTP Headers
//	Authorization (required)		An authorization token, obtained from b2_authorize_account.
//
//Request HTTP Message Body Parameters
//	bucketId (required)				The bucket containing the file to hide.
//	fileName (required)				The name of the file to hide.

/*
Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file will not be hidden. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
fileId
The unique identifier for this version of this file. Used with b2_get_file_info, b2_download_file_by_id, and b2_delete_file_version.

fileName
The name of this file, which can be used with b2_download_file_by_name.

action
Either "upload" or "hide". "upload" means a file that was uploaded to B2 Cloud Storage. "hide" means a file version marking the file as hidden, so that it will not show up in b2_list_file_names.

The result of b2_list_file_names will contain only "upload". The result of b2_list_file_versions may have both.

size
The number of bytes in the file.

Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

api_url = "" # Provided by b2_authorize_account
account_authorization_token = "" # Provided by b2_authorize_account
file_name = "" # The file name of the file you want to hide.
bucket_id = "" # The ID of bucket where the file you want to hide resides.
request = urllib2.Request(
	'%s/b2api/v1/b2_hide_file' % api_url,
	json.dumps({ 'bucketId' : bucket_id , 'fileName' : file_name }),
	headers = { 'Authorization': account_authorization_token }
	)
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
    "action" : "hide",
    "fileId" : "4_h4a48fe8875c6214145260818_f000000000000472a_d20140104_m032022_c001_v0000123_t0104",
    "fileName" : "typing_test.txt",
    "uploadTimestamp" : 1437815673000
}
*/
