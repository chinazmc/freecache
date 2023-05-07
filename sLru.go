package freecache

import "container/list"

type sLRU struct {
	sOneCap, sTwoCap int
	sOne, sTwo       *list.List
}

func newSLRU(sOneCap, sTwoCap int) *sLRU {
	return &sLRU{
		sOneCap: sOneCap,
		sTwoCap: sTwoCap,
		sOne:    list.New(),
		sTwo:    list.New(),
	}
}
func (s *sLRU) victim() []byte {
	s.sOne.re
}
