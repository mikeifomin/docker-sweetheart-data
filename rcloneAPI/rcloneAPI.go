package rcloneAPI

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	CMD    = "/bin/rclone"
	CONFIG = "~/.rclone.conf"
)

var PROVIDERS = [...]string{"B2", "S3", "DRIVE"}

type RClone struct {
	Cmd, Config string
}

func extractProviderConf(keyVal string) (provider, key, value string) {
	providersCheck := make(map[string]bool)
	for _, p := range PROVIDERS {
		providersCheck[p] = true
	}
	parts := strings.Split(keyVal, "_")
	if len(parts) > 0 && providersCheck[parts[0]] {
		provider = strings.ToLower(parts[0])
		//partsProvider := parts[1:]
		conf := strings.Join(parts[1:], "_")
		kv := strings.Split(conf, "=")
		key = strings.ToLower(kv[0])
		value = kv[1]
	}
	return
}

func parseConfig(lines []string) (result map[string][]string, err error) {
	for _, keyVal := range os.Environ() {
		provider, key, value := extractProviderConf(keyVal)
		if provider != "" {
			result[provider] = append(result[provider], key+" = "+value)
		}
	}
	return

}
func writeConfig(conf map[string][]string, filename string) error {
	out := []string{}
	for provider, lines := range conf {
		out = append(out, "["+provider+"]")
		for _, line := range lines {
			out = append(out, line)
		}
	}
	err := ioutil.WriteFile(filename, []byte(strings.Join(out, "\n")), 0644)
	return err
}
func run(args ...string) (result []byte, err error) {
	base := []string{"-c", CONFIG}
	all := append(base, args...)
	result, err = exec.Command(CMD, all...).Output()
	return
}
func EnsureConfig(envs []string) error {
	_, errR := ioutil.ReadFile(CONFIG)
	if errR == nil {
		log.Println("Config exists")
		return errR
	}
	conf, errP := parseConfig(envs)
	if errP != nil {
		return errP
	}
	errW := writeConfig(conf, CONFIG)
	if errW != nil {
		return errW
	}
	return nil
}
