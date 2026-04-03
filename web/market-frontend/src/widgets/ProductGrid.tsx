import ProductCard from "../elements/Card.tsx";

const products = [
    {
        name: "Ноутбук",
        description: "Мощный игровой ноутбук",
        price: "1500$",
        imageUrl: "https://via.placeholder.com/300x200",
    },
    {
        name: "Смартфон",
        description: "Современный смартфон с камерой 108MP",
        price: "800$",
        imageUrl: "https://via.placeholder.com/300x200",
    },
    {
        name: "Наушники",
        description: "Беспроводные наушники с шумоподавлением",
        price: "200$",
        imageUrl: "https://via.placeholder.com/300x200",
    },
    {
        name: "Часы",
        description: "Умные часы с большим дисплеем",
        price: "300$",
        imageUrl: "https://via.placeholder.com/300x200",
    },
];

export default function ProductGrid() {
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