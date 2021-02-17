package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

func AddTopic(title, category, content, author string, isMarkdown bool) error {
	o := orm.NewOrm()
	t := time.Now()
	topic := &Topic{
		Title:   title,
		Category: category,
		Content: content,
		Created: t,
		Updated: t,
		Author:  author,
		IsMarkdown: isMarkdown,
		ReplyTime: t,
	}
	_, err := o.Insert(topic)
	return err
}

func UpdateTopic(id, title, category, content string, isMarkdown bool) error {
	o := orm.NewOrm()
	tid, err := strconv.ParseInt(id, 10, 64)
	topic := &Topic{Id: tid}

	err = o.QueryTable(topic).Filter("id", id).One(topic)
	if err != nil {
		return err
	}

	topic.Title = title
	topic.Category = category
	topic.Content = content
	topic.IsMarkdown = isMarkdown
	topic.Updated = time.Now()

	_, err = o.Update(topic)
	return err
}

func DelTopic(id string) error {
	o := orm.NewOrm()
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tid}
	_, err = o.Delete(topic)
	return err
}

func FetchAllTopic(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	var err error
	qs := o.QueryTable("topic")
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func FetchTopicById(id string) (*Topic, error) {
	topicId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(Topic)
	o := orm.NewOrm()

	qs := o.QueryTable(topic)
	err = qs.Filter("id", topicId).One(topic)
	if err != nil {
		return topic, err
	}

	topic.Views++
	o.Update(topic)

	return topic, nil
}
