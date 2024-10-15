package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"todo"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	// Parsing command line flags
	add := flag.Bool("add", false, "Add task to the ToDo list")
	del := flag.Bool("del", false, "Delete Task from the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	pending := flag.Bool("listp", false, "Lists Item not completed")

	// Custom help function
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Define an items list
	l := &todo.List{}
	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	case *list:
		// List current to do items
		fmt.Print(l)
	case *pending:
		// List only uncompleted items
		l.Pending()
		fmt.Print(l)
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del:
		fmt.Println("Deleting Task...")
		t, err := func() (int, error) {
			t, err := getTask(os.Stdin, flag.Args()...)
			v, _ := strconv.Atoi(t)
			return v, err
		}()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err = l.Delete(t); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Item Deleted")

	case *add:
		// When any arguments (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// Concatenate all provided arguments with a space and
	// add to the list as an item
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new
// task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}
