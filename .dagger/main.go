package main

import (
	"context"
	"dagger/dagger-golang-binary-example/internal/dagger"
)

type DaggerGolangBinaryExample struct{}

func (m *DaggerGolangBinaryExample) BuildApp(source *dagger.Directory) *dagger.File {
	return dag.Go(dagger.GoOpts{
		Version: "1.23.1",
		Container: dag.
			Container().
			From("golang:1.23.1-alpine")},
	).Build(source)
}

func (m *DaggerGolangBinaryExample) Publish(ctx context.Context, source *dagger.Directory) (string, error) {
	app := m.BuildApp(source)
	return dag.Container().
		From("alpine:latest").
		WithFile("/usr/local/bin/app", app).
		WithRegistryAuth("ghcr.io", "username", dag.SetSecret("token", "actual-password")).
		Publish(ctx, "ghcr.io/rajatjindal/sample-image:v0.0.1")
}
