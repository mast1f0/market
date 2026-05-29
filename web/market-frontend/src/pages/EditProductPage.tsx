import { useState, useEffect, type FormEvent, type ChangeEvent } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { useAuth } from "../auth/useAuth.ts";
import { apiUrl, bearerHeaders, fetchJson } from "../lib/api.ts";
import { parsePrice } from "../lib/format.ts";
import { minioLinkUrl } from "../lib/endpoints.ts";
import ResolvedImage from "../elements/ResolvedImage.tsx";
import type { Category, Product } from "../types/catalog.ts";

type ProductForm = {
  name: string;
  description: string;
  price: number;
  category_id: number;
  image_url: string;
  stock: number;
};

export default function EditProductPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { profile } = useAuth();
  const [form, setForm] = useState<ProductForm | null>(null);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [forbidden, setForbidden] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState<{ type: "ok" | "err"; text: string } | null>(null);

  useEffect(() => {
    if (!id) {
      setLoading(false);
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        const [product, cats] = await Promise.all([
          fetchJson<Product>(`/products/${id}`),
          fetchJson<Category[]>("/categories").catch(() => [] as Category[]),
        ]);
        if (cancelled) return;

        if (profile && product.owner_id != null && product.owner_id !== profile.user_id) {
          setForbidden(true);
          return;
        }

        if (!cancelled) {
          setCategories(Array.isArray(cats) ? cats : []);
          setForm({
            name: product.name,
            description: product.description ?? "",
            price: product.price,
            category_id: product.category_id ?? 1,
            image_url: product.image_url ?? "",
            stock: product.stock ?? 0,
          });
        }
      } catch {
        if (!cancelled) setMessage({ type: "err", text: "Не удалось загрузить товар." });
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [id, profile]);

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm((prev) =>
      prev
        ? {
            ...prev,
            [name]:
              name === "price"
                ? parsePrice(value)
                : name === "category_id" || name === "stock"
                  ? Number(value) || 0
                  : value,
          }
        : prev
    );
  };

  const handleFile = async (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    e.target.value = "";
    if (!file) return;
    setMessage(null);
    setUploading(true);
    try {
      const body = new FormData();
      body.append("file", file);
      const res = await fetch(minioLinkUrl("/upload"), { method: "POST", body });
      const text = await res.text();
      if (!res.ok) {
        setMessage({ type: "err", text: text || `Ошибка загрузки (${res.status})` });
        return;
      }
      const data = JSON.parse(text) as { id?: string; link?: string };
      if (data.id) {
        setForm((prev) => (prev ? { ...prev, image_url: String(data.id) } : prev));
        setMessage({ type: "ok", text: "Изображение загружено." });
      } else if (data.link) {
        setForm((prev) => (prev ? { ...prev, image_url: String(data.link) } : prev));
        setMessage({ type: "ok", text: "Изображение загружено." });
      } else {
        setMessage({ type: "err", text: "Неожиданный ответ сервера." });
      }
    } catch {
      setMessage({ type: "err", text: "Не удалось загрузить изображение." });
    } finally {
      setUploading(false);
    }
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    if (!form || !id) return;
    if (form.price <= 0) {
      setMessage({ type: "err", text: "Укажите цену больше нуля." });
      return;
    }
    setMessage(null);
    setSubmitting(true);
    try {
      const res = await fetch(apiUrl(`/products/${id}`), {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          ...bearerHeaders(),
        },
        body: JSON.stringify({
          name: form.name,
          description: form.description,
          price: form.price,
          category_id: form.category_id,
          image_url: form.image_url,
          stock: form.stock,
        }),
      });
      const text = await res.text();
      if (res.status === 401 || res.status === 403) {
        setMessage({ type: "err", text: "Недостаточно прав для редактирования." });
        return;
      }
      if (!res.ok) {
        setMessage({ type: "err", text: text || `Ошибка ${res.status}` });
        return;
      }
      setMessage({ type: "ok", text: "Изменения сохранены." });
      window.setTimeout(() => navigate("/seller-panel"), 800);
    } catch {
      setMessage({ type: "err", text: "Не удалось сохранить товар." });
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div className="max-w-xl mx-auto p-6 md:p-8">
        <div className="h-8 w-40 bg-slate-200 rounded animate-pulse" />
      </div>
    );
  }

  if (forbidden) {
    return (
      <div className="max-w-xl mx-auto p-6 md:p-8">
        <p className="text-slate-600 mb-4">Вы можете редактировать только свои товары.</p>
        <Link to="/seller-panel" className="text-sm font-medium text-emerald-700 hover:text-emerald-800">
          ← Панель продавца
        </Link>
      </div>
    );
  }

  if (!form) {
    return (
      <div className="max-w-xl mx-auto p-6 md:p-8">
        <p className="text-red-600 mb-4">{message?.text ?? "Товар не найден"}</p>
        <Link to="/seller-panel" className="text-sm font-medium text-emerald-700 hover:text-emerald-800">
          ← Панель продавца
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-xl mx-auto p-6 md:p-8">
      <Link to="/seller-panel" className="text-sm font-medium text-emerald-700 hover:text-emerald-800">
        ← Панель продавца
      </Link>
      <div className="mt-6 bg-white rounded-2xl border border-slate-100 shadow-sm p-6 md:p-8">
        <h1 className="text-2xl font-bold text-slate-900">Редактировать товар</h1>

        {message ? (
          <p
            className={`mt-4 text-sm rounded-lg px-3 py-2 ${
              message.type === "ok" ? "bg-emerald-50 text-emerald-900" : "bg-red-50 text-red-800"
            }`}
          >
            {message.text}
          </p>
        ) : null}

        <form onSubmit={handleSubmit} className="mt-6 flex flex-col gap-4">
          <label className="block">
            <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Название</span>
            <input
              type="text"
              name="name"
              required
              value={form.name}
              onChange={handleChange}
              className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
            />
          </label>
          <label className="block">
            <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Описание</span>
            <textarea
              name="description"
              required
              rows={4}
              value={form.description}
              onChange={handleChange}
              className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
            />
          </label>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <label className="block">
              <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Цена (₽)</span>
              <input
                type="number"
                name="price"
                required
                min={0.01}
                step={0.01}
                inputMode="decimal"
                placeholder="999.99"
                value={form.price > 0 ? form.price : ""}
                onChange={handleChange}
                className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
              />
            </label>
            <label className="block">
              <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Категория</span>
              {categories.length > 0 ? (
                <select
                  name="category_id"
                  value={form.category_id}
                  onChange={handleChange}
                  className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
                >
                  {categories.map((c) => (
                    <option key={c.id} value={c.id}>
                      {c.name}
                    </option>
                  ))}
                </select>
              ) : (
                <input
                  type="number"
                  name="category_id"
                  required
                  min={1}
                  value={form.category_id || ""}
                  onChange={handleChange}
                  className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
                />
              )}
            </label>
          </div>

          <div className="rounded-xl border border-slate-100 bg-slate-50/80 p-4 space-y-3">
            <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Изображение</span>
            <input
              type="file"
              accept="image/*"
              disabled={uploading}
              onChange={handleFile}
              className="text-sm text-slate-700 file:mr-3 file:py-1.5 file:px-3 file:rounded-lg file:border-0 file:bg-emerald-600 file:text-white file:text-sm file:font-medium hover:file:bg-emerald-700"
            />
            {uploading ? <span className="text-sm text-slate-600">Загрузка…</span> : null}
            {form.image_url ? (
              <div className="w-40 rounded-lg overflow-hidden border border-slate-200 bg-white">
                <ResolvedImage imageRef={form.image_url} className="w-full h-28 object-cover" />
              </div>
            ) : null}
          </div>

          <label className="block">
            <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Остаток на складе</span>
            <input
              type="number"
              name="stock"
              min={0}
              value={form.stock}
              onChange={handleChange}
              className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
            />
          </label>
          <button
            type="submit"
            disabled={submitting}
            className="mt-2 px-4 py-2.5 rounded-lg bg-emerald-600 text-white font-medium hover:bg-emerald-700 disabled:opacity-60 transition-colors"
          >
            {submitting ? "Сохранение…" : "Сохранить"}
          </button>
        </form>
      </div>
    </div>
  );
}
