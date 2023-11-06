package linkedlist_test

import (
	"fmt"
	"go-dsa/structures/linkedlist"
	"reflect"
	"testing"
)

func TestSingly_InsertHead(t *testing.T) {
	list := linkedlist.NewSingly[int]()
	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	want := []int{1, 2, 3}
	got := []int{}
	cur := list.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_InsertTail(t *testing.T) {
	list := linkedlist.NewSingly[int]()
	list.InsertTail(3)
	list.InsertTail(2)
	list.InsertTail(1)

	want := []int{3, 2, 1}
	got := []int{}
	cur := list.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Transverse(t *testing.T) {
	list := linkedlist.NewSingly[int]()
	list.InsertHead(1)
	list.InsertTail(2)

	want := []int{1, 2}
	fmt.Println("I want", want)
	fmt.Print("I got ")
	list.Transverse()
}
func TestSingly_IsEmpty(t *testing.T) {
	list := linkedlist.NewSingly[int]()

	if !list.IsEmpty() {
		t.Errorf("list is not empty")
	}

	list.InsertHead(1)

	if list.IsEmpty() {
		t.Errorf("list is empty")
	}
}
