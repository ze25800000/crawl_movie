package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

import (
	"crawl_movie/models"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	var movieInfo models.MovieInfo
	//连接redis
	models.ConnectRedis("127.0.0.1:6379")
	//爬虫入口
	sUrl := `https://movie.douban.com`
	models.PutinQueue(sUrl)

	for {
		length := models.GetQueueLength()
		if length == 0 {
			break
		}
		sUrl = models.PopfromQueue()
		//判断sUrl是否被访问过
		if models.IsVisit(sUrl) {
			continue
		}

		rsp := httplib.Get(sUrl)
		sMovieHtml, err := rsp.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		// 记录电影信息
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharactors(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)

			models.AddMovie(&movieInfo)
		}

		//	提取该页面的索引连接
		urls := models.GetMovieUrls(sMovieHtml)

		for _, url := range urls {
			models.PutinQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}
		//sUrl 应当记录到访问队列set中
		models.AddToSet(sUrl)

		time.Sleep(time.Second)
	}
	c.Ctx.WriteString("end of crawl")
}
