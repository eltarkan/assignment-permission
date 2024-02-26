package main

import (
	"assignment-permission/config"
	"assignment-permission/models"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Group struct {
	GroupID string   `json:"group_id" bson:"group_id"`
	Role    string   `json:"role" bson:"role"`
	Members []string `json:"member_ids" bson:"member_ids"`
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

func CreateNewPermission(permission permission) (*mongo.InsertOneResult, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("permissions")

	record, err := GroupRepository.InsertOne(context.Background(), permission)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, nil
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

func GetRoleByID(roleID string) (*role, error) {
	var GroupRepository = config.DBConnection.Database(config.Config.MongodbName).Collection("roles")

	var role role
	err := GroupRepository.FindOne(context.Background(), map[string]string{"role_id": roleID}).Decode(&role)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &role, nil
}
