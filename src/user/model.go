package user

type User struct {
	ID             int    `json:"id"`
	LoginName      string `json:"loginName"`
	Fullname       string `json:"fullname"`
	Password       string `json:"password"`
	Mail           string `json:"mail"`
	Tel            string `json:"tel"`
	Act            string `json:"act"`
	JwtToken       string `json:"jwtToken"`
	Mrt            string `json:"mrt"`
	PermissionList []int  `json:"permissionList"`
	Roles          []Role `json:"roles"`
}

type Role struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}
type PasswdInfo struct {
	ID            int    `json:"id"`
	OldPassword   string `json:"oldPassword"`
	NewPassword   string `json:"newPassword"`
	AgainPassword string `json:"againPassword"`
}
