package mapper

type StructMapper interface {
}

type structMapper struct {
}

func NewStructMapper() StructMapper {
	return &structMapper{}
}
