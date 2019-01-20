package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Container holds data for docker container
type Container struct {
	Port string
	ID   string
}

// Settings Holds Port value printed from `docker inspect container_id`
type Settings struct {
	NetworkSettings struct {
		Ports struct {
			Two7017TCP []struct {
				HostIP   string `json:"HostIp"`
				HostPort string `json:"HostPort"`
			} `json:"27017/tcp"`
		} `json:"Ports"`
	} `json:"NetworkSettings"`
}

// SetContainer sets MONGO_HOST value
func SetContainer(c *Container) error {
	return os.Setenv("MONGO_HOST", fmt.Sprintf("mongodb://localhost:%s/api", c.Port))
}

// StartDB starts db container in docker
func StartDB() (*Container, error) {
	cmd := exec.Command("docker", "run", "-P", "-d", "mongo:3.4.5")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("starting mongo container: %v", err)
	}

	id := stdout.String()[:12]
	log.Println("DB Container Name:", id)

	cmd = exec.Command("docker", "inspect", id)
	stdout.Reset()
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("inspect container: %v", err)
	}

	var inspectPort []Settings
	if err := json.Unmarshal(stdout.Bytes(), &inspectPort); err != nil {
		return nil, fmt.Errorf("decoding json: %v", err)
	}

	c := Container{
		ID:   id,
		Port: inspectPort[0].NetworkSettings.Ports.Two7017TCP[0].HostPort,
	}

	log.Println("DB Port:", c.Port)

	return &c, nil
}

// StopDB stops and removes the specified container.
func StopDB(c *Container) error {
	if err := exec.Command("docker", "stop", c.ID).Run(); err != nil {
		return err
	}
	log.Println("Stopped:", c.ID)

	if err := exec.Command("docker", "rm", c.ID, "-v").Run(); err != nil {
		return err
	}
	log.Println("Removed:", c.ID)

	return nil
}
