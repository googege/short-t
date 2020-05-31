# short
> [中文](./README-cn.md)

URL shortening service.

## Usage
```go
package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/googege/short"
)

var (
	path    = "https://t.cn/"
	longURL = "https://github.com/googege/GOFamily/blob/master/%E5%9F%BA%E7%A1%80%E7%9F%A5%E8%AF%86/%E7%AE%97%E6%B3%95/%E7%AE%97%E6%B3%95%E9%A2%98/leetcode/1.md"
)

func main() {
	s := short.NewShort()
	db := short.NewMapDB()
	//
	shortURL, err := s.ShortAdd(longURL, db)
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println(shortURL)
	}
	//
	longURL, err = s.ShortFind(path+shortURL, db)
	// if http
	//http.Redirect(nil,nil,longURL,301)
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println("longURL:", longURL)
	}
	//
	shortURL, err = s.ShortFind("a", db)
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println("short: ", shortURL)
	}
	//
	if err = s.SetQR(path, 256, "/Users/googege/Desktop/test.png"); err != nil {
		glog.Error(err)
	}
}

```

## HERE
|items|description|
|:---:|:---:|
|Sponsor me|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/donate.png)|
|aliyun vps|[aliyun vps](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，China's largest cloud computing manufacturer, the world's top 5 cloud computing service providers|
|VPS|[vps](https://app.cloudcone.com/?ref=2525) fast and cheap just 2-3.75$/mo|
|WeChat public account|![p](https://raw.githubusercontent.com/googege/GOFamily/master/joinUsW.jpg)|
|My wechat (Please mark：form Github)|![p](https://raw.githubusercontent.com/googege/GOFamily/master/me.jpeg)|
|My bilibili|[b](https://space.bilibili.com/23170151)|
|MyYouTube|[YouTube](https://www.youtube.com/channel/UCM_-pFgD_HZDGD0yxfzguRQ?view_as=subscriber)|
