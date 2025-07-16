package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
)

type FFMPegTranscoder struct{}

func NewTranscoder() *FFMPegTranscoder {
	return &FFMPegTranscoder{}
}

func (f *FFMPegTranscoder) Transcode(fileName string) error {
	cmdStr := fmt.Sprintf("ffmpeg -i upload/%s.mp4 -c:v libx264 -c:a aac -f hls -hls_time 10 -hls_segment_filename \"segments/%s_%%03d.ts\" segments/%s", fileName, fileName, fileName)
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println(string(output))
	return nil
}
