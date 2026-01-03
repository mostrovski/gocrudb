package database

import "gocrudb/resource"

func GetSeedItems() []resource.Item {
	return []resource.Item{
		{Name: "Laptop", Stock: 10, Price: 999.99},
		{Name: "Smartphone", Stock: 20, Price: 699.99},
		{Name: "Headphones", Stock: 15, Price: 199.99},
		{Name: "Keyboard", Stock: 25, Price: 89.99},
		{Name: "Mouse", Stock: 30, Price: 49.99},
		{Name: "Monitor", Stock: 12, Price: 299.99},
		{Name: "Webcam", Stock: 18, Price: 79.99},
		{Name: "Printer", Stock: 7, Price: 149.99},
		{Name: "Tablet", Stock: 5, Price: 399.99},
		{Name: "Smartwatch", Stock: 14, Price: 249.99},
		{Name: "External Hard Drive", Stock: 8, Price: 119.99},
		{Name: "USB Flash Drive", Stock: 50, Price: 19.99},
		{Name: "Router", Stock: 6, Price: 89.99},
		{Name: "Projector", Stock: 3, Price: 499.99},
		{Name: "Bluetooth Speaker", Stock: 22, Price: 129.99},
		{Name: "Gaming Console", Stock: 11, Price: 499.99},
		{Name: "Camera", Stock: 4, Price: 599.99},
		{Name: "Fitness Tracker", Stock: 16, Price: 99.99},
		{Name: "Drone", Stock: 2, Price: 899.99},
		{Name: "VR Headset", Stock: 9, Price: 399.99},
	}
}
