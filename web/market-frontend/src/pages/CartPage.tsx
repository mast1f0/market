import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import CartItem from "../elements/CartItem.tsx";
import { formatRub } from "../lib/format.ts";
import { bearerHeaders } from "../lib/api.ts";
import { marketApiUrl } from "../lib/endpoints.ts";
import { normalizeCart, type CartDTO } from "../lib/cart.ts";
import { useAuth } from "../auth/useAuth.ts";

export default function CartPage() {
  const { token } = useAuth();
  const navigate = useNavigate();
  const [cart, setCart] = useState<CartDTO | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [checkoutBusy, setCheckoutBusy] = useState(false);

  useEffect(() => {
    if (!token) {
      setLoading(false);
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        const res = await fetch(marketApiUrl("/cart"), { headers: bearerHeaders() });
        if (!res.ok) throw new Error(await res.text());
        const raw = await res.json();
        if (!cancelled) setCart(normalizeCart(raw));
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
  }, [token]);

  const items = cart?.items ?? [];

  const handleRemove = async (id: number) => {
    const prev = cart;
    const item = cart?.items.find((i) => i.id === id);
    if (!item) return;

    setCart((p) => (p ? { ...p, items: p.items.filter((i) => i.id !== id) } : p));

    try {
      const res = await fetch(marketApiUrl("/cart/items"), {
        method: "DELETE",
        headers: { "Content-Type": "application/json", ...bearerHeaders() },
        body: JSON.stringify({ item_id: item.id }),
      });
      if (!res.ok) throw new Error(await res.text());
    } catch {
      setCart(prev);
    }
  };

  const handleChangeQuantity = async (id: number, newQuantity: number) => {
    if (newQuantity < 1) return;

    const prev = cart;
    setCart((p) =>
      p
        ? {
            ...p,
            items: p.items.map((i) => (i.id === id ? { ...i, quantity: newQuantity } : i)),
          }
        : p
    );

    try {
      const res = await fetch(marketApiUrl("/cart/items"), {
        method: "PUT",
        headers: { "Content-Type": "application/json", ...bearerHeaders() },
        body: JSON.stringify({ item_id: id, quantity: newQuantity }),
      });
      if (!res.ok) throw new Error(await res.text());
    } catch {
      setCart(prev);
    }
  };

  const handleCheckout = async () => {
    setCheckoutBusy(true);
    setError(null);
    try {
      const res = await fetch(marketApiUrl("/orders"), {
        method: "POST",
        headers: bearerHeaders(),
      });
      if (!res.ok) throw new Error(await res.text());
      setCart({ id: cart?.id ?? 0, items: [] });
      navigate("/profiles?ordered=1");
    } catch (e) {
      setError(e instanceof Error ? e.message : "Не удалось оформить заказ");
    } finally {
      setCheckoutBusy(false);
    }
  };

  const totalPrice = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

  if (!token) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl font-bold text-slate-900 mb-4">Корзина</h1>
        <p className="text-slate-600 mb-6">Войдите, чтобы просматривать корзину и оформлять заказы.</p>
        <Link
          to="/login"
          className="inline-flex px-4 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700"
        >
          Войти
        </Link>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <div className="h-8 w-40 bg-slate-200 rounded animate-pulse mb-6" />
      </div>
    );
  }

  if (error && items.length === 0) {
    return (
      <div className="max-w-3xl mx-auto p-6 text-red-600">
        {error}
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
        <Link to="/" className="mt-6 inline-block text-emerald-700 hover:text-emerald-800 text-sm font-medium">
          Перейти в каталог
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-3xl mx-auto p-6 md:p-8">
      <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-6">Корзина</h1>

      {error ? (
        <p className="mb-4 text-sm text-red-700 bg-red-50 border border-red-100 rounded-lg px-3 py-2">{error}</p>
      ) : null}

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
          Итого: <span className="text-emerald-700">{formatRub(totalPrice)}</span>
        </span>
        <button
          type="button"
          disabled={checkoutBusy}
          onClick={handleCheckout}
          className="px-6 py-2.5 rounded-lg bg-emerald-600 text-white hover:bg-emerald-700 transition disabled:opacity-60"
        >
          {checkoutBusy ? "Оформление…" : "Оформить заказ"}
        </button>
      </div>
    </div>
  );
}
