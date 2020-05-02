package gorky

import (
	"encoding/json"
	"os"
	"os/user"
	"path"
	"runtime"

	"github.com/guumaster/gorky/pkg/service"
)

func GetCurrentUser() (*user.User, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	if runtime.GOOS == "windows" {
		return u, nil
	}

	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser != "" {
		u, err = user.Lookup(sudoUser)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func CreateConfig(config *service.Config) error {
	keys := []string{
		"DESKTOP_SESSION",
		"XDG_CURRENT_DESKTOP",
		"XDG_RUNTIME_DIR",
	}

	f, err := os.Create(path.Join(config.Xdg.ConfigHome(), "env.json"))
	if err != nil {
		return err
	}

	s := map[string]string{}
	for _, key := range keys {
		s[key] = os.Getenv(key)
	}

	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
