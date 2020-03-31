package xc

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

// Interface
type Interface struct {
	datapoints map[byte]*Datapoint
	devices    map[int]*Device

	// tx command queue
	txCommandQueue chan request
	txSemaphore    *semaphore.Weighted

	// config command queue
	configCommandQueue chan request
	configMutex        sync.Mutex

	handler Handler
}

type request struct {
	command    []byte
	responseCh chan []byte
}

// Handler interface for receiving callbacks
type Handler interface {
	// Datapoint sent value
	StatusValue(datapoint *Datapoint, value int)
	// Datapoint switched on/off
	StatusBool(datapoint *Datapoint, on bool)
}

// Device returns the device with the specified serialNumber
func (i *Interface) Device(serialNumber int) *Device {
	return i.devices[serialNumber]
}

// Datapoint returns the requested datapoint
func (i *Interface) Datapoint(number int) *Datapoint {
	return i.datapoints[byte(number)]
}

// Init loads datapoints from the specified file and takes a handler which
// will get callbacks when events are received.
func (i *Interface) Init(filename string, handler Handler) error {
	i.datapoints = make(map[byte]*Datapoint)
	i.devices = make(map[int]*Device)
	i.handler = handler

	// Only allow four tx commands in parallel
	i.txSemaphore = semaphore.NewWeighted(4)
	i.txCommandQueue = make(chan request)

	i.configCommandQueue = make(chan request)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	r.Comma = '\t'
	r.FieldsPerRecord = 9

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		serialNo, err := strconv.Atoi(record[2])
		if err != nil {
			return err
		}
		datapoint, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		deviceType, err := strconv.Atoi(record[3])
		if err != nil {
			return err
		}
		channel, err := strconv.Atoi(record[4])
		if err != nil {
			return err
		}

		device, exists := i.devices[serialNo]
		if !exists {
			device = &Device{
				serialNumber: serialNo,
				deviceType:   DeviceType(deviceType),
				iface:        i,
			}
			i.devices[serialNo] = device
		}

		dp := &Datapoint{
			device:  device,
			name:    strings.Join(strings.Fields(strings.TrimSpace(record[1])), " "),
			number:  byte(datapoint),
			channel: channel,
		}
		device.Datapoints = append(device.Datapoints, dp)
		i.datapoints[byte(datapoint)] = dp
	}

	return nil
}