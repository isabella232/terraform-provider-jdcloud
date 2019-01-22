package jdcloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/client"
	vpcModels "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
	"time"
)

func resourceJDCloudEIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceJDCloudEIPCreate,
		Read:   resourceJDCloudEIPRead,
		Delete: resourceJDCloudEIPDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"eip_provider": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bandwidth_mbps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"elastic_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceJDCloudEIPCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*JDCloudConfig)
	elasticIpSpec := vpcModels.ElasticIpSpec{
		BandwidthMbps: d.Get("bandwidth_mbps").(int),
		Provider:      d.Get("eip_provider").(string),
		ChargeSpec:    &models.ChargeSpec{},
	}
	vpcClient := client.NewVpcClient(config.Credential)
	req := apis.NewCreateElasticIpsRequest(config.Region, MAX_EIP_COUNT, &elasticIpSpec)
	if _, ok := d.GetOk("elastic_ip_address"); ok {
		req.ElasticIpAddress = GetStringAddr(d, "elastic_ip_address")
	}

	return resource.Retry(20*time.Second, func() *resource.RetryError {

		resp, err := vpcClient.CreateElasticIps(req)

		if err == nil && resp.Error.Code == REQUEST_COMPLETED {
			d.SetId(resp.Result.ElasticIpIds[0])
			return nil
		}

		if connectionError(err) {
			return resource.RetryableError(formatConnectionErrorMessage())
		} else {
			return resource.NonRetryableError(formatErrorMessage(resp.Error, err))
		}
	})
}

func resourceJDCloudEIPRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*JDCloudConfig)
	req := apis.NewDescribeElasticIpRequest(config.Region, d.Id())
	vpcClient := client.NewVpcClient(config.Credential)
	resp, err := vpcClient.DescribeElasticIp(req)

	if err != nil {
		return fmt.Errorf("[ERROR] resourceJDCloudEIPRead failed %s ", err.Error())
	}
	if resp.Error.Code == RESOURCE_NOT_FOUND {
		d.SetId("")
		return nil
	}

	if resp.Error.Code != REQUEST_COMPLETED {
		return fmt.Errorf("[ERROR] resourceJDCloudEIPRead failed  code:%d staus:%s message:%s ", resp.Error.Code, resp.Error.Status, resp.Error.Message)
	}

	return nil
}

func resourceJDCloudEIPDelete(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*JDCloudConfig)
	elasticIpId := d.Id()
	rq := apis.NewDeleteElasticIpRequest(config.Region, elasticIpId)
	vpcClient := client.NewVpcClient(config.Credential)

	return resource.Retry(20*time.Second, func() *resource.RetryError {

		resp, err := vpcClient.DeleteElasticIp(rq)

		if err == nil && resp.Error.Code == REQUEST_COMPLETED {
			d.SetId("")
			return nil
		}

		if resp.Error.Code == RESOURCE_NOT_FOUND {
			d.SetId("")
			return nil
		}

		if connectionError(err) {
			return resource.RetryableError(formatConnectionErrorMessage())
		} else {
			return resource.NonRetryableError(formatErrorMessage(resp.Error, err))
		}
	})
}

/*
Log:
	Its been tested that you can not tell if an EIP has been detached or not
	Since there is no info on "status" about EIP
*/
