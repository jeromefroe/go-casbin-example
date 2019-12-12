package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("unable to create Casbin enforcer: %v", err)
	}

	aliceRoles, err := e.GetRolesForUser("alice")
	if err != nil {
		log.Fatalf("could not get roles for alice: %v", err)
	}
	alicePerms, err := e.GetImplicitPermissionsForUser("alice")
	if err != nil {
		log.Fatalf("could not get permissions for alice: %v", err)
	}
	fmt.Printf(
		"alice is a member of the following roles: %v, and her permissions are: %v\n",
		aliceRoles,
		alicePerms,
	)

	bobRoles, err := e.GetRolesForUser("bob")
	if err != nil {
		log.Fatalf("could not get roles for bob: %v", err)
	}
	bobPerms, err := e.GetImplicitPermissionsForUser("bob")
	if err != nil {
		log.Fatalf("could not get permissions for bob: %v", err)
	}
	fmt.Printf(
		"bob is a member of the following roles: %v, and his permissions are: %v\n",
		bobRoles,
		bobPerms,
	)
}
