package application

type VideoUseCase struct {
}

func NewVideoUseCase() *VideoUseCase {
	return &VideoUseCase{}
}

func (u *VideoUseCase) GetManifestPath(id string) (string, error) {
	return "", nil
}

func (u *VideoUseCase) GetSegmentPath(id, num string) (string, error) {
	return "", nil
}
