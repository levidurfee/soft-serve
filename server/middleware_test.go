package server

import (
	"os"
	"testing"

	sconfig "github.com/charmbracelet/soft-serve/config"
	"github.com/charmbracelet/soft-serve/internal/config"
	"github.com/charmbracelet/wish/testsession"
	"github.com/gliderlabs/ssh"
	"github.com/matryer/is"
)

var ()

func TestMiddleware(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll("testmiddleware")
	})
	is := is.New(t)
	appCfg, err := config.NewConfig(&sconfig.Config{
		Host:     "localhost",
		Port:     22223,
		RepoPath: "testmiddleware/repos",
		KeyPath:  "testmiddleware/key",
	})
	is.NoErr(err)
	_ = testsession.New(t, &ssh.Server{
		Handler: softMiddleware(appCfg)(func(s ssh.Session) {
			t.Run("TestCatConfig", func(t *testing.T) {
				_, err := s.Write([]byte("cat config/config.json"))
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			})
		}),
	}, nil)
}
