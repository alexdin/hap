package characteristic

const TypeHomeKitCameraActive = "21B"

type HomeKitCameraActive struct {
	*Bool
}

func NewHomeKitCameraActive() *HomeKitCameraActive {
	c := NewBool(TypeHomeKitCameraActive)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents, PermissionTimedWrite}
	c.SetValue(false)

	return &HomeKitCameraActive{c}
}
