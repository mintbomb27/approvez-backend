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

func CreateCampaign(c *fiber.Ctx) error {
	var data map[string]string
	var campaign models.Campaign
	var userID primitive.ObjectID

	err := CheckIfAuthorized(c, &userID)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	campaign.ID = primitive.NewObjectID()
	campaign.CreatedBy = userID
	campaign.Name = data["name"]
	campaign.Status = data["status"]
	campaign.CoverImage = data["coverImage"]
	campaign.TimeCreated = time.Now().UnixNano() / 1e6

	collection := database.MongoClient.Database("approvEZDB").Collection("campaigns")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, campaign)
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

func GetCampaigns(c *fiber.Ctx) error {
	var userID primitive.ObjectID

	err := CheckIfAuthorized(c, &userID)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	collection := database.MongoClient.Database("approvEZDB").Collection("campaigns")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	var campaigns []bson.M
	if err = cursor.All(ctx, &campaigns); err != nil {
		log.Fatal(err)
	}

	return c.JSON(campaigns)
}
