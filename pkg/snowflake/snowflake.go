package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(starTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", starTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//// ToDo github example sonwflake
//func main() {
//
//	// Create a new Node with a Node number of 1
//	node, err := snowflake.NewNode(1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Generate a snowflake ID.
//	id := node.Generate()
//
//	// Print out the ID in a few different ways.
//	fmt.Printf("Int64  ID: %d\n", id)
//	fmt.Printf("String ID: %s\n", id)
//	fmt.Printf("Base2  ID: %s\n", id.Base2())
//	fmt.Printf("Base64 ID: %s\n", id.Base64())
//
//	// Print out the ID's timestamp
//	fmt.Printf("ID Time  : %d\n", id.Time())
//
//	// Print out the ID's node number
//	fmt.Printf("ID Node  : %d\n", id.Node())
//
//	// Print out the ID's sequence number
//	fmt.Printf("ID Step  : %d\n", id.Step())
//
//	// Generate and print, all in one.
//	fmt.Printf("ID       : %d\n", node.Generate().Int64())
//}
