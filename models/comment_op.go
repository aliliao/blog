package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

func AddComment(tid, nickName, content, browserInfo, osInfo, ip string) error {
	if len(content) == 0 || len(nickName) == 0 {
		return ErrParamIsEmpty
	}

	o := orm.NewOrm()
	topicId, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	comment := &Comment{
		Tid:         topicId,
		NickName:    nickName,
		Content:     content,
		BrowserInfo: browserInfo,
		OsInfo:      osInfo,
		IpAddr:      ip,
		CommentTime: time.Now(),
	}

	_, err = o.Insert(comment)
	return err
}

func DelAComment(id string) error {
	o := orm.NewOrm()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	comment := &Comment{Id: cid}
	_, err = o.Delete(comment)

	return err
}

func FetchAllComments(tid string) ([]*Comment, error) {
	o := orm.NewOrm()
	topicId, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	comments := make([]*Comment, 0)
	_, err = o.QueryTable("comment").Filter("tid", topicId).All(&comments)
	return comments, err
}

func ThumbsUp(tid string) error {
	o := orm.NewOrm()
	topicId, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	comment := &Comment{Tid: topicId}
	err = o.QueryTable("comment").One(comment)
	if err != nil {
		return err
	}

	comment.LikesCount++
	_, err = o.Update(comment)
	return err
}
