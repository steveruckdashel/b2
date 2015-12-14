package b2

/*
b2_download_file_by_id
Downloads one file from B2.

The response contains the following headers, which contain the same information they did when the file was uploaded:

Content-Length
Content-Type
X-Bz-File-Id
X-Bz-File-Name
X-Bz-Content-Sha1
X-Bz-Info-*
HEAD requests are also supported, and work just like a GET, except that the body of the response is not included. All of the same headers, including Content-Length are returned.

If the bucket containing the file is set to require authorization, then you must supply the bucket's auth token in the Authorzation header.

Because errors can happen in network transmission, you should check the SHA1 of the data you receive against the SHA1 returned in the X-Bz-Content-Sha1 header.

Request
The URL to use for downloading is found on the Account page on the B2 web site, and looks like this, possible with different numbers in the host name: https://f001.backblaze.com/b2api/v1/b2_download_file_by_id.

As with normal API calls, the request information can either be posted as JSON, or put in the URL query parameters. Unlike the other API calls, the response is not JSON, but is the contents of the file.

Request HTTP Message Body Parameters

fileId
Required

The file ID that was returned from b2_upload_file. It can also be found using b2_list_files or b2_list_file_versions.

Example

GET /api/b2_download_file_by_id?fileId=4_h4a48fe8875c6214145260818_f000000000000472a_d20140104_m032022_c001_v0000123_t0104 HTTP/1.1
User-Agent: curl/7.41.0
Host: f001.backblaze.com
Authorization: 1_20100215141633_18d3a718d3a718d3a718d3a7_1d148b2427e9aff1364898aae20246802ec9733d

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file will not be downloaded. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
The response headers include the Content-Type that was specified when the file was uploaded. They also include the X-Bz-FileName and X-Bz-Content-Sha1 headers, plus X-Bz-Info-* headers for any custom file info that was provided with the upload. The X-Bz-FileName uses percent-encoding, as if it were a URL parameter.

Example

HTTP/1.1 200 OK
Content-Length: 46
Content-Type: text/plain
X-Bz-File-Id: 4_h4a48fe8875c6214145260818_f000000000000472a_d20140104_m032022_c001_v0000123_t0104
X-Bz-File-Name: typing-test.txt
X-Bz-Content-Sha1: bae5ed658ab3546aee12f23f36392f35dba1ebdd
X-Bz-Info-author: unknown

The quick brown fox jumped over the lazy dog.
Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

download_url = "" # Provided by authorize account
file_id = "" # The file ID of the file you want to download
print urllib2.urlopen(download_url + '/b2api/v1/b2_download_file_by_id?fileId=' + file_id)
Output

The quick brown fox jumped over the lazy dog.
*/
