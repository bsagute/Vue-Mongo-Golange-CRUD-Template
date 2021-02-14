package queue

import (
	"reflect"
	"testing"
)

func TestQueue_EnQueue(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		q    *Queue
		args args
	}{
		{
			name: "Queue_EnQueue",
			q:    &Queue{},
			args: args{
				val: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.EnQueue(tt.args.val)
		})
	}
}

func TestCreateNew(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Queue
		wantErr bool
	}{
		{
			name: "CreateNew",
			args: args{
				val: 150,
			},
			want: &Queue{
				QueueList: []interface{}{150},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNew(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name    string
		q       *Queue
		wantErr bool
	}{
		{
			name: "Queue_DeQueue",
			q: &Queue{
				QueueList: []interface{}{100, 150},
				Front:     100,
				Rear:      150,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.DeQueue(); (err != nil) != tt.wantErr {
				t.Errorf("Queue.DeQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueue_GetFront(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want interface{}
	}{
		{
			name: "queueFront",
			q: &Queue{
				QueueList: []interface{}{100, 150},
				Front:     100,
				Rear:      150,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetFront(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want bool
	}{
		{
			name: "queueEmpty",
			q: &Queue{
				QueueList: []interface{}{100, 150},
				Front:     100,
				Rear:      150,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
