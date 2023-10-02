package characteristic

const TypePeriodicSnapshotsActive = "225"

type PeriodicSnapshotsActive struct {
	*Bool
}

func NewPeriodicSnapshotsActive() *PeriodicSnapshotsActive {
	c := NewBool(TypePeriodicSnapshotsActive)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetValue(false)

	return &PeriodicSnapshotsActive{c}
}
