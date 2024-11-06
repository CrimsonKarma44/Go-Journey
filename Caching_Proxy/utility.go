package Caching_Proxy

import (
	"fmt"
	"os"
)

func ClearCache(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(fmt.Sprintf("%s/%s", dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}
