package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"

    "bytes"
	"io"
	
)

// func downloadFile(myBucket, myString, filename string) string {
// 	// The session the S3 Downloader will use
// 	sess := session.Must(session.NewSession())

// 	// Create a downloader with the session and default options
// 	downloader := s3manager.NewDownloader(sess)

// 	// Create a file to write the S3 Object contents to.
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file %q, %v", filename, err)
// 	}

// 	// Write the contents of S3 Object to the file
// 	n, err := downloader.Download(f, &s3.GetObjectInput{
// 		Bucket: aws.String(myBucket),
// 		Key:    aws.String(myString),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to download file, %v", err)
// 	}
// 	fmt.Printf("file downloaded, %d bytes\n", n)
// }


type S3Handler struct {
    Session *session.Session
    Bucket  string
}

func (h S3Handler) ReadFile(key string) (string, error) {
    results, err := s3.New(h.Session).GetObject(&s3.GetObjectInput{
        Bucket: aws.String(h.Bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        return "", err
    }
    defer results.Body.Close()

    buf := bytes.NewBuffer(nil)
    if _, err := io.Copy(buf, results.Body); err != nil {
        return "", err
    }
    return string(buf.Bytes()), nil
}