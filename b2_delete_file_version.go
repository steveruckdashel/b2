package b2

/*
b2_delete_file_version
Deletes one version of a file from B2.

If the version you delete is the latest version, and there are older versions, then the most recent older version will become the current version, and be the one that you'll get when downloading by name. See the File Versions page for more details.

Request
Request HTTP Headers

Authorization
required

The account authorization token returned by b2_authorize_account.

Request HTTP Message Body Parameters
fileName
required

The name of the file.

fileId
required

The ID of the file, as returned by b2_upload_file, b2_list_file_names, or b2_list_file_versions.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file version will not be deleted. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
fileId
The unique ID of the file version that was deleted.

fileName
The name of the file.

Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

api_url = "" # Provided by b2_authorize_account
account_authorization_token = "" # Provided by b2_authorize_account
file_name = "" # The name of the file you want to delete
file_id = "" # The fileId of the file you want to delete

request = urllib2.Request(
    '%s/b2api/v1/b2_delete_file_version' % api_url,
    json.dumps({ 'fileName' : file_name, 'fileId' : file_id }),
    headers = { 'Authorization': account_authorization_token }
    )
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
    "fileId" : "4_h4a48fe8875c6214145260818_f000000000000472a_d20140104_m032022_c001_v0000123_t0104",
    "fileName" : "typing_test.txt"
}
*/
