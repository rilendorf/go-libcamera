package libcamera

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func TakeStill(duration time.Duration, w io.Writer) error {
	return take("libcamera-still", duration, w)
}

func TakeJPEG(duration time.Duration, w io.Writer) error {
	return take("libcamera-jpeg", duration, w)
}

func TakeRAW(duration time.Duration, w io.Writer) error {
	return take("libcamera-raw", duration, w)
}

func take(c string, duration time.Duration, w io.Writer) error {
	if duration == 0 {
		return InvalidRunDuration
	}

	cmd := exec.Command(c,
		fmt.Sprintf("-t%d", int(duration.Seconds())),
		"-o-",
	)

	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Waiting")
	// wait for finish
	_, err = cmd.Process.Wait()

	return nil
}
