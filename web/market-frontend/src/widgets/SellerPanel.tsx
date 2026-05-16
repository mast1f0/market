import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { apiUrl, bearerHeaders, fetchJson } from "../lib/api.ts";
import { formatRub } from "../lib/format.ts";
import type { Product } from "../types/catalog.ts";
import ResolvedImage from "../elements/ResolvedImage.tsx";

export default function SellerPanel() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const data = await fetchJson<Product[]>("/products");
        if (!cancelled) {
          setProducts(Array.isArray(data) ? data : []);
          setError(null);
        }
      } catch (e) {
        if (!cancelled) {
          setProducts([]);
          setError(e instanceof Error ? e.message : "Ошибка загрузки");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();
    return () => {
      cancelled = true;
    };
  }, []);

  const handleDelete = async (id: number) => {
    if (!confirm("Удалить этот товар?")) return;
    try {
      const res = await fetch(apiUrl(`/products/${id}`), {
        method: "DELETE",
        headers: bearerHeaders(),
      });
      if (!res.ok) {
        const text = await res.text();
        throw new Error(text || `Ошибка ${res.status}`);
      }
      setProducts((prev) => prev.filter((p) => p.id !== id));
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось удалить");
    }
  };

  return (
    <div className="max-w-6xl mx-auto p-6 md:p-8">
      <div className="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-4 mb-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Панель продавца</h1>
        <Link
          to="/add"
          className="inline-flex justify-center px-4 py-2 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 transition-colors"
        >
          Добавить товар
        </Link>
      </div>

      {loading ? (
        <p className="text-slate-600">Загрузка…</p>
      ) : error ? (
        <p className="text-red-600 text-sm">{error}</p>
      ) : products.length === 0 ? (
        <p className="text-slate-600">Нет товаров.</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {products.map((product) => (
            <div
              key={product.id}
              className="bg-white rounded-xl border border-slate-100 shadow-sm overflow-hidden flex flex-col"
            >
              <ResolvedImage imageRef={product.image_url} className="h-44 w-full object-cover bg-slate-100" />
              <div className="p-4 flex flex-col flex-1">
                <h2 className="font-semibold text-slate-900">{product.name}</h2>
                <p className="text-sm text-slate-600 mt-1 max-h-10 overflow-hidden">{product.description}</p>
                <p className="text-emerald-700 font-semibold mt-3">{formatRub(product.price)}</p>
                <div className="mt-auto pt-4 flex gap-2">
                  <Link
                    to={`/product/${product.id}`}
                    className="flex-1 py-2 text-sm text-center rounded-lg border border-slate-200 text-slate-700 hover:bg-slate-50"
                  >
                    Открыть
                  </Link>
                  <button
                    type="button"
                    className="flex-1 py-2 text-sm rounded-lg border border-red-200 text-red-700 hover:bg-red-50"
                    onClick={() => handleDelete(product.id)}
                  >
                    Удалить
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
