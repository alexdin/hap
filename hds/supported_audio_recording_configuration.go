package hds

import "github.com/brutella/hap/rtp"

type SupportedAudioRecordingConfiguration struct {
	Codecs []rtp.AudioCodecConfiguration `tlv8:"1"`
}

func NewSupportedAudioRecordingConfiguration() SupportedAudioRecordingConfiguration {
	return SupportedAudioRecordingConfiguration{
		Codecs: []rtp.AudioCodecConfiguration{
			rtp.NewOpusAudioCodecConfiguration(),
			rtp.NewAacEldAudioCodecConfiguration(),
		},
	}
}
