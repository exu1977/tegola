package test

import "github.com/go-spatial/geom"

type layer struct {
	name     string
	geomType geom.Geometry
	srid     uint64
	mtag     *string
}

func (l layer) Name() string {
	return l.name
}

func (l layer) GeomType() geom.Geometry {
	return l.geomType
}

func (l layer) SRID() uint64 {
	return l.srid
}

func (l layer) ModificationTag() *string {
	return l.mtag
}
