package utils

import (
	"github.com/o1egl/govatar"
	"myBlog/setting"
	"net/url"
	"strconv"
	"time"
)

const (
	imgStory = "./img/"
)

var (
	conf = setting.Config
)

func GetAvatarByStr(username string) (imgUrl string, err error) {
	imgName := "Avatar_" + username + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"
	if err = govatar.GenerateFileForUsername(govatar.MALE, username, imgStory+imgName); err == nil {
		imgUrl = "http://" + conf.Host + ":" + strconv.Itoa(int(conf.Port)) + "/api/utils/img?imgName=" + url.QueryEscape(imgName)
	}
	return
}
