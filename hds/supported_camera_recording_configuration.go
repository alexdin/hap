package hds

const DefaultLength uint16 = 4000

const EventTriggerMotion = 0x01
const EventTriggerDoorbell = 0x02

const MediaContainerTypeMP4 = 0

type SupportedCameraRecordingConfiguration struct {
	PreBufferLength             uint16                      `tlv8:"1"`
	EventTrigger                byte                        `tlv8:"2"`
	MediaContainerConfiguration MediaContainerConfiguration `tlv8:"3"`
}

type MediaContainerConfiguration struct {
	MediaContainerType       uint8                    `tlv8:"1"`
	MediaContainerParameters MediaContainerParameters `tlv8:"2"`
}

type MediaContainerParameters struct {
	FragmentLength uint16 `tlv8:"1"`
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
