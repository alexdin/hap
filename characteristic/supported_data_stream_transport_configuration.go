package characteristic

const TypeSupportedDataStreamTransportConfiguration = "130"

type SupportedDataStreamTransportConfiguration struct {
	*Bytes
}

func NewSupportedDataStreamTransportConfiguration() *SupportedDataStreamTransportConfiguration {
	c := NewBytes(TypeSupportedDataStreamTransportConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionWriteResponse}
	c.SetValue([]byte{})
	return &SupportedDataStreamTransportConfiguration{c}
}
