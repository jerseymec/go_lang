package main

import "fmt"

const ArraySize = 7

// array with linkedlist to hold collisions
//HashTable structure
//BucketStructure
type bucketNode struct {
	key  string
	Next *bucketNode
}
type bucket struct {
	length int
	head   *bucketNode
}

type HashTable struct {
	array [ArraySize]*bucket
}

//Initialize the bucket in each slot of the Hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

//Insert a key into the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

//Find a key inside the hash table returns true or false
func (h HashTable) Search(Key string) bool {
	index := hash(Key)
	return h.array[index].search(Key)
}

//Delete a key from the hash table
func (h HashTable) Delete(Key string) {
	index := hash(Key)
	h.array[index].delete(Key)
}

//hash function to get index of Hash table
func hash(Key string) int {
	sum := 0
	for _, v := range Key {
		sum += int(v)
	}
	return sum % ArraySize
}

// insert a bucketNode into a bucketlist
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.Next = b.head
		b.head = newNode
		b.length++
	} else {
		err := fmt.Errorf("Key already exists, cannot insert multiples")
		fmt.Println(err.Error())
	}
}

//search the whole bucket list till you find a match
func (b bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		} else {
			currentNode = currentNode.Next
		}
	}
	return false
}

//delete a key from the bucketlist
func (b *bucket) delete(k string) {
	if b.head == nil {
		return

	}
	if b.head.key == k {
		b.head = b.head.Next
		b.length--
		return
	}
	currentNode := b.head

	for currentNode.Next != nil && currentNode.Next.key != k {
		currentNode = currentNode.Next
	}
	if currentNode.Next == nil {
		return
	}
	if currentNode.Next.key == k {
		currentNode.Next = currentNode.Next.Next
		b.length--
	}

}

func main() {

	hashicorp := Init()
	fmt.Println(hashicorp)

	//buck := &bucket{}
	list := []string{
		"Randy",
		"Eric",
		"Smith",
		"Dandy",
		"Tandy",
	}

	for _, v := range list {
		hashicorp.Insert(v)
	}
	//hashicorp.Insert("Randy")
	//hashicorp.Insert("Eric")
	//hashicorp.Insert("Smith")
	//hashicorp.Insert("Dandy")
	//hashicorp.Insert("Randy")
	//hashicorp.Insert("Tandy")

	//hashicorp.Search("Smith")
	//fmt.Println(buck.head.key)
	hashicorp.Delete("Tandy")
	//fmt.Println(hashicorp.array)
	fmt.Println(hashicorp.Search("Tandy"))

	hashicorp.Insert("Randy")
	//fmt.Println(hashicorp)

}
