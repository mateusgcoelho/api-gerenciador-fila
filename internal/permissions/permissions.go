package permissions

type Permission uint8

const (
	PermissionCreateUser   Permission = 1 << 0 // 00000001
	PermissionSeeAllUsers  Permission = 1 << 1 // 00000010
	PermissionCreateQueue  Permission = 1 << 2 // 00000100
	PermissionSeeAllQueues Permission = 1 << 3 // 00001000
)
