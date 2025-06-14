package infra

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

func OpenDefaultApp(filePath string) {
	cmd, err := evaluateOSSpecificCommand(filePath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err != nil {
		cobra.CheckErr(err)
	}

	runErr := cmd.Run()

	if runErr != nil {
		cobra.CheckErr(runErr)
	}
}

func evaluateOSSpecificCommand(filePath string) (*exec.Cmd, error) {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", filePath), nil
	case "linux":
		editor, err := getLinuxDefaultEditor()

		if err != nil {
			cobra.CheckErr(err)
		}

		return exec.Command(editor, filePath), nil
	case "windows":
		return exec.Command("start", filePath), nil
	default:
		return nil, errors.New("unsupported operational system")
	}
}

func getLinuxDefaultEditor() (string, error) {
	editor := retrieveEditorFromEnv()

	if editor != "" {
		return editor, nil
	}

	return "", errors.New("no default editor found")
}

func retrieveEditorFromEnv() string {
	if editor := os.Getenv("VISUAL"); editor != "" {
		return editor
	}

	return os.Getenv("EDITOR")
}
