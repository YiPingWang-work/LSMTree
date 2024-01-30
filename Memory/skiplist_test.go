package Memory

import (
	"log"
	"math/rand"
	"testing"
)

type cc struct {
	k int
}

func (c2 cc) Less(c Comparable) bool {
	return c2.k < c.(cc).k
}

func (c2 cc) Greater(c Comparable) bool {
	return c2.k > c.(cc).k
}

func (c2 cc) Equal(c Comparable) bool {
	return c2.k == c.(cc).k
}

func (c2 cc) Minit() Comparable {
	return cc{k: -1}
}

func (c2 cc) Maxit() Comparable {
	return cc{k: 1e9}
}

func TestQuickList(t *testing.T) {
	log.Println("====")
	for i := 0; i < 1000000; i++ {

	}
	log.Println("====")
	q := SkipList{}
	q.Init(cc{})
	log.Println("begin")
	for i := 0; i < 1000000; i++ {
		_ = q.Set(cc{rand.Intn(10000000)}, i)
	}
	log.Println("---")
	x := 0
	for i := 0; i < 1000000; i++ {
		if _, err := q.Get(cc{rand.Intn(100000)}); err != nil {
			x++
		}
		//fmt.Printf("%d ", h)
	}
	log.Println("---", x)
	log.Println(q.Get(cc{4}))
	for _, v := range q.heads {
		iter := v.first
		n := 0
		i := cc{-2}
		for iter != nil {
			if (iter.key).Less(i) {
				panic("err")
			} else {
				i = (iter.key).(cc)
			}
			//fmt.Printf("%v, ", iter.key)
			iter = iter.next
			n++
		}
		log.Println(n)
	}
}

func TestX(t *testing.T) {
	q := SkipList{}
	q.Init(cc{})
	for i := 0; i < 1000000; i++ {
		q.Set(cc{rand.Intn(1000000)}, i)
	}
	//for _, v := range q.heads {
	//	iter := v.first
	//	i := cc{-2}
	//	for iter != nil {
	//		if (iter.key).Less(i) {
	//			panic("err")
	//		} else {
	//			i = (iter.key).(cc)
	//		}
	//		fmt.Printf("%v, ", iter.key)
	//		iter = iter.next
	//	}
	//	fmt.Println()
	//}
}
