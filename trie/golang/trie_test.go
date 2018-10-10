package trie

import (
  "testing"
  "os"
)

var trie *Trie

func TestTrie_Add_Delete(t *testing.T) {
  trie.Add("浙江省", 1)
  trie.Add("浙江省杭州市", 1)
  if !trie.HasKey("浙江省"){
    t.Errorf("error")
  }
  if !trie.HasKey("浙江省杭州市") {
    t.Errorf("no such key")
  }
  ret := trie.Delete("浙江省")
  if !ret {
    t.Errorf("Delete Error")
    t.Fail()
  }
  if trie.childLen() != 1 {
    t.Errorf("Unexpect children size: expect 1 found %d", trie.childLen())
    t.Fail()
  }
  ret = trie.Delete("浙江省杭州市")
  if !ret {
   t.Errorf("Delete Error: expected true found %v", ret)
   t.Fail()
  }
  if trie.childLen() != 0 {
   t.Errorf("Unexpect children size: expect 0 found %d", trie.childLen())
   t.Fail()
  }
}

func TestTrie_PrefixSearchKey(t *testing.T) {
  trie.Add("浙江省", 1)
  trie.Add("浙江省杭州市", 1)
  trie.Add("浙江省杭州市移动", 1)
  ret := trie.PrefixSearchKey("浙", 0, 1)
  t.Log("current ret: ", ret)
  if len(ret) != 1 {
    t.Errorf("error count %d, %v, expect %d", len(ret), ret, 1)
  }
  ret = trie.PrefixSearchKey("浙", 0, 2)
  if len(ret) != 2 {
    t.Errorf("error count %d, %v, expect %d", len(ret), ret, 2)
  }
  ret = trie.PrefixSearchKey("浙", 0, 3)
  if len(ret) != 3 {
    t.Errorf("error count %d, %v, expect %d", len(ret), ret, 3)
  }
  ret = trie.PrefixSearchKey("浙", 1, 2)
  t.Log("current ret: ", ret)
  if len(ret) != 2 {
   t.Errorf("error count %d, %v, expect %d", len(ret), ret, 2)
  }
}

func TestMain(m *testing.M) {
  trie = NewTrie()
  trie.Add("浙江省", 1)
  trie.Add("浙江省", 1)

  code := m.Run()
  os.Exit(code)
}
