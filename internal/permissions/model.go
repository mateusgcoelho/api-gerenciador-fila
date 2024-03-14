package permissions

type Permission uint8

const (
	PermissionCreateUser Permission = 1 << 0 // 00000001
)
