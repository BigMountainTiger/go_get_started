package browser

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"song.com/go_get_started/launch_chromium/utils"
)

const browserd string = "chromium"
const appstats string = ".appstats"

type Chromium struct {
	App_built_time string
	Embeded        []byte
	cmd            *exec.Cmd
}

func (browser *Chromium) deployment_required() bool {

	fi, err := utils.File_info_in_home(browserd)
	if err != nil || !fi.IsDir() {
		return true
	}

	b, err := utils.Read_file_from_home(appstats)
	return err != nil || string(b) != browser.App_built_time
}

func (browser *Chromium) deploy() error {
	// https://golang.cafe/blog/golang-unzip-file-example.html

	utils.Remove_directory_from_home(browserd)

	dst, _ := utils.Home_path()
	embeded := browser.Embeded[:]
	archive, err := zip.NewReader(bytes.NewReader(embeded), int64(len(embeded)))
	if err != nil {
		return err
	}

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return errors.New("invalid file path in the zip file")
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}

			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		if err := dstFile.Close(); err != nil {
			return err
		}

		if fileInArchive.Close(); err != nil {
			return err
		}
	}

	utils.Save_file_to_home(appstats, []byte(browser.App_built_time))
	return nil
}

func (browser *Chromium) Launch(url string) error {

	if browser.deployment_required() {
		if err := browser.deploy(); err != nil {
			return err
		}
	}

	dst, _ := utils.Home_path()
	path := dst + "/" + browserd + "/chrome"
	if runtime.GOOS == "darwin" {
		path = dst + "/" + browserd + "/Chromium.app/Contents/MacOS/Chromium"
	}

	page := "-app=" + url
	cmd := exec.Command(path, page)
	browser.cmd = cmd

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (browser *Chromium) Wait_for_terminate() error {
	return browser.cmd.Wait()
}

func (browser *Chromium) Close() error {
	if browser.cmd == nil {
		return errors.New("no active browser instance to close")
	}

	return browser.cmd.Process.Kill()
}
