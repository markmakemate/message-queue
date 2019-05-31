package MQ

type job interface{
	do()
}
type worker struct{
	id int
	JobChannel chan job
}
type workerpool struct{
	WorkerChannel chan worker
}
func (w *worker) do(){

}