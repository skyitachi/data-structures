package trie

import (
  "fmt"
  "unicode/utf8"
)

type Node struct {
  key      rune
  term     bool
  value    interface{}
  size     int64
  children map[rune]*Node // 先不管顺序
}

func reverse(nodeList []*Node) {
  i := 0
  j := len(nodeList) - 1
  for ; i < j; {
    nodeList[i], nodeList[j] = nodeList[j], nodeList[i]
    i++
    j--
  }
}

type Trie struct {
  root *Node
  size int64
}

type RangeResponse struct {
  Total int64
  List []string
}

func NewNode(r rune) *Node {
  return &Node{
    key:      r,
    children: make(map[rune]*Node),
  }
}

func NewTrie() *Trie {
  return &Trie{
    root: &Node{
      children: make(map[rune]*Node),
    },
  }
}

func (t *Trie) Add(key string, value interface{}) {
  var pre []*Node
  cur := t.root
  for _, r := range key {
    pre = append(pre, cur)
    next, ok := cur.children[r]
    if !ok {
      cur.children[r] = NewNode(r)
      next = cur.children[r]
    }
    cur = next
  }
  if !cur.term {
    cur.term = true
    cur.size += 1
    for _, node := range pre {
      node.size += 1
    }
  }
}

func (t *Trie) Get(key string) (interface{}, bool) {
  cur := t.root
  for _, r := range key {
    next, ok := cur.children[r]
    if !ok {
      return nil, false
    }
    cur = next
  }
  return cur.value, true
}

func (t *Trie) HasKey(key string) bool {
  cur := t.root
  for _, r := range key {
    next, ok := cur.children[r]
    if !ok {
      return false
    }
    cur = next
  }
  return cur.term
}

func (t *Trie) HasPrefix(prefix string) bool {
  cur := t.root
  for _, r := range prefix {
    next, ok := cur.children[r]
    if !ok {
      return false
    }
    cur = next
  }
  return true
}

func (t *Trie) search(prefix string, start *Node, count int64, offset *int64) (ret []string) {
  if count <= 0 || start == nil {
    return
  }
  if start.term {
    if *offset == 0 {
      ret = append(ret, prefix)
      count -= 1
    } else {
      *offset -= 1
    }
  }
  for r, child := range start.children {
    newPrefix := prefix + string([]rune{r})
    result := t.search(newPrefix, child, count, offset)
    count -= int64(len(result))
    ret = append(ret, result...)
    if count <= 0 {
      return
    }
  }
  return
}

// range query support
func (t *Trie) PrefixSearchKey(prefix string, offset int64, limit int64) (ret []string) {
  if limit <= 0 || offset < 0 {
    return
  }
  cur := t.root
  for _, r := range prefix {
    next, ok := cur.children[r]
    if !ok {
      return ret
    }
    cur = next
  }
  if cur.size < offset {
    return ret
  }
  if cur.size < limit {
    limit = cur.size
  }
  if cur.term {
    ret = append(ret, string(prefix))
    limit -= 1
  }
  for limit > 0 {
    for r, n := range cur.children {
      if n.size < offset {
        offset -= n.size
      } else {
        if n.size - offset >= limit {
          result := t.search(prefix + string([]rune{r}), n, limit, &offset)
          limit = 0
          ret = append(ret, result...)
          return
        }
        result := t.search(prefix + string([]rune{r}), n, n.size - offset, &offset)
        limit -= n.size - offset
        ret = append(ret, result...)
      }

    }
  }
  return
}

// 删除内存
func (t *Trie) Delete(key string) bool {
  var prev []*Node
  l := utf8.RuneCountInString(key)
  cur := t.root
  prev = append(prev, cur)
  for _, r := range key {
    prev = append(prev, cur)
    next, ok := cur.children[r]
    if !ok {
      return false
    }
    cur = next
  }
  if cur.term {
    cur.term = false
  }
  prev = append(prev, cur)
  reverse(prev)
  for i := 0; i < l + 1; i++ {
    if i > 0 {
      child := prev[i - 1].key
      if prev[i - 1].size == 0 {
        delete(prev[i].children, child)
      }
    }
    prev[i].size -= 1
  }
  return true
}

func (t *Trie) Size() int64 {
  return t.root.size
}

func (t *Trie) Root() *Node {
  return t.root
}

func debug(t *Node) {
  if t == nil {
    return
  }
  fmt.Printf("current node %s, size %d\n", string([]rune{t.key}), t.size)
  for _, v := range t.children {
    debug(v)
  }
}

func (t *Trie) childLen() int64 {
  return int64(len(t.root.children))
}