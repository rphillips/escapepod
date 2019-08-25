package models

import (
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/rphillips/escapepod/pkg/sanitize"
)

type Episode struct {
	BaseModel
	PodcastRef  uint       `json:"podcast_ref"`
	GUID        string     `json:"guid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	Duration    string     `json:"duration"`
	Type        string     `json:"type"`
	ImageURL    string     `json:"image_url"`
	ImageTitle  string     `json:"image_title"`
	Published   *time.Time `json:"published"`
}

func NewEpisodeFromItem(item *gofeed.Item) *Episode {
	episode := &Episode{
		GUID:        item.GUID,
		Title:       sanitize.String(item.Title),
		Description: sanitize.String(item.Description),
		Published:   item.PublishedParsed,
	}
	if item.Image != nil {
		episode.ImageURL = item.Image.URL
		episode.ImageTitle = sanitize.String(item.Image.Title)
	}
	if item.ITunesExt != nil {
		episode.Duration = item.ITunesExt.Duration
	}
	if len(item.Enclosures) > 0 {
		episode.Type = item.Enclosures[0].Type
		episode.URL = item.Enclosures[0].URL
	}
	return episode
}
