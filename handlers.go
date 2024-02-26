package main

import (
	"assignment-permission/models"
	"assignment-permission/restapi/operations/health"
	permissionHandler "assignment-permission/restapi/operations/permission"
	roleHandler "assignment-permission/restapi/operations/role"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

func RoleGetRoleHandler(params roleHandler.GetRoleHandler) roleHandler.GetRoleHandlerFunc {
	return func(params roleHandler.GetRoleParams) middleware.Responder {
		fmt.Println(params.HTTPRequest.RequestURI)
		return middleware.NotImplemented("GET")
	}
}

func RolePostRoleHandler(params roleHandler.PostRoleHandler) roleHandler.PostRoleHandlerFunc {
	return func(params roleHandler.PostRoleParams) middleware.Responder {
		_, err := CreateRole(params.Body)
		if err != nil {
			return middleware.NotImplemented("POST")
		}

		//response := &group.PostGroupOK{
		//	Payload: &models.GroupCreateSuccessResponse{
		//		ID:   "1",
		//		Name: "group1",
		//	},
		//}
		//// return param body
		//return group.NewPostGroupOK().WithPayload(response.Payload)

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
		fmt.Println(params.HTTPRequest.RequestURI)
		return middleware.NotImplemented("GET")
	}
}

func PermissionPostPermissionHandler(params permissionHandler.PostPermissionHandler) permissionHandler.PostPermissionHandlerFunc {
	return func(params permissionHandler.PostPermissionParams) middleware.Responder {
		fmt.Println(params.Body)
		return middleware.NotImplemented("POST")
	}
}

func HealthHandler(params health.GetHandler) health.GetHandlerFunc {
	return func(params health.GetParams) middleware.Responder {
		fmt.Println(params.HTTPRequest.RequestURI)
		return middleware.NotImplemented("GET")
	}
}
