import { useState, useEffect } from "react";
import CategoryCard from "../elements/CategoryCard.tsx";
import { fetchJson } from "../lib/api.ts";
import type { Category } from "../types/catalog.ts";

type CategoriesGridProps = {
  onCategoryClick?: (category: Category) => void;
};

export default function CategoriesGrid({ onCategoryClick }: CategoriesGridProps) {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const data = await fetchJson<Category[]>("/categories");
        if (cancelled) return;
        setCategories(Array.isArray(data) ? data : []);
        setError(null);
      } catch (e) {
        if (!cancelled) {
          setError(e instanceof Error ? e.message : "Ошибка загрузки");
          setCategories([]);
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();
    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <div className="max-w-7xl mx-auto p-4 md:p-8">
      <div className="mb-6">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Категории</h1>
      </div>

      {loading ? (
        <p className="text-slate-600">Загрузка…</p>
      ) : error ? (
        <p className="mb-4 rounded-lg border border-red-200 bg-red-50 px-4 py-2 text-sm text-red-800">{error}</p>
      ) : categories.length === 0 ? (
        <p className="text-slate-600">Категории пока не заведены.</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          {categories.map((category) => (
            <CategoryCard
              key={category.id}
              name={category.name}
              onClick={() => onCategoryClick?.(category)}
            />
          ))}
        </div>
      )}
    </div>
  );
}
