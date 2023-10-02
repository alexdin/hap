package service

import "github.com/alexdin/hap/characteristic"

const TypeCameraOperatingMode = "21A"

type CameraOperatingMode struct {
	*S
	EventSnapshotsActive    *characteristic.EventSnapshotsActive
	HomeKitCameraActive     *characteristic.HomeKitCameraActive
	PeriodicSnapshotsActive *characteristic.PeriodicSnapshotsActive
}

func NewCameraOperatingMode() *CameraOperatingMode {
	s := CameraOperatingMode{}
	s.S = New(TypeCameraOperatingMode)

	s.EventSnapshotsActive = characteristic.NewEventSnapshotsActive()
	s.AddC(s.EventSnapshotsActive.C)

	s.HomeKitCameraActive = characteristic.NewHomeKitCameraActive()
	s.AddC(s.HomeKitCameraActive.C)

	s.PeriodicSnapshotsActive = characteristic.NewPeriodicSnapshotsActive()
	s.AddC(s.PeriodicSnapshotsActive.C)

	return &s
}
