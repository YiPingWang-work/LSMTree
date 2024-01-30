package LSMTree

import "LSMTree/Memory"

type LSMTree struct {
	active Memory.SkipList
	static Memory.SkipList
}
