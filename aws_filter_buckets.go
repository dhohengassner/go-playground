package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var bucket string

	flag.StringVar(&bucket, "b", "", "Bucket name.")
	flag.Parse()

	svc := s3.New(session.New())

	input := &s3.ListBucketsInput{}
	result, _ := svc.ListBuckets(input)

	fmt.Println(bucket)

	fmt.Println(result)
}
