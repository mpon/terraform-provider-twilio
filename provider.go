package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// This is the root of the provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
}
