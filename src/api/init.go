package api

import (
	v1 "src/api/v1"
)

// iInitialize all api versions (db connections, etc.)
func Init() {
	v1.Init()
}
