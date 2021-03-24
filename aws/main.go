package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"os"
	"strings"
)

var (
	s3session *s3.S3
)

const (
	//BUCKET_NAME = "test"
	BUCKET_NAME = "charlessobraj123"
	REGION      = "us-east-2"
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})))
}

func listBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		err := fmt.Errorf("Key already exists, cannot insert multiples")
		fmt.Println(err.Error())
	}
	return resp
}

func createBucket() (resp *s3.CreateBucketOutput) {
	resp, err := s3session.CreateBucket(&s3.CreateBucketInput{
		//ACL: aws.String(s3.BucketCannedACLPrivate),
		//ACL: aws.String(s3.BucketCannedACLPublicRead),
		Bucket: aws.String(BUCKET_NAME),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(REGION),
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println("Bucket Already Exists!")
				panic(err)
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println("Bucket exists and is owned by you! ")

			default:
				panic(err)

			}
		}
	}
	return resp
}

func uploadObjects(filename string) (resp *s3.PutObjectOutput) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Uploading Filename: ", filename)
	resp, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(strings.Split(filename, "/")[1]),
		//ACL: aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		panic(err)
	}
	return resp
}
func listObjects() (resp *s3.ListObjectsV2Output) {
	resp, err := s3session.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(BUCKET_NAME),
	})
	if err != nil {
		panic(err)
	}
	return resp
}

func getObject(filename string) {
	fmt.Println("Downloading ", filename)
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, body, 0644)

	if err != nil {
		panic(err)
	}
}

func deletefile(filename string) (resp *s3.DeleteObjectOutput) {

	fmt.Println("Deleting file : ", filename, " from bucket: ", BUCKET_NAME)
	resp, err := s3session.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})
	if err != nil {
		panic(err)
	}
	return resp
}

func main() {
	//fmt.Println("My Buckets are : ", listBuckets())
	//fmt.Println("Create Bucket Charlessobraj : ", createBucket())
	//fmt.Println("My Buckets are : ", listBuckets())
	//fmt.Println("Uploading objects to My Buckets: ")
	//uploadObjects("files/im1.png")
	//fmt.Println("Files uploaded : \n", listObjects())
	//fmt.Println("Dowloading files :")
	//getObject("im1.png")
	//deletefile("im1.png")
	//fmt.Println("Files uploaded : \n", listObjects())
	//
	//
	folder := "files"
	files, _ := ioutil.ReadDir(folder)
	fmt.Println(files)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			uploadObjects(folder + "/" + file.Name())
		}
	}
	fmt.Println("Files uploaded : \n", listObjects())
	fmt.Scanln()

	for _, object := range listObjects().Contents {
		getObject(*object.Key)
		deletefile(*object.Key)
	}

	//deletefile("im1.png")
	fmt.Println("Files uploaded : \n", listObjects())
}
