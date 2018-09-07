package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
	Create_time          string
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/movie?charset=utf8")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}
func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}
func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*rel="v:directedBy">(.*)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}
func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}

func GetMovieMainCharactors(movieHtml string) string {
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	mainCaracters := ""
	for _, v := range result {
		mainCaracters += v[1] + "/"
	}
	return mainCaracters
}
