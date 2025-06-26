package ffmpeg

type FFMPegTranscoder struct{}

func NewTranscoder() *FFMPegTranscoder {
	return &FFMPegTranscoder{}
}

func (f *FFMPegTranscoder) Transcode(fileName string) error {
	return nil
}
