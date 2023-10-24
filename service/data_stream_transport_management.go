package service

import "github.com/brutella/hap/characteristic"

const TypeDataStreamTransportManagement = "129"

type DataStreamTransportManagement struct {
	*S

	SetupDataStreamTransport                  *characteristic.SetupDataStreamTransport
	SupportedDataStreamTransportConfiguration *characteristic.SupportedDataStreamTransportConfiguration
	Version                                   *characteristic.Version
}

func NewDataStreamTransportManagement() *DataStreamTransportManagement {

	s := DataStreamTransportManagement{}
	s.S = New(TypeDataStreamTransportManagement)

	s.SetupDataStreamTransport = characteristic.NewSetupDataStreamTransport()
	s.AddC(s.SetupDataStreamTransport.C)

	s.SupportedDataStreamTransportConfiguration = characteristic.NewSupportedDataStreamTransportConfiguration()
	s.AddC(s.SupportedDataStreamTransportConfiguration.C)

	return &s
}
