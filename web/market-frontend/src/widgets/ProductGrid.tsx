import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import ProductCard from "../elements/Card.tsx";
import { fetchJson } from "../lib/api.ts";
import { useAddToCart } from "../hooks/useAddToCart.ts";
import type { Product } from "../types/catalog.ts";

export default function ProductGrid() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [addedId, setAddedId] = useState<number | null>(null);

  const navigate = useNavigate();
  const addToCart = useAddToCart();

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

  const handleAddToCart = async (product: Product) => {
    const ok = await addToCart(product.id);
    if (ok) {
      setAddedId(product.id);
      window.setTimeout(() => setAddedId((id) => (id === product.id ? null : id)), 2000);
    }
  };

  if (loading) {
    return (
      <div className="p-8 max-w-7xl mx-auto">
        <div className="h-8 w-48 bg-slate-200 rounded animate-pulse mb-8" />
      </div>
    );
  }

  return (
    <div className="p-6 md:p-8 max-w-7xl mx-auto min-h-[60vh]">
      <div className="mb-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Каталог</h1>
      </div>

      {error ? <div className="text-red-500 mb-4">{error}</div> : null}

      {products.length === 0 && !error ? (
        <p className="text-slate-600">Товаров пока нет.</p>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {products.map((product) => (
            <div key={product.id} className="relative">
              {addedId === product.id ? (
                <span className="absolute top-2 right-2 z-10 text-xs font-medium bg-emerald-600 text-white px-2 py-1 rounded-md">
                  В корзине
                </span>
              ) : null}
              <ProductCard
                name={product.name}
                description={product.description}
                price={product.price}
                imageUrl={product.image_url}
                onClick={() => navigate(`/product/${product.id}`)}
                onAddToCart={() => handleAddToCart(product)}
              />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
