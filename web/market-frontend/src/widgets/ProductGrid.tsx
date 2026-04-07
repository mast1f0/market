import { useState, useEffect } from "react";
import ProductCard from "../elements/Card.tsx";

type Product = {
    name: string;
    description: string;
    price: string;
    imageUrl: string;
};

const fetchProduct = async (): Promise<Product[]> => {
    try {
        const response = await fetch('http://localhost:8080/products');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        return [];
    }
};

export default function ProductGrid() {
    const [products, setProducts] = useState<Product[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loadProducts = async () => {
            const data = await fetchProduct();
            setProducts(data);
            setLoading(false);
        };

        loadProducts();
    }, []);

    if (loading) {
        return <div className="p-6">Загрузка...</div>;
    }

    return (
        <div className="p-6 bg-gray-100 min-h-screen">
            <h1 className="text-2xl font-bold mb-6">Наши товары</h1>

            <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {products.map((product, index) => (
                    <ProductCard
                        key={index}
                        name={product.name}
                        description={product.description}
                        price={product.price}
                        imageUrl={product.imageUrl}
                        onClick={() => alert(`Вы кликнули на ${product.name}`)}
                    />
                ))}
            </div>
        </div>
    );
}