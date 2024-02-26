package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"assignment-permission/restapi"
	"assignment-permission/restapi/operations"
)

func StartServer() {
	/*
	 * I move the swagger server to the server.go file
	 * because every time I run the swagger generate command the main.go file is overwritten.
	 */
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewPermissionsAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "A Permission application"
	parser.LongDescription = "Qubitro Task"
	server.ConfigureFlags()

	api.HealthGetHandler = HealthHandler(api.HealthGetHandler)

	api.RoleGetRoleHandler = RoleGetRoleHandler(api.RoleGetRoleHandler)
	api.RolePostRoleHandler = RolePostRoleHandler(api.RolePostRoleHandler)

	api.PermissionGetPermissionHandler = PermissionGetPermissionHandler(api.PermissionGetPermissionHandler)
	api.PermissionPostPermissionHandler = PermissionPostPermissionHandler(api.PermissionPostPermissionHandler)

	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	server.Port = 8080
	server.Host = "0.0.0.0"

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
