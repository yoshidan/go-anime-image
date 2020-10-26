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

	url := fmt.Sprintf("https://%s/?s=%s", config.Site, config.Keyword)
	fmt.Printf("start download from %s\n", url)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		fmt.Println("shutdown downloader...")
		os.Exit(1)
	}()

	scraping.NewScraper(config.Site).Execute(url)
}
