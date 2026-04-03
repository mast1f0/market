import { useState } from "react";

interface ProductForm {
    name: string;
    description: string;
    price: number;
    image: string;
    category: string;
}

export default function AddProductPage() {
    const [form, setForm] = useState<ProductForm>({
        name: "",
        description: "",
        price: 0,
        image: "",
        category: "",
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setForm(prev => ({
            ...prev,
            [name]: name === "price" ? Number(value) : value
        }));
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log("Товар добавлен:", form);
        alert(`Товар "${form.name}" добавлен!`);
        // Здесь можно добавить отправку на сервер
        setForm({ name: "", description: "", price: 0, image: "", category: "" });
    };

    return (
        <div className="max-w-2xl mx-auto p-6 bg-white shadow-md rounded">
            <h1 className="text-2xl font-bold mb-4">Добавить товар</h1>
            <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                <input
                    type="text"
                    name="name"
                    placeholder="Название товара"
                    value={form.name}
                    onChange={handleChange}
                    className="border p-2 rounded focus:outline-none focus:ring-2 focus:ring-purple-500"
                    required
                />
                <textarea
                    name="description"
                    placeholder="Описание товара"
                    value={form.description}
                    onChange={handleChange}
                    className="border p-2 rounded focus:outline-none focus:ring-2 focus:ring-purple-500"
                    required
                />
                <input
                    type="number"
                    name="price"
                    placeholder="Цена"
                    value={form.price}
                    onChange={handleChange}
                    className="border p-2 rounded focus:outline-none focus:ring-2 focus:ring-purple-500"
                    required
                />
                <input
                    type="text"
                    name="image"
                    placeholder="URL изображения"
                    value={form.image}
                    onChange={handleChange}
                    className="border p-2 rounded focus:outline-none focus:ring-2 focus:ring-purple-500"
                    required
                />
                <input
                    type="text"
                    name="category"
                    placeholder="Категория"
                    value={form.category}
                    onChange={handleChange}
                    className="border p-2 rounded focus:outline-none focus:ring-2 focus:ring-purple-500"
                    required
                />
                <button
                    type="submit"
                    className="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-700 font-bold"
                >
                    Добавить товар
                </button>
            </form>
        </div>
    );
}