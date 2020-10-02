package skillshare_dl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Downloader struct {
	cookie string
}

func parseCookie(cookie string) string {
	cookie = strings.TrimSpace(cookie)
	cookie = strings.Trim(cookie, `"`)
	return cookie
}

func NewDownloader(cookie string) *Downloader {
	return &Downloader{
		cookie: parseCookie(cookie),
	}
}

func (dl *Downloader) Cookie(cookie string) {
	dl.cookie = parseCookie(cookie)
}

func (dl *Downloader) GetInfo(id int) (*ClassInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.skillshare.com/classes/%d", id), nil)
	if err != nil {
		return &ClassInfo{}, err
	}
	req.Header.Add("Accept", "application/vnd.skillshare.class+json;,version=0.8")
	req.Header.Add("User-Agent", "Skillshare/5.3.0; Android 9.0.1")
	req.Header.Add("Referer", "https://www.skillshare.com/")
	req.Header.Add("Cookie", dl.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &ClassInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &ClassInfo{}, fmt.Errorf("cannot get video info, response with code %d", resp.StatusCode)
	}

	info := &ClassInfo{}

	err = json.NewDecoder(resp.Body).Decode(info)
	return info, err
}
