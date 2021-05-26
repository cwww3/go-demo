package main

import "os"

func writeManyFiles(filePaths []string) error {
	for _, path := range filePaths {
		// 闭包 让defer能及时执行
		if err := func() error {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			// 知道writeManyFiles函数退出 defer才会执行
			defer f.Close()

			_, err = f.WriteString("content")
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}
	return nil
}
