package main

import (
	//"fmt"
	"fmt"
	"lemin/graphs"
	readata "lemin/read_data"
	"lemin/rooms"
	"log"
	"os"
	"strings"
)

func main() {
	// ZAKARIA start

	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("[USAGE]: go run . example.txt")
		return
	}

	file, err := os.ReadFile(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	slice := strings.Split(string(file), "\n")

	var hold []string
	for _, v := range slice {
		if v == "" {
			continue
		}
		hold = append(hold, v)
	}

	readata.ExtractStartAndEnd(hold)
	EdjeList := readata.ExtractEdgeList(hold)
	
	AdjList := readata.EdgeListToAdjList(EdjeList)
	// list of all nodes
	var list []string
	for i := range AdjList{
		list = append(list, i)
	}
	//fmt.Prinln(list)

	myGraph := graphs.Graph{}
<<<<<<< HEAD
	 myGraph.NewGraph([]string{"start","r1","r2","r3","r4","r5","r6","r7","r8","r9","r10","r11","end"})
	//room_start := rooms.NewRoom("s",1,2)
	room_start := rooms.NewRoom("start",1,2)
	room_r1 := rooms.NewRoom("r1",1,2)
	room_r2 := rooms.NewRoom("r2",1,2)
	room_r3:= rooms.NewRoom("r3",1,2)
	room_r4 := rooms.NewRoom("r4",1,2)
	room_r5 := rooms.NewRoom("r5",1,2)
	room_r6 := rooms.NewRoom("r6",1,2)
	room_r7 := rooms.NewRoom("r7",1,2)
	room_r8 := rooms.NewRoom("r8",1,2)
	room_r9 := rooms.NewRoom("r9",1,2)
	room_r10 := rooms.NewRoom("r10",1,2)
	room_r11 := rooms.NewRoom("r11",1,2)
	room_end := rooms.NewRoom("end",1,2)

	 myGraph.AddRoom(room_start,room_r1)
	 myGraph.AddRoom(room_start,room_r2)
	 myGraph.AddRoom(room_start,room_r3)
	 myGraph.AddRoom(room_r1,room_r4)
	 myGraph.AddRoom(room_r2,room_r5) 
	 myGraph.AddRoom(room_r3,room_r6)
	 myGraph.AddRoom(room_r4,room_r7)
	 myGraph.AddRoom(room_r5,room_r8)
	 myGraph.AddRoom(room_r6,room_r9)
	 myGraph.AddRoom(room_r6,room_r5)
	 myGraph.AddRoom(room_r6,room_r10)
	 myGraph.AddRoom(room_r10,room_r11)
	 myGraph.AddRoom(room_r7,room_end)
	 myGraph.AddRoom(room_r8,room_end)
	 myGraph.AddRoom(room_r9,room_end)
	 myGraph.DFSExplore(room_start)
=======
	myGraph.NewGraph(list)
	myGraph.NewGraph([]string{"s", "0", "1", "2", "3", "4", "5", "6", "7", "8", "t"})
	//room_start := rooms.NewRoom("s",1,2)
	
 	room_s := rooms.NewRoom("s", 1, 2)
	room_0 := rooms.NewRoom("0", 1, 2)
	room_1 := rooms.NewRoom("1", 1, 2)
	room_2 := rooms.NewRoom("2", 1, 2)
	room_3 := rooms.NewRoom("3", 1, 2)
	room_4 := rooms.NewRoom("4", 1, 2)
	room_5 := rooms.NewRoom("5", 1, 2)
	room_6 := rooms.NewRoom("6", 1, 2)
	room_7 := rooms.NewRoom("7", 1, 2)
	room_8 := rooms.NewRoom("8", 1, 2)
	room_t := rooms.NewRoom("t", 1, 2) 

	myGraph.AddRoom(room_s, room_0)
	myGraph.AddRoom(room_s, room_1)
	myGraph.AddRoom(room_s, room_2)
	myGraph.AddRoom(room_0, room_3)
	myGraph.AddRoom(room_0, room_4)
	myGraph.AddRoom(room_1, room_4)
	myGraph.AddRoom(room_1, room_5)
	myGraph.AddRoom(room_2, room_3)
	myGraph.AddRoom(room_2, room_7)
	myGraph.AddRoom(room_3, room_6)
	myGraph.AddRoom(room_3, room_7)
	myGraph.AddRoom(room_4, room_6)
	myGraph.AddRoom(room_4, room_8)
	myGraph.AddRoom(room_4, room_5)
	myGraph.AddRoom(room_5, room_8)
	myGraph.AddRoom(room_6, room_t)
	myGraph.AddRoom(room_7, room_t)
	myGraph.AddRoom(room_8, room_t)
	myGraph.DFSExplore(room_s) 
>>>>>>> 2fb477f28408a056b95d2ab3c81b8b9d97b14574
	//myGraph.ShowAdjList()
	/*myGraph := NewGraph(11)
	myGraph.AddNode(0,4)
	myGraph.AddNode(0,6)
	myGraph.AddNode(1,3)
	myGraph.AddNode(4,3)
	myGraph.AddNode(5,2)
	myGraph.AddNode(3,5)
	myGraph.AddNode(4,2)
	myGraph.AddNode(2,1)
	myGraph.AddNode(7,6)
	myGraph.AddNode(7,2)
	myGraph.AddNode(7,4)
	myGraph.AddNode(6,5)
	myGraph.DFSsearch(0,5)*/

	// myGraph.ShowAdjList()
}
