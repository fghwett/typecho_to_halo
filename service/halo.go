package service

import (
	consolesdk "api.halo.run/apis/openapi-go-console"
	extensionsdk "api.halo.run/apis/openapi-go-extension"
	publicsdk "api.halo.run/apis/openapi-go-public"
	usersdk "api.halo.run/apis/openapi-go-user"
	"bytes"
	"context"
	"github.com/fghwett/typecho-to-halo/config"
	"github.com/google/uuid"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"log/slog"
	"os"
	"time"
)

const (
	apiVersion = `content.halo.run/v1alpha1`

	kindTag        = "Tag"
	kindPost       = "Post"
	kindCategory   = "Category"
	kindSinglePage = "SinglePage"

	contentRawType = "markdown"
	contentVersion = 1

	commentGroup   = "content.halo.run"
	commentVersion = "v1alpha1"

	visiblePublic   = "PUBLIC"
	visibleInternal = "INTERNAL"
	visiblePrivate  = "PRIVATE"
)

type Halo struct {
	conf *config.Halo

	//aggregated *aggregatedsdk.APIClient
	console   *consolesdk.APIClient
	extension *extensionsdk.APIClient
	public    *publicsdk.APIClient
	user      *usersdk.APIClient
}

func NewHalo(conf *config.Halo) *Halo {
	//aggregatedConf := aggregatedsdk.NewConfiguration()
	//aggregatedConf.Host = conf.Host
	//aggregatedConf.Scheme = conf.Schema
	//aggregatedConf.Debug = conf.Debug
	//aggregatedConf.AddDefaultHeader("Authorization", "Bearer "+conf.Token)
	//aggregatedClient := aggregatedsdk.NewAPIClient(aggregatedConf)

	consoleConf := consolesdk.NewConfiguration()
	consoleConf.Host = conf.Host
	consoleConf.Scheme = conf.Schema
	consoleConf.Debug = conf.Debug
	consoleConf.AddDefaultHeader("Authorization", "Bearer "+conf.Token)
	consoleClient := consolesdk.NewAPIClient(consoleConf)

	extensionConf := extensionsdk.NewConfiguration()
	extensionConf.Host = conf.Host
	extensionConf.Scheme = conf.Schema
	extensionConf.Debug = conf.Debug
	extensionConf.AddDefaultHeader("Authorization", "Bearer "+conf.Token)
	extensionClient := extensionsdk.NewAPIClient(extensionConf)

	publicConf := publicsdk.NewConfiguration()
	publicConf.Host = conf.Host
	publicConf.Scheme = conf.Schema
	publicConf.Debug = conf.Debug
	publicConf.AddDefaultHeader("Authorization", "Bearer "+conf.Token)
	publicClient := publicsdk.NewAPIClient(publicConf)

	userConf := usersdk.NewConfiguration()
	userConf.Host = conf.Host
	userConf.Scheme = conf.Schema
	userConf.Debug = conf.Debug
	userConf.AddDefaultHeader("Authorization", "Bearer "+conf.Token)
	userClient := usersdk.NewAPIClient(userConf)

	h := &Halo{
		conf: conf,
		//aggregated: aggregated,
		console:   consoleClient,
		extension: extensionClient,
		public:    publicClient,
		user:      userClient,
	}

	return h
}

func (halo *Halo) AddTag(name, slug string) (*extensionsdk.Tag, error) {
	result, _, err := halo.extension.TagV1alpha1API.CreateTag(context.Background()).Tag(extensionsdk.Tag{
		ApiVersion: apiVersion,
		Kind:       kindTag,
		Metadata: extensionsdk.Metadata{
			Name: uuid.New().String(),
		},
		Spec: extensionsdk.TagSpec{
			DisplayName: name,
			Slug:        slug,
		},
	}).Execute()

	return result, err
}

func (halo *Halo) AddCategory(name, slug, description string, priority int32) (*extensionsdk.Category, error) {
	result, _, err := halo.extension.CategoryV1alpha1API.CreateCategory(context.Background()).Category(extensionsdk.Category{
		ApiVersion: apiVersion,
		Kind:       kindCategory,
		Metadata: extensionsdk.Metadata{
			Name: uuid.New().String(),
		},
		Spec: extensionsdk.CategorySpec{
			Children:                      []string{},
			Cover:                         extensionsdk.PtrString(""),
			Description:                   extensionsdk.PtrString(description),
			HideFromList:                  extensionsdk.PtrBool(false),
			PostTemplate:                  extensionsdk.PtrString(""),
			PreventParentPostCascadeQuery: extensionsdk.PtrBool(false),
			DisplayName:                   name,
			Priority:                      priority,
			Slug:                          slug,
			Template:                      extensionsdk.PtrString(""),
		},
	}).Execute()

	return result, err
}

func (halo *Halo) AddCategoryChildren(name string, children []string) (*extensionsdk.Category, error) {
	result, _, err := halo.extension.CategoryV1alpha1API.PatchCategory(context.Background(), name).JsonPatchInner([]extensionsdk.JsonPatchInner{
		{
			AddOperation: &extensionsdk.AddOperation{
				Op:    "add",
				Path:  "/spec/children",
				Value: children,
			},
		},
	}).Execute()
	//result, _, err := halo.extension.CategoryV1alpha1API.UpdateCategory(context.Background(), category.Metadata.Name).Category(*category).Execute()
	return result, err
}

func (halo *Halo) AddAttachment(filename string) (*consolesdk.Attachment, error) {
	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	attachment, resp, err := halo.console.AttachmentV1alpha1ConsoleAPI.UploadAttachment(context.Background()).PolicyName(halo.conf.PolicyName).GroupName(halo.conf.GroupName).File(f).Execute()
	if err == nil {
		defer func() {
			if er := resp.Body.Close(); err != nil {
				slog.Error("关闭上传文件失败", slog.Any("err", er), slog.String("filename", filename))
			}
		}()
	}

	//slog.Info("上传文件结果", slog.Any("err", err), slog.String("filename", filename), slog.Any("attachment", attachment))

	return attachment, err
}

func (halo *Halo) AddPost(title, slug, content, typ, status string, order int32, created int64, allowComment bool, tags, categories []string) (*consolesdk.Post, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			extension.Typographer,
			extension.CJK,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return nil, err
	}

	publish := true
	if typ == "post_draft" {
		publish = false
	}

	var visible string
	switch status {
	case "publish":
		visible = visiblePublic
	case "hidden":
		visible = visiblePrivate
	case "private":
		visible = visiblePrivate
	case "waiting":
		visible = visiblePublic
		publish = false
	default:
		visible = visiblePrivate
	}

	post, _, err := halo.console.PostV1alpha1ConsoleAPI.DraftPost(context.Background()).PostRequest(consolesdk.PostRequest{
		Content: &consolesdk.ContentUpdateParam{
			Content: consolesdk.PtrString(buf.String()),
			Raw:     consolesdk.PtrString(content),
			RawType: consolesdk.PtrString(contentRawType),
			Version: consolesdk.PtrInt64(contentVersion),
		},
		Post: consolesdk.Post{
			ApiVersion: apiVersion,
			Kind:       kindPost,
			Metadata: consolesdk.Metadata{
				Name: uuid.New().String(),
			},
			Spec: consolesdk.PostSpec{
				Title:        title,
				Slug:         slug,
				Deleted:      false,
				Publish:      publish,
				PublishTime:  consolesdk.PtrTime(time.Unix(created, 0)),
				Pinned:       false,
				AllowComment: allowComment,
				Visible:      visible,
				Priority:     order,
				Excerpt: consolesdk.Excerpt{
					AutoGenerate: true,
				},
				Categories: categories,
				Tags:       tags,
				HtmlMetas:  []map[string]string{},
			},
		},
	}).Execute()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (halo *Halo) AddPage(title, slug, content, typ, status string, order int32, created int64, allowComment bool) (*consolesdk.SinglePage, error) {

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			extension.Typographer,
			extension.CJK,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return nil, err
	}

	publish := true
	if typ == "post_draft" {
		publish = false
	}

	var visible string
	switch status {
	case "publish":
		visible = visiblePublic
	case "hidden":
		visible = visiblePrivate
	case "private":
		visible = visiblePrivate
	case "waiting":
		visible = visiblePublic
		publish = false
	default:
		visible = visiblePrivate
	}

	page, _, err := halo.console.SinglePageV1alpha1ConsoleAPI.DraftSinglePage(context.Background()).SinglePageRequest(consolesdk.SinglePageRequest{
		Content: consolesdk.ContentUpdateParam{
			Content: consolesdk.PtrString(buf.String()),
			Raw:     consolesdk.PtrString(content),
			RawType: consolesdk.PtrString(contentRawType),
			Version: consolesdk.PtrInt64(contentVersion),
		},
		Page: consolesdk.SinglePage{
			ApiVersion: apiVersion,
			Kind:       kindSinglePage,
			Metadata: consolesdk.Metadata{
				Name: uuid.New().String(),
			},
			Spec: consolesdk.SinglePageSpec{
				AllowComment: allowComment,
				Deleted:      false,
				Excerpt: consolesdk.Excerpt{
					AutoGenerate: true,
				},
				HtmlMetas:   []map[string]string{},
				Pinned:      false,
				Priority:    order,
				Publish:     publish,
				PublishTime: consolesdk.PtrTime(time.Unix(created, 0)),
				Slug:        slug,
				Title:       title,
				Visible:     visible,
			},
		},
	}).Execute()

	return page, err
}

func (halo *Halo) AddComment(postName, ownerName, ownerEmail, ownerWebsite, content, status string, isPost bool) (*consolesdk.Comment, error) {

	if status == "spam" {
		return nil, nil
	}

	// todo 判断评论是否审核通过 未通过时 使用public提供的接口
	kind := kindPost
	if !isPost {
		kind = kindSinglePage
	}

	comment, _, err := halo.console.CommentV1alpha1ConsoleAPI.CreateComment(context.Background()).CommentRequest(consolesdk.CommentRequest{
		AllowNotification: consolesdk.PtrBool(true),
		Content:           content,
		Raw:               content,
		Owner: &consolesdk.CommentEmailOwner{
			DisplayName: consolesdk.PtrString(ownerName),
			Email:       consolesdk.PtrString(ownerEmail),
			Website:     consolesdk.PtrString(ownerWebsite),
		},
		SubjectRef: consolesdk.Ref{
			Group:   consolesdk.PtrString(commentGroup),
			Kind:    consolesdk.PtrString(kind),
			Name:    postName,
			Version: consolesdk.PtrString(commentVersion),
		},
	}).Execute()

	return comment, err
}

func (halo *Halo) AddReply(commentName, ownerName, ownerEmail, ownerWebsite, content, quote string) (*consolesdk.Reply, error) {

	reply, _, err := halo.console.CommentV1alpha1ConsoleAPI.CreateReply(context.Background(), commentName).ReplyRequest(consolesdk.ReplyRequest{
		AllowNotification: consolesdk.PtrBool(true),
		Owner: &consolesdk.CommentEmailOwner{
			DisplayName: consolesdk.PtrString(ownerName),
			Email:       consolesdk.PtrString(ownerEmail),
			Website:     consolesdk.PtrString(ownerWebsite),
		},
		QuoteReply: consolesdk.PtrString(quote),
		Content:    content,
		Raw:        content,
	}).Execute()

	return reply, err
}
