package rtp

type StreamConfiguration struct {
	Command SessionControlCommand `tlv8:"1"`
	Video   VideoParameters       `tlv8:"2"`
	Audio   AudioParameters       `tlv8:"3"`
}

const (
	SessionControlCommandTypeEnd         byte = 0
	SessionControlCommandTypeStart            = 1
	SessionControlCommandTypeSuspend          = 2
	SessionControlCommandTypeResume           = 3
	SessionControlCommandTypeReconfigure      = 4
)

type SessionControlCommand struct {
	Identifier []byte `tlv8:"1"` // docu says 16 bytes but overall size of SessionControlCommand is also 16 bytes?
	Type       byte   `tlv8:"2"`
}

type VideoParameters struct {
	CodecType   byte                 `tlv8:"1,optional"`
	CodecParams VideoCodecParameters `tlv8:"2,optional"`
	Attributes  VideoCodecAttributes `tlv8:"3,optional"`
	RTP         RTPParams            `tlv8:"4,optional"`
}

type AudioParameters struct {
	CodecType    byte                 `tlv8:"1,optional"`
	CodecParams  AudioCodecParameters `tlv8:"2,optional"`
	RTP          RTPParams            `tlv8:"3,optional"`
	ComfortNoise bool                 `tlv8:"4,optional"`
}

type RTPParams struct {
	PayloadType             uint8   `tlv8:"1,optional"`
	Ssrc                    uint32  `tlv8:"2,optional"`
	Bitrate                 uint16  `tlv8:"3,optional"`
	Interval                float32 `tlv8:"4,optional"` // MinimumRTCP interval
	ComfortNoisePayloadType uint8   `tlv8:"5,optional"` // only for audio
	MTU                     uint16  `tlv8:"6,optional"` // only for video
}
