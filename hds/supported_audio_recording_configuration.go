package hds

import "github.com/brutella/hap/rtp"

// Recording audio codec types (HKSV spec, different from RTP streaming codec types).
const (
	RecordingAudioCodecType_AAC_LC  byte = 0
	RecordingAudioCodecType_AAC_ELD byte = 1
)

type SupportedAudioRecordingConfiguration struct {
	Codecs []rtp.AudioCodecConfiguration `tlv8:"1"`
}

func NewSupportedAudioRecordingConfiguration() SupportedAudioRecordingConfiguration {
	return SupportedAudioRecordingConfiguration{
		Codecs: []rtp.AudioCodecConfiguration{
			// HKSV recording requires AAC-LC at 32kHz (not AAC-ELD or Opus)
			{
				Type: RecordingAudioCodecType_AAC_LC,
				Parameters: rtp.AudioCodecParameters{
					Channels:   1,
					Bitrate:    rtp.AudioCodecBitrateVariable,
					Samplerate: rtp.AudioCodecSampleRate32Khz,
				},
			},
		},
	}
}
