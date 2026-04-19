import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import ProductCard from "../elements/Card.tsx";
import { fetchJson } from "../lib/api.ts";
import type { Product } from "../types/catalog.ts";

export default function ProductGrid() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();

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
          setError(e instanceof Error ? e.message : "Не удалось загрузить каталог");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();
    return () => {
      cancelled = true;
    };
  }, []);

  if (loading) {
    return (
      <div className="p-8 max-w-7xl mx-auto">
        <div className="h-8 w-48 bg-slate-200 rounded animate-pulse mb-8" />
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {[1, 2, 3, 4].map((i) => (
            <div key={i} className="rounded-xl border border-slate-100 overflow-hidden bg-white">
              <div className="h-48 bg-slate-200 animate-pulse" />
              <div className="p-4 space-y-2">
                <div className="h-4 bg-slate-200 rounded animate-pulse" />
                <div className="h-3 bg-slate-100 rounded animate-pulse w-4/5" />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="p-6 md:p-8 max-w-7xl mx-auto min-h-[60vh]">
      <div className="mb-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900 tracking-tight">Каталог</h1>
        <p className="text-slate-600 mt-1 text-sm md:text-base">Выберите товар, чтобы открыть карточку</p>
      </div>

      {error ? (
        <div className="rounded-lg border border-amber-200 bg-amber-50 text-amber-900 px-4 py-3 text-sm mb-6">
          {error}
          <span className="block mt-1 text-amber-800/80">
            Убедитесь, что API запущен (например, на порту 8080) и что вы открыли сайт через{" "}
            <code className="text-xs bg-amber-100/80 px-1 rounded">npm run dev</code> — тогда запросы идут через прокси{" "}
            <code className="text-xs bg-amber-100/80 px-1 rounded">/api</code>.
          </span>
        </div>
      ) : null}

      {!error && products.length === 0 ? (
        <p className="text-slate-600">Пока нет товаров в каталоге.</p>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {products.map((product) => (
            <ProductCard
              key={product.id}
              name={product.name}
              description={product.description}
              price={product.price}
              imageUrl={product.image_url}
              onClick={() => navigate(`/product/${product.id}`)}
            />
          ))}
        </div>
      )}
    </div>
  );
}
