package models

import (
	"github.com/beego/beego/v2/core/logs"
	"os"
	"path"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME   = "data/blog.db"
	_DB_DRIVER = "sqlite3"
)

func init() {
	registerDB()

	mode, _ := config.String("runmode")
	if mode == "prod" {
		logs.Info("Run as prod mode...")
		orm.Debug = false
		orm.RunSyncdb("default", false, false)
	} else {
		logs.Info("Run as dev mode...")
		orm.Debug = true
		orm.RunSyncdb("default", false, true)
	}
}

// 分类
type Category struct {
	Id              int64
	Cid             int64
	Uid             int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Tid             int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	IsMarkdown      bool
	Category        string
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

// 用户
type User struct {
	Id       int64
	UserName string
	NickName string
	Pwd      string
	IsAdmin  bool   // 是否管理员()
	LoginIp  string // 登入IP
	//TopicIdList    string // 文章
	//VideoIdList    string // 视频
	//MusicIdList    string // 音乐
	//ImageIdList    string // 图片
	//CategoryIdList string // 分类
}

// 评论
type Comment struct {
	Id          int64
	Tid         int64
	NickName    string
	Content     string    `orm:"size(2000)"`
	LikesCount  int64     // 点赞数
	CommentTime time.Time `orm:"index"`
	IpAddr      string
	BrowserInfo string
	OsInfo      string
}

func registerDB() {
	if _, err := os.Stat(_DB_NAME); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(User), new(Comment))
	orm.RegisterDriver(_DB_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _DB_DRIVER, _DB_NAME)
}
