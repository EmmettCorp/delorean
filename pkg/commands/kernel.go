package commands

import "syscall"

// KernelVersion returns current kernel version.
func KernelVersion() (string, error) {
	u := syscall.Utsname{}
	err := syscall.Uname(&u)
	if err != nil {
		return "", err
	}

	var byteString [65]byte
	var i int
	for ; u.Release[i] != 0; i++ {
		byteString[i] = uint8(u.Release[i])
	}

	return string(byteString[:i]), nil
}
