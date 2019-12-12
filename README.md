# Go Casbin Example

This repository contains an example of using the [Go Casbin client] to implement a simple
[role-based access control] system for two users, Alice and Bob.

```shell
$ go run main.go
alice is a member of the following roles: [role:admin], and her permissions are: [[role:admin data read] [role:admin data write]]
bob is a member of the following roles: [role:user], and his permissions are: [[role:user data read]]
```

[go casbin client]: https://github.com/casbin/casbin
[role-based access control]: https://en.wikipedia.org/wiki/Role-based_access_control
