package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// TempFileName generates a temporary filename for use in testing or whatever
func TempFileName(extension string) (string, string) {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	fileName := hex.EncodeToString(randBytes) + "." + extension
	p, err := os.Getwd()
	check(err)
	return filepath.Join(p, fileName), fileName
}

// CreateTempFile creates a temp file with certain extension
func CreateTempFile(content, extension string) (string, error) {
	fullPath, fileName := TempFileName(extension)
	f, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	f.WriteString(content)
	return fileName, nil
}

func compileCode(req Request) ContainerOutput {

	fileName, err := CreateTempFile(req.Code, "c")
	check(err)

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	check(err)
	currentDir, err := os.Getwd()
	check(err)
	fmt.Println(fileName)

	// cmd := make([]string, 4)
	// switch req.LangID {
	// case 1:
	// 	cmd = append(cmd, "gcc", fileName, "&", "./a.out")
	// default:
	// 	cmd = append(cmd, "echo", "'Fuckboi'")
	// }

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      "weirdwiz/oj-micro",
		Tty:        true,
		WorkingDir: "/code",
		Cmd:        []string{"gcc", fileName},
	}, &container.HostConfig{Binds: []string{currentDir + "/temp/:/code"}}, nil, "")

	check(err)
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	check(err)
	_, err = cli.ContainerWait(ctx, resp.ID)
	check(err)
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	check(err)
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)

	return ContainerOutput{
		Output:  buf.String(),
		RunTime: 0, //TODO
		Memory:  0, //TODO
	}
}
