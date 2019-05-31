package Cache

import "block";
type queue struct{
	var(
		data nil
		topic byte
		partition uint32
		maxsize int
	)
	func New(size int){
		data := make(chan block, size)
		maxsize := size
	}
	func init(Topic byte, Partition uint32){
		topic := Topic
		partition := Partition
	}
	func push(x block){
		x.getOffset()++
		data <- x
	}
	func front() block{
		var x block
		x <- data
		return x
	}
	func size() int{
		return len(data)
	}
	func getTopic() byte{
		return topic
	}
	func getPartition() uint32{
		return partition
	}
}