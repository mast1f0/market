import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import ProductCard from "../elements/Card.tsx";
import { fetchJson } from "../lib/api.ts";
import { useAddToCart } from "../hooks/useAddToCart.ts";
import type { Product } from "../types/catalog.ts";

function normalizeProduct(raw: unknown): Product {
  const data = (raw ?? {}) as Record<string, unknown>;
  return {
    id: Number(data.id ?? 0),
    owner_id: data.owner_id != null ? Number(data.owner_id) : undefined,
    name: String(data.name ?? ""),
    description: String(data.description ?? ""),
    price: Number(data.price ?? 0),
    category_id: data.category_id != null ? Number(data.category_id) : undefined,
    image_url: String(data.image_url ?? data.imageUrl ?? data.ImageURL ?? ""),
    stock: data.stock != null ? Number(data.stock) : undefined,
    created_at: data.created_at != null ? String(data.created_at) : undefined,
  };
}

export default function CategoryPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const addToCart = useAddToCart();

  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;

    let cancelled = false;

    (async () => {
      try {
        const data = await fetchJson<unknown[]>(`/categories/${id}`);
        if (!cancelled) {
          setProducts(Array.isArray(data) ? data.map(normalizeProduct) : []);
          setError(null);
        }
      } catch (e) {
        if (!cancelled) {
          setProducts([]);
          setError(e instanceof Error ? e.message : "Не удалось загрузить товары");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [id]);

  if (loading) {
    return (
      <div className="p-8 max-w-7xl mx-auto">
        <div className="h-8 w-56 bg-slate-200 rounded animate-pulse mb-8" />
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
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900 tracking-tight">Товары категории</h1>
      </div>

      {error ? (
        <div className="rounded-lg border border-amber-200 bg-amber-50 text-amber-900 px-4 py-3 text-sm mb-6">
          {error}
        </div>
      ) : null}

      {!error && products.length === 0 ? (
        <p className="text-slate-600">В этой категории пока нет товаров.</p>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {products.map((product) => (
            <div key={product.id} className="h-full">
              <ProductCard
                name={product.name}
                description={product.description}
                price={product.price}
                imageUrl={product.image_url}
                onClick={() => navigate(`/product/${product.id}`)}
                onAddToCart={() => addToCart(product.id)}
              />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
