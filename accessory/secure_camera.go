package accessory

import (
	"github.com/brutella/hap/service"
)

// SecureCamera provides RTP video streaming.
type SecureCamera struct {
	*A
	CameraOperatingMode       *service.CameraOperatingMode
	StreamManagement          *service.CameraRTPStreamManagement
	CameraRecordingManagement *service.CameraRecordingManagement
}

// NewSecureCamera returns an IP camera accessory.
func NewSecureCamera(info Info) *SecureCamera {
	a := SecureCamera{}
	a.A = New(info, TypeIPCamera)

	a.CameraOperatingMode = service.NewCameraOperatingMode()
	a.AddS(a.CameraOperatingMode.S)

	a.StreamManagement = service.NewCameraRTPStreamManagement()
	a.AddS(a.StreamManagement.S)

	a.CameraRecordingManagement = service.NewCameraRecordingManagement()
	a.AddS(a.CameraRecordingManagement.S)

	return &a
}