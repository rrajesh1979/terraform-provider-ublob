package ublob

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"

	//aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	//gcp-sdk-go
	"cloud.google.com/go/storage"

	//azure-sdk-go
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func resourceUblob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBlobCreate,
		ReadContext:   resourceBlobRead,
		UpdateContext: resourceBlobUpdate,
		DeleteContext: resourceBlobDelete,
		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"storage_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_account": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_account_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBlobCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucket := d.Get("bucket").(string)
	cloud := d.Get("cloud").(string)
	region := d.Get("region").(string)
	storageClass := d.Get("storage_class").(string)
	projectID := d.Get("project_id").(string)
	storageAccount := d.Get("storage_account").(string)
	storageAccountKey := d.Get("storage_account_key").(string)

	bucketURL := ""

	if cloud == "AWS" {
		bucketURL = "s3://" + bucket + "?" + region
		log.Println("bucketURL : {}", bucketURL)
		diags = awsCreateBucket(region, bucket, diags)
	} else if cloud == "GCP" {
		bucketURL = "gs://" + bucket
		log.Println("bucketURL : {}", bucketURL)
		diags = gcpCreateBucket(projectID, storageClass, region, bucket, diags)
	} else if cloud == "AZURE" {
		bucketURL = storageAccount + "/" + bucket
		log.Println("bucketURL : {}", bucketURL)
		diags = azureCreateBucket(storageAccount, storageAccountKey, bucket, diags)
	}

	d.SetId(bucketURL)

	resourceBlobRead(ctx, d, m)

	return diags
}

func resourceBlobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//bucketURL := d.Id()

	return diags
}

func resourceBlobUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceBlobRead(ctx, d, m)
}

func resourceBlobDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucket := d.Get("bucket").(string)
	cloud := d.Get("cloud").(string)
	region := d.Get("region").(string)
	storageAccount := d.Get("storage_account").(string)
	storageAccountKey := d.Get("storage_account_key").(string)

	bucketURL := d.Id()

	if cloud == "AWS" {
		bucketURL = "s3://" + bucket + "?" + region
		log.Println("bucketURL : {}", bucketURL)
		diags = awsDeleteBucket(region, bucket, diags, bucketURL)
	} else if cloud == "GCP" {
		bucketURL = "gs://" + bucket
		log.Println("bucketURL : {}", bucketURL)
		diags = gcpDeleteBucket(bucket, diags)
	} else if cloud == "AZURE" {
		bucketURL = storageAccount + "/" + bucket
		log.Println("bucketURL : {}", bucketURL)
		diags = azureDeleteBucket(storageAccount, storageAccountKey, bucket, diags)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func awsDeleteBucket(region string, bucket string, diags diag.Diagnostics, bucketURL string) diag.Diagnostics {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	}
	result, err := svc.DeleteBucket(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete bucket",
			Detail:   "Unable to delete bucket " + bucketURL,
		})

	}
	log.Println(result)
	return diags
}

func awsCreateBucket(region string, bucket string, diags diag.Diagnostics) diag.Diagnostics {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				log.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				log.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create bucket",
			Detail:   "Unable to create bucket " + bucket,
		})
	}
	log.Println(result.String())
	return diags
}

func gcpCreateBucket(projectID string, storageClass string, region string, bucketName string, diags diag.Diagnostics) diag.Diagnostics {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("storage.NewClient: {}", err)
	}
	defer func(client *storage.Client) {
		err := client.Close()
		if err != nil {
			log.Println("Error closing client")
		}
	}(client)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	storageClassAndLocation := &storage.BucketAttrs{
		StorageClass: storageClass,
		Location:     region,
	}
	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, projectID, storageClassAndLocation); err != nil {
		log.Println("Bucket create error", bucketName, err)
	}
	log.Println("Created bucket with storage class", bucketName, storageClassAndLocation.Location, storageClassAndLocation.StorageClass)

	return diags
}

func gcpDeleteBucket(bucketName string, diags diag.Diagnostics) diag.Diagnostics {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("Error creating storage.NewClient {}", err)
	}
	defer func(client *storage.Client) {
		err := client.Close()
		if err != nil {
			log.Println("Error closing client")
		}
	}(client)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucket := client.Bucket(bucketName)
	if err := bucket.Delete(ctx); err != nil {
		log.Println("Bucket.Delete Error: {} {}", bucketName, err)
	}
	log.Println("Deleted bucket {}", bucketName)

	return diags
}

func azureCreateBucket(storageAccount string, storageAccountKey string, bucket string, diags diag.Diagnostics) diag.Diagnostics {

	cred, err := azblob.NewSharedKeyCredential(storageAccount, storageAccountKey)
	if err != nil {
		log.Fatal(err)
	}

	service, err := azblob.NewServiceClient("https://"+storageAccount+".blob.core.windows.net/", cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	container := service.NewContainerClient(bucket)

	_, err = container.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return diags
}

func azureDeleteBucket(storageAccount string, storageAccountKey string, bucket string, diags diag.Diagnostics) diag.Diagnostics {
	cred, err := azblob.NewSharedKeyCredential(storageAccount, storageAccountKey)
	if err != nil {
		log.Fatal(err)
	}

	service, err := azblob.NewServiceClient("https://"+storageAccount+".blob.core.windows.net/", cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	container := service.NewContainerClient(bucket)

	_, err = container.Delete(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return diags
}
