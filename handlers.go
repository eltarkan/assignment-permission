package main

import (
	"assignment-permission/models"
	"assignment-permission/restapi/operations/health"
	permissionHandler "assignment-permission/restapi/operations/permission"
	roleHandler "assignment-permission/restapi/operations/role"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"strings"
)

func RoleGetRoleHandler(params roleHandler.GetRoleHandler) roleHandler.GetRoleHandlerFunc {
	return func(params roleHandler.GetRoleParams) middleware.Responder {
		roles, err := ListAllRoles()
		if err != nil {
			return middleware.NotImplemented("GET")
		}

		log.Println(roles)
		var roleListResponse models.RoleListAllResponse
		for _, _role := range *roles {
			roleListResponse = append(roleListResponse, &_role)
		}

		response := &roleHandler.GetRoleOK{
			Payload: roleListResponse,
		}
		return roleHandler.NewGetRoleOK().WithPayload(response.Payload)
	}
}

func RolePostRoleHandler(params roleHandler.PostRoleHandler) roleHandler.PostRoleHandlerFunc {
	return func(params roleHandler.PostRoleParams) middleware.Responder {
		_, err := CreateRole(params.Body)
		if err != nil {
			return middleware.NotImplemented("Something went wrong!")
		}
		response := &roleHandler.PostRoleOK{
			Payload: &models.RoleCreateSuccessResponse{
				ID:          params.Body.ID,
				Name:        params.Body.Name,
				Description: params.Body.Description,
			},
		}
		return roleHandler.NewPostRoleOK().WithPayload(response.Payload)
	}
}

func PermissionGetPermissionHandler(params permissionHandler.GetPermissionHandler) permissionHandler.GetPermissionHandlerFunc {
	return func(params permissionHandler.GetPermissionParams) middleware.Responder {

		var userID string

		// totally wrong implementation but I need some time
		if len(strings.Split(params.HTTPRequest.RequestURI, "=")) > 1 {
			userID = strings.Split(params.HTTPRequest.RequestURI, "=")[1]
		}

		log.Println(userID)
		log.Println("#######")
		userPermissions, err := GetPermissionByUserID(userID)
		if err != nil {
			return middleware.NotImplemented("Something went wrong! Please don't do it again :/")
		}

		log.Println(userPermissions)

		var response models.PermissionListAll
		for _, _permission := range *userPermissions {
			response = append(response, &_permission)
		}

		return permissionHandler.NewGetPermissionOK().WithPayload(response)
	}
}

func PermissionPostPermissionHandler(params permissionHandler.PostPermissionHandler) permissionHandler.PostPermissionHandlerFunc {
	return func(params permissionHandler.PostPermissionParams) middleware.Responder {
		_, err := CreateNewPermission(params.Body)
		if err != nil {
			return middleware.NotImplemented("Something went wrong!")
		}
		response := &permissionHandler.PostPermissionOK{
			Payload: &models.PermissionCreateSuccessResponse{
				Resource: params.Body.Resource,
				Groups:   params.Body.Groups,
				Users:    params.Body.Users,
			},
		}
		return permissionHandler.NewPostPermissionOK().WithPayload(response.Payload)
	}
}

func HealthHandler(params health.GetHandler) health.GetHandlerFunc {
	return func(params health.GetParams) middleware.Responder {
		fmt.Println(params.HTTPRequest.RequestURI)
		return middleware.NotImplemented("GET")
	}
}
