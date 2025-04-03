package handlers

import (
	"context"
	"fiberpress-api/database"
	"fiberpress-api/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllUsers (Admin Only)
func GetAllUsers(c *fiber.Ctx) error {
	userCollection := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{"isDeleted": false})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return c.JSON(users)
}

// GetSingleUser (Admin Only)
func GetSingleUser(c *fiber.Ctx) error {
	userCollection := database.DB.Collection("users")
	userID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID, "isDeleted": false}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// GetProfile (Authenticated User)
func GetProfile(c *fiber.Ctx) error {
	userCollection := database.DB.Collection("users")
	userID := c.Locals("userId").(string)

	objID, err := primitive.ObjectIDFromHex(userID)

	fmt.Println(userID, objID.Hex())
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID, "isDeleted": false}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Remove password before sending response
	user.Password = ""

	return c.JSON(user)
}


// UpdateProfile (Authenticated User)
func UpdateProfile(c *fiber.Ctx) error {
	userCollection := database.DB.Collection("users")
	userID := c.Locals("userId").(string)

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Extract validated data from middleware
	updateData := c.Locals("validatedData").(*models.UpdateProfile)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":      updateData.Name,
			"phone":     updateData.Phone,
			"updatedAt": time.Now(),
		},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objID, "isDeleted": false}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}


// DeleteProfile (Authenticated User - Soft Delete)
func DeleteProfile(c *fiber.Ctx) error {
	userCollection := database.DB.Collection("users")
	userID := c.Locals("userId").(string)

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"isDeleted": true,
			"updatedAt": time.Now(),
		},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile deleted successfully"})
}
