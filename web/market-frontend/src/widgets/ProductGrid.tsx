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

  const addToCart = async (product: Product) => {
    try {
      await fetchJson(
          "/cart/items",
          {
            method: "POST",
            body: JSON.stringify({
              product_id: product.id,
              quantity: 1,
            }),
          },
          { auth: true }
      );
    } catch (e) {
      console.error("Add to cart failed", e);
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
          <h1 className="text-2xl md:text-3xl font-bold text-slate-900">
            Каталог
          </h1>
        </div>

        {error ? (
            <div className="text-red-500 mb-4">{error}</div>
        ) : null}

        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {products.map((product) => (
              <ProductCard
                  key={product.id}
                  name={product.name}
                  description={product.description}
                  price={product.price}
                  imageUrl={product.image_url}
                  onClick={() => navigate(`/product/${product.id}`)}
                  onAddToCart={() => addToCart(product)}
              />
          ))}
        </div>
      </div>
  );
}