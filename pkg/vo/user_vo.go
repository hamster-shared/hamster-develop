package vo

type UserVo struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	AvatarUrl  string `json:"avatarUrl"`
	HtmlUrl    string `json:"htmlUrl"`
	Token      string `json:"token"`
	FirstState int    `json:"firstState"`
}
