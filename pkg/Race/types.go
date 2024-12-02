package race

import (
	class "github.com/CyTechNomad/pc-session-aid/pkg/Class"
	stats "github.com/CyTechNomad/pc-session-aid/pkg/Stats"
)

type Race interface {
	GetName() string
	GetFavoredClass() class.Class
	GetSize() stats.Size
	GetSpeed() uint
}

type race struct {
	Name         string
	Size         stats.Size
	Speed        uint
	favoredClass class.Class
}
