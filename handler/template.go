package handler

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func replaceConfig(path string, config DevOps) string {
	if b, err := ioutil.ReadFile(path); err != nil {
		return ""
	} else {
		content := string(b)
		content = strings.ReplaceAll(content, "${image-group}", config.ImageGroup)
		content = strings.ReplaceAll(content, "${image-name}", config.Image)
		content = strings.ReplaceAll(content, "${git-group}", config.GitGroup)
		content = strings.ReplaceAll(content, "${git-name}", config.GitName)
		content = strings.ReplaceAll(content, "${ns-name}", config.Project)
		content = strings.ReplaceAll(content, "${app-name}", config.App)
		content = strings.ReplaceAll(content, "${deploy-name}", config.Workload)
		content = strings.ReplaceAll(content, "${deploy-desc}", config.Desc)
		content = strings.ReplaceAll(content, "${config-key}", config.Workload)
		content = strings.ReplaceAll(content, "${ns-name-prod}", strings.ReplaceAll(config.Project, "-test", "-prod"))

		return content
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func pathExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func copy(source, dest string, config DevOps) {
	if rd, err := ioutil.ReadDir(source); err != nil {

	} else {
		for _, fi := range rd {
			source_path := path.Join(source, fi.Name())
			if !pathExists(dest) {
				os.MkdirAll(dest, 0666)
			}
			dest_path := path.Join(dest, fi.Name())
			if fi.IsDir() {
				copy(source_path, dest_path, config)
			} else {
				content := replaceConfig(source_path, config)
				// log.Printf("%s\n%s", source_path, content)
				if f, err := os.Create(dest_path); err != nil {
					log.Printf("create file error:%s", err.Error())
				} else {
					_, err := io.WriteString(f, content)
					if err != nil {
						log.Printf("write file error:%s", err.Error())
					}
					defer f.Close()
				}
			}
		}
	}
}

func zipTemplate(source, dest string, config DevOps) (string, error) {
	path := fmt.Sprintf("%s.zip", dest)
	log.Printf("zip path: %s", path)
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	err = zipFile(source, "", zipWriter, config)
	if err != nil {
		log.Printf("zip template err:%s", err.Error())
		return "", err
	}
	return path, nil
}

func zipFile(source, prefix string, writer *zip.Writer, config DevOps) error {
	file, err := os.Open(source)
	if err != nil {
		return nil
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		files, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range files {
			source_path := path.Join(source, fi.Name())
			prefix_path := path.Join(prefix, fi.Name())
			zipFile(source_path, prefix_path, writer, config)
		}
	} else {
		content := replaceConfig(source, config)
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = prefix
		writer, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		writer.Write([]byte(content))
	}
	return nil
}
