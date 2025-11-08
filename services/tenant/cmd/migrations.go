package main

import "embed"

//go:embed resources/migrations/*.sql
var migrations embed.FS

