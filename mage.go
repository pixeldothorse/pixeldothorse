// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// the version of Go to use in docker images.
	goVersion = "1.9.4"
)

// Generate runs all relevant code generation tasks.
func Generate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "statik", "-src", "./public", "-f")
	shouldWork(ctx, nil, filepath.Join(wd, "internal", "database", "migrations"), "go-bindata", "-pkg=dmigrations", "-o=../dmigrations/bindata.go", ".")

	fmt.Println("reran code generation")
}

// Travis runs initial setup needed for travis, then a full build and test cycle.
func Travis() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create database
	shouldWork(ctx, nil, wd, "psql", "-c", "CREATE DATABASE test;", "-U", "postgres")

	os.Setenv("DATABASE_URL", "postgres://postgres:hunter2@127.0.0.1/test?sslmode=disable")
	os.Setenv("PATH", os.Getenv("PATH")+":"+os.Getenv("HOME")+"/.local/bin")

	fmt.Println("[-] building horseville...")
	Generate()
	Build()

	fmt.Println("[-] testing horseville...")
	Test()
}

// Test runs horseville's test suite.
func Test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "go", "test", "-v", "./...")
}

// Build builds the command code into binaries in ./bin.
func Build() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	os.Mkdir("bin", 0777)

	outd := filepath.Join(wd, "bin")
	cmds := []string{
		"horsevilled",
	}

	for _, c := range cmds {
		shouldWork(ctx, nil, outd, "go", "build", "../cmd/"+c)
		fmt.Println("built ./bin/" + c)
	}
}

// Dep reruns `dep`.
func Dep() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "dep", "ensure", "-update")
	shouldWork(ctx, nil, wd, "dep", "prune")
}

// Docker creates the docker image xena/horseville with the horseville server.
func Docker() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "docker", "pull", "xena/alpine")
	shouldWork(ctx, nil, wd, "docker", "pull", "xena/go:"+goVersion)
	shouldWork(ctx, nil, wd, "docker", "build", "-t", "horseville/horsevilled", ".")
}

// Run starts an instance of horsevilled with default configuration and no
// authentication for debugging.
func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("building docker images")
	Docker()

	fmt.Println("Starting docker compose")

	defer shouldWork(ctx, nil, wd, "docker-compose", "down")
	shouldWork(ctx, nil, wd, "docker-compose", "up")
}
