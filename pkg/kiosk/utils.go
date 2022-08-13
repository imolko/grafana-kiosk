package kiosk

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// GenerateURL constructs URL with appropriate parameters for kiosk mode
func GenerateURL(anURL string, kioskMode string, autoFit bool, isPlayList bool) string {
	u, _ := url.ParseRequestURI(anURL)
	q, _ := url.ParseQuery(u.RawQuery)

	switch kioskMode {
	case "tv": // TV
		q.Set("kiosk", "tv") // no sidebar, topnav without buttons
		log.Printf("KioskMode: TV")
	case "full": // FULLSCREEN
		q.Set("kiosk", "1") // sidebar and topnav always shown
		log.Printf("KioskMode: Fullscreen")
	case "disabled": // FULLSCREEN
		log.Printf("KioskMode: Disabled")
	default: // disabled
		q.Set("kiosk", "1") // sidebar and topnav always shown
		log.Printf("KioskMode: Fullscreen")
	}
	// a playlist should also go inactive immediately
	if isPlayList {
		q.Set("inactive", "1")
	}
	u.RawQuery = q.Encode()
	if autoFit {
		u.RawQuery += "&autofitpanels"
	}
	return u.String()
}

func CopyFolder(source, destination string) error {
	var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	return err
}
