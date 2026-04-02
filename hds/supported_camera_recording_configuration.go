package hds

// DefaultLength is the default prebuffer/fragment length in milliseconds.
const DefaultLength = 4000

const EventTriggerMotion uint64 = 0x01
const EventTriggerDoorbell uint64 = 0x02

const MediaContainerTypeMP4 = 0

type SupportedCameraRecordingConfiguration struct {
	PreBufferLength             uint32                      `tlv8:"1"` // 4 bytes LE per HAP spec
	EventTrigger                uint64                      `tlv8:"2"` // 8 bytes LE per HAP spec
	MediaContainerConfiguration MediaContainerConfiguration `tlv8:"3"`
}

type MediaContainerConfiguration struct {
	MediaContainerType       uint8                    `tlv8:"1"`
	MediaContainerParameters MediaContainerParameters `tlv8:"2"`
}

type MediaContainerParameters struct {
	FragmentLength uint32 `tlv8:"1"` // 4 bytes LE per HAP spec
}

func NewSupportedCameraRecordingConfiguration() SupportedCameraRecordingConfiguration {
	return SupportedCameraRecordingConfiguration{
		PreBufferLength: DefaultLength,
		EventTrigger:    EventTriggerMotion,
		MediaContainerConfiguration: MediaContainerConfiguration{
			MediaContainerType: MediaContainerTypeMP4,
			MediaContainerParameters: MediaContainerParameters{
				FragmentLength: DefaultLength,
			},
		},
	}
}
