package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Article struct {
	TypeOf                 string   `json:"type_of"`
	ID                     int32    `json:"id"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
	CoverImage             string   `json:"cover_image,omitempty"`
	ReadablePublishDate    string   `json:"readable_publish_date"`
	SocialImage            string   `json:"social_image"`
	TagList                []string `json:"tag_list"`
	Tags                   string   `json:"tags"`
	Slug                   string   `json:"slug"`
	Path                   string   `json:"path"`
	URL                    string   `json:"url"`
	CanonicalURL           string   `json:"canonical_url"`
	CommentsCount          int32    `json:"comments_count"`
	PositiveReactionsCount int32    `json:"positive_reactions_count"`
	PublicReactionsCount   int32    `json:"public_reactions_count"`
	CreatedAt              string   `json:"created_at"`
	EditedAt               string   `json:"edited_at,omitempty"`
	CrosspostedAt          string   `json:"crossposted_at,omitempty"`
	PublishedAt            string   `json:"published_at"`
	LastCommentAt          string   `json:"last_comment_at"`
	PublishedTimestamp     string   `json:"published_timestamp"`
	User                   *User    `json:"user"`
	ReadingTimeMinutes     int32    `json:"reading_time_minutes"`
}

type State string

const (
	Fresh  State = "fresh"
	Rising State = "rising"
	All    State = "all"
)

type ArticleQueryParams struct {
	Page         int32  `url:"page,omitempty"`
	PerPage      int32  `url:"per_page,omitempty"`
	Tag          string `url:"tag,omitempty"`
	Tags         string `url:"tags,omitempty"`
	TagsExclude  string `url:"tags_exclude,omitempty"`
	Username     string `url:"username,omitempty"`
	State        State  `url:"state,omitempty"`
	Top          int32  `url:"top,omitempty"`
	CollectionID int32  `url:"collection_id,omitempty"`
}

// GetPublishedArticles allows client to retrieve a list of articles
func (c *Client) GetPublishedArticles(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}