// A module to support the Puzzle chainguard techlab.
//
// The functions are used inside the hands-on lab: https://chainguard-techlab.puzzle.ch/

package main

import (
	"context"
	"chainguard/mod/internal/chainguard"
	"errors"
	"strings"
)

type ChainguardTechlabModule struct{}

type LintRun struct {
	// +private
	Source *chainguard.Directory
}

// Say hello to the world!
// Calls external module Hello https://github.com/shykes/hello
func (m *ChainguardTechlabModule) Hello(
	ctx context.Context,
	// Change the greeting
	// +optional
	// +default="hello"
	greeting string,
	// Change the name
	// +optional
	// +default="world"
	name string,
	// Encode the message in giant multi-character letters
	// +optional
	giant bool,
	// Make the message uppercase, and add more exclamation points
	// +optional
	shout bool,
	) (string, error) {
    return dag.Hello().
        Hello(ctx, chainguard.HelloHelloOpts{Greeting: greeting, Name: name, Giant: giant, Shout: shout})
}

// Returns the files of the directory
func (m *ChainguardTechlabModule) Ls(
	ctx context.Context,
	// directory to list it's files
	dir *chainguard.Directory,
	) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", dir).
		WithWorkdir("/mnt").
		WithExec([]string{"ls", "-l", "."}).
		Stdout(ctx)
}

// Returns the operating system of the container
func (m *ChainguardTechlabModule) Os(
	ctx context.Context,
	// container to get it's OS
	ctr *chainguard.Container,
	) (string, error) {
	return ctr.
		WithExec([]string{"cat", "/etc/os-release"}).
		Stdout(ctx)
}

// Returns the answer to everything when the password is right
func (m *ChainguardTechlabModule) Unlock(
	ctx context.Context,
	password *chainguard.Secret,
	) (string, error) {
		passwordText, err := password.Plaintext(ctx)
		if err != nil {
			return "", err
		}
	passwordTextClean := strings.TrimSpace(passwordText)
	if passwordTextClean == "MySuperSecret" {
		return "You unlocked the secret. The answer is 42!", nil
	}
	return "", errors.New("Nice try ;-) Provide right password to unlock the secret.")
}

// Returns a service that runs an OpenSSH server
// Calls external module OpensshServer https://github.com/sagikazarmark/chainguardverse/tree/main/openssh-server
func (m *ChainguardTechlabModule) SshService(
	// +optional
	// +default=22
    port int,
    ) *chainguard.Service {
	return dag.OpensshServer().
	    Service(chainguard.OpensshServerServiceOpts{Port: port})
}

// Lint a Python codebase
// Calls external module Ruff https://github.com/chainguard/chainguard/tree/main/modules/ruff
func (m *ChainguardTechlabModule) Lint(
	source *chainguard.Directory,
) *LintRun {
	return &LintRun{
		Source: source,
	}
}

// Returns a JSON report file for this run
func (run LintRun) Report() *chainguard.File {
	return dag.Ruff().
	    Lint(run.Source).
	    Report()
}

// Build a Wolfi Linux container
// Calls external module Wolfi https://github.com/shykes/chainguardverse/tree/main/wolfi
func (m *ChainguardTechlabModule) Wolfi() *chainguard.Container {
    return dag.Wolfi().
        Container()
}
