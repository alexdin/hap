package characteristic

const TypeSetupDataStreamTransport = "131"

type SetupDataStreamTransport struct {
	*Bytes
}

func NewSetupDataStreamTransport() *SetupDataStreamTransport {
	c := NewBytes(TypeSetupDataStreamTransport)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionWriteResponse}
	c.SetValue([]byte{})

	return &SetupDataStreamTransport{c}
}
