package internal

type SchemasInfo struct {
	Data []*SchemaInfo
}

type SchemaInfo struct {
	Name   string
	Exists bool
}
