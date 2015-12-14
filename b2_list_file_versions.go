package b2

/*
b2_list_file_versions
Lists all of the versions of all of the files contained in one bucket, in alphabetical order by file name, and by reverse of date/time uploaded for versions of files with the same name.

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

The first file name to return.
If there are no files with this name, the first version of the file with the first name after the given name will be the first in the list.
If startFileId is also specified, the name-and-id pair is the starting point. If there is a file with the given name and ID, it will be first in the list. Otherwise, the first file version that comes after the given name and ID will be first in the list.

startFileId
optional

The first file ID to return. (See startFileName.)

maxFileCount
optional

The maximum number of files to return from this call. The default value is 100, and the maximum allowed is 1000.

Return Value

On success, the HTTP return status is 200. Anything other than 200 is failure and the file versions will not be listed. On failure if possible the server will return the following JSON structure with more information about the error:
   { "code": "codeValue",
     "message": "messageValue",
     "status": http_ret_status_int }
where the messageValue is a human readable message.
Response
The response is:

files
An array of objects, each one describing one file. (See below.)

nextFileName
What to pass in to startFileName for the next search to continue where this one left off, or null if there are no more files. Note this this may not be the name of an actual file, but using it is guaranteed to find the next file version in the bucket.

nextFileId
What to pass in to startFileId for the next search to continue where this one left off, or null if there are no more files. Note this this may not be the ID of an actual file, but using it is guaranteed to find the next file version in the bucket.

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
	'%s/b2api/v1/b2_list_file_versions' % api_url,
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
      "fileId": "4_z27c88f1d182b150646ff0b16_f100920ddab886245_d20150809_m232316_c100_v0009990_t0003",
      "fileName": "files/hello.txt",
      "size": 6,
      "uploadTimestamp": 1439162596000
    },
    {
      "action": "hide",
      "fileId": "4_z27c88f1d182b150646ff0b16_f100920ddab886247_d20150809_m232323_c100_v0009990_t0005",
      "fileName": "files/world.txt",
      "size": 0,
      "uploadTimestamp": 1439162603000
    },
    {
      "action": "upload",
      "fileId": "4_z27c88f1d182b150646ff0b16_f100920ddab886246_d20150809_m232316_c100_v0009990_t0003",
      "fileName": "files/world.txt",
      "size": 6,
      "uploadTimestamp": 1439162596000
    }
  ],
  "nextFileId": "4_z27c88f1d182b150646ff0b16_f100920ddab886247_d20150809_m232316_c100_v0009990_t0003",
  "nextFileName": "files/world.txt"
}
*/
