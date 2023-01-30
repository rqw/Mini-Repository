package repository

type Repository struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	DiskPath     string   `json:"diskPath"`              // 磁盘路径
	Mode         int      `json:"mode" default:"4"`      // 仓库模式 0 无效 2 仅可写 4 仅可读 6 可读写
	Cache        bool     `json:"cache" default:"false"` // 是否缓存镜像文件, 默认不缓存
	PublicAccess bool     `json:"publicAccess"`          //是否可以公开访问
	Mirror       []string `json:"mirror"`                // 镜像地址, 会先尝试在本地加载, 如果加载失败, 会尝试从镜像依次读取
}
type Component struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	IsDir   bool   `json:"isDir"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"modTime"`
}
