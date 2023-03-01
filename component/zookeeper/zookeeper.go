package zookeeper

import (
	"encoding/json"
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

// Zk is
type Zk interface {
	Create(path string, data interface{}, flag int32) error
	Get(path string) (string, error)
	Set(path string, data interface{}) error
	Del(path string) error
	Notify(path string, data interface{}) error
}

type Zookeeper struct {
	servers        []string
	sessionTimeout time.Duration
}

func New(opts ...Option) (*Zookeeper, error) {
	op := setDefault()
	for _, o := range opts {
		if o != nil {
			o(&op)
		}
	}
	if len(op.servers) <= 0 {
		return nil, fmt.Errorf("must config zk servers")
	}
	zk := &Zookeeper{
		servers:        op.servers,
		sessionTimeout: op.sessionTimeout,
	}
	return zk, nil
}

func (z *Zookeeper) getClient() (*zk.Conn, error) {
	conn, _, err := zk.Connect(z.servers, z.sessionTimeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Create
// @param path string 节点路径
// @param data interface 节点数据
// @param flags int32 节点类型 0:永久，除非手动删除; zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除;zk.FlagSequence  = 2:会自动在节点后面添加序号;3:Ephemeral和Sequence，即，短暂且自动添加序号
func (z *Zookeeper) Create(path string, data interface{}, flags int32) error {
	_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	conn, err := z.getClient()
	defer conn.Close()
	if err != nil {
		return err
	}
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	_, err = conn.Create(path, _data, flags, acls)
	if err != nil {
		return fmt.Errorf("create node:%s failed, err:%v", path, err)
	}
	return nil
}

func (z *Zookeeper) Get(path string) (string, error) {
	conn, err := z.getClient()
	defer conn.Close()
	if err != nil {
		return "", err
	}
	data, _, err := conn.Get(path)
	if err != nil {
		return "", fmt.Errorf("get node:%s failed, err:%v", path, err)
	}
	return string(data), nil
}

func (z *Zookeeper) Set(path string, data interface{}) error {
	conn, err := z.getClient()
	defer conn.Close()
	if err != nil {
		return err
	}
	_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, sate, _ := conn.Get(path)
	_, err = conn.Set(path, _data, sate.Version)
	if err != nil {
		return fmt.Errorf("modify node:%s failed, err:%v", path, err)
	}
	return nil
}

func (z *Zookeeper) Del(path string) error {
	conn, err := z.getClient()
	defer conn.Close()
	if err != nil {
		return err
	}
	_, sate, _ := conn.Get(path)
	err = conn.Delete(path, sate.Version)
	if err != nil {
		return fmt.Errorf("del node:%s failed, err:%v", path, err)
	}
	return nil
}

//func (z *Zookeeper) Watch(path string, callback func(event zk.Event)) error {
//	options := zk.WithEventCallback(callback)
//	conn, _, err := zk.Connect(z.servers, z.sessionTimeout, options)
//	defer conn.Close()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (z *Zookeeper) Notify(path string, data interface{}) error {
	conn, err := z.getClient()
	defer conn.Close()
	if err != nil {
		return err
	}
	_data := []byte{}
	if data != nil {
		_data, err = json.Marshal(data)

		if err != nil {
			return err
		}
	}

	exist, _, err := conn.Exists(path)
	if err != nil {
		return fmt.Errorf("notify node:%s failed, err:%v", path, err)
	}
	if !exist {
		acls := zk.WorldACL(zk.PermAll)
		_, err = conn.Create(path, _data, 0, acls)
		if err != nil {
			return fmt.Errorf("notify node:%s failed, err:%v", path, err)
		}
		return nil
	}
	_, sate, _ := conn.Get(path)
	_, err = conn.Set(path, _data, sate.Version)
	if err != nil {
		return fmt.Errorf("notify node:%s failed, err:%v", path, err)
	}
	return nil
}
