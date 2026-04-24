import { useEffect, useState } from "react";
import CartItem from "../elements/CartItem.tsx";
import { formatRub } from "../lib/format.ts";
import { marketApiUrl } from "../lib/endpoints.ts";

// interface ProductDTO {
//   id: number;
//   name: string;
//   price: number;
//   image_url: string;
// }

interface CartItemDTO {
  id: number;
  product_id: number;
  quantity: number;
  price: number;
  name: string;
  image_url: string;
}

interface CartDTO {
  id: number;
  items: CartItemDTO[];
}

const token = () => localStorage.getItem("market_access_token");

export default function CartPage() {
  const [cart, setCart] = useState<CartDTO | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // ===================== LOAD CART =====================
  useEffect(() => {
    let cancelled = false;

    (async () => {
      try {
        const res = await fetch(marketApiUrl("/cart"), {
          headers: {
            Authorization: `Bearer ${token()}`,
          },
        });

        if (!res.ok) throw new Error(await res.text());

        const raw = await res.json();

        const normalized: CartDTO = {
          id: raw.id ?? raw.ID,
          items: (raw.items ?? raw.Items ?? []).map((item: any) => ({
            id: item.id ?? item.ID,
            product_id: item.product_id ?? item.ProductID,
            quantity: item.quantity ?? item.Quantity,

            price:
                item.price ??
                item.price_snapshot ??
                item.PriceSnapshot ??
                item.product?.price ??
                item.Product?.price ??
                0,

            name:
                item.name ??
                item.product?.name ??
                item.Product?.name ??
                item.Product?.Name ??
                "Без названия",

            image_url:
                item.image_url ??
                item.product?.image_url ??
                item.Product?.image_url ??
                item.Product?.ImageURL ??
                "",
          })),
        };

        if (!cancelled) {
          setCart(normalized);
        }
      } catch (e) {
        if (!cancelled) {
          setError(e instanceof Error ? e.message : "Ошибка загрузки корзины");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, []);

  const items = cart?.items ?? [];

  // ===================== DELETE ITEM =====================
  const handleRemove = async (id: number) => {
    const prev = cart;
    const item = cart?.items.find((i) => i.id === id);
    if (!item) return;

    // optimistic UI
    setCart((p) =>
        p
            ? { ...p, items: p.items.filter((i) => i.id !== id) }
            : p
    );

    try {
      const res = await fetch(marketApiUrl("/cart/items"), {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token()}`,
        },
        body: JSON.stringify({
          item_id: item.id,
        }),
      });

      if (!res.ok) throw new Error(await res.text());
    } catch (e) {
      console.error("Remove failed", e);
      setCart(prev); // rollback
    }
  };

  // ===================== UPDATE QUANTITY =====================
  const handleChangeQuantity = async (id: number, newQuantity: number) => {
    if (newQuantity < 1) return;

    const prev = cart;

    // optimistic update
    setCart((p) =>
        p
            ? {
              ...p,
              items: p.items.map((i) =>
                  i.id === id ? { ...i, quantity: newQuantity } : i
              ),
            }
            : p
    );

    try {
      const res = await fetch("http://localhost:8080/cart/items", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token()}`,
        },
        body: JSON.stringify({
          item_id: id,
          quantity: newQuantity,
        }),
      });

      if (!res.ok) throw new Error(await res.text());
    } catch (e) {
      console.error("Update failed", e);
      setCart(prev); // rollback
    }
  };

  // ===================== TOTAL =====================
  const totalPrice = items.reduce(
      (sum, item) => sum + item.price * item.quantity,
      0
  );

  // ===================== UI STATES =====================
  if (loading) {
    return (
        <div className="max-w-3xl mx-auto p-6 md:p-8">
          <div className="h-8 w-40 bg-slate-200 rounded animate-pulse mb-6" />
        </div>
    );
  }

  if (error) {
    return (
        <div className="max-w-3xl mx-auto p-6 text-red-600">
          Ошибка: {error}
        </div>
    );
  }

  if (items.length === 0) {
    return (
        <div className="max-w-3xl mx-auto p-6 md:p-8">
          <h1 className="text-2xl font-bold mb-6">Корзина</h1>
          <div className="rounded-xl border border-dashed border-slate-200 bg-white p-10 text-center text-slate-600">
            Корзина пуста
          </div>
        </div>
    );
  }

  // ===================== RENDER =====================
  return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-6">
          Корзина
        </h1>

        <div className="flex flex-col">
          {items.map((item) => (
              <CartItem
                  key={item.id}
                  id={item.id}
                  name={item.name}
                  price={item.price}
                  quantity={item.quantity}
                  image={item.image_url}
                  onRemove={handleRemove}
                  onChangeQuantity={handleChangeQuantity}
              />
          ))}
        </div>

        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mt-8 p-5 rounded-xl bg-white border border-slate-100 shadow-sm">
        <span className="text-lg font-semibold text-slate-900">
          Итого:{" "}
          <span className="text-emerald-700">
            {formatRub(totalPrice)}
          </span>
        </span>

          <button className="px-6 py-2.5 rounded-lg bg-emerald-600 text-white hover:bg-emerald-700 transition">
            Оформить заказ
          </button>
        </div>
      </div>
  );
}