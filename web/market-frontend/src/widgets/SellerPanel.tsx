import { useState } from "react";

interface Product {
    id: number;
    name: string;
    description: string;
    price: number;
    image: string;
    category: string;
    ownerId: number; // айди пользователя-продавца
}

// Пример JWT данные
const jwt = {
    role: "seller", // "admin" или "seller"
    userId: 2,
};

// Тестовые товары
const initialProducts: Product[] = [
    { id: 1, name: "Телефон", description: "Смартфон 2024", price: 29999, image: "https://via.placeholder.com/150", category: "Электроника", ownerId: 2 },
    { id: 2, name: "Ноутбук", description: "Ноутбук i5", price: 59999, image: "https://via.placeholder.com/150", category: "Электроника", ownerId: 3 },
    { id: 3, name: "Книга", description: "Фантастика", price: 499, image: "https://via.placeholder.com/150", category: "Книги", ownerId: 2 },
];

export default function SellerPanel() {
    const [products, setProducts] = useState<Product[]>(initialProducts);

    // Фильтруем товары в зависимости от роли
    const visibleProducts = jwt.role === "admin"
        ? products
        : products.filter(p => p.ownerId === jwt.userId);

    const handleDelete = (id: number) => {
        if (confirm("Удалить этот товар?")) {
            setProducts(products.filter(p => p.id !== id));
        }
    };

    return (
        <div className="max-w-5xl mx-auto p-6">
            <h1 className="text-3xl font-bold mb-6">Панель {jwt.role === "admin" ? "администратора" : "продавца"}</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
                {visibleProducts.map(product => (
                    <div key={product.id} className="border rounded shadow p-4 flex flex-col">
                        <img src={product.image} alt={product.name} className="mb-2 h-40 object-cover rounded"/>
                        <h2 className="font-bold text-lg">{product.name}</h2>
                        <p className="text-sm text-gray-600 mb-2">{product.description}</p>
                        <p className="font-semibold mb-2">{product.price} ₽</p>
                        <p className="text-xs text-gray-500 mb-2">Категория: {product.category}</p>
                        <div className="mt-auto flex gap-2">
                            <button
                                onClick={() => alert(`Редактировать товар "${product.name}"`)}
                                className="flex-1 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
                            >
                                Редактировать
                            </button>
                            <button
                                onClick={() => handleDelete(product.id)}
                                className="flex-1 py-1 bg-red-500 text-white rounded hover:bg-red-600"
                            >
                                Удалить
                            </button>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}