package hds

import "github.com/brutella/hap/rtp"

type SupportedAudioRecordingConfiguration struct {
	Codecs []rtp.AudioCodecConfiguration `tlv8:"1"`
}

func NewSupportedAudioRecordingConfiguration() SupportedAudioRecordingConfiguration {
	return SupportedAudioRecordingConfiguration{
		Codecs: []rtp.AudioCodecConfiguration{
			// HKSV recording requires AAC-LC at 32kHz (not AAC-ELD or Opus)
			{
				Type: rtp.AudioCodecType_AAC_ELD,
				Parameters: rtp.AudioCodecParameters{
					Channels:   1,
					Bitrate:    rtp.AudioCodecBitrateVariable,
					Samplerate: rtp.AudioCodecSampleRate32Khz,
				},
			},
		},
	}
}
