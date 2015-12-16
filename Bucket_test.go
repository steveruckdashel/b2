package b2

import (
	"os"
	"testing"
	"log"
	"flag"
)

var b2 *B2
var bucketId string
var bucketName string = "5891c9ba3b431ce032169bf1a39abbf2"

func TestMain(m *testing.M) {
	flag.Parse()
	
	accountID := os.Getenv("B2_ID")
	secretKey := os.Getenv("B2_KEY")
	t_b2, err := Authorize(accountID, secretKey)
	if err != nil || t_b2 == nil {
		log.Print("Unable to Authorize")
	}
	b2 = t_b2
	
	os.Exit(m.Run())
}

func TestCreateBucket(t *testing.T) {
	if b2 == nil {
		t.Skip("Preconditions (valid Authorization) not met")
	}
	bucket, err := b2.CreateBucket(bucketName, "allPrivate")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", bucket)
	bucketId = bucket.BucketID
}

func TestListBuckets(t *testing.T) {
	if b2 == nil {
		t.Skip("Preconditions (valid Authorization) not met")
	}
	if len(bucketId) == 0 {
		t.Skip("Preconditions (added bucket) not met")
	}
	buckets, err := b2.ListBuckets()
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", buckets)
	
	contains := false
	for _, b := range buckets.Buckets {
		if b.BucketID == bucketId {
			contains = contains || true
		}
	}
}

func TestUpdateBucket(t *testing.T) {
	if b2 == nil {
		t.Skip("Preconditions (valid Authorization) not met")
	}
	if len(bucketId) == 0 {
		t.Skip("Preconditions (added bucket) not met")
	}
	bucket, err := b2.UpdateBucket(bucketId, "allPublic")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", bucket)
	if bucket.BucketType != "allPublic" {
		t.Fatal("Bucket Type not updated to Public")
	}
}

func TestDeleteBucket(t *testing.T) {
	if b2 == nil {
		t.Skip("Preconditions (valid Authorization) not met")
	}
	if len(bucketId) == 0 {
		t.Skip("Preconditions (added bucket) not met")
	}
	bucket, err := b2.DeleteBucket(bucketId)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", bucket)
	if (bucket.BucketID != bucketId) {
		t.Fatal("BucketID mismatch")
	}
}
