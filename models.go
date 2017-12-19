package main

import (
	"database/sql"
	"time"
)

type MezzanineUser struct {
	ID          int64     `db:"id,omitempty"`
	Username    string    `db:"username"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	IsStaff     bool      `db:"is_staff"`
	IsActive    bool      `db:"is_active"`
	IsSuperUser bool      `db:"is_superuser"`
	LastLogin   time.Time `db:"last_login"`
	JoinDate    time.Time `db:"date_joined"`
}

type MezzanineBlogCategory struct {
	ID     int64          `db:"id,omitempty"`
	SiteID int64          `db:"site_id"`
	Name   string         `db:"title"`
	Slug   sql.NullString `db:"slug"`
}

type MezzanineBlogPost struct {
	// Maps the "Name" property to the "name" column
	// of the "birthday" table.
	ID                     int64          `db:"id,omitempty"`
	CommentCount           int64          `db:"comments_count"`
	Keywords               string         `db:"keywords_string"`
	RatingCount            int64          `db:"rating_count"`
	RatingAverage          float64        `db:"rating_average"`
	SiteID                 int64          `db:"site_id"`
	Title                  string         `db:"title"`
	Slug                   sql.NullString `db:"slug"`
	MetaTitle              sql.NullString `db:"_meta_title"`
	Description            string         `db:"description"`
	IsGeneratedDescription bool           `db:"gen_description"`
	Status                 int64          `db:"status"`
	PublishDate            *time.Time     `db:"publish_date"`
	ExpiryDate             *time.Time     `db:"expiry_date"`
	ShortURL               sql.NullString `db:"short_url"`
	InSitemap              bool           `db:"in_sitemap"`
	Content                string         `db:"content"`
	AuthorID               int64          `db:"user_id"`
	AllowComments          bool           `db:"allow_comments"`
	FeaturedImage          sql.NullString `db:"featured_image"`
	RatingSummary          int64          `db:"rating_sum"`
	Created                *time.Time     `db:"created"`
	Updated                *time.Time     `db:"updated"`
}

type MezzanineBlogPostCategory struct {
	ID             int64 `db:"id,omitempty"`
	BlogPostID     int64 `db:"blogpost_id"`
	BlogCategoryID int64 `db:"blogcategory_id"`
}

type MezzanineBlogPostRelatedPost struct {
	ID             int64 `db:"id,omitempty"`
	FromBlogPostID int64 `db:"from_blogpost_id"`
	ToBlogPostID   int64 `db:"to_blogpost_id"`
}

type GhostUser struct {
	ID              string         `db:"id"`
	Name            string         `db:"name"`
	Slug            string         `db:"slug"`
	AccessToken     sql.NullString `db:"ghost_auth_access_token"`
	AuthID          sql.NullString `db:"ghost_auth_id"`
	Password        string         `db:"password"`
	Email           string         `db:"email"`
	ProfileImage    sql.NullString `db:"profile_image"`
	CoverImage      sql.NullString `db:"cover_image"`
	Biography       sql.NullString `db:"bio"`
	Website         sql.NullString `db:"website"`
	Location        sql.NullString `db:"location"`
	Facebook        sql.NullString `db:"facebook"`
	Twitter         sql.NullString `db:"twitter"`
	Accessibility   sql.NullString `db:"accessibility"`
	Status          string         `db:"status"`
	Locale          sql.NullString `db:"locale"`
	Visibility      string         `db:"visibility"`
	MetaTitle       sql.NullString `db:"meta_title"`
	MetaDescription sql.NullString `db:"meta_description"`
	Tour            sql.NullString `db:"tour"`
	LastSeen        *time.Time     `db:"last_seen"`
	CreatedAt       time.Time      `db:"created_at"`
	CreatedBy       string         `db:"created_by"`
	UpdatedAt       *time.Time     `db:"updated_at"`
	UpdatedBy       sql.NullString `db:"updated_by"`
}

type GhostBlogPost struct {
	ID                   string         `db:"id"`
	UUID                 string         `db:"uuid"`
	Title                string         `db:"title"`
	Slug                 string         `db:"slug"`
	Mobiledoc            sql.NullString `db:"mobiledoc"`
	HTML                 sql.NullString `db:"html"`
	AMP                  sql.NullString `db:"amp"`
	PlainText            sql.NullString `db:"plaintext"`
	FeatureImage         sql.NullString `db:"feature_image"`
	Featured             bool           `db:"featured"`
	Page                 bool           `db:"page"`
	Status               string         `db:"status"`
	Locale               sql.NullString `db:"locale"`
	Visibility           string         `db:"visibility"`
	MetaTitle            sql.NullString `db:"meta_title"`
	MetaDescription      sql.NullString `db:"meta_description"`
	AuthorID             string         `db:"author_id"`
	CreatedAt            time.Time      `db:"created_at"`
	CreatedBy            string         `db:"created_by"`
	UpdatedAt            *time.Time     `db:"updated_at"`
	UpdatedBy            sql.NullString `db:"updated_by"`
	PublishedAt          time.Time      `db:"published_at"`
	PublishedBy          sql.NullString `db:"published_by"`
	CustomExcerpt        sql.NullString `db:"custom_excerpt"`
	CodeInjectionHead    sql.NullString `db:"codeinjection_head"`
	CodeInjectionFoot    sql.NullString `db:"codeinjection_foot"`
	OpenGraphImage       sql.NullString `db:"og_image"`
	OpenGraphTitle       sql.NullString `db:"og_title"`
	OpenGraphDescription sql.NullString `db:"og_description"`
	TwitterImage         sql.NullString `db:"twitter_image"`
	TwitterTitle         sql.NullString `db:"twitter_title"`
	TwitterDescription   sql.NullString `db:"twitter_description"`
	CustomTemplate       sql.NullString `db:"custom_template"`
}

type GhostBlogPostTag struct {
	ID              string         `db:"id"`
	Name            string         `db:"name"`
	Slug            string         `db:"slug"`
	Description     sql.NullString `db:"description"`
	FeatureImage    sql.NullString `db:"feature_image"`
	ParentID        sql.NullString `db:"parent_id"`
	Visibility      string         `db:"visibility"`
	MetaTitle       sql.NullString `db:"meta_title"`
	MetaDescription sql.NullString `db:"meta_description"`
	CreatedAt       time.Time      `db:"created_at"`
	CreatedBy       string         `db:"created_by"`
	UpdatedAt       *time.Time     `db:"updated_at"`
	UpdatedBy       sql.NullString `db:"updated_by"`
}

type GhostBlogPostTagPointer struct {
	ID        string `db:"id"`
	PostID    string `db:"post_id"`
	TagID     string `db:"tag_id"`
	SortOrder uint32 `db:"sort_order"`
}
