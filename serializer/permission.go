package serializer

// 树状权限结构体
type TreePermission struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Path     string            `json:"path"`
	Method   string            `json:"method"`
	Category string            `json:"category"`
	ParentID uint64            `json:"parent_id"`
	Childen  []*TreePermission `json:"childen"`
}
