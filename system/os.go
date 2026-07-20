package system

import (
	"os/exec"
)

type (
	Option func(*exec.Cmd)
)

func IsWindows() bool { _ = "STUB: not implemented"; return false }

func IsLinux() bool { _ = "STUB: not implemented"; return false }

func IsMac() bool { _ = "STUB: not implemented"; return false }

func GetOsEnv(key string) string { _ = "STUB: not implemented"; return "" }

func SetOsEnv(key, value string) error { _ = "STUB: not implemented"; return nil }

func RemoveOsEnv(key string) error { _ = "STUB: not implemented"; return nil }

func CompareOsEnv(key, comparedEnv string) bool { _ = "STUB: not implemented"; return false }

func ExecCommand(command string, opts ...Option) (stdout, stderr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func byteToString(data []byte, charset string) string { _ = "STUB: not implemented"; return "" }

func GetOsBits() int { _ = "STUB: not implemented"; return 0 }

func StartProcess(command string, args ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func StopProcess(pid int) error { _ = "STUB: not implemented"; return nil }

func KillProcess(pid int) error { _ = "STUB: not implemented"; return nil }

type ProcessInfo struct {
	PID                int
	CPU                string
	Memory             string
	State              string
	User               string
	Cmd                string
	Threads            []string
	IOStats            string
	StartTime          string
	ParentPID          int
	NetworkConnections string
}

func GetProcessInfo(pid int) (*ProcessInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func parseProcessInfo(output []byte, pid int) (*ProcessInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getThreadsInfo(pid int) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getIOStats(pid int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getProcessStartTime(pid int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getParentProcess(pid int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func getNetworkConnections(pid int) (string, error) { _ = "STUB: not implemented"; return "", nil }
