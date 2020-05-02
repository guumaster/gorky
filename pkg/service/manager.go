package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/kardianos/service"
	"gopkg.in/natefinch/lumberjack.v2"
)

type manager struct {
	config     *Config
	execPath   string
	cmd        *exec.Cmd
	log        service.Logger
	out        string
	runChannel chan struct{}
}

func newManager(config *Config) (*manager, error) {
	execPath, err := exec.LookPath(config.Binary)
	if err != nil {
		return nil, fmt.Errorf("failed to find executable %q: %w", config.Binary, err)
	}

	outputFile := path.Join(config.Xdg.CacheHome(), fmt.Sprintf("%s.log", config.Binary))

	return &manager{
		execPath:   execPath,
		config:     config,
		out:        outputFile,
		runChannel: make(chan struct{}),
	}, nil
}

func (s *manager) Start(ss service.Service) error {
	_ = s.log.Info("Starting ", s.config.DisplayName)
	defer s.Stop(ss) // nolint:errcheck

	s.forceStop()

	for range s.runChannel {
		s.run()
	}

	return nil
}

func (s *manager) Stop(_ service.Service) error {
	_ = s.log.Info("Stopping ", s.config.DisplayName)
	s.forceStop()

	return nil
}

func (s *manager) run() {
	_ = s.log.Info("Running ", s.config.DisplayName)
	cmd := exec.Command(s.execPath, s.config.Args...) // nolint:gosec

	log.SetOutput(&lumberjack.Logger{
		Filename:   s.out,
		MaxSize:    1, // mb
		MaxBackups: 7,
		MaxAge:     1, // days
		Compress:   true,
	})

	cmd.Env = append(os.Environ(), s.config.Env...)
	cmd.Dir = s.config.User.HomeDir
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err := cmd.Run()
	if err != nil {
		_ = s.log.Error("running error: ", err)
	}
}

func (s *manager) forceStop() {
	if s.cmd != nil && s.cmd.ProcessState != nil && !s.cmd.ProcessState.Exited() {
		_ = s.cmd.Process.Kill()
	}
}
