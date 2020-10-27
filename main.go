package main

import (
	"com.github.yoshidan/go-anime-image/scraping"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Keyword string `default:"" envconfig:"SCRAPING_KEYWORD"`
	Site    string `default:"tsundora.com" envconfig:"SCRAPING_SITE"`
}

func main() {
	config := &Config{}
	envconfig.Process("", config)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		fmt.Println("shutdown downloader...")
		os.Exit(1)
	}()

	var site scraping.Site
	if config.Site == "tsundora.com" {
		site = scraping.NewTsundora(config.Keyword)
	}else {
		site = scraping.NewWallpaperboys(config.Keyword)
	}
	site.Download()

}

