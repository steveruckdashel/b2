package b2

/*
b2_get_upload_url
Gets an URL to use for uploading files.

When you upload a file to B2, you must call b2_get_upload_url first to get the URL for uploading directly to the place where the file will be stored.

TODO: Describe how you know when to get a new upload URL.

Request
Request HTTP Headers

Authorization
required

An account authorization token, obtained from b2_authorize_account.

Request HTTP Message Body Parameters

bucketId
required

The ID of the bucket that you want to upload to.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and no upload url will be returned. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
bucketId
The unique ID of the bucket.

uploadUrl
The URL that can be used to upload files to this bucket, see b2_upload_file.

authorizationToken
The authorizationToken that must be used when uploading files to this bucket, see b2_upload_file.
Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

api_url = "" # Provided by b2_authorize_account
account_authorization_token = "" # Provided by b2_authorize_account
bucket_id = "" # The ID of the bucket you want to upload your file to
request = urllib2.Request(
	'%s/b2api/v1/b2_get_upload_url' % api_url,
	json.dumps({ 'bucketId' : bucket_id }),
	headers = { 'Authorization': account_authorization_token }
	)
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
    "bucketId" : "4a48fe8875c6214145260818",
    "uploadUrl" : "https://pod-000-1005-03.backblaze.com/b2api/v1/b2_upload_file?cvt=c001_v0001005_t0027&bucket=4a48fe8875c6214145260818",
    "authorizationToken" : "2_20151009170037_f504a0f39a0f4e657337e624_9754dde94359bd7b8f1445c8f4cc1a231a33f714_upld"
}
*/
