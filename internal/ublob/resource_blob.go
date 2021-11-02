package ublob

import (
	"context"
	//"gocloud.dev/blob"

	//"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"io/ioutil"
	"log"
	//"os"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBlobCreate,
		ReadContext:   resourceBlobRead,
		UpdateContext: resourceBlobUpdate,
		DeleteContext: resourceBlobDelete,
		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBlobCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucket := d.Get("bucket").(string)
	//cloud := d.Get("cloud").(string)
	//region := d.Get("region").(string)
	cloud := "AWS"
	region := "us-east-1"

	bucketURL := ""

	if cloud == "AWS" {
		bucketURL = "s3://" + bucket + "?" + region
		log.Println("bucketURL : {}", bucketURL)
	}

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
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(result)

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

	return diags
}
