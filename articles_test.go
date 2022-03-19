package dev

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestGetPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	t.Run("page limit", func(t *testing.T) {
		if os.Getenv("TEST_SKIP") != "" {
			t.Skip()
		}
		pageLimit := int32(10)
		articles, err := c.GetPublishedArticles(ArticleQueryParams{Page: 1, PerPage: pageLimit})
		assert.NoErrorf(t, err, "Error fetching articles: %w", err)
		assert.EqualValues(t, pageLimit, len(articles))
	})

	t.Run("articles with tag", func(t *testing.T) {
		if os.Getenv("TEST_SKIP") != "" {
			t.Skip()
		}
		articles, err := c.GetPublishedArticles(ArticleQueryParams{PerPage: 1, Tag: "golang"})
		assert.NoErrorf(t, err, "Error fetching articles: %w", err)

		assert.Containsf(t, articles[0].Tags, "go", "Expected tags to contain given tag, got: %s", articles[0].Tags)
	})

	t.Run("articles with tag", func(t *testing.T) {
		if os.Getenv("TEST_SKIP") != "" {
			t.Skip()
		}
		username := os.Getenv("TEST_USERNAME")
		assert.NotEmpty(t, username)

		articles, err := c.GetPublishedArticles(ArticleQueryParams{PerPage: 1, Username: username})
		assert.NoErrorf(t, err, "Error fetching articles: %w", err)
		assert.Equal(t, username, articles[0].User.Username)
	})
}

func TestLifecycleUnpublishedArticle(t *testing.T) {
	if os.Getenv("TEST_SKIP") != "" {
		t.Skip()
	}
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	gofakeit.Seed(time.Now().UnixNano())

	title := gofakeit.HipsterSentence(10)
	markdownCreate := gofakeit.HipsterParagraph(2, 5, 20, "\n")
	payloadCreate := ArticleBodySchema{}
	payloadCreate.Article.Title = title
	payloadCreate.Article.BodyMarkdown = markdownCreate
	payloadCreate.Article.Published = false
	payloadCreate.Article.Tags = []string{"golang"}

	articleCr, err := c.CreateArticle(payloadCreate, nil)
	assert.NoErrorf(t, err, "Error trying to create article: %w", err)

	assert.Equal(t, title, articleCr.Title)
	assert.Equal(t, markdownCreate, articleCr.BodyMarkdown)
	assert.Equal(t, len(payloadCreate.Article.Tags), len(articleCr.Tags))
	articleID := strconv.Itoa(int(articleCr.ID))

	payloadUpdate := ArticleBodySchema{}
	markdownUpdate := gofakeit.HipsterParagraph(2, 5, 20, "\n")
	payloadUpdate.Article.Title = title
	payloadUpdate.Article.BodyMarkdown = markdownUpdate
	payloadUpdate.Article.Published = false
	payloadUpdate.Article.Tags = []string{"golang", "discuss"}

	articleUpd, err := c.UpdateArticle(articleID, payloadUpdate, nil)
	assert.NoError(t, err)
	assert.Equal(t, title, articleUpd.Title)
	assert.Equal(t, markdownUpdate, articleUpd.BodyMarkdown)
	assert.Equal(t, len(payloadUpdate.Article.Tags), len(articleUpd.Tags))
}

func TestLifecyclePublishedArticle(t *testing.T) {
	if os.Getenv("TEST_SKIP") != "" {
		t.Skip()
	}
	username := os.Getenv("TEST_USERNAME")
	assert.NotEmpty(t, username)

	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	gofakeit.Seed(time.Now().UnixNano())

	title := gofakeit.HipsterSentence(10)
	markdownCreate := gofakeit.HipsterParagraph(2, 5, 10, "\n")
	payloadCreate := ArticleBodySchema{}
	payloadCreate.Article.Title = title
	payloadCreate.Article.BodyMarkdown = markdownCreate
	payloadCreate.Article.Published = true
	payloadCreate.Article.Tags = []string{"golang"}

	articleCr, err := c.CreateArticle(payloadCreate, nil)
	assert.NoError(t, err)

	assert.Equal(t, title, articleCr.Title)
	assert.Equal(t, markdownCreate, articleCr.BodyMarkdown)
	assert.Equal(t, len(payloadCreate.Article.Tags), len(articleCr.Tags))

	articleID := strconv.Itoa(int(articleCr.ID))
	articleByID, err := c.GetPublishedArticleByID(articleID)
	assert.NoErrorf(t, err, "Error fetching article by ID: %w", err)
	assert.NotNil(t, articleByID)

	a, err := strconv.Atoi(articleID)
	assert.NoErrorf(t, err, "Error converting string to int: %w", err)
	assert.EqualValues(t, a, articleByID.ID)

	articleByPath, err := c.GetPublishedArticleByPath(username, articleCr.Slug)
	assert.NoErrorf(t, err, "Error fetching article by path: %w", err)
	assert.EqualValues(t, articleCr.Slug, articleByPath.Slug)
	assert.NotNil(t, articleByPath)
}

func TestGetPublishedArticlesSorted(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	articles, err := c.GetPublishedArticlesSorted(ArticleQueryParams{Page: 1, PerPage: 10})
	assert.NoErrorf(t, err, "Error fetching articles: %w", err)

	if len(articles) >= 2 {
		t1, err := parseUTCDate(articles[0].PublishedAt)
		assert.NoErrorf(t, err, "Error parsing UTC date: %w", err)
		t2, err := parseUTCDate(articles[1].PublishedAt)
		assert.NoErrorf(t, err, "Error parsing UTC date: %w", err)
		assert.Greater(t, t1, t2)
	}
}

func TestGetUserArticles(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	username := os.Getenv("TEST_USERNAME")
	assert.NotEmpty(t, username)

	perPage := int32(2)
	articles, err := c.GetUserArticles(ArticleQueryParams{Page: 1, PerPage: perPage})
	assert.NoErrorf(t, err, "Error fetching articles: %w", err)

	assert.EqualValues(t, perPage, len(articles))
	assert.EqualValues(t, username, articles[0].User.Username)
}

func TestGetUserPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	username := os.Getenv("TEST_USERNAME")
	assert.NotEmpty(t, username)

	articles, err := c.GetUserPublishedArticles(ArticleQueryParams{Page: 1, PerPage: 5})
	assert.NoErrorf(t, err, "Error fetching articles: %w", err)

	for _, v := range articles {
		assert.False(t, v.Published, "Expected result to contain published articles")
	}
}

func TestGetUserUnPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	articles, err := c.GetUserUnPublishedArticles(ArticleQueryParams{Page: 1, PerPage: 5})
	assert.NoErrorf(t, err, "Error fetching articles: %w", err)
	assert.NotNil(t, articles)
	assert.Equal(t, articles[0].User.Username, os.Getenv("TEST_USERNAME"))

	for _, v := range articles {
		assert.False(t, v.Published)
	}
}

func TestGetArticlesWithVideo(t *testing.T) {
	c, err := NewTestClient()
	assert.NoErrorf(t, err, "Failed to create TestClient: %w", err)

	articles, err := c.GetArticlesWithVideo(ArticleQueryParams{Page: 1, PerPage: 5})
	assert.NoErrorf(t, err, "Error fetching articles: %w", err)

	for _, v := range articles {
		assert.Equal(t, "video_article", v.TypeOf)
	}
}
