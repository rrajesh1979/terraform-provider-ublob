package ublob

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"ublob_meta": dataSourceMeta(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ublob_blob": resourceUblob(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

//TODO: Add code here to initialize the provider
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return nil, diags
}
