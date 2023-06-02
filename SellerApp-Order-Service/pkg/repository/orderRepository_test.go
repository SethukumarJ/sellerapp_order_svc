package repository_test

// func TestFetchOrder(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	repo := repository.NewOrderRepository(gormDB)

// 	// Mock the database response
// 	rows := sqlmock.NewRows([]string{"id", "status", "total"}).
// 		AddRow("order1", "pending", 10.5).
// 		AddRow("order2", "completed", 20.0)

// 	mock.ExpectQuery("SELECT (.+) FROM orders WHERE (.+) LIMIT (.+) OFFSET (.+)").
// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
// 		WillReturnRows(rows)
// 	mock.ExpectQuery("SELECT count(1) FROM orders").WillReturnRows(sqlmock.NewRows([]string{"count(1)"}).AddRow(2))

// 	filter := domain.Filter{Status: "pending", MinTotal: 0, MaxTotal: 0, SortBy: "", SortOrder: ""}
// 	pagination := utils.Filter{Page: 1, PageSize: 10}

// 	orders, metadata, err := repo.FetchOrder(context.Background(), 123, filter, pagination)
// 	assert.NoError(t, err)
// 	assert.Len(t, orders, 2)
// 	assert.Equal(t, 2, metadata.TotalRecords)
// }

// func TestUpdateOrder(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	repo := repository.NewOrderRepository(gormDB)

// 	// Mock the database response
// 	mock.ExpectExec("UPDATE orders SET").WithArgs("order1", "completed").WillReturnResult(sqlmock.NewResult(0, 1))

// 	orderID, err := repo.UpdateOrder(context.Background(), "order1", "completed")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "order1", orderID)
// }

// func TestFindItem(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	repo := repository.NewOrderRepository(gormDB)

// 	// Mock the database response
// 	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
// 		AddRow("item1", "Item 1", 10.5)

// 	mock.ExpectQuery("SELECT * FROM items").WithArgs("item1").WillReturnRows(rows)

// 	item, err := repo.FindItem(context.Background(), "item1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "item1", item.ID)
// 	assert.Equal(t, "Item 1", item.Description)
// 	assert.Equal(t, 10.5, item.Price)
// }

// func TestCreateItem(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	repo := repository.NewOrderRepository(gormDB)

// 	// Mock the database response
// 	mock.ExpectExec("INSERT INTO items").WillReturnResult(sqlmock.NewResult(0, 1))

// 	item := domain.Item{ID: "item1", Description: "Item 1", Price: 10.5}
// 	itemID, err := repo.CreateItem(context.Background(), item)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "item1", itemID)
// }

// func TestCreateOrder(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	repo := repository.NewOrderRepository(gormDB)

// 	// Mock the database response
// 	mock.ExpectExec("INSERT INTO orders").WillReturnResult(sqlmock.NewResult(0, 1))

// 	order := domain.Order{ID: "order1", Status: "pending", Total: 10.5}
// 	orderID, err := repo.CreateOrder(context.Background(), order)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "order1", orderID)
// }
