package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	dbName      = "data/beeblog.sql"
	mysqlDriver = "mysql"
)

type Category struct {
	Id              int64
	Title           string    `orm:"null"`
	Created         time.Time `orm:"index;null"`
	Views           int64     `orm:"index;null"`
	TopicTime       time.Time `orm:"index;null"`
	TopicCount      int64     `orm:"null"`
	TopicLastUserId int64     `orm:"null"`
}

type Topic struct {
	Id              int64
	Uid             int64     `orm:"null"`
	Title           string    `orm:"null"`
	Content         string    `orm:"size(5000);null"`
	Category        string    `orm:"null"`
	Attachment      string    `orm:"null"`
	Labels          string    `orm:"null"`
	Created         time.Time `orm:"index;null"`
	Updated         time.Time `orm:"index;null"`
	Views           int64     `orm:"index;null"`
	Author          string    `orm:"null"`
	ReplayTime      time.Time `orm:"index;null"`
	ReplyCount      int64     `orm:"null"`
	ReplyLastUserId int64     `orm:"null"`
}

type Comment struct {
	Id          int64
	Tid         int64
	Name        string
	Content     string    `orm:"size(1000)"`
	CreatedTime time.Time `orm:"index"`
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	c := &Comment{Id: ridNum}
	o := orm.NewOrm()
	var tid int64
	if o.Read(c) == nil {
		tid = c.Tid
		_, err = o.Delete(c)
		if err != nil {
			return err
		}
	}
	replies := make([]*Comment, 0)
	_, err = o.QueryTable("comment").Filter("tid", tid).OrderBy("-createdtime").All(&replies)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tid}
	if o.Read(topic) == nil {
		if len(replies) == 0 {
			topic.ReplyCount = 0
			topic.ReplayTime = time.Now()
		} else {
			topic.ReplayTime = replies[0].CreatedTime
			topic.ReplyCount = int64(len(replies))
		}

		_, err = o.Update(topic)
	}
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	replies = make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return
}

func AddReplay(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:         tidNum,
		Name:        nickname,
		Content:     content,
		CreatedTime: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyCount++
		topic.ReplayTime = time.Now()
		_, err = o.Update(topic)
	}
	return err
}

func ModifyTopic(tid, title, label, category, content, attachment string) error {

	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	var oldCate, oldAttch string
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttch = topic.Attachment
		topic.Title = title
		topic.Content = content
		topic.Labels = label
		topic.Updated = time.Now()
		topic.Category = category
		topic.Attachment = attachment
		_, err = o.Update(topic)
	}
	//更新分类统计
	if category != oldCate {
		c := &Category{Title: category}
		if o.Read(c, "title") == nil {
			c.TopicCount++
			_, err = o.Update(c)
		}
		c1 := &Category{Title: oldCate}
		if o.Read(c1, "title") == nil {
			c1.TopicCount--
			_, err = o.Update(c1)
		}
	}

	//删除旧的附件
	if len(oldAttch) > 0 {
		err := os.Remove(path.Join("attachment",oldAttch))
		if err != nil {
			return err
		}
	}
	return err
}

func AddTopic(title, category, label, content, attachment string) error {
	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	//空格作为多个标签的分隔符

	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Labels:     label,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
	_, err := o.Insert(topic)
	c := &Category{Title: category}
	if o.Read(c, "title") == nil {
		c.TopicCount++
		_, err = o.Update(c)
	}
	return err
}

func GetAllTopics(label, cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func RegisterDB() {
	if !com.IsExist(dbName) {
		os.MkdirAll(path.Dir(dbName), os.ModePerm)
		os.Create(dbName)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(mysqlDriver, orm.DRMySQL)
	host, _ := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.String("db::port")
	dbname, _ := beego.AppConfig.String("db::databaseName")
	username, _ := beego.AppConfig.String("db::username")
	pwd, _ := beego.AppConfig.String("db::password")
	dbcon := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&loc=Local"
	fmt.Println(dbcon)
	orm.RegisterDataBase("default", mysqlDriver, dbcon)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		c := &Category{Title: topic.Category}
		if o.Read(c, "title") == nil {
			c.TopicCount--
			_, err = o.Update(c)
		}
	}
	_, err = o.Delete(topic)
	return err
}

func GetTopic(isModify bool, tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	if !isModify {
		topic.Views++
		_, err = o.Update(topic)
	}
	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)
	return topic, err
}
