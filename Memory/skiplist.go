package Memory

import (
	"errors"
	"math/rand"
	"time"
)

type Comparable interface {
	Less(c Comparable) bool
	Greater(c Comparable) bool
	Equal(c Comparable) bool
	Minit() Comparable
	Maxit() Comparable
}

type SkipList struct {
	top   *node
	heads []pair
	n     int
	h     int
}

type node struct {
	next *node
	down *node
	key  Comparable
	val  interface{}
}

type pair struct {
	first  *node
	second *node
}

func (s *SkipList) Init(item Comparable) {
	s.n, s.h = 0, 0
	s.top = &node{
		next: &node{
			next: nil,
			down: nil,
			key:  item.Maxit(),
		},
		down: nil,
		key:  item.Minit(),
	}
	s.heads = append(s.heads, pair{s.top, s.top.next})
	rand.Seed(int64(time.Now().Nanosecond()))
}

func (s *SkipList) Get(key Comparable) (interface{}, error) {
	iter := s.top
	for {
		if iter == nil {
			return nil, errors.New("can not find the key")
		}
		if key.Equal(iter.key) {
			return iter.val, nil
		}
		if key.Greater(iter.key) {
			if !key.Less(iter.next.key) {
				iter = iter.next
			} else {
				iter = iter.down
			}
		}
	}
}

func (s *SkipList) Set(key Comparable, val interface{}) error {
	if _, err := s.Get(key); err == nil {
		return errors.New("key exist")
	}
	s.n++
	nh, h := rand.Int(), 0
	for h = 0; h <= s.h; h++ {
		if nh%(1<<h) != 0 {
			h--
			break
		}
	}
	var last *node = nil
	if h > s.h {
		top := &node{
			next: &node{
				next: nil,
				down: s.heads[s.h].second,
				key:  key.Maxit(),
			},
			down: s.heads[s.h].first,
			key:  key.Minit(),
		}
		s.top = top
		s.heads = append(s.heads, pair{top, top.next})
		s.h++
	}
	iter := s.top
	nowh := s.h
	for {
		if iter == nil {
			return nil
		}
		if key.Greater(iter.key) {
			if key.Greater(iter.next.key) {
				iter = iter.next
			} else if key.Less(iter.next.key) {
				// 创建节点
				if nowh <= h {
					item := node{next: iter.next, key: key, val: val}
					iter.next = &item
					if last != nil {
						last.down = &item
					}
					last = &item
				}
				iter = iter.down
				nowh--
			} else {
				panic("?")
			}
		}
	}
}

func (s *SkipList) Remove(key Comparable) error {
	if _, err := s.Get(key); err != nil {
		return err
	}
	iter := s.top
	for {
		if iter == nil {
			return nil
		}
		if key.Greater(iter.key) {
			if key.Greater(iter.next.key) {
				iter = iter.next
			} else if key.Less(iter.next.key) {
				iter = iter.down
			} else {
				iter.next = iter.next.next
				iter = iter.down
			}
		}
	}
}

func (s *SkipList) Size() int {
	return s.n
}
