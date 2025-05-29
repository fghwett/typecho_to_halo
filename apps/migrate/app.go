package migrate

import (
	consolesdk "api.halo.run/apis/openapi-go-console"
	extensionsdk "api.halo.run/apis/openapi-go-extension"
	"fmt"
	"github.com/fghwett/typecho-to-halo/config"
	"github.com/fghwett/typecho-to-halo/pkg/typecho/model"
	"github.com/fghwett/typecho-to-halo/service"
	"github.com/spf13/cast"
	"log/slog"
	"strings"
	"sync"
)

type App struct {
	conf *config.Config

	typecho     *service.Typecho
	halo        *service.Halo
	fileManager *service.FileManager

	typechoAttachments []*model.TypechoContents

	haloTagMap      sync.Map
	haloCategoryMap sync.Map
	haloFileMap     sync.Map
	haloPostMap     sync.Map
	haloPageMap     sync.Map
	haloCommentMap  sync.Map
	haloReplyMap    sync.Map
}

func NewApp(conf *config.Config) (*App, error) {
	a := &App{
		conf: conf,
	}

	if t, e := service.NewTypecho(conf.Typecho); e != nil {
		return nil, e
	} else {
		a.typecho = t
	}

	slog.Info("配置信息", slog.Any("conf", conf))

	if f, e := service.NewFileManager(conf.FileManager); e != nil {
		return nil, e
	} else {
		slog.Info("文件管理器", slog.Any("conf", conf.FileManager))
		a.fileManager = f
	}

	a.halo = service.NewHalo(conf.Halo)

	return a, nil
}

func (app *App) Run() error {
	actions := []*Action{
		NewAction("迁移标签", app.migrateTags),
		NewAction("迁移分类", app.migrateCategories),
		NewAction("迁移附件", app.migrateAttachments),
		NewAction("迁移文章", app.migratePosts),
		NewAction("迁移页面", app.migratePages),
		NewAction("迁移评论", app.migrateComments),
	}

	return app.doActions(actions)
}

func (app *App) migrateTags() error {
	typechoTags, err := app.typecho.GetTags()
	if err != nil {
		return err
	}

	var tag *extensionsdk.Tag
	for _, typechoTag := range typechoTags {
		if tag, err = app.halo.AddTag(cast.ToString(typechoTag.Name), cast.ToString(typechoTag.Slug)); err != nil {
			return err
		}
		app.haloTagMap.Store(typechoTag.Mid, tag)
	}

	return nil
}

func (app *App) migrateCategories() error {
	typechoCategories, err := app.typecho.GetCategories()
	if err != nil {
		return err
	}

	var category *extensionsdk.Category
	for _, typechoCategory := range typechoCategories {
		if category, err = app.halo.AddCategory(cast.ToString(typechoCategory.Name), cast.ToString(typechoCategory.Slug), cast.ToString(typechoCategory.Description), cast.ToInt32(typechoCategory.Order_)); err != nil {
			return err
		}
		app.haloCategoryMap.Store(typechoCategory.Mid, category)
	}

	cache := make(map[uint32][]string)
	for _, typechoCategory := range typechoCategories {
		parentId := cast.ToUint32(typechoCategory.Parent)
		if parentId == 0 {
			continue
		}

		mid := typechoCategory.Mid
		value, ok := app.haloCategoryMap.Load(mid)
		if !ok {
			continue
		}
		childCategory := value.(*extensionsdk.Category)
		list, exist := cache[parentId]
		if exist {
			list = append(list, childCategory.Metadata.Name)
			cache[parentId] = list
			continue
		}
		cache[parentId] = []string{childCategory.Metadata.Name}
	}
	slog.Info("分类父级", slog.Any("cache", cache))

	for mid, children := range cache {
		value, ok := app.haloCategoryMap.Load(mid)
		if !ok {
			continue
		}
		parentCategory := value.(*extensionsdk.Category)
		parentCategory.Spec.Children = children

		if category, err = app.halo.AddCategoryChildren(parentCategory.Metadata.Name, children); err != nil {
			return err
		}
		app.haloCategoryMap.Store(mid, category)
	}

	return nil
}

func (app *App) migrateAttachments() error {
	typechoFiles, err := app.typecho.GetAttachments()
	if err != nil {
		return err
	}
	slog.Info("文件数量", slog.Any("count", len(typechoFiles)))

	app.typechoAttachments = typechoFiles

	var attachment *consolesdk.Attachment
	for _, typechoFile := range typechoFiles {

		link := app.typecho.GetAttachmentUrl(typechoFile)

		var filename *string
		if filename, err = app.fileManager.DownFile(link); err != nil {
			return err
		}
		slog.Info("下载文件", slog.String("filename", cast.ToString(filename)))

		if attachment, err = app.halo.AddAttachment(cast.ToString(filename)); err != nil {
			return err
		}

		app.haloFileMap.Store(typechoFile.Cid, attachment)

		// 删除文件
		//if err = app.fileManager.DeleteFile(cast.ToString(filename)); err != nil {
		//	return err
		//}

	}

	return nil
}

func (app *App) migratePosts() error {
	typechoPosts, err := app.typecho.GetPosts()
	if err != nil {
		return err
	}

	var relationShips []*model.TypechoRelationships
	if relationShips, err = app.typecho.GetRelationShips(); err != nil {
		return err
	}

	var post *consolesdk.Post
	for _, typechoPost := range typechoPosts {
		var tags, categories []string
		for _, relationShip := range relationShips {
			if relationShip.Cid != typechoPost.Cid {
				continue
			}
			value, ok := app.haloCategoryMap.Load(relationShip.Mid)
			if ok {
				haloCategory := value.(*extensionsdk.Category)
				categories = append(categories, haloCategory.Metadata.Name)
				continue
			}
			value, ok = app.haloTagMap.Load(relationShip.Mid)
			if !ok {
				continue
			}
			haloTag := value.(*extensionsdk.Tag)
			tags = append(tags, haloTag.Metadata.Name)
		}

		// 处理文章中的附件信息
		content := cast.ToString(typechoPost.Text)
		for _, typechoAttachment := range app.typechoAttachments {
			value, ok := app.haloFileMap.Load(typechoAttachment.Cid)
			if !ok {
				continue
			}
			typechoFileUrl := app.typecho.GetAttachmentUrl(typechoAttachment)

			haloAttachment := value.(*consolesdk.Attachment)
			//haloFileUrl := fmt.Sprintf("%s://%s/upload/%s", app.conf.Halo.Schema, app.conf.Halo.Host, haloAttachment.Spec.GetDisplayName())
			haloFileUrl := fmt.Sprintf("/upload/%s", haloAttachment.Spec.GetDisplayName())

			content = strings.ReplaceAll(content, typechoFileUrl, haloFileUrl)
		}
		content = strings.ReplaceAll(content, `<!--markdown-->`, ``)

		if post, err = app.halo.AddPost(
			cast.ToString(typechoPost.Title),
			cast.ToString(typechoPost.Slug),
			content,
			cast.ToString(typechoPost.Type),
			cast.ToString(typechoPost.Status),
			cast.ToInt32(typechoPost.Order_),
			cast.ToInt64(typechoPost.Created),
			cast.ToBool(typechoPost.AllowComment),
			tags,
			categories,
		); err != nil {
			return err
		}

		app.haloPostMap.Store(typechoPost.Cid, post)
	}

	return nil
}

func (app *App) migratePages() error {
	typechoPages, err := app.typecho.GetPage()
	if err != nil {
		return err
	}

	var haloPage *consolesdk.SinglePage
	for _, typechoPage := range typechoPages {

		content := cast.ToString(typechoPage.Text)
		for _, typechoAttachment := range app.typechoAttachments {
			value, ok := app.haloFileMap.Load(typechoAttachment.Cid)
			if !ok {
				continue
			}
			typechoFileUrl := app.typecho.GetAttachmentUrl(typechoAttachment)

			haloAttachment := value.(*consolesdk.Attachment)
			//haloFileUrl := fmt.Sprintf("%s://%s/upload/%s", app.conf.Halo.Schema, app.conf.Halo.Host, haloAttachment.Spec.GetDisplayName())
			haloFileUrl := fmt.Sprintf("/upload/%s", haloAttachment.Spec.GetDisplayName())

			content = strings.ReplaceAll(content, typechoFileUrl, haloFileUrl)
		}

		if haloPage, err = app.halo.AddPage(
			cast.ToString(typechoPage.Title),
			cast.ToString(typechoPage.Slug),
			content,
			cast.ToString(typechoPage.Type),
			cast.ToString(typechoPage.Status),
			cast.ToInt32(typechoPage.Order_),
			cast.ToInt64(typechoPage.Created),
			cast.ToBool(typechoPage.AllowComment),
		); err != nil {
			return err
		}

		app.haloPageMap.Store(typechoPage.Cid, haloPage)
	}

	return nil
}

func (app *App) migrateComments() error {
	allTypechoComments, err := app.typecho.GetComments()
	if err != nil {
		return err
	}

	var (
		typechoComments       []*model.TypechoComments
		typechoReplyMap       = make(map[uint32][]*model.TypechoComments)
		typechoReplyParentMap = make(map[uint32]uint32)
	)

	for _, typechoComment := range allTypechoComments {
		parentCommentId := cast.ToUint32(typechoComment.Parent)
		typechoReplyParentMap[typechoComment.Coid] = parentCommentId
		if parentCommentId == 0 {
			typechoComments = append(typechoComments, typechoComment)
			continue
		}
		list, ok := typechoReplyMap[parentCommentId]
		if !ok {
			typechoReplyMap[parentCommentId] = []*model.TypechoComments{
				typechoComment,
			}
			continue
		}
		list = append(list, typechoComment)
		typechoReplyMap[parentCommentId] = list
	}

	var nextIds []uint32
	for _, typechoComment := range typechoComments {
		contentId := cast.ToUint32(typechoComment.Cid)

		var contentName string
		var isPost bool
		if value, ok := app.haloPostMap.Load(contentId); ok {
			haloPost := value.(*consolesdk.Post)
			contentName = haloPost.Metadata.Name
			isPost = true
		} else if value, ok = app.haloPageMap.Load(contentId); ok {
			haloPage := value.(*consolesdk.SinglePage)
			contentName = haloPage.Metadata.Name
		}
		if contentName == "" {
			continue
		}

		var haloComment *consolesdk.Comment
		if haloComment, err = app.halo.AddComment(
			contentName,
			cast.ToString(typechoComment.Author),
			cast.ToString(typechoComment.Mail),
			cast.ToString(typechoComment.URL),
			cast.ToString(typechoComment.Text),
			cast.ToString(typechoComment.Status),
			isPost,
		); err != nil {
			return err
		}

		if haloComment == nil {
			continue
		}

		nextIds = append(nextIds, typechoComment.Coid)
		app.haloCommentMap.Store(typechoComment.Coid, haloComment)
	}

	for {
		var newNextIds []uint32

		for _, typechoCommentId := range nextIds {
			typechoReplyList, ok := typechoReplyMap[typechoCommentId]
			if !ok {
				continue
			}
			for _, typechoReply := range typechoReplyList {
				parentCommentId := cast.ToUint32(typechoReply.Parent)
				var commentName, quoteComment string
				if value, exist := app.haloCommentMap.Load(parentCommentId); exist {
					haloComment := value.(*consolesdk.Comment)
					commentName = haloComment.Metadata.Name
				} else if value, exist = app.haloReplyMap.Load(parentCommentId); exist {
					haloReply := value.(*consolesdk.Reply)
					quoteComment = haloReply.Metadata.Name
					// 找出第一条评论
					var parentId uint32
					childrenId := parentCommentId
					for {
						parentId, ok = typechoReplyParentMap[childrenId]
						if !ok {
							parentId = childrenId
							break
						}
						if parentId == 0 {
							parentId = childrenId
							break
						}
						childrenId = parentId
					}
					if value, exist = app.haloCommentMap.Load(parentId); exist {
						commentName = value.(*consolesdk.Comment).Metadata.Name
					}
				}
				if commentName == "" {
					continue
				}
				var haloReply *consolesdk.Reply
				if haloReply, err = app.halo.AddReply(
					commentName,
					cast.ToString(typechoReply.Author),
					cast.ToString(typechoReply.Mail),
					cast.ToString(typechoReply.URL),
					cast.ToString(typechoReply.Text),
					quoteComment,
				); err != nil {
					return err
				}
				app.haloReplyMap.Store(typechoReply.Coid, haloReply)
				newNextIds = append(newNextIds, typechoReply.Coid)
			}
			delete(typechoReplyMap, typechoCommentId)

		}

		if len(typechoReplyMap) == 0 || len(newNextIds) == 0 {
			break
		}
		nextIds = newNextIds
	}

	return nil
}

type Action struct {
	Name string
	Do   func() error
}

func NewAction(name string, do func() error) *Action {
	return &Action{
		Name: name,
		Do:   do,
	}
}

func (app *App) doActions(actions []*Action) error {
	for _, action := range actions {
		slog.Info("执行操作", slog.String("action", action.Name))
		if err := action.Do(); err != nil {
			slog.Error("执行操作失败", slog.String("action", action.Name), slog.Any("err", err))
			return err
		}
		slog.Info("执行操作成功", slog.String("action", action.Name))
	}
	return nil
}
