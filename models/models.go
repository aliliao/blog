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
	_DB_NAME    = "data/blog.db"
	_DB_DRIVER  = "sqlite3"
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
	Id             int64
	Uid            int64
	UserName       string
	NickName       string
	Pwd            string
	IsAdmin        bool   // 是否管理员
	LoginIpList    string // 登入IP列表 逗号分隔
	TopicIdList    string // 文章ID列表 逗号分隔
	VideoIdList    string // 视频ID列表 逗号分隔
	MusicIdList    string // 音乐ID列表 逗号分隔
	ImageIdList    string // 图片ID列表 逗号分隔
	CategoryIdList string // 分类ID列表 逗号分隔
}

// 评论
type Comment struct {
	Id         int64
	Tid        int64  // topic id
	Uid        int64  // 注册用户的user id
	Content    string // 评论内容
	LikesCount int64  // 点赞数

}

func registerDB() {
	if _, err := os.Stat(_DB_NAME); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(User))
	orm.RegisterDriver(_DB_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _DB_DRIVER, _DB_NAME)
}
