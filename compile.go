package main

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func compileCode(req Request) ContainerOutput {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      "golang",
		Tty:        true,
		WorkingDir: "/code",
		Cmd:        []string{"go", "run", "1.go"},
	}, &container.HostConfig{Binds: []string{"/home/weirdwiz/go/src/github.com/weirdwiz/rand:/code/"}}, nil, "")
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	_, err = cli.ContainerWait(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)

	return ContainerOutput{
		Output:  buf.String(),
		RunTime: 0, //TODO
		Memory:  0, //TODO
	}
}
