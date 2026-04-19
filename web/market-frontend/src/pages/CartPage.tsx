import { useState } from "react";
import CartItem from "../elements/CartItem.tsx";
import { formatRub } from "../lib/format.ts";

interface CartLine {
  id: number;
  name: string;
  price: number;
  quantity: number;
  image: string;
}

export default function CartPage() {
  const [cart, setCart] = useState<CartLine[]>([
    { id: 1, name: "Футболка", price: 1990, quantity: 2, image: "https://via.placeholder.com/100" },
    { id: 2, name: "Кроссовки", price: 8990, quantity: 1, image: "https://via.placeholder.com/100" },
    { id: 3, name: "Рюкзак", price: 3490, quantity: 1, image: "https://via.placeholder.com/100" },
  ]);

  const handleRemove = (id: number) => {
    setCart(cart.filter((item) => item.id !== id));
  };

  const handleChangeQuantity = (id: number, newQuantity: number) => {
    if (newQuantity < 1) return;
    setCart(cart.map((item) => (item.id === id ? { ...item, quantity: newQuantity } : item)));
  };

  const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

  return (
    <div className="max-w-3xl mx-auto p-6 md:p-8">
      <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-2">Корзина</h1>
      <p className="text-slate-600 text-sm mb-8">Демо-данные: оформление заказа пока не подключено к API.</p>

      {cart.length === 0 ? (
        <div className="rounded-xl border border-dashed border-slate-200 bg-white p-10 text-center text-slate-600">
          Корзина пуста
        </div>
      ) : (
        <>
          <div className="flex flex-col">{cart.map((item) => (
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
            ))}</div>

          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mt-8 p-5 rounded-xl bg-white border border-slate-100 shadow-sm">
            <span className="text-lg font-semibold text-slate-900">
              Итого: <span className="text-emerald-700">{formatRub(totalPrice)}</span>
            </span>
            <button
              type="button"
              className="px-6 py-2.5 rounded-lg bg-emerald-600 text-white font-medium hover:bg-emerald-700 transition-colors"
            >
              Оформить заказ
            </button>
          </div>
        </>
      )}
    </div>
  );
}
