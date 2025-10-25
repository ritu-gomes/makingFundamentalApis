package repo

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func generateInitialProducts(r *productRepo) {
	prod1 := &Product{
		Id:          1,
		Title:       "Wireless Mouse",
		Description: "Ergonomic wireless mouse with USB receiver",
		Price:       25.99,
		ImageUrl:    "https://m.media-amazon.com/images/I/61LtuGzXeaL._AC_SL1500_.jpg",
	}

	prod2 := &Product{
		Id:          2,
		Title:       "Mechanical Keyboard",
		Description: "RGB backlit keyboard with clicky switches",
		Price:       59.99,
		ImageUrl:    "https://cdn.thewirecutter.com/wp-content/media/2024/04/mechanicalkeyboards-2048px-1353-2x1-1.jpg?width=2048&quality=75&crop=2:1&auto=webp",
	}

	prod3 := &Product{
		Id:          3,
		Title:       "USB-C Hub",
		Description: "7-in-1 USB-C hub with HDMI, USB, and Ethernet",
		Price:       39.50,
		ImageUrl:    "https://www.ryans.com/storage/products/main/ugreen-type-c-male-to-quad-usb-female-meter-hub-11646050628.webp",
	}

	prod4 := &Product{
		Id:          4,
		Title:       "Full HD Webcam",
		Description: "1080p webcam with built-in mic and autofocus",
		Price:       49.00,
		ImageUrl:    "https://diamu.com.bd/wp-content/uploads/2021/07/Logitech-C920-HD-Pro-Webcam-768x768.jpg.webp",
	}

	r.productList = append(r.productList, prod1, prod2, prod3, prod4)

}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.Id = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.Id == productId {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productId int) error {
	var temptList []*Product
	for _, p := range r.productList {
		if p.Id != productId {
			temptList = append(temptList, p)
		}
	}

	r.productList = temptList
	return nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for idx, prod := range r.productList {
		if prod.Id == p.Id {
			r.productList[idx] = &p
		}
	}
	return &p, nil
}