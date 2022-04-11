package dev

import (
	"os"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestGetPublishedListings(t *testing.T) {
	c, err := NewTestClient()
	assert.NoError(t, err, "Failed to create TestClient")

	listings, err := c.GetPublishedListings(ListingQueryParams{PerPage: 5, Category: "cfp"})
	assert.NoError(t, err, "Error fetching articles")
	assert.Equalf(t, "listing", listings[0].TypeOf, "Expected field 'type_of' to be 'listing', got '%s'", listings[0].TypeOf)
	assert.Equalf(t, ListingCategoryCfp, listings[0].Category, "Expected field 'category' to be 'cfp', got '%s'", listings[0].Category)
}

func TestGetPublishedListingsByCategory(t *testing.T) {
	c, err := NewTestClient()
	assert.NoError(t, err, "Failed to create TestClient")

	listings, err := c.GetPublishedListingsByCategory("cfp", ListingQueryParams{PerPage: 5})
	assert.NoError(t, err, "Error fetching articles")

	for _, v := range listings {
		assert.Equalf(t, ListingCategoryCfp, v.Category, "Expected catrgory to be 'cfp', instead got '%s'", v.Category)
	}
}

func TestCreateListing_published(t *testing.T) {
	if os.Getenv("TEST_SKIP") != "" {
		t.Skip()
	}
	c, err := NewTestClient()
	assert.NoError(t, err, "Failed to create TestClient")

	title := gofakeit.Sentence(3)
	markdownCreate := gofakeit.Paragraph(1, 2, 5, "\n")
	category := ListingCategoryCfp
	tags := []string{"python", "linux"}

	payload := ListingBodySchema{}
	payload.Listing.Title = title
	payload.Listing.BodyMarkdown = markdownCreate
	payload.Listing.Category = category
	payload.Listing.Tags = tags

	listingCreateResp, err := c.CreateListing(payload, nil)
	assert.NoError(t, err)
	assert.NotNil(t, listingCreateResp)
	assert.Equal(t, "listing", listingCreateResp.TypeOf)
	assert.Equal(t, title, listingCreateResp.Title)
	assert.Equal(t, markdownCreate, listingCreateResp.BodyMarkdown)
	assert.Equal(t, category, listingCreateResp.Category)
	assert.Equal(t, len(tags), len(listingCreateResp.Tags))
	assert.NotEmpty(t, listingCreateResp.CreatedAt)
	assert.True(t, listingCreateResp.Published)
	assert.NotNil(t, listingCreateResp.User)

	listingReadResp, err := c.GetListingByID(strconv.Itoa(int(listingCreateResp.ID)))
	assert.NoError(t, err)
	assert.NotNil(t, listingReadResp)
	assert.EqualValues(t, listingCreateResp.ID, listingReadResp.ID)
	assert.Equal(t, "listing", listingReadResp.TypeOf)
	assert.Equal(t, title, listingReadResp.Title)
	assert.EqualValues(t, markdownCreate, listingReadResp.BodyMarkdown)
	assert.Equal(t, category, listingReadResp.Category)
	assert.Equal(t, len(tags), len(listingReadResp.Tags))
	assert.NotEmpty(t, listingReadResp.CreatedAt)
	assert.True(t, listingReadResp.Published)
	assert.NotNil(t, listingReadResp.User)
}

func TestCreateListing_draft(t *testing.T) {
	if os.Getenv("TEST_SKIP") != "" {
		t.Skip()
	}
	c, err := NewTestClient()
	assert.NoError(t, err, "Failed to create TestClient")

	title := gofakeit.Sentence(3)
	markdownCreate := gofakeit.Paragraph(1, 2, 5, "\n")
	category := ListingCategoryCfp
	tags := []string{"python"}

	payload := ListingBodySchema{}
	payload.Listing.Title = title
	payload.Listing.BodyMarkdown = markdownCreate
	payload.Listing.Category = category
	payload.Listing.Tags = tags
	payload.Listing.Action = "draft"

	listingCreateResp, err := c.CreateListing(payload, nil)
	assert.NoError(t, err)
	assert.NotNil(t, listingCreateResp)
	assert.Equal(t, "listing", listingCreateResp.TypeOf)
	assert.Equal(t, title, listingCreateResp.Title)
	assert.Equal(t, markdownCreate, listingCreateResp.BodyMarkdown)
	assert.Equal(t, category, listingCreateResp.Category)
	assert.Equal(t, len(tags), len(listingCreateResp.Tags))
	assert.NotEmpty(t, listingCreateResp.CreatedAt)
	assert.False(t, listingCreateResp.Published)
	assert.NotNil(t, listingCreateResp.User)

	listingReadResp, err := c.GetListingByID(strconv.Itoa(int(listingCreateResp.ID)))
	assert.NoError(t, err)
	assert.NotNil(t, listingReadResp)
	assert.EqualValues(t, listingCreateResp.ID, listingReadResp.ID)
	assert.Equal(t, "listing", listingReadResp.TypeOf)
	assert.Equal(t, title, listingReadResp.Title)
	assert.EqualValues(t, markdownCreate, listingReadResp.BodyMarkdown)
	assert.Equal(t, category, listingReadResp.Category)
	assert.Equal(t, len(tags), len(listingReadResp.Tags))
	assert.NotEmpty(t, listingReadResp.CreatedAt)
	assert.False(t, listingReadResp.Published)
	assert.NotNil(t, listingReadResp.User)
}
