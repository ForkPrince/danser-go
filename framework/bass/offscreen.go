package bass

/*
#include <stdint.h>
#include "bass.h"
#include "bassmix.h"
*/
import "C"
import (
	"unsafe"
)

func GetMixerRequiredBufferSize(seconds float64) int {
	if masterMixer == 0 {
		return int(float64(48000) * 2 * 4 * seconds)
	}
	return int(C.BASS_ChannelSeconds2Bytes(masterMixer, C.double(seconds)))
}

func ProcessMixer(buffer []byte) {
	if masterMixer == 0 {
		for i := range buffer {
			buffer[i] = 0
		}
		return
	}
	C.BASS_ChannelGetData(masterMixer, unsafe.Pointer(&buffer[0]), C.DWORD(len(buffer)))
}
