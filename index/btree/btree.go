package btree

type nodeType int

const (
    nodeType1 nodeType = 0
    nodeType2 nodeType = 1
    nodeType3 nodeType = 2
    nodeType4 nodeType = 3
)

type btree struct {
    root  *node
    order int
}

type node struct {
    key   Value
    val   Value
    typ   nodeType
    nodes nodes
    left  *node
    right *node
}

type nodes []*node

func (ns nodes) len() int {
    return ns.len()
}

func (ns *nodes) find(key, val Value) (n, l, r *node) {
    if ns.len() == 0 {
        return nil, nil, nil
    }
    for _, cur := range *ns {
        if cur.key.Compare(key) == 1{
            return cur,
        }
    }
}

type Value interface {
    Val() interface{}
    Compare(o Value) int
}

func New(order int) *btree {
    return &btree{
        order: order,
    }
}

func (b *btree) Insert(key, val Value) *node {
    if b.root == nil {
        b.root = &node{key: key, val: val}
        return b.root
    }
    
}
