package hds

import "github.com/brutella/hap/rtp"

const Bitrate uint16 = 2000

type SupportedVideoRecordingConfiguration struct {
	Configurations VideoCodecConfiguration `tlv8:"1"`
}

type VideoCodecConfiguration struct {
	CodecType   byte                       `tlv8:"1"`
	CodecParams VideoCodecParameters       `tlv8:"2"`
	Attributes  []rtp.VideoCodecAttributes `tlv8:"3"`
}

type VideoCodecParameters struct {
	ProfileID byte   `tlv8:"1"`
	Level     byte   `tlv8:"2"`
	Bitrate   uint16 `tlv8:"3"`
	//iFrameInterval uint16 `tlv8:"4"`
}

func DefaultSupportedVideoRecordingConfiguration() SupportedVideoRecordingConfiguration {
	return SupportedVideoRecordingConfiguration{
		Configurations: NewVideoCodecConfiguration(),
	}
}

func NewVideoCodecConfiguration() VideoCodecConfiguration {
	return VideoCodecConfiguration{
		CodecType: rtp.VideoCodecType_H264,
		CodecParams: VideoCodecParameters{
			ProfileID: rtp.VideoCodecProfileHigh,
			Level:     rtp.VideoCodecLevel4,
			Bitrate:   Bitrate,
			//iFrameInterval: DefaultLength,
		},
		Attributes: []rtp.VideoCodecAttributes{
			{1920, 1080, 30}, // 1080p
			{1280, 720, 30},  // 720p
			{640, 360, 30},
			{480, 270, 30},
			{320, 180, 30},
			{1280, 960, 30},
			{1024, 768, 30}, // XVGA
			{640, 480, 30},  // VGA
			{480, 360, 30},  // HVGA
			{320, 240, 15},  // QVGA (Apple Watch)
		},
	}
}
