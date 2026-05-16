import { useState, useEffect, type FormEvent, type ChangeEvent } from "react";
import { Link } from "react-router-dom";
import { apiUrl, bearerHeaders, fetchJson } from "../lib/api.ts";
import { minioLinkUrl } from "../lib/endpoints.ts";
import ResolvedImage from "../elements/ResolvedImage.tsx";
import type { Category } from "../types/catalog.ts";

type ProductForm = {
  name: string;
  description: string;
  price: number;
  category_id: number;
  image_url: string;
  stock: number;
};

const initial: ProductForm = {
  name: "",
  description: "",
  price: 0,
  category_id: 1,
  image_url: "",
  stock: 0,
};

export default function AddProductPage() {
  const [form, setForm] = useState<ProductForm>(initial);
  const [categories, setCategories] = useState<Category[]>([]);
  const [submitting, setSubmitting] = useState(false);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState<{ type: "ok" | "err"; text: string } | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const data = await fetchJson<Category[]>("/categories");
        if (!cancelled && Array.isArray(data) && data.length > 0) {
          setCategories(data);
          setForm((prev) => ({ ...prev, category_id: data[0].id }));
        }
      } catch {
        /* категории необязательны для формы */
      }
    })();
    return () => {
      cancelled = true;
    };
  }, []);

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]:
        name === "price" || name === "category_id" || name === "stock" ? Number(value) || 0 : value,
    }));
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
        setForm((prev) => ({ ...prev, image_url: String(data.id) }));
        setMessage({ type: "ok", text: "Изображение загружено." });
      } else if (data.link) {
        setForm((prev) => ({ ...prev, image_url: String(data.link) }));
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
    setMessage(null);
    setSubmitting(true);
    try {
      const res = await fetch(apiUrl("/products"), {
        method: "POST",
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
        setMessage({ type: "err", text: "Недостаточно прав. Войдите как продавец." });
        return;
      }
      if (!res.ok) {
        setMessage({ type: "err", text: text || `Ошибка ${res.status}` });
        return;
      }
      setMessage({ type: "ok", text: "Товар создан." });
      setForm((prev) => ({ ...initial, category_id: prev.category_id }));
    } catch {
      setMessage({ type: "err", text: "Не удалось создать товар." });
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="max-w-xl mx-auto p-6 md:p-8">
      <Link to="/" className="text-sm font-medium text-emerald-700 hover:text-emerald-800">
        ← В каталог
      </Link>
      <div className="mt-6 bg-white rounded-2xl border border-slate-100 shadow-sm p-6 md:p-8">
        <h1 className="text-2xl font-bold text-slate-900">Добавить товар</h1>

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
                min={1}
                value={form.price || ""}
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
            {submitting ? "Сохранение…" : "Создать товар"}
          </button>
        </form>
      </div>
    </div>
  );
}
