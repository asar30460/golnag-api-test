package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var client_id int = -1

type Req struct {
	ClientID int `json:"client_id"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Object struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
	DeletedTime string `json:"deletedTime"`
	Id          string `json:"id"`
	ExternalId  string `json:"externalId"`
	Type        string `json:"type"`
	Password    string `json:"password"`
}

type WebhookData struct {
	Action       string       `json:"action"`
	ClientIp     string       `json:"clientIp"`
	CreatedTime  string       `json:"createdTime"`
	ExtendedUser ExtendedUser `json:"extendedUser"`
	Id           int          `json:"id"`
	IsTriggered  bool         `json:"isTriggered"`
	Language     string       `json:"language"`
	Method       string       `json:"method"`
	Name         string       `json:"name"`
	Object       interface{}  `json:"object"`
	Organization string       `json:"organization"`
	Owner        string       `json:"owner"`
	RequestUri   string       `json:"requestUri"`
	Response     interface{}  `json:"response"`
	StatusCode   int          `json:"statusCode"`
	User         string       `json:"user"`
}

type ExtendedUser struct {
	AccessKey           string            `json:"accessKey"`
	AccessSecret        string            `json:"accessSecret"`
	AccessToken         string            `json:"accessToken"`
	Address             []interface{}     `json:"address"`
	Adfs                string            `json:"adfs"`
	Affiliation         string            `json:"affiliation"`
	Alipay              string            `json:"alipay"`
	Amazon              string            `json:"amazon"`
	Apple               string            `json:"apple"`
	Auth0               string            `json:"auth0"`
	Avatar              string            `json:"avatar"`
	AvatarType          string            `json:"avatarType"`
	Azuread             string            `json:"azuread"`
	Azureadb2c          string            `json:"azureadb2c"`
	Baidu               string            `json:"baidu"`
	Balance             int               `json:"balance"`
	Battlenet           string            `json:"battlenet"`
	Bilibili            string            `json:"bilibili"`
	Bio                 string            `json:"bio"`
	Birthday            string            `json:"birthday"`
	Bitbucket           string            `json:"bitbucket"`
	Box                 string            `json:"box"`
	Casdoor             string            `json:"casdoor"`
	Cloudfoundry        string            `json:"cloudfoundry"`
	CountryCode         string            `json:"countryCode"`
	CreatedIp           string            `json:"createdIp"`
	CreatedTime         string            `json:"createdTime"`
	Currency            string            `json:"currency"`
	Custom              string            `json:"custom"`
	Dailymotion         string            `json:"dailymotion"`
	Deezer              string            `json:"deezer"`
	DeletedTime         string            `json:"deletedTime"`
	Digitalocean        string            `json:"digitalocean"`
	Dingtalk            string            `json:"dingtalk"`
	Discord             string            `json:"discord"`
	DisplayName         string            `json:"displayName"`
	Douyin              string            `json:"douyin"`
	Dropbox             string            `json:"dropbox"`
	Education           string            `json:"education"`
	Email               string            `json:"email"`
	EmailVerified       bool              `json:"emailVerified"`
	Eveonline           string            `json:"eveonline"`
	ExternalId          string            `json:"externalId"`
	FaceIds             interface{}       `json:"faceIds"`
	Facebook            string            `json:"facebook"`
	FirstName           string            `json:"firstName"`
	Fitbit              string            `json:"fitbit"`
	Gender              string            `json:"gender"`
	Gitea               string            `json:"gitea"`
	Gitee               string            `json:"gitee"`
	Github              string            `json:"github"`
	Gitlab              string            `json:"gitlab"`
	Google              string            `json:"google"`
	Groups              interface{}       `json:"groups"`
	Hash                string            `json:"hash"`
	Heroku              string            `json:"heroku"`
	Homepage            string            `json:"homepage"`
	Id                  string            `json:"id"`
	IdCard              string            `json:"idCard"`
	IdCardType          string            `json:"idCardType"`
	Influxcloud         string            `json:"influxcloud"`
	Infoflow            string            `json:"infoflow"`
	Instagram           string            `json:"instagram"`
	Intercom            string            `json:"intercom"`
	Invitation          string            `json:"invitation"`
	InvitationCode      string            `json:"invitationCode"`
	IsAdmin             bool              `json:"isAdmin"`
	IsDefaultAvatar     bool              `json:"isDefaultAvatar"`
	IsDeleted           bool              `json:"isDeleted"`
	IsForbidden         bool              `json:"isForbidden"`
	IsOnline            bool              `json:"isOnline"`
	Kakao               string            `json:"kakao"`
	Karma               int               `json:"karma"`
	Language            string            `json:"language"`
	Lark                string            `json:"lark"`
	LastName            string            `json:"lastName"`
	LastSigninIp        string            `json:"lastSigninIp"`
	LastSigninTime      string            `json:"lastSigninTime"`
	LastSigninWrongTime string            `json:"lastSigninWrongTime"`
	Lastfm              string            `json:"lastfm"`
	Ldap                string            `json:"ldap"`
	Line                string            `json:"line"`
	Linkedin            string            `json:"linkedin"`
	Location            string            `json:"location"`
	Mailru              string            `json:"mailru"`
	ManagedAccounts     interface{}       `json:"managedAccounts"`
	Meetup              string            `json:"meetup"`
	Metamask            string            `json:"metamask"`
	MfaEmailEnabled     bool              `json:"mfaEmailEnabled"`
	MfaPhoneEnabled     bool              `json:"mfaPhoneEnabled"`
	Microsoftonline     string            `json:"microsoftonline"`
	Name                string            `json:"name"`
	Naver               string            `json:"naver"`
	NeedUpdatePassword  bool              `json:"needUpdatePassword"`
	Nextcloud           string            `json:"nextcloud"`
	Okta                string            `json:"okta"`
	Onedrive            string            `json:"onedrive"`
	Oura                string            `json:"oura"`
	Owner               string            `json:"owner"`
	Password            string            `json:"password"`
	PasswordSalt        string            `json:"passwordSalt"`
	PasswordType        string            `json:"passwordType"`
	Patreon             string            `json:"patreon"`
	Paypal              string            `json:"paypal"`
	PermanentAvatar     string            `json:"permanentAvatar"`
	Permissions         interface{}       `json:"permissions"`
	Phone               string            `json:"phone"`
	PreHash             string            `json:"preHash"`
	PreferredMfaType    string            `json:"preferredMfaType"`
	Properties          map[string]string `json:"properties"`
	Qq                  string            `json:"qq"`
	Ranking             int               `json:"ranking"`
	RecoveryCodes       interface{}       `json:"recoveryCodes"`
	Region              string            `json:"region"`
	Roles               interface{}       `json:"roles"`
	Salesforce          string            `json:"salesforce"`
	Score               int               `json:"score"`
	Shopify             string            `json:"shopify"`
	SigninWrongTimes    int               `json:"signinWrongTimes"`
	SignupApplication   string            `json:"signupApplication"`
	Slack               string            `json:"slack"`
	Soundcloud          string            `json:"soundcloud"`
	Spotify             string            `json:"spotify"`
	Steam               string            `json:"steam"`
	Strava              string            `json:"strava"`
	Stripe              string            `json:"stripe"`
	Tag                 string            `json:"tag"`
	Tiktok              string            `json:"tiktok"`
	Title               string            `json:"title"`
	TotpSecret          string            `json:"totpSecret"`
	Tumblr              string            `json:"tumblr"`
	Twitch              string            `json:"twitch"`
	Twitter             string            `json:"twitter"`
	Type                string            `json:"type"`
	Typetalk            string            `json:"typetalk"`
	Uber                string            `json:"uber"`
	UpdatedTime         string            `json:"updatedTime"`
	Vk                  string            `json:"vk"`
	Web3onboard         string            `json:"web3onboard"`
	WebauthnCredentials interface{}       `json:"webauthnCredentials"`
	Wechat              string            `json:"wechat"`
	Wecom               string            `json:"wecom"`
	Weibo               string            `json:"weibo"`
	Wepay               string            `json:"wepay"`
	Xero                string            `json:"xero"`
	Yahoo               string            `json:"yahoo"`
	Yammer              string            `json:"yammer"`
	Yandex              string            `json:"yandex"`
	Zoom                string            `json:"zoom"`
}
