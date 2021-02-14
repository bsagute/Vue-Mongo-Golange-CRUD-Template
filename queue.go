package queue

import "errors"

//Queue - struct of queue
type Queue struct {
	QueueList []interface{}
	Front     interface{}
	Rear      interface{}
}

//EnQueue - Method to add value in queue list
func (q *Queue) EnQueue(val interface{}) {

	temp, _ := CreateNew(val)

	q.QueueList = append(q.QueueList, temp.QueueList...)
}

//CreateNew - Function to create new queue value
func CreateNew(val interface{}) (*Queue, error) {
	return &Queue{
		QueueList: []interface{}{val},
	}, nil
}

//DeQueue - Method to deQueue value from queue list
func (q *Queue) DeQueue() error {
	if len(q.QueueList) > 0 {
		q.QueueList = q.QueueList[1:]
		return nil
	}
	return errors.New("Queue is empty")
}

//GetFront - Method to return front value from queue list
func (q *Queue) GetFront() interface{} {
	if len(q.QueueList) > 0 {
		return q.QueueList[0]
	}
	return nil
}

//IsEmpty - Method to check whether queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.QueueList) == 0
}
