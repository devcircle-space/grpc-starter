package main

import (
	"context"
	"log"
	"time"

	"content-service/src/db"
	contentpb "content-service/src/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type contentServer struct {
	contentpb.UnimplementedContentServiceServer
}

func (s *contentServer) GetContentById(ctx context.Context, req *contentpb.GetContentByIdRequest) (*contentpb.GetContentByIdResponse, error) {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Printf("MongoDB connection error: %v", err)
		return nil, err
	}
	collection := client.Database("Manoyukti").Collection("Content")
	var doc bson.M
	err = collection.FindOne(ctx, bson.M{"_id": parseObjectID(req.GetId())}).Decode(&doc)
	if err != nil {
		log.Printf("MongoDB findOne error: %v", err)
		return nil, err
	}

	// Map MongoDB document to gRPC Content message
	content := &contentpb.Content{}
	if oid, ok := doc["_id"]; ok {
		switch v := oid.(type) {
		case primitive.ObjectID:
			content.XId = v.Hex()
		case string:
			content.XId = v
		}
	}
	if title, ok := doc["title"]; ok {
		if t, ok := title.(string); ok {
			content.Title = t
		}
	}
	if desc, ok := doc["description"]; ok {
		if d, ok := desc.(string); ok {
			content.Description = d
		}
	}
	if createdAt, ok := doc["createdAt"]; ok {
		switch v := createdAt.(type) {
		case primitive.DateTime:
			content.CreatedAt = timestamppb.New(v.Time())
		case int64:
			content.CreatedAt = timestamppb.New(primitive.DateTime(v).Time())
		case string:
			if t, err := primitiveParseDateTimeFromString(v); err == nil {
				content.CreatedAt = timestamppb.New(t.Time())
			}
		}
	}
	if updatedAt, ok := doc["updatedAt"]; ok {
		switch v := updatedAt.(type) {
		case primitive.DateTime:
			content.UpdatedAt = timestamppb.New(v.Time())
		case int64:
			content.UpdatedAt = timestamppb.New(primitive.DateTime(v).Time())
		case string:
			if t, err := primitiveParseDateTimeFromString(v); err == nil {
				content.UpdatedAt = timestamppb.New(t.Time())
			}
		}
	}

	return &contentpb.GetContentByIdResponse{Content: content}, nil
}

func (s *contentServer) ListContent(ctx context.Context, req *contentpb.ListContentRequest) (*contentpb.ListContentResponse, error) {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Printf("MongoDB connection error: %v", err)
		return nil, err
	}
	collection := client.Database("Manoyukti").Collection("Content")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("MongoDB find error: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var contents []*contentpb.Content
	for cur.Next(ctx) {
		var doc bson.M
		if err := cur.Decode(&doc); err != nil {
			log.Printf("Decode error: %v", err)
			continue
		}
		// Map MongoDB document to gRPC Content message
		content := &contentpb.Content{}
		if oid, ok := doc["_id"]; ok {
			switch v := oid.(type) {
			case primitive.ObjectID:
				content.XId = v.Hex()
			case string:
				content.XId = v
			}
		}
		if title, ok := doc["title"]; ok {
			if t, ok := title.(string); ok {
				content.Title = t
			}
		}
		if desc, ok := doc["description"]; ok {
			if d, ok := desc.(string); ok {
				content.Description = d
			}
		}
		if createdAt, ok := doc["createdAt"]; ok {
			switch v := createdAt.(type) {
			case primitive.DateTime:
				content.CreatedAt = timestamppb.New(v.Time())
			case int64:
				content.CreatedAt = timestamppb.New(primitive.DateTime(v).Time())
			case string:
				if t, err := primitiveParseDateTimeFromString(v); err == nil {
					content.CreatedAt = timestamppb.New(t.Time())
				}
			}
		}
		if updatedAt, ok := doc["updatedAt"]; ok {
			switch v := updatedAt.(type) {
			case primitive.DateTime:
				content.UpdatedAt = timestamppb.New(v.Time())
			case int64:
				content.UpdatedAt = timestamppb.New(primitive.DateTime(v).Time())
			case string:
				if t, err := primitiveParseDateTimeFromString(v); err == nil {
					content.UpdatedAt = timestamppb.New(t.Time())
				}
			}
		}
		contents = append(contents, content)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}
	return &contentpb.ListContentResponse{Contents: contents}, nil
}

// Helper to parse string to ObjectID if possible
func parseObjectID(id string) interface{} {
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		return oid
	}
	return id
}

// Helper for string to DateTime (if needed)
func primitiveParseDateTimeFromString(s string) (primitive.DateTime, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return 0, err
	}
	return primitive.NewDateTimeFromTime(t), nil
}
