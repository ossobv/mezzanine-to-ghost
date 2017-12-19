package main

import (
	"database/sql"
	"encoding/json"
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
		newPost.ID = fmt.Sprintf("%x", string(bson.NewObjectId())) // This is what Ghost uses
		fmt.Printf("newPost.ID: %s", newPost.ID)

		newPost.UUID = fmt.Sprintf("%s", uuid.NewV4()) // This is what Ghost uses
		newPost.Title = oldPost.Title
		newPost.MetaTitle = oldPost.MetaTitle
		newPost.Slug = oldPost.Slug.String

		// TODO: generate AMP and HTML maybe?
		jsonifiedContent, err := json.Marshal(oldPost.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Wrap markdown in Mobiledoc representation
		// From Ghost 1.0, the posts have a mobiledoc field instead of a markdown field.
		// To save editable content, you'll need to wrap any markdown in the MobileDoc representation
		wrappedContent := fmt.Sprintf(`
{
  "version":"0.3.1",
  "markups":[],
  "atoms":[],
  "cards":[
    ["card-markdown", {
      "cardName":"card-markdown",
      "markdown":%s
    }]
  ],
  "sections":[[10,0]]
}
`, jsonifiedContent)

		newPost.Mobiledoc = sql.NullString{String: wrappedContent, Valid: len(oldPost.Content) > 0}
		newPost.MetaDescription = sql.NullString{String: oldPost.Description, Valid: len(oldPost.Description) > 0}
		newPost.CustomExcerpt = sql.NullString{String: oldPost.Description, Valid: len(oldPost.Description) > 0}

		// TODO: find out which timezone Mezzanine uses and convert accordingly to UTC
		if oldPost.Created != nil {
			newPost.CreatedAt = *oldPost.Created
		} else {
			newPost.CreatedAt = *oldPost.PublishDate
		}

		newPost.UpdatedAt = oldPost.Updated

		newPost.PublishedAt = *oldPost.PublishDate

		// TODO: Convert status
		newPost.Status = "published"

		if _, err = ghostPosts.Insert(&newPost); err != nil {
			fmt.Println(err.Error())
			return
		

		fmt.Printf("Imported blog %s", newPost.Title)
		fmt.Println()

		// Create tags beloning to this post
		// TODO: also combine/merge Mezzanine categories
		oldTags := strings.Split(oldPost.Keywords, " ")

		for keyword := range oldTags {
			newTag := GhostBlogPostTag{
				ID: fmt.Sprintf("%x", string(bson.NewObjectId()))
				UUID: fmt.Sprintf("%s", uuid.NewV4())
				Name: keyword,
				Slug: keyword				
			}
		}
	}
}
