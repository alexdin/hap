package characteristic

import "github.com/brutella/hap/tlv8"

const TypeSupportedDataStreamTransportConfiguration = "130"

type SupportedDataStreamTransportConfiguration struct {
	*Bytes
}

// hdsTransportType is the inner TLV: {TransportType(1): 0x00 = TCP}
type hdsTransportType struct {
	TransportType uint8 `tlv8:"1"`
}

// hdsTransferTransportConfig is the outer TLV wrapper per HAP spec:
// {TransferTransportConfiguration(1): {TransportType(1): 0x00}}
type hdsTransferTransportConfig struct {
	Config hdsTransportType `tlv8:"1"`
}

func NewSupportedDataStreamTransportConfiguration() *SupportedDataStreamTransportConfiguration {
	c := NewBytes(TypeSupportedDataStreamTransportConfiguration)
	c.Format = FormatTLV8
	c.Permissions = []string{PermissionRead, PermissionWrite, PermissionWriteResponse}
	cfg := hdsTransferTransportConfig{Config: hdsTransportType{TransportType: 0}}
	if b, err := tlv8.Marshal(cfg); err == nil {
		c.SetValue(b)
	}
	return &SupportedDataStreamTransportConfiguration{c}
}
