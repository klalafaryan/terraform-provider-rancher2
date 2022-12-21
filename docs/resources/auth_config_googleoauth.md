---
page_title: "rancher2_auth_config_googleoauth Resource"
---

# rancher2\_auth\_config\_googleoauth

Provides a Rancher v2 Auth Config GoogleOauth resource. This can be used to configure and enable Auth Config GoogleOauth for Rancher v2 RKE clusters and retrieve their information.

In addition to the built-in local auth, only one external auth config provider can be enabled at a time.

## Example Usage

```hcl
# Create a new rancher2 Auth Config GoogleOauth
resource "rancher2_auth_config_googleoauth" "googleoauth" {
  admin_email = "<GOOGLEOAUTH_ADMIN_EMAIL>"
  hostname = "<GOOGLEOAUTH_HOSTNAME>"
  oauth_credential = "<GOOGLEOAUTH_OAUTH_CREDENTIAL>"
  service_account_credential = "<GOOGLEOAUTH_SERVICE_ACCOUNT_CREDENTIAL>"
  nested_group_membership_enabled = true
}
```