package rancher2

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

// Flatteners

func flattenAuthConfigGoogleOauth(d *schema.ResourceData, in *managementClient.GoogleOauthConfig) error {
	d.SetId(AuthConfigGoogleOauthName)
	d.Set("name", AuthConfigGoogleOauthName)
	d.Set("type", managementClient.GoogleOauthConfigType)
	d.Set("access_mode", in.AccessMode)

	err := d.Set("allowed_principal_ids", toArrayInterface(in.AllowedPrincipalIDs))
	if err != nil {
		return err
	}

	d.Set("enabled", in.Enabled)

	err = d.Set("annotations", toMapInterface(in.Annotations))
	if err != nil {
		return err
	}
	err = d.Set("labels", toMapInterface(in.Labels))
	if err != nil {
		return err
	}

	d.Set("admin_email", in.AdminEmail)
	d.Set("domain", in.Hostname)
	d.Set("nested_group_membership_enabled", in.NestedGroupMembershipEnabled)
	d.Set("oauth_credentials", in.OauthCredential)
	d.Set("service_account_credentials", in.ServiceAccountCredential)

	return nil
}

// Expanders

func expandAuthConfigGoogleOauth(in *schema.ResourceData) (*managementClient.GoogleOauthConfig, error) {
	obj := &managementClient.GoogleOauthConfig{}
	if in == nil {
		return nil, fmt.Errorf("expanding %s Auth Config: Input ResourceData is nil", AuthConfigGoogleOauthName)
	}

	obj.Name = AuthConfigGoogleOauthName
	obj.Type = managementClient.GoogleOauthConfigType

	if v, ok := in.Get("access_mode").(string); ok && len(v) > 0 {
		obj.AccessMode = v
	}

	if v, ok := in.Get("allowed_principal_ids").([]interface{}); ok && len(v) > 0 {
		obj.AllowedPrincipalIDs = toArrayString(v)
	}

	if (obj.AccessMode == "required" || obj.AccessMode == "restricted") && len(obj.AllowedPrincipalIDs) == 0 {
		return nil, fmt.Errorf("expanding %s Auth Config: allowed_principal_ids is required on access_mode %s", AuthConfigGoogleOauthName, obj.AccessMode)
	}

	if v, ok := in.Get("enabled").(bool); ok {
		obj.Enabled = v
	}

	if v, ok := in.Get("annotations").(map[string]interface{}); ok && len(v) > 0 {
		obj.Annotations = toMapString(v)
	}

	if v, ok := in.Get("labels").(map[string]interface{}); ok && len(v) > 0 {
		obj.Labels = toMapString(v)
	}

	if v, ok := in.Get("admin_email").(string); ok && len(v) > 0 {
		obj.AdminEmail = v
	}

	if v, ok := in.Get("domain").(string); ok && len(v) > 0 {
		obj.Hostname = v
	}

	if v, ok := in.Get("nested_group_membership_enabled").(bool); ok {
		obj.NestedGroupMembershipEnabled = v
	}

	if v, ok := in.Get("oauth_credentials").(string); ok && len(v) > 0 {
		obj.OauthCredential = v
	}

	if v, ok := in.Get("service_account_credentials").(string); ok && len(v) > 0 {
		obj.ServiceAccountCredential = v
	}

	return obj, nil
}
