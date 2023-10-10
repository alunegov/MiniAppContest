package main

// InMemoryRepo is a in-memory implementation of Repo
type InMemoryRepo struct {
	goods []Item
	//orders [][]OrderItem
	orderCnt int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		goods: []Item{
			{1, "Pen", 10, "https://miniappcontest.work.gd/images/pen.svg", "Pen"},
			{2, "Pineapple", 20, "https://miniappcontest.work.gd/images/pineapple.svg", "Pineapple"},
			{3, "Apple", 30, "https://miniappcontest.work.gd/images/apple.svg", "Apple"},
		},
		//orders: make([][]OrderItem, 0),
		orderCnt: 0,
	}
}

func (t *InMemoryRepo) LoadItems() []Item {
	return t.goods
}

func (t *InMemoryRepo) FindItem(id int) *Item {
	for _, it := range t.goods {
		if it.Id == id {
			return &it
		}
	}
	return nil
}

func (t *InMemoryRepo) StoreOrder(items []OrderItem) int {
	// TODO: persist order
	t.orderCnt++
	return t.orderCnt
}
