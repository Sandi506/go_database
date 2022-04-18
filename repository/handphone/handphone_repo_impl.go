package handphone

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type handphoneRepositoryImpl struct {
	DB *sql.DB
}

func NewHandphoneRepository(db *sql.DB) HandphoneRepository {
	return &handphoneRepositoryImpl{db}
}
func (repository *handphoneRepositoryImpl) Insert(ctx context.Context, handphone entity.Handphone) (entity.Handphone, error) {
	script := "INSERT INTO handphone( merk, harga) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, handphone.Merk, handphone.Harga)
	if err != nil {
		return handphone, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return handphone, err
	}
	handphone.Id = int32(id)
	return handphone, nil
}
func (repository *handphoneRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Handphone, error) {
	script := "SELECt id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	handphone := entity.Handphone{}
	if err != nil {
		return handphone, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&handphone.Id, &handphone.Merk, &handphone.Harga)
		return handphone, nil
	} else {
		//tidak ada
		return handphone, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
}

func (repository *handphoneRepositoryImpl) FindAll(ctx context.Context) ([]entity.Handphone, error) {
	script := "SELECt id, merk, harga FROM handphone "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var handphones []entity.Handphone
	for rows.Next() {
		handphone := entity.Handphone{}
		rows.Scan(&handphone.Id, &handphone.Merk, &handphone.Harga)
		handphones = append(handphones, handphone)
	}
	return handphones, nil
}

func (repository *handphoneRepositoryImpl) Update(ctx context.Context, id int32, handphone entity.Handphone) (entity.Handphone, error) {
	//TODO implement me
	script := "SELECT id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return handphone, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE handphone SET merk = ?, harga = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, handphone.Merk, handphone.Harga, id)
		if err != nil {
			return handphone, err
		}
		handphone.Id = id
		return handphone, nil
	} else {
		// no
		return handphone, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *handphoneRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM handphone WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
