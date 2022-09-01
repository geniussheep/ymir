package pkg

import "fmt"

func CheckIntArgs(name string, value int) error {
	if value <= 0 {
		return fmt.Errorf("the args:%s is error", name)
	}
	return nil
}

func CheckStringArgs(name string, value string) error {
	if len(value) <= 0 {
		return fmt.Errorf("the args:%s is empty", name)
	}
	return nil
}

func CheckArrayArgs(name string, value []string) error {
	if value == nil || len(value) <= 0 {
		return fmt.Errorf("the args:%s is empty", name)
	}
	return nil
}
