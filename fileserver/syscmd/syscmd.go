package fileserver

import (
	"fmt"
	"os/exec"
)

func Get(groupname, pathname string) []byte {
	// get a file from a group's directory
	cmd := exec.Command("cat", "storage/"+groupname+"/"+pathname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return []byte("")
	}
	return stdout
}

func Add_File(groupname, filename string) {
	// add files to a group's directory
	// err := os.WriteFile("storage/"+filename, []byte("Hello, World!\n"), 0660)
	cmd := exec.Command("touch", "storage/"+groupname+"/"+filename)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("created file ", filename, stdout)
}

func Add_Dir(groupname, dir string) {
	// add directory to a group
	cmd := exec.Command("mkdir", "-p", "storage/"+groupname+"/"+dir)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("added dir ", groupname, stdout)
}

func Add_Group(groupname string) {
	// create group directory
	cmd := exec.Command("mkdir", "-p", "storage/"+groupname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error(), "error")
		return
	}
	fmt.Println("added group ", groupname, stdout)
}

func Delete_File(groupname, filename string) {
	// delete a file from a group's directory
	cmd := exec.Command("rm", "storage/"+groupname+"/"+filename)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("deleted file ", filename, stdout)
}

func Delete_Dir(groupname, dir string) {
	// add directory to a group
	cmd := exec.Command("rm", "-r", "storage/"+groupname+"/"+dir)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("removed dir ", dir, stdout)
}

func Delete_Group(groupname string) {
	// delete all directories within the group
	cmd := exec.Command("rm", "-r", "storage/"+groupname)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("removed group ", groupname, stdout)
}
