package b2

/*
b2_upload_file
Uploads one file to B2, returning its unique file ID.

Request
The upload request is a POST. The file name and other parameters are in the request headers, and the file to be uploaded is the request body.

URL Path

Use the b2_get_upload_url operation to get a URL you can use to upload files. The URL it returns will contain your bucket ID and the upload destination, and will look something like this:

https://pod-000-1007-13.backblaze.com/b2api/v1/b2_upload_file/4a48fe8875c6214145260818/c001_v0001007_t0042
Request HTTP Headers

Authorization
required

An upload authorization token, from b2_get_upload_url.

X-Bz-File-Name
required

The name of the file, in percent-encoded UTF-8. See Files for requirements on file names. See String Encoding.

Content-Type
required

The MIME type of the content of the file, which will be returned in the Content-Type header when downloading the file. Use the Content-Type b2/x-auto to automatically set the stored Content-Type post upload. In the case where a file extension is absent or the lookup fails, the Content-Type is set to application/octet-stream. The Content-Type mappings can be purused here.

Content-Length
required

The number of bytes in the file being uploaded. Note that this header is required; you cannot leave it out and just use chunked encoding.

X-Bz-Content-Sha1
required

The SHA1 checksum of the content of the file. B2 will check this when the file is uploaded, to make sure that the file arrived correctly. It will be returned in the X-Bz-Content-Sha1 header when the file is downloaded.

X-Bz-Info-src_last_modified_millis
optional

If the original source of the file being uploaded has a last modified time concept, Backblaze recommends using this spelling of one of your ten X-Bz-Info-* headers (see below). Using a standard spelling allows different B2 clients and the B2 web user interface to interoperate correctly. The value should be a base 10 number which represents a UTC time when the original source file was last modified. It is a base 10 number of milliseconds since midnight, January 1, 1970 UTC. This fits in a 64 bit integer such as the type "long" in the programming language Java. It is intended to be compatible with Java's time long. For example, it can be passed directly into the Java call Date.setTime(long time).

X-Bz-Info-*
optional

Up to 10 of these headers may be present. The * part of the header name is replace with the name of a custom field in the file information stored with the file, and the value is an arbitrary UTF-8 string, percent-encoded. The same info headers sent with the upload will be returned with the download.

The following HTTP headers must not be included in the b2_upload_file request:

Cache-Control
Content-Disposition
Content-Encoding
Content-Language
Content-Location
Content-Language
Content-Range
Expires
Request HTTP Message Body Parameters

There are no JSON parameters allowed. The file to be uploaded is the message body and is not encoded in any way. It is not URL encoded. It is not MIME encoded.

Return Value

On success, the HTTP return status is 200 and the file is inside B2. Anything other than 200 is failure and the file will not be in B2.

If the failure returns an HTTP status code in the range 400 through 499, it means that there is a problem with your request. See the error code and message for information on what went wrong. This is an example of a returned error:

{
    "code": "storage_cap_exceeded",
    "message": "Cannot upload files, storage cap exceeded",
    "status": 403
}

Failure codes in the range 500 through 599 mean that the storage pod is having trouble accepting your data. In this case you must call b2_get_upload_url to get a new uploadUrl and a new authorizationToken. The reason for this is that the individual storage pod you are uploading into may be full of data and thus will never accept any more data for now, or it has crashed, or has been placed into maintenance and will not accept data, etc. No matter what the reason, you must use b2_get_upload_url to get a new uploadUrl and new authorizationToken in order to retry the upload.

Response
fileId
The unique identifier for this version of this file. Used with b2_get_file_info, b2_download_file_by_id, and b2_delete_file_version.

fileName
The name of this file, which can be used with b2_download_file_by_name.

accountId
Your account ID.

bucketId
The bucket that the file is in.

contentLength
The number of bytes stored in the file.

contentSha1
The SHA1 of the bytes stored in the file.

contentType
The MIME type of the file.

fileInfo
The custom information that was uploaded with the file. This is a JSON object, holding the name/value pairs that were uploaded with the file.

Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2
import hashlib

upload_url = "" # Provided by b2_get_upload_url
upload_authorization_token = "" # Provided by b2_get_upload_url
file_data = "Now, I am become Death, the destroyer of worlds."
file_name = "oppenheimer_says.txt"
content_type = "text/plain"
sha1_of_file_data = hashlib.sha1(file_data).hexdigest()

headers = {
    'Authorization' : upload_authorization_token,
    'X-Bz-File-Name' :  file_name,
    'Content-Type' : content_type,
    'X-Bz-Content-Sha1' : sha1_of_file_data
    }
request = urllib2.Request(upload_url, file_data, headers)

response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
    "fileId" : "4_h4a48fe8875c6214145260818_f000000000000472a_d20140104_m032022_c001_v0000123_t0104",
    "fileName" : "typing_test.txt",
    "accountId" : "d522aa47a10f",
    "bucketId" : "4a48fe8875c6214145260818",
    "contentLength" : 46,
    "contentSha1" : "bae5ed658ab3546aee12f23f36392f35dba1ebdd",
    "contentType" : "text/plain",
    "fileInfo" : {
       "author" : "unknown"
    }
}
*/
