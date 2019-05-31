package block;

type block struct{
	var (
		data[32] byte
		offset int
	)
	func get() *byte{
		return data
	}
	func set(x *byte) {
		data = x
	}
	func getOffset(){
		return offset
	}
}