package characteristic

const TypeRecordingAudioActive = "226"

type RecordingAudioActive struct {
	*Int
}

func NewRecordingAudioActive() *RecordingAudioActive {
	c := NewInt(TypeRecordingAudioActive)
	c.Format = FormatUInt8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionEvents}
	c.SetValue(0)

	return &RecordingAudioActive{c}
}