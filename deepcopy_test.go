package deepcopy_test

import (
	"testing"

	"github.com/iqbalmind/go-deepcopy"
)

// Define a test struct with nested types
type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address *Address
	Hobbies []string
	Props   map[string]interface{}
}

func TestDeepCopy(t *testing.T) {
	// Case 1: Simple struct with nested slice and map
	original := Person{
		Name:    "Dewi",
		Age:     28,
		Address: &Address{Street: "Jl. Braga No. 99", City: "Bandung"},
		Hobbies: []string{"ngopi di Dago", "hiking ke Tangkuban Perahu"},
		Props:   map[string]interface{}{"is_active": true, "score": 95},
	}

	cloned, err := deepcopy.DeepCopy(original)
	if err != nil {
		t.Fatalf("DeepCopy failed: %v", err)
	}

	clonedPerson, ok := cloned.(Person)
	if !ok {
		t.Fatalf("Cloned value is not a Person")
	}

	// Modify the cloned value
	clonedPerson.Name = "Budi"
	clonedPerson.Age = 35
	clonedPerson.Address.Street = "Jl. Asia Afrika No. 45"
	clonedPerson.Hobbies[0] = "jalan-jalan ke Lembang"
	clonedPerson.Props["score"] = 100

	// Verify that the original is unchanged (deep copy successful)
	if original.Name != "Dewi" {
		t.Errorf("Original Name was modified: got %s, want Dewi", original.Name)
	}
	if original.Age != 28 {
		t.Errorf("Original Age was modified: got %d, want 28", original.Age)
	}
	if original.Address.Street != "Jl. Braga No. 99" {
		t.Errorf("Original Address.Street was modified: got %s, want Jl. Braga No. 99", original.Address.Street)
	}
	if original.Hobbies[0] != "ngopi di Dago" {
		t.Errorf("Original Hobbies[0] was modified: got %s, want ngopi di Dago", original.Hobbies[0])
	}
	if original.Props["score"].(int) != 95 {
		t.Errorf("Original Props[\"score\"] was modified: got %v, want 95", original.Props["score"])
	}

	// Case 2: Nil pointer
	var nilPtr *Person
	clonedNil, err := deepcopy.DeepCopy(nilPtr)
	if err != nil {
		t.Fatalf("DeepCopy of nil pointer failed: %v", err)
	}
	if clonedNil != nil {
		t.Errorf("Cloned nil pointer is not nil")
	}

	// Case 3: Slice of pointers
	originalSlice := []*Address{
		{Street: "Jl. Setiabudi", City: "Bandung"},
		{Street: "Jl. Cipaganti", City: "Bandung"},
	}
	clonedSlice, err := deepcopy.DeepCopy(originalSlice)
	if err != nil {
		t.Fatalf("DeepCopy of slice of pointers failed: %v", err)
	}
	clonedAddresses, ok := clonedSlice.([]*Address)
	if !ok {
		t.Fatalf("Cloned value is not a slice of pointers")
	}
	clonedAddresses[0].Street = "Jl. Riau"
	if originalSlice[0].Street == "Jl. Riau" {
		t.Errorf("Original slice of pointers was modified")
	}

	// Case 4: Array
	originalArray := [3]string{"Cimahi", "Lembang", "Ciwidey"}
	clonedArrayIface, err := deepcopy.DeepCopy(originalArray)
	if err != nil {
		t.Fatalf("DeepCopy of array failed: %v", err)
	}
	clonedArray := clonedArrayIface.([3]string)
	clonedArray[0] = "Garut"
	if originalArray[0] == "Garut" {
		t.Errorf("Original array was modified")
	}

	// Case 5: Interface value
	var originalIface interface{} = &Address{Street: "Jl. Merdeka", City: "Bandung"}
	clonedIface, err := deepcopy.DeepCopy(originalIface)
	if err != nil {
		t.Fatalf("DeepCopy of interface failed: %v", err)
	}
	clonedAddr, ok := clonedIface.(*Address)
	if !ok {
		t.Fatalf("Cloned interface is not *Address")
	}
	clonedAddr.Street = "Jl. Gatot Subroto"
	if (originalIface.(*Address)).Street == "Jl. Gatot Subroto" {
		t.Errorf("Original interface value was modified")
	}

	// Case 6: Channel
	originalChan := make(chan int, 2)
	clonedChanIface, err := deepcopy.DeepCopy(originalChan)
	if err != nil {
		t.Fatalf("DeepCopy of channel failed: %v", err)
	}
	clonedChan, ok := clonedChanIface.(chan int)
	if !ok {
		t.Fatalf("Cloned value is not a channel")
	}
	if &clonedChan == &originalChan {
		t.Errorf("Cloned channel points to the same channel")
	}
}
