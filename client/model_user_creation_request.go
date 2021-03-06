/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type UserCreationRequest struct {
	// Status of a user
	Active bool `json:"active,omitempty"`
	// User addresses
	Addresses []AddressInfo `json:"addresses,omitempty"`
	// User email addresses
	Emails []EmailInfo `json:"emails"`
	// External identifier of a user
	ExternalId string `json:"externalId,omitempty"`
	// Internal identifier of a user
	Id   string   `json:"id,omitempty"`
	Name NameInfo `json:"name"`
	// User phone numbers
	PhoneNumbers []PhoneNumberInfoRequest `json:"phoneNumbers,omitempty"`
	Photos       []PhotoInfo              `json:"photos,omitempty"`
	// Specification links
	Schemas                                           []string       `json:"schemas"`
	Urnietfparamsscimschemasextensionenterprise20User EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
	// User mailbox. Must be same as work type email address
	UserName string `json:"userName"`
}
