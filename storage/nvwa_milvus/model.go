package nvwa_milvus

import "github.com/milvus-io/milvus-sdk-go/v2/entity"

const (
	IdColumnName         = "id"
	embeddingColumnName  = "embedding"
	importanceColumnName = "importance"
	contentColumnName    = "content"
)

var (
	idField = &entity.Field{
		Name:        IdColumnName,
		PrimaryKey:  true,
		AutoID:      false,
		Description: "",
		DataType:    entity.FieldTypeInt64,
		TypeParams:  nil,
		IndexParams: nil,
	}
	embeddingField = &entity.Field{
		Name:        embeddingColumnName,
		PrimaryKey:  false,
		AutoID:      false,
		Description: "",
		DataType:    entity.FieldTypeFloatVector,
		TypeParams:  map[string]string{"dim": "1536"},
		IndexParams: nil,
	}
	importanceField = &entity.Field{
		Name:        importanceColumnName,
		PrimaryKey:  false,
		AutoID:      false,
		Description: "",
		DataType:    entity.FieldTypeFloat,
		TypeParams:  nil,
		IndexParams: nil,
	}
	contentField = &entity.Field{
		Name:        contentColumnName,
		PrimaryKey:  false,
		AutoID:      false,
		Description: "",
		DataType:    entity.FieldTypeVarChar,
		TypeParams:  map[string]string{"max_length": "4096"},
		IndexParams: nil,
	}
	memoryFields = []*entity.Field{idField, embeddingField, importanceField, contentField}
)
