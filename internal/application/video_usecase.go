package application

import "st-ember.github.com/streamingservice/internal/infra/ffmpeg"

type VideoUseCase struct {
	transcoder *ffmpeg.FFMPegTranscoder
}

func NewVideoUseCase(f *ffmpeg.FFMPegTranscoder) *VideoUseCase {
	return &VideoUseCase{transcoder: f}
}

func (u *VideoUseCase) GetManifestPath(id string) (string, error) {
	return "", nil
}

func (u *VideoUseCase) GetSegmentPath(id, num string) (string, error) {
	return "", nil
}

func (u *VideoUseCase) CreateResource() error {
	return nil
}
