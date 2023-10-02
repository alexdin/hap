package characteristic

const TypeEventSnapshotsActive = "223"

type EventSnapshotsActive struct {
	*Bool
}

func NewEventSnapshotsActive() *EventSnapshotsActive {
	c := NewBool(TypeEventSnapshotsActive)
	c.Format = FormatBool
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents, PermissionTimedWrite}
	c.SetValue(false)

	return &EventSnapshotsActive{c}
}
