package b2

/*
b2_download_file_by_name
Downloads one file by providing the name of the bucket and the name of the file.

The base URL to use comes from the b2_authorize_account call, and looks something like https://f345.backblaze.com. The "f" in the URL stands for "file", and the number is the cluster number that your account is in. To this base, you add your bucket name, a "/", and then the name of the file. The file name may itself include more "/" characters.

If you have a bucket named "photos", and a file called "cute/kitten.jpg", then the URL for downloading that file would be: https://f345.backblaze.com/file/photos/cute/kitten.jpg.

Request
Request HTTP Headers

Authorization
optional

An account authorization token, obtained from b2_authorize_account. This is required if the bucket containing the file is not public. It is optional for buckets that are public.

Request HTTP Message Body Parameters

This is a GET request, so you cannot post any parameters, and none are accepted in the URL query string.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file will not be downloaded. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
The response headers include the Content-Type that was specified when the file was uploaded. They also include the X-Bz-FileName and X-Bz-Content-Sha1 headers, plus X-Bz-Info-* headers for any custom file info that was provided with the upload. The file-name uses percent-encoding, as if it were a URL parameter.

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

import urllib2

DOWNLOAD_URL = "..." # Comes from b2_authorize_account
BUCKET_NAME = "..." # The name of your bucket (not the ID)
FILE_NAME = "..." # The name of the file in the bucket

url = DOWNLOAD_URL + '/file/' + BUCKET_NAME + '/' + FILE_NAME
print urllib2.urlopen(url).read()

# You will need to use the account authorization token if your bucket's type is allPrivate.

DOWNLOAD_URL = "..." # Comes from b2_authorize_account
BUCKET_NAME = "any_name_you_pick" # 50 char max: letters, digits, “-“ and “_”
FILE_NAME = "..." # The name of the file in the bucket
ACCOUNT_AUTHORIZATION_TOKEN = "..." # Comes from the b2_authorize_account call

url = DOWNLOAD_URL + '/file/' + BUCKET_NAME + '/' + FILE_NAME

headers = {
    'Authorization': ACCOUNT_AUTHORIZATION_TOKEN
    }

request = urllib2.Request(url, None, headers)
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

The quick brown fox jumped over the lazy dog.
*/
