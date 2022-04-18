package barang

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestBarangInsert(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()
	barang := entity.Barang{
		Name: "Barang luar",
		Stok: 506,
	}
	result, err := barangRepository.Insert(ctx, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByIdBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindById(context.Background(), 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(barang)
}

func TestFindAllBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, barang := range barang {
		fmt.Println(barang)
	}
}
func TestUpdate(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	barang := entity.Barang{
		Name: "Barang Baru",
		Stok: 300,
	}
	result, err := barangRepository.Update(ctx, 8, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := barangRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
