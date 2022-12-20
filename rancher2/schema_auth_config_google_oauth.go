package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const AuthConfigGoogleOauthName = "googleOauth"

//Schemas

func authConfigGoogleOauthFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"admin_email": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nested_group_membership_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"oauth_credentials": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			StateFunc: TrimSpace,
		},
		"service_account_credentials": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			StateFunc: TrimSpace,
		},
	}

	for k, v := range authConfigFields() {
		s[k] = v
	}

	return s
}
