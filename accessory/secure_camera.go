package accessory

import (
	"github.com/brutella/hap/service"
)

// SecureCamera provides RTP video streaming.
type SecureCamera struct {
	*A
	CameraOperatingMode           *service.CameraOperatingMode
	StreamManagement              *service.CameraRTPStreamManagement
	CameraRecordingManagement     *service.CameraRecordingManagement
	MotionSensor                  *service.MotionSensor
	DataStreamTransportManagement *service.DataStreamTransportManagement
}

// NewSecureCamera returns an IP camera accessory.
func NewSecureCamera(info Info) *SecureCamera {
	a := SecureCamera{}
	a.A = New(info, TypeIPCamera)

	// MotionSensor and DataStreamTransportManagement must be in a.Ss so they get
	// IIDs assigned, but are only exposed via linked references (not primary).
	a.MotionSensor = service.NewMotionSensor()
	a.AddS(a.MotionSensor.S)

	a.DataStreamTransportManagement = service.NewDataStreamTransportManagement()
	a.AddS(a.DataStreamTransportManagement.S)

	a.CameraRecordingManagement = service.NewCameraRecordingManagement()
	a.CameraRecordingManagement.AddS(a.MotionSensor.S)
	a.CameraRecordingManagement.AddS(a.DataStreamTransportManagement.S)
	a.AddS(a.CameraRecordingManagement.S)

	a.StreamManagement = service.NewCameraRTPStreamManagement()
	a.AddS(a.StreamManagement.S)

	// CameraOperatingMode is the primary service and links to the two main
	// camera services so the Home app treats this as a single HKSV accessory.
	a.CameraOperatingMode = service.NewCameraOperatingMode()
	a.CameraOperatingMode.S.Primary = true
	a.CameraOperatingMode.AddS(a.StreamManagement.S)
	a.CameraOperatingMode.AddS(a.CameraRecordingManagement.S)
	a.AddS(a.CameraOperatingMode.S)

	return &a
}
