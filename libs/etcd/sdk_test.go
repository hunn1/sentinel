package etcd

import (
	"fmt"
	"log"

	"testing"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

func TestNewEtcdClient(t *testing.T) {
	cli, err := NewEtcdClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	key1, value1 := "testkey1", "value"

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	_, err = cli.Put(ctx, key1, value1)
	cancel()
	if err != nil {
		log.Println("Put failed. ", err)
	} else {
		log.Printf("Put {%s:%s} succeed\n", key1, value1)
	}

	//
	ctx, cancel = context.WithTimeout(context.Background(), RequestTimeout)
	resp, err := cli.Get(ctx, key1)
	cancel()
	if err != nil {
		log.Println("Get failed. ", err)
		return
	}

	for _, kv := range resp.Kvs {
		log.Printf("Get {%s:%s} \n", kv.Key, kv.Value)
	}

	done := make(chan bool)

	go func() {
		wch := cli.Watch(context.Background(), "", clientv3.WithPrefix())

		for item := range wch {
			for _, ev := range item.Events {
				log.Printf("Type:%s, key:%s, value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	go func() {
		for cnt := 0; cnt < 11; cnt++ {
			value := fmt.Sprintf("%s%d", "value", cnt)
			_, err = cli.Put(context.Background(), key1, value)
			if err != nil {
				log.Println("Put failed. ", err)
			} else {
				log.Printf("Put {%s:%s} succeed\n", key1, value)
			}
		}
	}()

	<-done

	log.Println("Done!")
}
