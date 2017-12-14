package main

import (
	"database/sql"
	"fmt"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"upper.io/db.v3/mysql"
)

func main() {
	config, err := mezzanineDBConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mezzanineDB, err := mysql.Open(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer mezzanineDB.Close()

	config, err = ghostDBConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ghostDB, err := mysql.Open(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ghostDB.Close()

	q := mezzanineDB.SelectFrom("blog_blogpost")
	iter := q.Iterator()
	ghostPosts := ghostDB.Collection("posts")

	var oldPost MezzanineBlogPost
	newPost := GhostBlogPost{}

	for iter.Next(&oldPost) {
		newPost.ID = string(bson.NewObjectId()) // This is what Ghost uses
		newPost.UUID = uuid.NewV4().String()    // This is what Ghost uses
		newPost.Title = oldPost.Title
		newPost.Slug = oldPost.Slug.String

		// TODO: generate AMP and HTML maybe?
		// TODO: Convert Markdown to Mobiledoc?

		newPost.PlainText = sql.NullString{String: oldPost.Content, Valid: len(oldPost.Content) > 0}
		newPost.MetaDescription = sql.NullString{String: oldPost.Description, Valid: len(oldPost.Description) > 0}

		if oldPost.Created != nil {
			newPost.CreatedAt = *oldPost.Created
		}
		newPost.UpdatedAt = oldPost.Updated

		// TODO: Convert status

		if _, err = ghostPosts.Insert(&newPost); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Imported blog %s", newPost.Title)
		fmt.Println()
	}
}
