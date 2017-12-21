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
