package main

import (
	"github.com/fghwett/typecho-to-halo/apps/migrate"
	"github.com/fghwett/typecho-to-halo/config"
	"github.com/golang-cz/devslog"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"log/slog"
	"os"
)

func init() {
	opts := &devslog.Options{
		HandlerOptions: &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		},
		TimeFormat:        "[15:04:05]",
		DebugColor:        devslog.Magenta,
		StringerFormatter: true,
	}

	logger := slog.New(devslog.NewHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func main() {
	start()
}

func start() {
	if err := run(); err != nil {
		slog.Error("运行失败", slog.Any("err", err))
		os.Exit(1)
	}
}

func run() error {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[Migrate] ")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./_tmp/app.log",
		MaxSize:    100,   // 单个文件最大大小（MB）
		MaxBackups: 3,     // 保留旧文件的最大数量
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩旧文件
	})

	conf, err := config.New()
	if err != nil {
		return err
	}

	a, e := migrate.NewApp(conf)
	if e != nil {
		return e
	}

	return a.Run()
}
