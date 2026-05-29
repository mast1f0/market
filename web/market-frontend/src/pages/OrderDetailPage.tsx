import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import OrderItemRow from "../elements/OrderItemRow.tsx";
import OrderStatusBadge from "../elements/OrderStatusBadge.tsx";
import OrderStatusPicker from "../elements/OrderStatusPicker.tsx";
import OrderStatusTimeline from "../elements/OrderStatusTimeline.tsx";
import { bearerHeaders } from "../lib/api.ts";
import { formatRub } from "../lib/format.ts";
import { marketApiUrl } from "../lib/endpoints.ts";
import { formatOrderDate, normalizeOrder, updateOrderStatus, type OrderDTO } from "../lib/orders.ts";
import { useAuth } from "../auth/useAuth.ts";

export default function OrderDetailPage() {
  const { id } = useParams<{ id: string }>();
  const { token, profile } = useAuth();
  const [order, setOrder] = useState<OrderDTO | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [statusDraft, setStatusDraft] = useState("");
  const [statusSaving, setStatusSaving] = useState(false);
  const [statusError, setStatusError] = useState<string | null>(null);

  useEffect(() => {
    if (!token || !id) {
      setLoading(false);
      return;
    }

    let cancelled = false;

    (async () => {
      try {
        const res = await fetch(marketApiUrl(`/orders/${id}`), { headers: bearerHeaders() });
        if (!res.ok) throw new Error(await res.text());
        const raw = await res.json();
        if (!cancelled) {
          const normalized = normalizeOrder(raw as Record<string, unknown>);
          setOrder(normalized);
          setStatusDraft(normalized.status);
        }
      } catch (e) {
        if (!cancelled) {
          setError(e instanceof Error ? e.message : "Не удалось загрузить заказ");
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [token, id]);

  if (!token) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl font-bold text-slate-900 mb-4">Заказ</h1>
        <p className="text-slate-600 mb-6">Войдите, чтобы просматривать заказы.</p>
        <Link
          to="/login"
          className="inline-flex px-4 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700"
        >
          Войти
        </Link>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8 space-y-4">
        <div className="h-4 w-32 bg-slate-200 rounded animate-pulse" />
        <div className="h-24 bg-slate-100 rounded-xl animate-pulse" />
        <div className="h-40 bg-slate-100 rounded-xl animate-pulse" />
      </div>
    );
  }

  if (error || !order) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <p className="text-red-700 mb-4">{error ?? "Заказ не найден"}</p>
        <Link to="/profiles" className="text-sm font-medium text-emerald-700 hover:text-emerald-800">
          ← К профилю
        </Link>
      </div>
    );
  }

  const role = profile?.role ?? "buyer";
  const canManageStatus =
    (role === "seller" || role === "admin") &&
    (role === "admin" || order.user_id === profile?.user_id);

  const handleStatusSave = async () => {
    if (!order || statusDraft === order.status) return;
    setStatusSaving(true);
    setStatusError(null);
    try {
      await updateOrderStatus(order.id, statusDraft);
      setOrder((prev) => (prev ? { ...prev, status: statusDraft } : prev));
    } catch (e) {
      setStatusError(e instanceof Error ? e.message : "Не удалось обновить статус");
    } finally {
      setStatusSaving(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto p-6 md:p-8">
      <Link to="/profiles" className="text-sm font-medium text-emerald-700 hover:text-emerald-800 mb-6 inline-block">
        ← Мои заказы
      </Link>

      <div className="flex flex-wrap items-start justify-between gap-4 mb-6">
        <div>
          <h1 className="text-2xl md:text-3xl font-bold text-slate-900">Заказ №{order.id}</h1>
          <p className="mt-1 text-slate-500 text-sm">{formatOrderDate(order.created_at)}</p>
          {role === "admin" && order.user_id !== profile?.user_id ? (
            <p className="mt-1 text-xs text-violet-700">Покупатель: ID {order.user_id}</p>
          ) : null}
        </div>
        <OrderStatusBadge status={order.status} size="md" />
      </div>

      <div className="mb-8">
        <OrderStatusTimeline status={order.status} />
      </div>

      <ul className="space-y-4">
        {order.items.map((item) => (
          <OrderItemRow key={item.id} item={item} />
        ))}
      </ul>

      {canManageStatus ? (
        <div className="mt-8 p-5 rounded-xl bg-white border border-slate-100 shadow-sm">
          <h2 className="text-sm font-semibold text-slate-900 mb-1">Изменить статус</h2>
          <p className="text-xs text-slate-500 mb-4">Выберите новый этап обработки заказа</p>
          {statusError ? (
            <p className="mb-4 text-sm text-red-700 bg-red-50 border border-red-100 rounded-lg px-3 py-2">
              {statusError}
            </p>
          ) : null}
          <OrderStatusPicker
            value={statusDraft}
            onChange={setStatusDraft}
            disabled={statusSaving}
          />
          <div className="mt-4 flex justify-end">
            <button
              type="button"
              disabled={statusSaving || statusDraft === order.status}
              onClick={handleStatusSave}
              className="px-5 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 disabled:opacity-60 transition-colors"
            >
              {statusSaving ? "Сохранение…" : "Сохранить статус"}
            </button>
          </div>
        </div>
      ) : null}

      <div className="mt-8 flex justify-end p-5 rounded-xl bg-white border border-slate-100 shadow-sm">
        <p className="text-lg font-semibold text-slate-900">
          Итого: <span className="text-emerald-700">{formatRub(order.total_price)}</span>
        </p>
      </div>
    </div>
  );
}
