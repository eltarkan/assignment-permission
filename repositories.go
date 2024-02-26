package main

import (
	"assignment-permission/config"
	"assignment-permission/models"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Group struct {
	GroupID string   `json:"group_id" bson:"group_id"`
	Role    string   `json:"role" bson:"role"`
	Members []string `json:"member_ids" bson:"member_ids"`
}

type Permission struct {
	Resource string            `bson:"resource"`
	Users    []userPermission  `bson:"users"`
	Groups   []groupPermission `bson:"groups"`
}

func CreateNewGroup(group Group) (*mongo.InsertOneResult, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("groups")

	record, err := GroupRepository.InsertOne(context.Background(), group)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, nil
}

func GetGroupByID(groupID string) (*Group, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("groups")

	var group Group
	err := GroupRepository.FindOne(context.Background(), map[string]string{"group_id": groupID}).Decode(&group)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &group, nil
}

func GetPermissionByResourceName(resourceName string) (*Group, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("permissions")

	var group Group
	err := GroupRepository.FindOne(context.Background(), map[string]string{"resource": resourceName}).Decode(&group)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &group, nil
}

func CreateNewPermission(permission *models.PermissionCreateRequest) (*mongo.InsertOneResult, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("permissions")

	record, err := GroupRepository.InsertOne(context.Background(), permission)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, nil
}

func GetPermissionByUserID(userID string) (*[]Permission, error) {
	var PermissionRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("permissions")

	var group []Permission
	var filter bson.M
	if userID != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"users.user_id": userID},
				{"groups.members": bson.M{"$elemMatch": bson.M{"$eq": userID}}},
			},
		}
	}
	cursor, err := PermissionRepository.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = cursor.All(context.Background(), &group)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &group, nil
}

func CreateRole(role *models.RoleCreateRequest) (*mongo.InsertOneResult, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("roles")

	record, err := GroupRepository.InsertOne(context.Background(), role)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, nil
}

func ListAllRoles() (*[]role, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("roles")

	var roles []role
	cursor, err := GroupRepository.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = cursor.All(context.Background(), &roles)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &roles, nil
}
