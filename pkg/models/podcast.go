package models

import (
	"bufio"
	"bytes"
	"image"
	"image/jpeg"

	// import png to support png
	_ "image/png"

	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/mmcdole/gofeed"

	"github.com/rphillips/escapepod/pkg/sanitize"
)

type Podcast struct {
	BaseModel

	Author      string  `json:"author" gorm:""`
	Description string  `json:"description"`
	FeedURL     *string `json:"feed_url" gorm:"not null;unique"`
	ImageTitle  string  `json:"image_title" gorm:""`
	ImageURL    string  `json:"image_url" gorm:""`
	Link        string  `json:"link"`
	Title       string  `json:"title" gorm:"not_null"`

	ImageResized200 []byte `json:"-"`

	Episodes []*Episode `json:"episodes" gorm:"ForeignKey:PodcastRef"`
}

func (p *Podcast) Crawl() error {
	parser := gofeed.NewParser()
	var feed *gofeed.Feed
	if strings.HasPrefix(*p.FeedURL, "file://") {
		fileData, err := ioutil.ReadFile(strings.TrimPrefix(*p.FeedURL, "file://"))
		if err != nil {
			return err
		}
		feed, err = parser.ParseString(string(fileData))
		if err != nil {
			return err
		}
	} else {
		var err error
		feed, err = parser.ParseURL(*p.FeedURL)
		if err != nil {
			return err
		}
	}
	return p.PopulateFromFeed(feed)
}

func (p *Podcast) PopulateFromFeed(feed *gofeed.Feed) error {
	p.Title = sanitize.String(feed.Title)
	p.ImageURL = feed.Image.URL
	p.ImageTitle = sanitize.String(feed.Image.Title)
	p.Link = feed.Link
	p.Description = sanitize.String(feed.Description)
	if feed.Author != nil {
		p.Author = feed.Author.Name
	}
	for _, item := range feed.Items {
		if ep := NewEpisodeFromItem(item); ep != nil {
			p.Episodes = append(p.Episodes, NewEpisodeFromItem(item))
		}
	}
	img, err := getImage(feed.Image.URL)
	if err != nil {
		return err
	}
	p.ImageResized200, err = resizeImage(img, 200, 200)
	return err
}

func getImage(url string) (image.Image, error) {
	imageBytes, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer imageBytes.Body.Close()
	img, _, err := image.Decode(imageBytes.Body.(io.Reader))
	return img, err
}

func resizeImage(img image.Image, width, height int) ([]byte, error) {
	dstImage200 := imaging.Resize(img, width, height, imaging.Lanczos)
	var out bytes.Buffer
	if err := jpeg.Encode(bufio.NewWriter(&out), dstImage200, &jpeg.Options{Quality: 90}); err != nil {
		return []byte{}, err
	}
	return out.Bytes(), nil
}
