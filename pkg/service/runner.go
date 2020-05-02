package service

import (
	"errors"
	"fmt"
	"log"
	"os/user"
	"time"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/kardianos/service"
)

// Runner contains config to run binaries as system service
type Runner struct {
	manager *manager
	service service.Service
}

type Config struct {
	Binary, DisplayName, Description string

	User *user.User
	Xdg  *xdg.XDG

	Args []string
	Env  []string
}

func New(config *Config) (*Runner, error) {
	m, err := newManager(config)
	if err != nil {
		return nil, err
	}

	prg := &Runner{manager: m}

	err = prg.newService(config)
	if err != nil {
		return nil, err
	}

	err = prg.setupLogger()
	if err != nil {
		return nil, err
	}

	return prg, nil
}

func (r *Runner) ManageService(action string) error {
	err := service.Control(r.service, action)
	if err != nil {
		return fmt.Errorf("valid actions: %q\n%w", service.ControlAction, err)
	}

	return nil
}

func (r *Runner) RepeatAfter(d time.Duration) {
	ticker := time.NewTicker(d)

	go func() {
		r.manager.runChannel <- struct{}{}
		for range ticker.C {
			r.manager.runChannel <- struct{}{}
		}
	}()

	r.start()
}

func (r *Runner) newService(config *Config) error {
	s, err := service.New(r.manager, &service.Config{
		Name:             config.Binary,
		DisplayName:      config.DisplayName,
		Description:      config.Description,
		UserName:         config.User.Username,
		WorkingDirectory: config.User.HomeDir,
	})
	if err != nil {
		return err
	}

	r.service = s

	return nil
}

func (r *Runner) setupLogger() error {
	errs := make(chan error, 5)

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Println(err)
			}
		}
	}()

	l, err := r.service.Logger(errs)
	if err != nil {
		return err
	}

	r.manager.log = l

	return nil
}

func (r *Runner) start() {
	status, err := r.service.Status()
	if err != nil && !errors.Is(err, service.ErrNotInstalled) {
		return
	}

	// Only starts if is installed and stopped
	if status == service.StatusRunning {
		_ = r.service.Stop()
	}

	err = r.service.Run()
	if err != nil {
		_ = r.manager.log.Error(err)
		_ = r.service.Stop()

		return
	}
}
