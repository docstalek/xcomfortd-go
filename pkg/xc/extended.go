package xc

import (
	"encoding/binary"
	"io"
)

type stickDplReader struct {
	i        *Interface
	position uint32
}

func (d *stickDplReader) Read(p []byte) (n int, err error) {
	var data []byte
	if d.position == 0 {
		if data, err = d.i.sendExtendedCommand([]byte{MCI_ET_REQU_DPL, 0, 0, 0, 0, 0, 0}); err != nil {
			return 0, err
		}
		if data[0] != MCI_ET_SEND_DPL {
			return 0, ErrUnexpectedReponse
		}
	} else {
		address := []byte{0, 0, 0, 0, 10, 0}
		binary.LittleEndian.PutUint32(address, d.position)
		if data, err = d.i.sendExtendedCommand(append([]byte{MCI_ET_RD}, address...)); err != nil {
			return 0, err
		}
		if data[0] != MCI_ET_REPLY {
			return 0, ErrUnexpectedReponse
		}
	}

	copied := copy(p, data[7:])
	d.position = binary.LittleEndian.Uint32(data[1:5]) + uint32(copied)

	return copied, nil
}

func (d *stickDplReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		d.position = uint32(offset)
	case io.SeekCurrent:
		d.position += uint32(offset)
	default:
	}
	return int64(d.position), nil
}

func (i *Interface) RequestDPL() error {
	i.extendedMutex.Lock()
	defer i.extendedMutex.Unlock()

	devs, dps, err := i.dplReader(&stickDplReader{i, 0})
	if err != nil {
		return err
	}

	i.setupChan <- datapoints{devs, dps}

	return nil
}
