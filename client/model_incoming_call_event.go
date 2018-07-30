/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

import (
	"time"
)

type IncomingCallEvent struct {
	Aps ApsInfo `json:"aps,omitempty"`
	// Event filter URI
	Event string `json:"event,omitempty"`
	// Universally unique identifier of a notification
	Uuid string `json:"uuid,omitempty"`
	// Internal identifier of a subscription
	SubscriptionId string `json:"subscriptionId,omitempty"`
	// The datetime of a call action in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Internal identifier of an extension
	ExtensionId string `json:"extensionId,omitempty"`
	// Calling action, for example 'StartRing'
	Action string `json:"action,omitempty"`
	// Identifier of a call session
	SessionId string `json:"sessionId,omitempty"`
	// Identifier of a server
	ServerId string `json:"serverId,omitempty"`
	// Phone number of a caller
	From string `json:"from,omitempty"`
	// Caller name
	FromName string `json:"fromName,omitempty"`
	// Phone number of a callee
	To string `json:"to,omitempty"`
	// Callee name
	ToName string `json:"toName,omitempty"`
	// Unique identifier of a session
	Sid string `json:"sid,omitempty"`
	// SIP proxy registration name
	ToUrl string `json:"toUrl,omitempty"`
	// User data
	SrvLvl string `json:"srvLvl,omitempty"`
	// User data
	SrvLvlExt string `json:"srvLvlExt,omitempty"`
	// File containing recorded caller name
	RecUrl string `json:"recUrl,omitempty"`
	// Notification lifetime value in seconds, the default value is 30 seconds
	PnTtl int32 `json:"pn_ttl,omitempty"`
}