package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kennygrant/sanitize"
	"github.com/remeh/sizedwaitgroup"
	"github.com/sirupsen/logrus"

	ssdl "github.com/iochen/skillshare-dl"
)

type idList []int

func (l *idList) String() string {
	return "video id list"
}

func (l *idList) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*l = append(*l, v)
	return nil
}

func main() {
	var list idList
	cf := flag.String("cookie", "cookie.txt", "the file stored cookie")
	flag.Var(&list, "id", "video id")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*cf)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Video ID List:", list)

	for i := range list {
		fmt.Printf("=========== Downloading %d ===========\n", i)
		logrus.Error(Download(string(bytes), list[i]))
		fmt.Printf("=========== Video %d has been downloaded! ===========\n", i)
	}
}

func Download(cookie string, id int) error {
	dl := ssdl.NewDownloader(cookie)
	classInfo, err := dl.GetInfo(id)
	if err != nil {
		return err
	}

	sessions := classInfo.AllSessions()

	dir := sanitize.BaseName(classInfo.Title)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	wg := sizedwaitgroup.New(5)
	for i := range sessions {
		wg.Add()
		session := *sessions[i]
		go func(i int, session ssdl.Session) {
			defer wg.Done()
			video, err := session.Video("3695997568001", "BCpkADawqM2OOcM6njnM7hf9EaK6lIFlqiXB0iWjqGWUQjU7R8965xUvIQNqdQbnDTLz0IAO7E6Ir2rIbXJtFdzrGtitoee0n1XXRliD-RH9A-svuvNW9qgo3Bh34HEZjXjG4Nml4iyz3KqF")
			if err != nil {
				logrus.Error(err)
				return
			}

			title := sanitize.BaseName(strings.ReplaceAll(session.Title, "/", `-`))
			path := fmt.Sprintf("%s/%d.%s.mp4", dir, i, title)
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				fmt.Println("EXISTED:", path)
				return
			}

			err = FetchVideo(path, video)
			if err != nil {
				logrus.Error(err)
				return
			} else {
				fmt.Println("FINISHED:", path)
			}
		}(i, session)
	}
	wg.Wait()

	return nil
}

func FetchVideo(path string, video []*ssdl.Video) error {
	var url string
	for i := range video {
		if video[i].Container == "MP4" {
			url = video[i].Src
		}
	}
	if url == "" {
		return errors.New("cannot get available video src")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
