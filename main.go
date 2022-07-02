package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	inp := bufio.NewReader(os.Stdin)
	var fsel int
	var wsel int
	var path string
	clear()
	for {
		fmt.Print("This Program is Simple CLI Notepad\n----------\n1. Write 2. Read 3. Edit 4. Delete 5. Files List 6. Clear 7. Exit\nEnter>")
		fmt.Scanln(&fsel)
		if fsel == 1 {
			for {
				fmt.Print("\nDo you want to save on Downloads Folder?\n1. Yes 2. No 3. Back\nEnter>")
				fmt.Scanln(&wsel)
				if wsel == 1 {
					fmt.Print("\nEnter file name\nEnter>")
					fmt.Scanln(&path)
					if len(path) < 5 {
						path = path + ".txt"

					} else if len(path) >= 5 {
						pathex := path[len(path)-4:]
						if pathex != ".txt" {
							path = path + ".txt"
						}
					}
					f1, err := os.OpenFile("/Users/User/Downloads/"+path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
					checkError(err)
					defer f1.Close()
					fmt.Print("\n----------\n?help : Command list\n----------\n\n")
					for {
						fmt.Print("Enter>")
						str, _ := inp.ReadString('\n')
						str = strings.TrimSpace(str)
						if str == "stopnote" {
							break
						} else if str == "?help" {
							fmt.Print("----------\nstopnote : Write Stop\n----------\n")
						} else {
							fmt.Fprintf(f1, str)
							fmt.Fprint(f1, "\n")
						}
					}
					fmt.Print("\nSave\n\n")
					break
				} else if wsel == 2 {
					fmt.Print("\nEnter Full Path (ex. /Users/User/Test.txt)\nEnter>")
					fmt.Scanln(&path)
					pathex := path[len(path)-4:]
					if pathex != ".txt" {
						path = path + ".txt"
					}
					f1, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
					checkError(err)
					defer f1.Close()
					fmt.Print("\n")
					for {
						fmt.Print("Enter>")
						str, _ := inp.ReadString('\n')
						str = strings.TrimSpace(str)
						if str == "stopnote" {
							break
						} else if str == "?help" {
							fmt.Print("----------\nstopnote : Write Stop\n----------\n")
						} else {
							fmt.Fprintf(f1, str)
							fmt.Fprint(f1, "\n")
						}
					}
					fmt.Print("\nSave")
					break
				} else if wsel == 3 {
					fmt.Print("\n")
					break
				}
			}
		} else if fsel == 2 {
			fmt.Print("\nEnter Full Path (ex. /Users/User/Test.txt) or Back\nEnter>")
			fmt.Scanln(&path)
			if path == "Back" || path == "back" || path == "BACK" {
				fmt.Print("\n")
				continue
			}

			input, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			lines := strings.Split(string(input), "\n")

			fmt.Print("\n")

			for i := range lines {
				fmt.Print(i, " : ", lines[i]+"\n")
			}

			fmt.Print("\n")
		} else if fsel == 3 {
			var seli int
			var rest string
			var cnt = -1

			fmt.Print("\nEnter Full Path (ex. /Users/User/Test.txt) or Back\nEnter>")
			fmt.Scanln(&path)
			if path == "Back" || path == "back" || path == "BACK" {
				fmt.Print("\n")
				continue
			}

			fmt.Print("\n")
			input, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			lines := strings.Split(string(input), "\n")

			for i := range lines {
				fmt.Print(i, " : ", lines[i]+"\n")
				cnt++
			}

			for {
				fmt.Print("\nSelect Line\nEnter>")
				fmt.Scanln(&seli)
				if seli <= cnt && seli >= 0 {
					break
				} else {
					fmt.Print("\nTry again1\n")
				}
			}
			fmt.Print("\nEnter New String (Line Delete : Enter)\nEnter>")
			fmt.Scanln(&rest)
			lines[seli] = rest

			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(path, []byte(output), 0644)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("\nFinish Edit\n\n")
		} else if fsel == 4 {
			fmt.Print("\nEnter Full Path (ex. /Users/User/Test.txt) or Back\nEnter>")
			fmt.Scanln(&path)
			if path == "Back" || path == "back" || path == "BACK" {
				fmt.Print("\n")
				continue
			}

			err3 := os.Remove(path)
			if err3 != nil {
				panic(err3)
			}

			fmt.Print("\nFile Delete\n\n")
		} else if fsel == 5 {
			var lsel int
			for {
				fmt.Print("\nDo you want to list of Downloads Files?\n1. Yes 2. No 3. Back\nEnter>")
				fmt.Scanln(&lsel)
				if lsel == 1 {
					fmt.Print("\n")
					cmd := exec.Command("ls", "-1")
					cmd.Dir = "/Users/User/Downloads"
					cmd.Stdout = os.Stdout
					cmd.Run()
					fmt.Print("\n")
					break
				} else if lsel == 2 {
					fmt.Print("\nEnter Full Path (ex. /Users/User/Test.txt) or Back\nEnter>")
					fmt.Scanln(&path)
					fmt.Print("\n")
					cmd := exec.Command("ls", "-1")
					cmd.Dir = path
					cmd.Stdout = os.Stdout
					cmd.Run()
					fmt.Print("\n")
					break
				} else if lsel == 3 {
					fmt.Print("\n")
					break
				} else {
					fmt.Print("\nTry again\n")
				}
			}
		} else if fsel == 6 {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		} else if fsel == 7 {
			fmt.Print("\nGoodbye.\n")
			break
		} else {
			fmt.Print("\nTry again\n\n")
		}
	}
}
