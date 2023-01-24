package permission

// Permission 权限信息
type Permission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"` //1-菜单、2-仓库权限、3-操作权限
	ValidRegex  string `json:"validRegex"`
}

type PermissionCache struct {
	MaxId          int           `json:"maxId"`
	PermissionList []*Permission `json:"permissionList"`
}
