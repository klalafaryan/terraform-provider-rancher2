package rancher2

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

const (
	testAccRancher2AuthConfigGoogleOauthType   = "rancher2_auth_config_google_oauth"
	testAccRancher2AuthConfigGoogleOauthConfig = `
resource "` + testAccRancher2AuthConfigGoogleOauthType + `" "google_oauth" {
  admin_email = "admin_email@gmail.com"
  domain = "domain.com"
  nested_group_membership_enabled = true
  oauth_credentials = "test_credentials"
  service_account_credentials = "test_credentials_service_account"
}
`
	testAccRancher2AuthConfigGoogleOauthUpdateConfig = `
resource "` + testAccRancher2AuthConfigGoogleOauthType + `" "google_oauth" {
  admin_email = "admin_email@gmail.com"
  domain = "domain-updated.com"
  nested_group_membership_enabled = true
  oauth_credentials = "test_credentials-UPDATED"
  service_account_credentials = "test_credentials_service_account"
}
 `
)

func TestAccRancher2AuthConfigGoogleOauth_basic(t *testing.T) {
	var authConfig *managementClient.AuthConfig

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRancher2AuthConfigGoogleOauthDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRancher2AuthConfigGoogleOauthConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2AuthConfigExists(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, authConfig),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "name", AuthConfigGoogleOauthName),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "admin_email", "admin_email@gmail.com"),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "domain", "domain.com"),
				),
			},
			{
				Config: testAccRancher2AuthConfigGoogleOauthUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2AuthConfigExists(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, authConfig),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "name", AuthConfigGoogleOauthName),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "admin_email", "admin_email@gmail.com"),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "domain", "domain-updated.com"),
				),
			},
			{
				Config: testAccRancher2AuthConfigGoogleOauthConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2AuthConfigExists(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, authConfig),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "name", AuthConfigGoogleOauthName),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "admin_email", "admin_email@gmail.com"),
					resource.TestCheckResourceAttr(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, "domain", "domain.com")),
			},
		},
	})
}

func TestAccRancher2AuthConfigGoogleOauth_disappears(t *testing.T) {
	var authConfig *managementClient.AuthConfig

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRancher2AuthConfigGoogleOauthDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRancher2AuthConfigGoogleOauthConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2AuthConfigExists(testAccRancher2AuthConfigGoogleOauthType+"."+AuthConfigGoogleOauthName, authConfig),
					testAccRancher2AuthConfigDisappears(authConfig, testAccRancher2AuthConfigGoogleOauthType),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckRancher2AuthConfigGoogleOauthDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccRancher2AuthConfigGoogleOauthType {
			continue
		}
		client, err := testAccProvider.Meta().(*Config).ManagementClient()
		if err != nil {
			return err
		}

		auth, err := client.AuthConfig.ByID(rs.Primary.ID)
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return err
		}

		if auth.Enabled == true {
			err = client.Post(auth.Actions["disable"], nil, nil)
			if err != nil {
				return fmt.Errorf("[ERROR] Disabling Auth Config %s: %s", auth.ID, err)
			}
		}
		return nil
	}
	return nil
}
