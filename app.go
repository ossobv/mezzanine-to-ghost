package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ossobv/ghostapi"
	"github.com/satori/go.uuid"
	"github.com/thanhpk/randstr"
	"gopkg.in/mgo.v2/bson"
	"upper.io/db.v3"
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

	now := time.Now()
	apiUser := ghostapi.Client{
		ID:        fmt.Sprintf("%x", string(bson.NewObjectId())),
		UUID:      fmt.Sprintf("%s", UUID.NewV4()),
		Name:      "mezzanine-to-ghost",
		Slug:      "mezzanine-to-ghost",
		Secret:    "lie7teCa",
		Status:    "enabled",
		Type:      "ua",
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: "1",
	}

	// Add API user for this application to Ghost database if not exist
	if matches, err := ghostDB.Collection("clients").Find(db.Cond{"secret": apiUser.Secret, "name": apiUser.Name}).Count(); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		if _, err := ghostDB.Collection("clients").Insert(&apiUser); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	gstapi := ghostapi.Login(&apiUser)

	// Key = User ID
	// Value = User object
	mezzanineUsers := map[int64]MezzanineUser{}

	var mznUser MezzanineUser

	// Retrieve all Mezzanine users
	iter := mezzanineDB.SelectFrom("auth_user").Iterator()
	for iter.Next(&mznUser) {
		mezzanineUsers[mznUser.ID] = mznUser
	}

	// Key = Mezzanine User ID
	// Value = User object
	ghostUsers := map[int64]GhostUser{}

	var gstUser GhostUser

	// Create new Ghost users if not exists
	for mznUserID, mznUser := range mezzanineUsers {
		// If user with the same slug/username already exists, we retrieve that user
		if matches, err := ghostDB.Collection("users").Find(db.Cond{"slug": mznUser.Username}).Count(); matches > 0 || err != nil {
			if err != nil {
				iter.Close()
				fmt.Println(err.Error())
				return
			}

			ghostDB.Collection("users").Find(db.Cond{"slug": mznUser.Username}).One(&gstUser)

			ghostUsers[mznUser.ID] = gstUser
			continue
		}

		gstUser.ID = fmt.Sprintf("%x", string(bson.NewObjectId()))
		gstUser.Name = fmt.Sprintf("%s %s", mznUser.FirstName, mznUser.LastName)
		gstUser.Slug = mznUser.Username

		// Generate random password
		gstUser.Password = randstr.String(16)
		gstUser.Email = mznUser.Email
		gstUser.CreatedAt = mznUser.JoinDate
		gstUser.LastSeen = &mznUser.LastLogin
		gstUser.Visibility = "public"

		// Generate random password

		// First save user
		// gstapi.CreateUser(&gstUser)
		if _, err := ghostDB.Collection("users").Insert(&gstUser); err != nil {
			iter.Close()

			fmt.Println(err.Error())
			return
		}

		// Print password to console
		fmt.Println(fmt.Sprintf("Ghost user password for %s: %s", mznUser.Username, gstUser.Password))

		ghostUsers[mznUserID] = gstUser
	}

	var oldPost MezzanineBlogPost
	newPost := GhostBlogPost{}

	lpr := mezzanineDB.SelectFrom("blog_blogpost").Iterator()
	for lpr.Next(&oldPost) {
		oldAuthor := mezzanineUsers[oldPost.AuthorID]
		newAuthor := ghostUsers[oldAuthor.ID]

		newPost.AuthorID = newAuthor.ID

		newPost.ID = fmt.Sprintf("%x", string(bson.NewObjectId())) // This is what Ghost uses
		fmt.Printf("newPost.ID: %s", newPost.ID)

		newPost.UUID = fmt.Sprintf("%s", uuid.NewV4()) // This is what Ghost uses
		newPost.Title = oldPost.Title
		newPost.MetaTitle = oldPost.MetaTitle
		newPost.Slug = oldPost.Slug.String

		// TODO: generate AMP and HTML maybe?
		jsonifiedContent, err := json.Marshal(oldPost.Content)
		if err != nil {
			iter.Close()

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
		var createdAt time.Time

		if oldPost.Created != nil {
			createdAt = *oldPost.Created
		} else {
			createdAt = *oldPost.PublishDate
		}

		newPost.CreatedAt = createdAt
		newPost.CreatedBy = newAuthor.ID
		newPost.UpdatedAt = oldPost.Updated
		newPost.UpdatedBy = sql.NullString{String: newAuthor.ID, Valid: len(newAuthor.ID) > 0}
		newPost.PublishedAt = *oldPost.PublishDate

		// TODO: Convert status
		newPost.Status = "published"
		newPost.Visibility = "public"

		// gstapi.CreatePost(&newPost)
		if _, err = ghostDB.Collection("posts").Insert(&newPost); err != nil {
			iter.Close()

			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Imported blog %s", newPost.Title)
		fmt.Println()

		// Create tags beloning to this post
		// TODO: also combine/merge Mezzanine categories
		oldTags := strings.Split(oldPost.Keywords, " ")

		var newTag GhostBlogPostTag
		var newTagPointer GhostBlogPostTagPointer

		for _, keyword := range oldTags {
			// If tag already exists, we re-use that tag -- else we create one
			// if gstapi.TagExists(keyword) { } else { }
			if matches, err := ghostDB.Collection("tags").Find(db.Cond{"name": keyword}).Count(); matches > 0 || err != nil {
				if err != nil {
					iter.Close()
					fmt.Println(err.Error())
					return
				}

				ghostDB.Collection("tags").Find(db.Cond{"name": keyword}).One(&newTag)
				continue
			} else {

				newTag.ID = fmt.Sprintf("%x", string(bson.NewObjectId()))
				newTag.Name = keyword
				newTag.Slug = keyword
				newTag.Visibility = "public"
				newTag.CreatedAt = createdAt
				newTag.CreatedBy = newAuthor.ID
				newTag.UpdatedAt = oldPost.Updated
				newTag.UpdatedBy = sql.NullString{String: newAuthor.ID, Valid: len(newAuthor.ID) > 0}

				// gstapi.CreateTag(&newTag)
				if _, err := ghostDB.Collection("tags").Insert(&newTag); err != nil {
					iter.Close()

					fmt.Println(err.Error())
					return
				}
			}

			newTagPointer.ID = fmt.Sprintf("%x", string(bson.NewObjectId()))
			newTagPointer.PostID = newPost.ID
			newTagPointer.TagID = newTag.ID

			// gstapi.CreateTag()
			if _, err := ghostDB.Collection("posts_tags").Insert(&newTagPointer); err != nil {
				iter.Close()

				fmt.Println(err.Error())
				return
			}
		}

		// gstapi.CreatePost(&newPost)
	}
}
