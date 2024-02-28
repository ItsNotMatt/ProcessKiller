package main

import (
	"fmt"
	"log"
    "github.com/shirou/gopsutil/v3/process"
)

func main() {
    disallowList := []string{"steam.exe", "nvcontainer.exe", "NVIDIA Share.exe", "NVDisplay.Container.exe", "python.exe", "Unity Hub.exe", "msedge.exe", "nordvpn-service.exe", "DriverSupport.exe", "lghub_updater.exe", "updatechecker.exe", "NordVPN", "Driver Support"};

    processes, err := process.Processes()
    if err != nil {
        log.Fatal("Unable to get running processes")
    }

    for _, process := range processes {
        name, _ := process.Name()
        if err == nil {
            fmt.Println("Process: ", name)
            for _, entry := range disallowList {
                if name == entry {
                    fmt.Println("Killing process: ", name)
                    process.Kill()
                }
            }
        }
    }
}
