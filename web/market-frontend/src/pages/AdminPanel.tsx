import { useEffect, useState, type FormEvent } from "react";
import { Link } from "react-router-dom";
import { apiUrl, bearerHeaders, fetchJson } from "../lib/api.ts";
import { updateOrderStatus } from "../lib/orders.ts";
import type { Category } from "../types/catalog.ts";
import AdminUsers from "../widgets/AdminUsers.tsx";
import AdminOrderSearch from "../widgets/AdminOrderSearch.tsx";
import OrderStatusBadge from "../elements/OrderStatusBadge.tsx";
import OrderStatusPicker from "../elements/OrderStatusPicker.tsx";

export default function AdminPanel() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [newName, setNewName] = useState("");
  const [creating, setCreating] = useState(false);

  const [editingId, setEditingId] = useState<number | null>(null);
  const [editName, setEditName] = useState("");
  const [savingId, setSavingId] = useState<number | null>(null);

  const [orderIdInput, setOrderIdInput] = useState("");
  const [orderStatus, setOrderStatus] = useState<string>("processing");
  const [orderUpdating, setOrderUpdating] = useState(false);
  const [orderMessage, setOrderMessage] = useState<{ type: "ok" | "err"; text: string } | null>(null);

  const loadCategories = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await fetchJson<Category[]>("/categories");
      setCategories(Array.isArray(data) ? data : []);
    } catch (e) {
      setCategories([]);
      setError(e instanceof Error ? e.message : "Ошибка загрузки категорий");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadCategories();
  }, []);

  const handleCreate = async (e: FormEvent) => {
    e.preventDefault();
    const name = newName.trim();
    if (!name) return;
    setCreating(true);
    try {
      const res = await fetch(apiUrl("/categories"), {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          ...bearerHeaders(),
        },
        body: JSON.stringify({ name }),
      });
      const text = await res.text();
      if (!res.ok) throw new Error(text || `Ошибка ${res.status}`);
      setNewName("");
      await loadCategories();
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось создать категорию");
    } finally {
      setCreating(false);
    }
  };

  const startEdit = (cat: Category) => {
    setEditingId(cat.id);
    setEditName(cat.name);
  };

  const cancelEdit = () => {
    setEditingId(null);
    setEditName("");
  };

  const handleSave = async (id: number) => {
    const name = editName.trim();
    if (!name) return;
    setSavingId(id);
    try {
      const res = await fetch(apiUrl(`/categories/${id}`), {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          ...bearerHeaders(),
        },
        body: JSON.stringify({ name }),
      });
      const text = await res.text();
      if (!res.ok) throw new Error(text || `Ошибка ${res.status}`);
      setEditingId(null);
      await loadCategories();
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось сохранить");
    } finally {
      setSavingId(null);
    }
  };

  const handleDelete = async (id: number) => {
    if (!confirm("Удалить категорию? Товары с этой категорией могут стать недоступны.")) return;
    try {
      const res = await fetch(apiUrl(`/categories/${id}`), {
        method: "DELETE",
        headers: bearerHeaders(),
      });
      if (!res.ok) {
        const text = await res.text();
        throw new Error(text || `Ошибка ${res.status}`);
      }
      await loadCategories();
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось удалить");
    }
  };

  const handleOrderStatus = async (e: FormEvent) => {
    e.preventDefault();
    const id = Number(orderIdInput);
    if (!id || id < 1) {
      setOrderMessage({ type: "err", text: "Укажите корректный номер заказа." });
      return;
    }
    setOrderUpdating(true);
    setOrderMessage(null);
    try {
      await updateOrderStatus(id, orderStatus);
      setOrderMessage({ type: "ok", text: `Статус заказа №${id} обновлён.` });
    } catch (e) {
      setOrderMessage({
        type: "err",
        text: e instanceof Error ? e.message : "Не удалось обновить статус",
      });
    } finally {
      setOrderUpdating(false);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-6 md:p-8">
      <div className="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-4 mb-8">
        <div>
          <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Панель администратора</h1>
          <p className="mt-1 text-sm text-slate-600">Пользователи, категории и заказы</p>
        </div>
        <Link
          to="/seller-panel"
          className="inline-flex justify-center px-4 py-2 rounded-lg border border-slate-200 text-slate-800 text-sm font-medium hover:bg-slate-50"
        >
          Товары
        </Link>
      </div>

      <AdminUsers />

      <AdminOrderSearch />

      <section className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6 mb-8">
        <h2 className="text-lg font-semibold text-slate-900 mb-4">Категории</h2>

        <form onSubmit={handleCreate} className="flex flex-col sm:flex-row gap-2 mb-6">
          <input
            type="text"
            value={newName}
            onChange={(e) => setNewName(e.target.value)}
            placeholder="Название новой категории"
            required
            className="flex-1 border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
          />
          <button
            type="submit"
            disabled={creating}
            className="px-4 py-2 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 disabled:opacity-60"
          >
            {creating ? "…" : "Добавить"}
          </button>
        </form>

        {loading ? (
          <p className="text-slate-600 text-sm">Загрузка…</p>
        ) : error ? (
          <p className="text-red-600 text-sm">{error}</p>
        ) : categories.length === 0 ? (
          <p className="text-slate-600 text-sm">Категорий пока нет.</p>
        ) : (
          <ul className="divide-y divide-slate-100">
            {categories.map((cat) => (
              <li key={cat.id} className="py-3 flex flex-col sm:flex-row sm:items-center gap-3">
                {editingId === cat.id ? (
                  <>
                    <input
                      type="text"
                      value={editName}
                      onChange={(e) => setEditName(e.target.value)}
                      className="flex-1 border border-slate-200 rounded-lg px-3 py-2 text-sm"
                    />
                    <div className="flex gap-2 shrink-0">
                      <button
                        type="button"
                        disabled={savingId === cat.id}
                        onClick={() => handleSave(cat.id)}
                        className="px-3 py-1.5 text-sm rounded-lg bg-emerald-600 text-white hover:bg-emerald-700 disabled:opacity-60"
                      >
                        Сохранить
                      </button>
                      <button
                        type="button"
                        onClick={cancelEdit}
                        className="px-3 py-1.5 text-sm rounded-lg border border-slate-200 hover:bg-slate-50"
                      >
                        Отмена
                      </button>
                    </div>
                  </>
                ) : (
                  <>
                    <span className="flex-1 font-medium text-slate-900">{cat.name}</span>
                    <span className="text-xs text-slate-400 sm:mr-2">id {cat.id}</span>
                    <div className="flex gap-2 shrink-0">
                      <button
                        type="button"
                        onClick={() => startEdit(cat)}
                        className="px-3 py-1.5 text-sm rounded-lg border border-slate-200 hover:bg-slate-50"
                      >
                        Изменить
                      </button>
                      <button
                        type="button"
                        onClick={() => handleDelete(cat.id)}
                        className="px-3 py-1.5 text-sm rounded-lg border border-red-200 text-red-700 hover:bg-red-50"
                      >
                        Удалить
                      </button>
                    </div>
                  </>
                )}
              </li>
            ))}
          </ul>
        )}
      </section>

      <section className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6">
        <h2 className="text-lg font-semibold text-slate-900 mb-4">Статус заказа</h2>

        {orderMessage ? (
          <p
            className={`mb-4 text-sm rounded-lg px-3 py-2 ${
              orderMessage.type === "ok" ? "bg-emerald-50 text-emerald-900" : "bg-red-50 text-red-800"
            }`}
          >
            {orderMessage.text}
          </p>
        ) : null}

        <form onSubmit={handleOrderStatus} className="flex flex-col gap-5">
          <label className="block max-w-xs">
            <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">№ заказа</span>
            <input
              type="number"
              min={1}
              required
              value={orderIdInput}
              onChange={(e) => setOrderIdInput(e.target.value)}
              className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2"
            />
          </label>

          <div>
            <div className="flex flex-wrap items-center justify-between gap-2 mb-3">
              <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Новый статус</span>
              <OrderStatusBadge status={orderStatus} size="sm" />
            </div>
            <OrderStatusPicker value={orderStatus} onChange={setOrderStatus} disabled={orderUpdating} />
          </div>

          <button
            type="submit"
            disabled={orderUpdating}
            className="self-start px-5 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 disabled:opacity-60"
          >
            {orderUpdating ? "Обновление…" : "Применить статус"}
          </button>
        </form>
      </section>
    </div>
  );
}
