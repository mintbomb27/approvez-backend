package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	CampaignID primitive.ObjectID `json:"campaignID"`
	Timestamp  int64              `json:"timestamp"`
	LastUpdate int64              `json:"lastUpdate"`
	PostType   string             `json:"postType"`
	Texts      []ContentType      `json:"texts"`
	Images     []ContentType      `json:"images"`
	Deadline   int64              `json:"deadline"`
	Release    int64              `json:"release"`
	Creator    string             `json:"creator"`
	Status     string             `json:"status"`
	Comments   []Comment          `json:"comments"`
}

type ContentType struct {
	RevisionID primitive.ObjectID `json:"revisionID"`
	Iteration  int                `json:"iteration"`
	Content    string             `json:"content"`
	Author     string             `json:"author"`
}

type Comment struct {
	CommentID  primitive.ObjectID `json:"commentID"`
	RevisionID primitive.ObjectID `json:"revisionID"`
	Text       string             `json:"text"`
	Timestamp  int64              `json:"timestamp"`
	Author     string             `json:"author"`
	Type       string             `json:"type"`
}
