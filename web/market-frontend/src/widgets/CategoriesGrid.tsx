import { useState, useEffect } from "react";
import CategoryCard from "../elements/CategoryCard.tsx";
import { fetchJson } from "../lib/api.ts";
import type { Category } from "../types/catalog.ts";

const EMPTY_FALLBACK: Category[] = [];

type CategoriesGridProps = {
  /** Показываются, если с сервера пришёл пустой список или запрос не удался. */
  fallbackCategories?: Category[];
  onCategoryClick?: (category: Category) => void;
};

export default function CategoriesGrid({ fallbackCategories, onCategoryClick }: CategoriesGridProps) {
  const fallbacks = fallbackCategories ?? EMPTY_FALLBACK;
  const [categories, setCategories] = useState<Category[]>([]);
  const [usedFallback, setUsedFallback] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const data = await fetchJson<Category[]>("/categories");
        if (cancelled) return;
        const list = Array.isArray(data) ? data : [];
        if (list.length > 0) {
          setCategories(list);
          setUsedFallback(false);
          setError(null);
        } else if (fallbacks.length > 0) {
          setCategories(fallbacks);
          setUsedFallback(true);
          setError(null);
        } else {
          setCategories([]);
          setUsedFallback(false);
          setError(null);
        }
      } catch (e) {
        if (cancelled) return;
        setError(e instanceof Error ? e.message : "Ошибка загрузки");
        setCategories(fallbacks);
        setUsedFallback(fallbacks.length > 0);
      }
    })();
    return () => {
      cancelled = true;
    };
  }, [fallbacks]);

  return (
    <div className="max-w-7xl mx-auto p-4 md:p-8">
      <div className="mb-6">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Категории</h1>
        <p className="text-slate-600 mt-1 text-sm md:text-base">Разделы каталога</p>
      </div>

      {error ? (
        <p className="mb-4 rounded-lg border border-amber-200 bg-amber-50 px-4 py-2 text-sm text-amber-900">
          Не удалось загрузить категории с сервера — {error}
          {usedFallback ? " Показан локальный список." : ""}
        </p>
      ) : null}

      {usedFallback && !error ? (
        <p className="mb-4 text-sm text-slate-500">С сервера категорий нет — показан демонстрационный список.</p>
      ) : null}

      {categories.length === 0 ? (
        <p className="text-slate-600">Категории пока не заведены.</p>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          {categories.map((category) => (
            <CategoryCard
              key={category.id}
              name={category.name}
              description={category.description}
              onClick={() => onCategoryClick?.(category)}
            />
          ))}
        </div>
      )}
    </div>
  );
}
