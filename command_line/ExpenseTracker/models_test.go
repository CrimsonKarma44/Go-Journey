package ExpenseTracker

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := ListExpense{}
	amount := 2000
	description := "New Task"
	l.Add(description, amount)
	if l[len(l)-1].Id != 1 {
		t.Errorf("Expected %q, got %q instead.", description, l[0].Description)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := ListExpense{}
	l2 := ListExpense{}
	description := "New Task"
	amount := 2000
	l1.Add(description, amount)
	if l1[0].Description != description {
		t.Errorf("Expected %q, got %q instead.", description, l1[0].Description)
	}
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	if l1[0].Description != l2[0].Description {
		t.Errorf("Task %q should match %q task.", l1[0].Description, l2[0].Description)
	}
	if l1[0].Amount != l2[0].Amount {
		t.Errorf("Task %q should match %q task.", l1[0].Amount, l2[0].Amount)
	}
}

func TestSum(t *testing.T) {
	ls := ListExpense{}
	description := "New Task"
	description2 := "Another Task"
	amount := 2000
	amount2 := 3000
	ls.Add(description, amount)
	ls.Add(description2, amount2)
	if value := ls.Summation(); value != amount+amount2 {
		t.Errorf("Expected %v, got %v instead.", amount2+amount2, value)
	}
}
