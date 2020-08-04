package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func getConee() *zk.Conn {
	conn, _, err := zk.Connect([]string{"172.17.0.2:2181"}, time.Second*3)
	if err != nil {
		fmt.Println("创建zookeeper连接失败", err)
		return nil
	}
	return conn
}

func main() {
	con := getConee()
	defer con.Close()

	//var path = "/zk_test_go"
	//var data = []byte("hello")
	//var flags int32 = 0
	// permission
	//var acls = zk.WorldACL(zk.PermAll)

	// create
	/*p, err_create := con.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Println("created:", p)*/

	// get
	/*v, s, err := con.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("value of path[%s]=[%s].\n", path, v)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))

	// exist
	exist, s, err := con.Exists(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("path[%s] exist[%t]\n", path, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))

	// update
	var new_data = []byte("zk_test_new_value")
	s, err = con.Set(path, new_data, s.Version)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("update state:\n")
	fmt.Printf("%s\n", ZkStateStringFormat(s))

	*/
	ZkChildWatchTest()
}

func ZkChildWatchTest() {
	fmt.Printf("ZkChildWatchTest")

	var hosts = []string{"172.17.0.2:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// try create root path
	var root_path = "/test_root"

	// check root path exist
	exist, _, err := conn.Exists(root_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exist {
		fmt.Printf("try create root path: %s\n", root_path)
		var acls = zk.WorldACL(zk.PermAll)
		p, err := conn.Create(root_path, []byte("root_value"), 0, acls)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("root_path: %s create\n", p)
	}

	// try create child node
	cur_time := time.Now().Unix()
	ch_path := fmt.Sprintf("%s/ch_%d", root_path, cur_time)
	var acls = zk.WorldACL(zk.PermAll)
	p, err := conn.Create(ch_path, []byte("ch_value"), zk.FlagEphemeral, acls)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ch_path: %s create\n", p)

	// watch the child events
	children, s, child_ch, err := conn.ChildrenW(root_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("root_path[%s] child_count[%d]\n", root_path, len(children))
	for idx, ch := range children {
		fmt.Printf("%d, %s\n", idx, ch)
	}

	fmt.Printf("watch children result state[%s]\n", ZkStateString(s))

	for {
		select {
		case ch_event := <-child_ch:
			{
				fmt.Println("path:", ch_event.Path)
				fmt.Println("type:", ch_event.Type.String())
				fmt.Println("state:", ch_event.State.String())

				if ch_event.Type == zk.EventNodeCreated {
					fmt.Printf("has node[%s] detete\n", ch_event.Path)
				} else if ch_event.Type == zk.EventNodeDeleted {
					fmt.Printf("has new node[%d] create\n", ch_event.Path)
				} else if ch_event.Type == zk.EventNodeDataChanged {
					fmt.Printf("has node[%d] data changed", ch_event.Path)
				}
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func ZkStateString(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d, Mzxid: %d, Ctime: %d, Mtime: %d, Version: %d, Cversion: %d, Aversion: %d, EphemeralOwner: %d, DataLength: %d, NumChildren: %d, Pzxid: %d",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}

func ZkStateStringFormat(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}
