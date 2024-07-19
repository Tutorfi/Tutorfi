package main

import (
	"fmt"
	"os/exec"
)

func get(groupname, pathname string) []byte {
	// get a file from a group's directory
	cmd := exec.Command("cat", "storage/"+groupname+"/"+pathname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return []byte("")
	}
	return stdout
}

func add_file(groupname, filename string) {
	// add files to a group's directory
	// err := os.WriteFile("storage/"+filename, []byte("Hello, World!\n"), 0660)
	cmd := exec.Command("cat", "> storage/"+groupname+"/"+filename)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("created file ", stdout)
}

func add_dir(groupname, dir string) {
	// add directory to a group
	cmd := exec.Command("mkdir", "-p storage/"+groupname+"/"+dir)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("added dir ", stdout)
}

func add_group(groupname string) {
	// create group directory
	cmd := exec.Command("mkdir", "-p storage/"+groupname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("added group ", stdout)
}

func delete_file(groupname, filename string) {
	// delete a file from a group's directory
	cmd := exec.Command("cat", "> storage/"+groupname+"/"+filename)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("deleted file ", stdout)
}

func delete_dir(groupname, dir string) {
	// add directory to a group
	cmd := exec.Command("rmdir", "storage/"+groupname+"/"+dir)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("removed dir ", stdout)
}

func delete_group(groupname string) {
	// delete all directories within the group
	cmd := exec.Command("rm", "-r storage/"+groupname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("removed group ", stdout)
}
