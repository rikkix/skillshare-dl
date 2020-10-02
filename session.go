package skillshare_dl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Session struct {
	ID             int    `json:"id"`
	ParentClassSku int    `json:"parent_class_sku"`
	Title          string `json:"title"`
	VideoHashedID  string `json:"video_hashed_id"`
}

type videoSourceInfo struct {
	Sources []*Video `json:"sources"`
}

type Video struct {
	Codecs      string `json:"codecs,omitempty"`
	ExtXVersion string `json:"ext_x_version,omitempty"`
	Src         string `json:"src"`
	Type        string `json:"type,omitempty"`
	Profiles    string `json:"profiles,omitempty"`
	AvgBitrate  int    `json:"avg_bitrate,omitempty"`
	Codec       string `json:"codec,omitempty"`
	Container   string `json:"container,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Height      int    `json:"height,omitempty"`
	Width       int    `json:"width,omitempty"`
}

func (s *Session) Video(account, pk string) ([]*Video, error) {
	vid := strings.Split(s.VideoHashedID, ":")[1]
	url := fmt.Sprintf("https://edge.api.brightcove.com/playback/v1/accounts/%s/videos/%s", account, vid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", fmt.Sprintf("application/json;pk=%s", pk))
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0")
	req.Header.Add("Origin", "https://www.skillshare.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vsi := &videoSourceInfo{}
	err = json.NewDecoder(resp.Body).Decode(vsi)
	if err != nil {
		return nil, err
	}

	return vsi.Sources, nil
}
