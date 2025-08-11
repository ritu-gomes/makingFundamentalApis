package database

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	prod1 := Product{
		Id:          1,
		Title:       "Wireless Mouse",
		Description: "Ergonomic wireless mouse with USB receiver",
		Price:       25.99,
		ImageUrl:    "https://m.media-amazon.com/images/I/61LtuGzXeaL._AC_SL1500_.jpg",
	}

	prod2 := Product{
		Id:          2,
		Title:       "Mechanical Keyboard",
		Description: "RGB backlit keyboard with clicky switches",
		Price:       59.99,
		ImageUrl:    "https://cdn.thewirecutter.com/wp-content/media/2024/04/mechanicalkeyboards-2048px-1353-2x1-1.jpg?width=2048&quality=75&crop=2:1&auto=webp",
	}

	prod3 := Product{
		Id:          3,
		Title:       "USB-C Hub",
		Description: "7-in-1 USB-C hub with HDMI, USB, and Ethernet",
		Price:       39.50,
		ImageUrl:    "https://www.ryans.com/storage/products/main/ugreen-type-c-male-to-quad-usb-female-meter-hub-11646050628.webp",
	}

	prod4 := Product{
		Id:          4,
		Title:       "Full HD Webcam",
		Description: "1080p webcam with built-in mic and autofocus",
		Price:       49.00,
		ImageUrl:    "https://diamu.com.bd/wp-content/uploads/2021/07/Logitech-C920-HD-Pro-Webcam-768x768.jpg.webp",
	}

	// prod5 := Product{
	// 	Id:          5,
	// 	Title:       "Bluetooth Headphones",
	// 	Description: "Wireless over-ear headphones with deep bass",
	// 	Price:       79.99,
	// 	ImageUrl:    "https://sonysmart.com.bd/public/uploads/all/G7qeCeK42taREYyfLLxQppDfnH4L0UKvj23z9QDW.jpg",
	// }
	ProductList = []Product{prod1, prod2, prod3, prod4}

}