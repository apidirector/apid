package config

//##############################################################################
// Content Type
//##############################################################################

type ContentType struct {
	Kind       ContentTypeKind
	Info       ContentTypeInfo
	Attributes []ContentTypeAttribute
}

//##############################################################################
// ContentTypeKind
//##############################################################################

type ContentTypeKind string

const (
	ContentTypeKindUnknown ContentTypeKind = "unknown"
	ContentTypeKindJSON    ContentTypeKind = "json"
	ContentTypeKindYAML    ContentTypeKind = "yaml"
)

//##############################################################################
// ContentTypeInfo
//##############################################################################

type ContentTypeInfo struct {
	SingularName string
	PluralName   string
	Description  string
}

//##############################################################################
// ContentTypeAttribute
//##############################################################################

type ContentTypeAttribute struct {
	Name     string
	Type     ContentTypeAttributeType
	Private  bool
	Required bool
}

//##############################################################################
// ContentTypeAttributeType
//##############################################################################

type ContentTypeAttributeType string

const (
	ContentTypeAttributeTypeUnknown ContentTypeAttributeType = "unknown"

	// string types
	ContentTypeAttributeTypeString      ContentTypeAttributeType = "string"
	ContentTypeAttributeTypeText        ContentTypeAttributeType = "text"
	ContentTypeAttributeTypeRichText    ContentTypeAttributeType = "rich_text"
	ContentTypeAttributeTypeEnumeration ContentTypeAttributeType = "enumeration"
	ContentTypeAttributeTypeEmail       ContentTypeAttributeType = "email"
	ContentTypeAttributeTypePassword    ContentTypeAttributeType = "password"
	ContentTypeAttributeTypeURL         ContentTypeAttributeType = "url"

	// number types
	ContentTypeAttributeTypeInt     ContentTypeAttributeType = "int"
	ContentTypeAttributeTypeBigint  ContentTypeAttributeType = "bigint"
	ContentTypeAttributeTypeFloat   ContentTypeAttributeType = "float"
	ContentTypeAttributeTypeDecimal ContentTypeAttributeType = "decimal"

	// boolean types
	ContentTypeAttributeTypeBool ContentTypeAttributeType = "bool"

	// time types
	ContentTypeAttributeTypeTime      ContentTypeAttributeType = "time"
	ContentTypeAttributeTypeDate      ContentTypeAttributeType = "date"
	ContentTypeAttributeTypeDateTime  ContentTypeAttributeType = "date_time"
	ContentTypeAttributeTypeTimestamp ContentTypeAttributeType = "timestamp"
)
