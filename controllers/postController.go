package controllers

import (
	"approvez-backend/database"
	"approvez-backend/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *fiber.Ctx) error {
	var userID primitive.ObjectID
	// var data map[string]string
	var post models.Post

	err := CheckIfAuthorized(c, &userID)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if err := c.BodyParser(&post); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	collection := database.MongoClient.Database("approvEZDB").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Finding User name
	var creator models.User
	usersCollection := database.MongoClient.Database("approvEZDB").Collection("users")
	err = usersCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&creator)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	post.Creator = creator.Name
	post.ID = primitive.NewObjectID()
	newRevision := primitive.NewObjectID()
	post.Timestamp = time.Now().Unix()
	if post.Images != nil && post.Texts != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "both images and texts can't be empty",
		})
	}

	if post.Images != nil {
		// Multiple iterations of the same revision
		var image models.ContentType
		for index, value := range post.Images {
			image.Author = creator.Name
			image.RevisionID = newRevision
			image.Iteration = index + 1
			image.Content = value.Content
			post.Images = append(post.Images, image)
		}
	}
	if post.Texts != nil {
		// Multiple iterations of the same revision
		var text models.ContentType
		for index, value := range post.Texts {
			text.Author = creator.Name
			text.RevisionID = newRevision
			text.Iteration = index + 1
			text.Content = value.Content
			post.Texts = append(post.Texts, text)
		}
	}
	if post.Comments != nil {
		var comment models.Comment
		comment.Author = creator.Name
		comment.CommentID = primitive.NewObjectID()
		comment.Timestamp = time.Now().Unix()
		comment.RevisionID = newRevision
		comment.Text = post.Comments[0].Text
		comment.Type = post.Comments[0].Type
		post.Comments = append(post.Comments, comment)
	}

	_, err = collection.InsertOne(ctx, post)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdatePost(c *fiber.Ctx) error {
	return nil
}

func DeletePost(c *fiber.Ctx) error {
	return nil
}

func GetPosts(c *fiber.Ctx) error {
	var userID primitive.ObjectID
	err := CheckIfAuthorized(c, &userID)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	campaignID, err := primitive.ObjectIDFromHex(c.Query("campaignID"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "campaignID unable to parse",
		})
	}

	collection := database.MongoClient.Database("approvEZDB").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var posts []bson.M
	cursor, err := collection.Find(ctx, bson.M{"campaignid": campaignID})
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// for cursor.Next(ctx) {
	// 	var post models.Post
	// 	var creator models.User

	// 	if err:= cursor.Decode(&post); err != nil {
	// 		c.Status(fiber.StatusInternalServerError)
	// 		return c.JSON(fiber.Map{
	// 			"message": err,
	// 		})
	// 	}

	// collection := database.MongoClient.Database("approvEZDB").Collection("users")
	// err := collection.FindOne(ctx, bson.M{"campaignid": campaignID}).Decode(&creator)
	// if err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": err,
	// 	})
	// post.Creator = creator.Name

	// }

	// }

	if err = cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}

	// for index, post := range posts {
	// 	err := json.Unmarshal([])
	// 	creator, _ := primitive.ObjectIDFromHex(post["creator"])

	// }

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func GetPost(c *fiber.Ctx) error {
	return nil
}
