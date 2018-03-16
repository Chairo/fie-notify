package notifier

import (
	"file-notify/configer"
	"file-notify/file_manager"
	"log"
	s "strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/howeyc/fsnotify"
)

var client *redis.Client
var config *configer.Config

type Notifier struct {
	PConfiger configer.Configer
	PWatcher  *fsnotify.Watcher        //文件监听
	PPath     string                   //监听目录
	PFile     file_manager.FileManager //文件管理
	PPubSub   *redis.PubSub
}

//
func (this *Notifier) SetConfiger(configer configer.Configer) *Notifier {
	this.PConfiger = configer
	return this
}

// 更新文件后发送Redis文件通知
func (this *Notifier) Notify() {
	this.PWatcher.Watch(config.Source)
	this.Do()
}

func (this *Notifier) Update(file string) {
	content := this.PFile.ReadFile(file)
	file = s.Replace(file, "\\", "/", -1)
	this.PFile.UpdateFile(s.Replace(file, config.Source, config.File, -1), content)
}

func (this *Notifier) Do() {
	go func() {
		for {
			select {
			case w := <-this.PWatcher.Event:
				if w.IsModify() {
					client.Publish(config.Channel, w.Name)
					continue
				}
			case err := <-this.PWatcher.Error:
				log.Fatalln(err)
			}
		}
	}()
}

func (this *Notifier) Init() {
	client = this.PConfiger.GetClient()
	config = this.PConfiger.GetConfig()
	pubsub := client.Subscribe(config.Channel)
	pubsub.ReceiveTimeout(time.Second)
	this.PPubSub = pubsub
}

// 初始化监控对象
func NewNotifier() *Notifier {
	w, _ := fsnotify.NewWatcher()
	notify := &Notifier{PConfiger: configer.NewRedisConfiger(), PFile: file_manager.NewLocalFileManager(), PWatcher: w}
	notify.Init()
	return notify
}
