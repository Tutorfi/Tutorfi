package main

// go test -v

import (
	syscmd "fileserver/syscmd"
	"testing"
)

func TestSyscmd_1(t *testing.T) {
	syscmd.Add_Group("test1")
	syscmd.Add_File("test1", "example.txt")
	// fmt.Println(syscmd.Get("test1", "example.txt"))
	syscmd.Delete_File("test1", "example.txt")
	syscmd.Delete_Group("test1")
}

func TestSyscmd_2(t *testing.T) {
	syscmd.Add_Group("test2")
	syscmd.Add_File("test2", "example.txt")
	// fmt.Println(syscmd.Get("test2", "example.txt"))
	syscmd.Delete_Group("test2")
}

func TestSyscmd_3(t *testing.T) {
	syscmd.Add_Group("test3")
	syscmd.Add_Dir("test3", "folder")
	syscmd.Add_File("test3", "folder/example.txt")
	// fmt.Println(syscmd.Get("test3", "folder/example.txt"))
	syscmd.Delete_Group("test3")
}

func TestSyscmd_Overall(t *testing.T) {
	syscmd.Add_Group("testing")
	syscmd.Add_Dir("testing", "folder")
	syscmd.Add_File("testing", "folder/file.txt")

	syscmd.Delete_File("testing", "folder/file.txt")
	syscmd.Delete_Dir("testing", "folder")
	syscmd.Delete_Group("testing")
}
