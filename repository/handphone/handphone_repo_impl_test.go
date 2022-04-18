package handphone

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestHandphoneInsert(t *testing.T) {
	handphonerepository := NewHandphoneRepository(go_database.GetConnection())

	ctx := context.Background()
	handphone := entity.Handphone{
		Merk:  "SAMSUNG",
		Harga: 2100000,
	}
	result, err := handphonerepository.Insert(ctx, handphone)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func TestFindById(t *testing.T) {
	handphonerepository := NewHandphoneRepository(go_database.GetConnection())

	handphone, err := handphonerepository.FindById(context.Background(), 4)
	if err != nil {
		panic(err)
	}

	fmt.Println(handphone)
}

func TestFindAll(t *testing.T) {
	handphonerepository := NewHandphoneRepository(go_database.GetConnection())

	handphone, err := handphonerepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, handphone := range handphone {
		fmt.Println(handphone)
	}
}

func TestUpdate(t *testing.T) {
	handphoneRepository := NewHandphoneRepository(go_database.GetConnection())
	ctx := context.Background()
	handphone := entity.Handphone{
		Merk:  "Halow",
		Harga: 300000,
	}
	result, err := handphoneRepository.Update(ctx, 1, handphone)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	handphonerepository := NewHandphoneRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := handphonerepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
