package user

type User struct {
	ID             int    `json:"id"`
	LoginName      string `json:"loginName"`
	Fullname       string `json:"fullname"`
	Password       string `json:"password"`
	Token          string `json:"token"`
	PermissionList []int  `json:"permissionList"`
	JwtToken       string `json:"jwtToken"`
}
type Permission struct {
	ID         int    `json:"id"`
	Type       int    `json:"type"`
	ValidRegex string `json:"validRegex"`
}