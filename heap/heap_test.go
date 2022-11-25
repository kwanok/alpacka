package heap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestNode struct {
	Value    int
	Distance int
}

type HeapTestSuite struct {
	suite.Suite
	heap Heap[TestNode]
}

func (suite *HeapTestSuite) SetupTest() {
	suite.heap = New[TestNode](func(target, with TestNode) bool {
		return target.Distance > with.Distance
	})
}

func (suite *HeapTestSuite) TestHeap() {
	suite.heap.Push(TestNode{Value: 1, Distance: 5})
	suite.heap.Push(TestNode{Value: 2, Distance: 3})
	suite.heap.Push(TestNode{Value: 3, Distance: 2})
	suite.heap.Push(TestNode{Value: 4, Distance: 7}) // 두번째로큼
	suite.heap.Push(TestNode{Value: 5, Distance: 1})
	suite.heap.Push(TestNode{Value: 6, Distance: 9}) // 제일큼
	suite.heap.Push(TestNode{Value: 7, Distance: 1})

	suite.heap.PrintTree()

	node, err := suite.heap.Pop()
	suite.Require().NoError(err)
	suite.Require().Equal(6, node.Value)

	suite.heap.PrintTree()

	node, err = suite.heap.Pop()
	suite.Require().NoError(err)
	suite.Require().Equal(4, node.Value)
}

func TestHeapTestSuite(t *testing.T) {
	suite.Run(t, new(HeapTestSuite))
}
