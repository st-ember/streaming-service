package application

import (
	"fmt"

	diskstorage "st-ember.github.com/streamingservice/internal/infra/disk_storage"
	"st-ember.github.com/streamingservice/internal/infra/ffmpeg"
)

type VideoUseCase struct {
	transcoder  *ffmpeg.FFMPegTranscoder
	diskstorage *diskstorage.DiskStorage
}

func NewVideoUseCase(f *ffmpeg.FFMPegTranscoder, s *diskstorage.DiskStorage) *VideoUseCase {
	return &VideoUseCase{
		transcoder:  f,
		diskstorage: s,
	}
}

func (u *VideoUseCase) GetManifestPath(id string) (string, error) {
	return "", nil
}

func (u *VideoUseCase) GetSegmentPath(id, num string) (string, error) {
	return "", nil
}

func (u *VideoUseCase) CreateResource(input UploadInput) error {
	storedName, err := u.diskstorage.Store(input.File, input.Extension)
	if err != nil {
		return fmt.Errorf("Error when storing file: %w", err)
	}

	if err = u.transcoder.Transcode(storedName); err != nil {
		return fmt.Errorf("Error transcoding with ffmpeg: %w", err)
	}

	return nil
}
