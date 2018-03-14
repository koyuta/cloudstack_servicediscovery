package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xanzy/go-cloudstack/cloudstack"
)

type fileSDConfigs []fileSDConfig

type fileSDConfig struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

func run() error {
	client := cloudstack.NewAsyncClient(*endpoint, *apiKey, *secretKey, false)
	service := cloudstack.NewVirtualMachineService(client)

	params := &cloudstack.ListVirtualMachinesParams{}
	res, err := service.ListVirtualMachines(params)
	if err != nil {
		return err
	}

	targetGroups := strings.Split(*groups, ",")

	targets := []string{}
	for _, m := range res.VirtualMachines {
		for _, g := range targetGroups {
			if m.Group == g {
				targets = append(targets, fmt.Sprintf("%s:%d", m.Nic[0].Ipaddress, *port))
			}
		}
	}
	config := fileSDConfig{Targets: targets, Labels: map[string]string{}}

	if *labels != "" {
		for _, label := range strings.Split(*labels, ",") {
			l := strings.Split(label, ":")
			config.Labels[l[0]] = l[1]
		}
	}

	data, err := json.Marshal(fileSDConfigs{config})
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "  "); err != nil {
		return err
	}

	return writeFileAtomic(*filename, out.Bytes())
}

func writeFileAtomic(filename string, data []byte) error {
	tmpFilename := fmt.Sprintf("%s.%s", filename, "new")
	if err := ioutil.WriteFile(tmpFilename, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmpFilename, filename)
}
