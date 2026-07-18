package etc

// type FastPermissions2 struct {

// 	//starter:component

// 	_as func(rbac.PermissionDTO) //starter:as(".")

// }

// // // ListRegistrations implements [permissions.Registry].
// // func (inst *FastPermissions2) ListRegistrations() []*permissions.Registration {

// // 	r1 := &permissions.Registration{}
// // 	r1.Method = "PUT"
// // 	r1.Path = "/api/v1/my/products/:id"
// // 	r1.Enabled = true
// // 	r1.Priority = 999
// // 	r1.Roles = "admin,root"

// // 	return []*permissions.Registration{r1}
// // }

// func (inst *FastPermissions2) _impl()  subjects.GetCurrent() {
// 	return inst
// }
