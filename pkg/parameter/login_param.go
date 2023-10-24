package parameter

import "github.com/google/go-github/v55/github"

type LoginParam struct {
	Code         string `json:"code"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type MetaMaskLoginParam struct {
	Address string `json:"address"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type AuthParam struct {
	Code         string `json:"code"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	UserId       int64  `json:"userId"`
}

type GithubWebHook struct {
	Action       string      `json:"action"`
	Installation interface{} `json:"installation"`
	Sender       SenderUser  `json:"sender"`
}

type GithubWebHookInstall struct {
	Action       string              `json:"action"`
	Installation github.Installation `json:"installation"`
}

type SenderUser struct {
	AvatarUrl         string `json:"avatar_url"`
	EventsUrl         string `json:"events_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	GravatarId        string `json:"gravatar_id"`
	Id                int64  `json:"id"`
	Login             string `json:"login"`
	NodeId            string `json:"node_id"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	ReposUrl          string `json:"repos_url"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	Url               string `json:"url"`
}

type InstallParam struct {
	Code string `json:"code"`
}
