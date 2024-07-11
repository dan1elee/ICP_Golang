package enums

// 用户权限
const (
	STUDENT uint = 1
	TEACHER uint = (1 << 1) | 1
	ADMIN   uint = (1 << 2) | (1 << 1) | 1
)

// 信息级别
const (
	LEVELNORMAL uint = 1
	LEVELSECRET uint = 1 << 1
	LEVELTOP    uint = 1 << 2
)
