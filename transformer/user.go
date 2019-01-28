package transformer

import (
	"fmt"
	"kurapika/extractor"
	"kurapika/utils"
)

func SQLUser(u extractor.User) error {
	t, err := utils.FetchID("(SELECT IFNULL((SELECT id from dim_user WHERE `source_id`=? LIMIT 1), 0))", u.ID.Hex())
	if err != nil {
		panic(err)
	}
	fmt.Print("Processing UserID #%s\n", u.ID.Hex())

	if t != 0 {
		return nil
	}
	fmt.Printf("New data found! Insertin UserID #%s...\n", u.ID.Hex())
	extracted, err := utils.Extract(u.CreatedAt)
	if err != nil {
		panic(err)
	}

	dateID, err := utils.DateDim(extracted)
	if err != nil {
		panic(err)
	}
	if t == 0 {
		t, err = baseUser(u.ID.Hex(), u.Source)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func baseUser(ID, source string) (int, error) {
	p := []interface{}{ID, source}
	q := `INSERT INTO dim_user (
		source_id,
		source
	) VALUES (
		?,
		?,
	);`
	utils.ExecQuery(q, p)

	t, err := utils.FetchID("SELECT id FROM dim_user WHERE `source_id`=?", ID)
	if err != nil {
		panic(err)
	}
	return t, nil
}
