// +build !windows

package freeport

import "golang.org/x/sys/unix"

func systemLimit() (uint64, error) {
	var limit unix.Rlimit
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &limit)
	return limit.Cur, err
}
