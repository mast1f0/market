import { useState } from "react";
import CartItem from "../elements/CartItem.tsx";
interface Product {
    id: number;
    name: string;
    price: number;
    quantity: number;
    image: string;
}

export default function CartPage() {
    const [cart, setCart] = useState<Product[]>([
        { id: 1, name: "Футболка", price: 20, quantity: 2, image: "https://via.placeholder.com/100" },
        { id: 2, name: "Кроссовки", price: 75, quantity: 1, image: "https://via.placeholder.com/100" },
        { id: 3, name: "Рюкзак", price: 40, quantity: 3, image: "https://via.placeholder.com/100" },
    ]);

    const handleRemove = (id: number) => {
        setCart(cart.filter(item => item.id !== id));
    };

    const handleChangeQuantity = (id: number, newQuantity: number) => {
        if (newQuantity < 1) return;
        setCart(cart.map(item => item.id === id ? { ...item, quantity: newQuantity } : item));
    };

    const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

    return (
        <div className="max-w-4xl mx-auto p-6">
            <h1 className="text-3xl font-bold mb-6">Корзина</h1>

            {cart.length === 0 ? (
                <p className="text-gray-600">Ваша корзина пуста</p>
            ) : (
                <>
                    <div className="flex flex-col">
                        {cart.map(item => (
                            <CartItem
                                key={item.id}
                                id={item.id}
                                name={item.name}
                                price={item.price}
                                quantity={item.quantity}
                                image={item.image}
                                onRemove={handleRemove}
                                onChangeQuantity={handleChangeQuantity}
                            />
                        ))}
                    </div>

                    <div className="flex justify-between items-center mt-6 p-4 bg-gray-100 rounded">
                        <span className="text-xl font-semibold">Итого: ${totalPrice.toFixed(2)}</span>
                        <button className="px-6 py-2 bg-purple-600 text-white font-bold rounded hover:bg-purple-700">
                            Оформить заказ
                        </button>
                    </div>
                </>
            )}
        </div>
    );
}