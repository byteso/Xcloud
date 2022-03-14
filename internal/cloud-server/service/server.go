package service

import (
	"fmt"
	"time"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func ServerInfo() (types.ServerInfo, error) {
	var response types.ServerInfo
	cpu, _ := cpu.Percent(time.Second, false)
	memory, _ := mem.VirtualMemory()
	disk, _ := disk.Usage("/")
	tem, _ := host.SensorsTemperatures()
	fmt.Println(memory.Total)
	fmt.Println(cpu)
	fmt.Println(disk.Total)

	fmt.Println(tem)

	response.CentralProcessingUnit.Percent = cpu[0]
	response.Memory.Total = memory.Total
	response.Memory.Free = memory.Free
	response.Memory.UsedPercent = memory.UsedPercent
	response.Disk.Total = float64(disk.Total)
	response.Disk.Free = float64(disk.Free)
	return response, nil
}
