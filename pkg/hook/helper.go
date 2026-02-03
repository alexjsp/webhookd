package hook

import (
	"errors"
	"os"
	"path"
	"strings"
)

// ResolveScript is resolving the target script.
func ResolveScript(dir, name, defaultExt string) (string, error) {
	scriptPath := path.Join(dir, name)
	// Validate script path
	scriptPath = path.Clean(scriptPath)
	if !strings.HasPrefix(scriptPath, dir) {
		return "", errors.New("invalid script path: " + name)
	}
	// Check if the script exists
	if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
		// Try to add the default extension if not provided
		if path.Ext(name) == "" {
			scriptPathWithExt := scriptPath + "." + defaultExt
			if _, err := os.Stat(scriptPathWithExt); errors.Is(err, os.ErrNotExist) {
				return "", errors.New("script not found: " + scriptPath + "[." + defaultExt + "]")
			} else {
				return scriptPathWithExt, nil
			}
		}
		return "", errors.New("script not found: " + scriptPath)
	} else {
		return scriptPath, nil
	}
}
