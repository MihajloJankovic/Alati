package Dao2

import (
	"encoding/json"
	"fmt"
	ps "github.com/MihajloJankovic/Alati/Dao"
	"github.com/hashicorp/consul/api"
	"os"
)

type Dao2 struct {
	cli *api.Client
}

func New() (*Dao2, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Dao2{
		cli: client,
	}, nil
}

func (pss *Dao2) GetGroup(id string) (*ps.ConfigGroup, error) {
	kv := pss.cli.KV()

	data, _, err := kv.List(constructKey(id), nil)
	if err != nil {
		return nil, err
	}

	post := &ps.ConfigGroup{}
	for _, pair := range data {
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}

	}
	return post, nil
}

func (pss *Dao2) GetAllGroups() ([]*ps.ConfigGroup, error) {
	kv := pss.cli.KV()
	data, _, err := kv.List(posts, nil)
	if err != nil {
		return nil, err
	}

	posts := []*ps.ConfigGroup{}
	for _, pair := range data {
		post := &ps.ConfigGroup{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (pss *Dao2) DeleteGroup(id string, version string) (map[string]string, error) {
	kv := pss.cli.KV()
	_, err := kv.DeleteTree(constructKey(id), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (pss *Dao2) CreateGroup(post *ps.ConfigGroup) (*ps.ConfigGroup, error) {
	kv := pss.cli.KV()

	rid, a := generateKey()
	post.Id = a

	data, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: rid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}
func (pss *Dao2) SaveGroup(post *ps.ConfigGroup) (*ps.ConfigGroup, error) {
	kv := pss.cli.KV()

	data, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: constructKey(post.Id), Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pss *Dao2) GetPostsByLabels(id string, version string, labels string) ([]*ps.ConfigGroup, error) {
	kv := pss.cli.KV()

	data, _, err := kv.List(constructKey(id), nil)
	if err != nil {
		return nil, err
	}

	posts := []*ps.ConfigGroup{}

	for _, pair := range data {
		post := &ps.ConfigGroup{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}
