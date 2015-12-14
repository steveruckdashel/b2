package b2

/*
b2_list_file_names
Lists the names of all files in a bucket, starting at a given name.

This call returns at most 1000 file names, but it can be called repeatedly to scan through all of the file names in a bucket. Each time you call, it returns an "endFileName" that can be used as the starting point for the next call.

There may be many file versions for the same name, but this call will return each name only once. If you want all of the versions, use b2_list_file_versions instead.

To go through all of the file names in a bucket, use a loop like this:

startFileName = null
while true:
    response = b2_list_file_names(bucketId = ..., startFileName = startFileName)
    for file in response.files:
        process_one_file_name(file)
    if response.endFileName == null:
         break
    startFileName = response.nextFileName
Request
Request HTTP Headers

Authorization
required

The account authorization token returned by b2_authorize_account.

Request HTTP Message Body Parameters

bucketId
required

The bucket to look for file names in.

startFileName
optional

The first file name to return. If there is a file with this name, it will be returned in the list. If not, the first file name after this the first one after this name.

maxFileCount
optional

The maximum number of files to return from this call. The default value is 100, and the maximum allowed is 1000.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file names will not be listed. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
The response is:

files
An array of objects, each one describing one file. (See below.)

nextFileName
What to pass in to startFileName for the next search to continue where this one left off, or null if there are no more files. Note this this may not be the name of an actual file, but using it is guaranteed to find the next file in the bucket.

And each of the files is:

fileId
The unique identifier for this version of this file. Used with b2_get_file_info, b2_download_file_by_id, and b2_delete_file_version.

fileName
The name of this file, which can be used with b2_download_file_by_name.

action
Either "upload" or "hide". "upload" means a file that was uploaded to B2 Cloud Storage. "hide" means a file version marking the file as hidden, so that it will not show up in b2_list_file_names.

The result of b2_list_file_names will contain only "upload". The result of b2_list_file_versions may have both.

size
The number of bytes in the file.

uploadTimestamp
This is a UTC time when this file was uploaded. It is a base 10 number of milliseconds since midnight, January 1, 1970 UTC. This fits in a 64 bit integer such as the type "long" in the programming language Java. It is intended to be compatible with Java's time long. For example, it can be passed directly into the java call Date.setTime(long time).

Sample Code
cUrlJavaPythonSwiftRubyC#PHP
Code

import json
import urllib2

api_url = "" # Provided by b2_authorize_account
account_authorization_token = "" # Provided by b2_authorize_account
bucket_id = "" # The ID of the bucket you are querying
request = urllib2.Request(
	'%s/b2api/v1/b2_list_file_names' % api_url,
	json.dumps({ 'bucketId' : bucket_id }),
	headers = { 'Authorization': account_authorization_token }
	)
response = urllib2.urlopen(request)
response_data = json.loads(response.read())
response.close()
Output

{
  "files": [
    {
      "action": "upload",
      "fileId": "4_z27c88f1d182b150646ff0b16_f1004ba650fe24e6b_d20150809_m012853_c100_v0009990_t0000",
      "fileName": "files/hello.txt",
      "size": 6,
      "uploadTimestamp": 1439083733000
    },
    {
      "action": "upload",
      "fileId": "4_z27c88f1d182b150646ff0b16_f1004ba650fe24e6c_d20150809_m012854_c100_v0009990_t0000",
      "fileName": "files/world.txt",
      "size": 6,
      "uploadTimestamp": 1439083734000
    }
  ],
  "nextFileName": null
}
*/
