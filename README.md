## typecho迁移到hugo

使用gorm获取typecho数据库中的原始数据，通过openapi接口创建对应的halo记录。


### 已实现

- 标签
- （多级）分类
- 附件
- 文章
  - 文章状态
- 页面
  - 附件链接更新
- （多级）评论
  - 过滤垃圾评论

> 只针对typecho官方数据库结构进行操作

### 使用前操作

1. `halo`博客搭建
   1. 在`个人中心-个人令牌`生成PAT令牌。
   2. 在`halo`中删除不需要的`标签` `分类` `文章` `页面` `评论` `附件`
   3. 取消勾选 `设置` `评论设置` `仅允许注册用户评论`
   4. 取消勾选 `设置` `通知设置` `启用邮件通知器` 防止导入评论时发送邮件
   5. `应用市场` 安装 `Vditor 编辑器`，然后点击`插件` `启用` 
   6. `插件` `评论组件` `头像设置` `启用第三方头像` ，`头像策略` `匿名&无头像用户`
   7. `备份` `创建备份` 可以下载备份防止程序出现异常
2. `typecho`操作
   1. 数据备份
   2. 获取可访问的数据库信息


### 开发

**必须环境：**

- [Git](https://git-scm.com/downloads)
- [Golang](https://golang.google.cn/dl)
- [Taskfile](https://taskfile.dev/installation)

环境配置

```shell
# 拉取仓库
git clone https://github.com/fghwett/typecho-to-halo.git

# 下载openapi协议文件
task down-json

# 生成openapi sdk
task gen-sdk

# 拷贝配置文件并修改
task gen-config

# 生成数据库 sdk
taks gen-db
```

快捷命令

```shell
# 查看所有支持命令
task --list

# 迁移操作
task run

# 更多命令为调试相关 可自行探索
```

