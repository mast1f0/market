import { Link, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { fetchJson } from "../lib/api.ts";
import { formatRub } from "../lib/format.ts";
import type { Product } from "../types/catalog.ts";
import ResolvedImage from "../elements/ResolvedImage.tsx";

export default function ProductPage() {
  const { id } = useParams();

  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) {
      setLoading(false);
      return;
    }
    let cancelled = false;
    (async () => {
      try {
        const data = await fetchJson<Product>(`/products/${id}`);
        if (!cancelled) {
          setProduct(data);
          setError(null);
        }
      } catch (e) {
        if (!cancelled) {
          setProduct(null);
          setError(e instanceof Error ? e.message : "Ошибка загрузки");
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
      <div className="max-w-4xl mx-auto p-6 md:p-8">
        <div className="h-4 w-32 bg-slate-200 rounded animate-pulse mb-6" />
        <div className="h-64 md:h-80 bg-slate-200 rounded-xl animate-pulse mb-6" />
        <div className="h-8 w-2/3 bg-slate-200 rounded animate-pulse mb-4" />
        <div className="h-4 w-full bg-slate-100 rounded animate-pulse" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="max-w-4xl mx-auto p-6 md:p-8">
        <Link to="/" className="text-emerald-700 hover:text-emerald-800 text-sm font-medium">
          ← Назад в каталог
        </Link>
        <p className="mt-6 text-red-600">{error}</p>
      </div>
    );
  }

  if (!product) {
    return (
      <div className="max-w-4xl mx-auto p-6 md:p-8">
        <Link to="/" className="text-emerald-700 hover:text-emerald-800 text-sm font-medium">
          ← Назад в каталог
        </Link>
        <p className="mt-6 text-slate-600">Товар не найден</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto p-6 md:p-8 pb-16">
      <Link to="/" className="inline-block text-emerald-700 hover:text-emerald-800 text-sm font-medium mb-6">
        ← Назад в каталог
      </Link>

      <div className="bg-white rounded-2xl border border-slate-100 shadow-sm overflow-hidden">
        <ResolvedImage
          imageRef={product.image_url}
          className="w-full h-64 md:h-96 object-cover bg-slate-100"
        />

        <div className="p-6 md:p-8">
          <h1 className="text-2xl md:text-3xl font-bold text-slate-900">{product.name}</h1>
          <p className="text-2xl font-semibold text-emerald-700 mt-4">{formatRub(product.price)}</p>

          {product.description ? (
            <p className="text-slate-600 mt-6 leading-relaxed whitespace-pre-wrap">{product.description}</p>
          ) : null}

          <dl className="mt-8 grid gap-2 text-sm text-slate-600 border-t border-slate-100 pt-6">
            {product.category_id != null ? (
              <div className="flex gap-2">
                <dt className="text-slate-500 shrink-0">Категория</dt>
                <dd className="font-medium text-slate-800">№{product.category_id}</dd>
              </div>
            ) : null}
            {product.stock != null ? (
              <div className="flex gap-2">
                <dt className="text-slate-500 shrink-0">Остаток</dt>
                <dd className="font-medium text-slate-800">{product.stock} шт.</dd>
              </div>
            ) : null}
          </dl>

          <div className="mt-8 flex flex-wrap gap-3">
            <button
              type="button"
              className="px-5 py-2.5 rounded-lg bg-emerald-600 text-white font-medium hover:bg-emerald-700 transition-colors"
            >
              В корзину
            </button>
            <Link
              to="/cart"
              className="px-5 py-2.5 rounded-lg border border-slate-200 text-slate-800 font-medium hover:bg-slate-50 transition-colors"
            >
              Перейти в корзину
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
