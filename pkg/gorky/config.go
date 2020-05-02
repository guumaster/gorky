package gorky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/guumaster/gorky/pkg/service"
	"github.com/guumaster/gorky/pkg/xdg"
)

func MakeConfig() *service.Config {
	appDirs, err := xdg.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &service.Config{
		Binary:      "gorky",
		DisplayName: "gorky",
		Description: "Gorky Wallpaper",
		Xdg:         appDirs,
	}

	u, err := GetCurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	cfg.User = u

	cfgDir := appDirs.ConfigHome()
	f, _ := os.Open(path.Join(cfgDir, "env.json"))
	b, _ := ioutil.ReadAll(f)

	var envs map[string]string
	_ = json.Unmarshal(b, &envs)

	for k, v := range envs {
		cfg.Env = append(cfg.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return cfg
}
