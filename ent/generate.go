package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./gen --feature privacy,entql,schema/snapshot,sql/schemaconfig,sql/modifier,sql/upsert ./schema
