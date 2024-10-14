package main

import (
	
	"lem-in/graphs"

	//readata "lem-in/readdata"
	"lem-in/rooms"
	//"log"
	//"os"
	"lem-in/paths"
	"strings"
)

var (
	ArrayRooms = make([]*rooms.Room, 0)
)

func getRoom(name string) *rooms.Room {
	for _, v := range ArrayRooms {
		if v != nil && v.GetName() == name {
			return v
		}
	}
	return nil
}

// sourceArray Contain All Rooms Recieved from FileName.txt
// And CreateRooms will define every Room In Local Variable ->  ArrayRooms that Already defined
func CreateRooms(graph *graphs.Graph, sourceArray []string) {
	for _, name := range sourceArray {
		myRoom := rooms.NewRoom(name, 1, 2)
		ArrayRooms = append(ArrayRooms, myRoom)
	}
	graph.NewGraph(sourceArray)
}

// LinkRoomsTogether Take arg = graph where we will add Rooms
// and arg2 = RoomsLinkedSource will recieve it from File like that ["roomX-roomY","roomA-roomB","r...-r..."]...
// then we will link them Together
func LinkRoomsTogether(graph *graphs.Graph, RoomsLinkedSource []string) {
	for _, link := range RoomsLinkedSource {
		roomsLinked := strings.Split(link, "-")
		graph.AddRoom(getRoom(roomsLinked[0]), getRoom(roomsLinked[1]))
	}
}

func main() {
	// ZAKARIA start
	/*
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
		}*/
	//fmt.Prinln(list)
	/*var arr = []string{"start", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8", "r9", "r10", "r11", "end"}
	var linksRooms = []string{"start-r1", "start-r2", "start-r3", "r1-r4", "r2-r5",
		"r3-r6", "r4-r7", "r5-r8", "r6-r5", "r6-r9", "r6-r10", "r10-r11",
		"r7-end", "r8-end", "r9-end"}
	myGraph := graphs.Graph{}
	CreateRooms(&myGraph, arr)
	LinkRoomsTogether(&myGraph, linksRooms)
	myGraph.DFSExplore(ArrayRooms[0],"end")
	paths.PrintData(myGraph.Paths)*/
	var arr = []string{"s", "0", "1", "2", "3", "4", "5", "6", "7", "8",  "t"}
	var linksRooms = []string{"s-0", "s-1", "s-2", "0-3","0-4","1-4", "1-5",
		"2-3", "2-7", "3-6", "3-7", "4-5", "4-6", "4-8","5-8","6-t", "7-t", "8-t"}
	myGraph := graphs.Graph{}
	CreateRooms(&myGraph, arr)
	LinkRoomsTogether(&myGraph, linksRooms)
	myGraph.DFSExplore(ArrayRooms[0],"t")
	//paths.PrintData(myGraph.Paths)
	groupedPaths := paths.GroupUniquePaths(myGraph.Paths)[0]
	paths.GetAllPaths(groupedPaths)
	paths.MakeAntsInPlaces(10)
	paths.ShowPathList()
}
