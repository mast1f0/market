import { useEffect, useMemo, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import OrderCard from "../elements/OrderCard.tsx";
import OrderSearchBar from "../elements/OrderSearchBar.tsx";
import { useAuth } from "../auth/useAuth.ts";
import { filterOrders } from "../lib/order-search.ts";
import { fetchOrders, type OrderDTO } from "../lib/orders.ts";
import { LOGIN_DISPLAY_KEY } from "../lib/token.ts";

const ROLE_LABELS: Record<string, string> = {
  buyer: "Покупатель",
  seller: "Продавец",
  admin: "Администратор",
};

function roleLabel(role: string): string {
  return ROLE_LABELS[role] ?? role;
}

export default function ProfilePage() {
  const { token, profile, profileLoading, logout } = useAuth();
  const [searchParams, setSearchParams] = useSearchParams();
  const [orders, setOrders] = useState<OrderDTO[]>([]);
  const [ordersLoading, setOrdersLoading] = useState(false);
  const [ordersError, setOrdersError] = useState<string | null>(null);
  const [orderQuery, setOrderQuery] = useState("");

  const justOrdered = searchParams.get("ordered") === "1";
  const displayLogin = localStorage.getItem(LOGIN_DISPLAY_KEY);

  useEffect(() => {
    if (!token) return;

    let cancelled = false;
    setOrdersLoading(true);
    setOrdersError(null);

    (async () => {
      try {
        const list = await fetchOrders();
        if (!cancelled) setOrders(list);
      } catch (e) {
        if (!cancelled) {
          setOrdersError(e instanceof Error ? e.message : "Не удалось загрузить заказы");
        }
      } finally {
        if (!cancelled) setOrdersLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [token]);

  useEffect(() => {
    if (!justOrdered) return;
    const t = window.setTimeout(() => {
      searchParams.delete("ordered");
      setSearchParams(searchParams, { replace: true });
    }, 6000);
    return () => window.clearTimeout(t);
  }, [justOrdered, searchParams, setSearchParams]);

  if (!token) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-2">Профиль</h1>
        <p className="text-slate-600 mb-6">Войдите, чтобы управлять аккаунтом и смотреть заказы.</p>
        <Link
          to="/login"
          className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700"
        >
          Войти
        </Link>
      </div>
    );
  }

  if (profileLoading) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <div className="h-10 w-48 bg-slate-200 rounded animate-pulse mb-8" />
        <div className="h-32 bg-slate-100 rounded-2xl animate-pulse" />
      </div>
    );
  }

  const role = profile?.role ?? "buyer";
  const isSeller = role === "seller" || role === "admin";
  const isAdmin = role === "admin";
  const initial = (displayLogin?.[0] ?? "U").toUpperCase();

  const filteredOrders = useMemo(() => filterOrders(orders, orderQuery), [orders, orderQuery]);

  return (
    <div className="max-w-3xl mx-auto p-6 md:p-8">
      {justOrdered ? (
        <p className="mb-6 text-sm text-emerald-800 bg-emerald-50 border border-emerald-100 rounded-lg px-4 py-3">
          Заказ оформлен. Детали — ниже в списке заказов.
        </p>
      ) : null}

      <div className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6 md:p-8 mb-8">
        <div className="flex items-start gap-4">
          <span className="flex h-14 w-14 shrink-0 items-center justify-center rounded-2xl bg-emerald-600 text-xl font-bold text-white">
            {initial}
          </span>
          <div className="min-w-0 flex-1">
            <h1 className="text-2xl md:text-3xl font-bold text-slate-900 truncate">
              {displayLogin ?? `Пользователь #${profile?.user_id ?? "—"}`}
            </h1>
            <p className="mt-1 text-sm text-slate-500">ID {profile?.user_id ?? "—"}</p>
            <span className="mt-3 inline-flex rounded-full bg-slate-100 px-3 py-0.5 text-xs font-medium text-slate-700">
              {roleLabel(role)}
            </span>
          </div>
        </div>

        <div className="mt-6 flex flex-wrap gap-3 pt-6 border-t border-slate-100">
          <Link
            to="/cart"
            className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-slate-200 text-slate-800 text-sm font-medium hover:bg-slate-50"
          >
            Корзина
          </Link>
          <Link
            to="/categories"
            className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-slate-200 text-slate-800 text-sm font-medium hover:bg-slate-50"
          >
            Каталог
          </Link>
          {isSeller ? (
            <>
              <Link
                to="/seller-panel"
                className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700"
              >
                Панель товаров
              </Link>
              <Link
                to="/add"
                className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-slate-200 text-slate-800 text-sm font-medium hover:bg-slate-50"
              >
                Новый товар
              </Link>
            </>
          ) : null}
          {isAdmin ? (
            <Link
              to="/admin"
              className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-violet-200 text-violet-800 text-sm font-medium hover:bg-violet-50"
            >
              Админ-панель
            </Link>
          ) : null}
          <button
            type="button"
            onClick={logout}
            className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-red-200 text-red-700 text-sm font-medium hover:bg-red-50"
          >
            Выйти
          </button>
        </div>
      </div>

      <section>
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-4">
          <h2 className="text-lg font-semibold text-slate-900">Мои заказы</h2>
          {orders.length > 0 ? (
            <p className="text-xs text-slate-500 shrink-0">
              {orderQuery ? `${filteredOrders.length} из ${orders.length}` : `${orders.length} всего`}
            </p>
          ) : null}
        </div>

        {orders.length > 0 ? (
          <OrderSearchBar value={orderQuery} onChange={setOrderQuery} className="mb-4" />
        ) : null}

        {ordersLoading ? (
          <div className="space-y-3">
            {[1, 2].map((i) => (
              <div key={i} className="h-28 rounded-xl bg-slate-100 animate-pulse" />
            ))}
          </div>
        ) : ordersError ? (
          <p className="text-sm text-red-700 bg-red-50 border border-red-100 rounded-lg px-4 py-3">{ordersError}</p>
        ) : orders.length === 0 ? (
          <div className="rounded-xl border border-dashed border-slate-200 bg-white p-10 text-center">
            <p className="text-slate-600 mb-4">Заказов пока нет</p>
            <Link
              to="/"
              className="inline-flex px-4 py-2 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700"
            >
              Перейти в каталог
            </Link>
          </div>
        ) : filteredOrders.length === 0 ? (
          <div className="rounded-xl border border-dashed border-slate-200 bg-white p-8 text-center">
            <p className="text-slate-600">Ничего не найдено по запросу «{orderQuery}»</p>
          </div>
        ) : (
          <div className="space-y-3">
            {filteredOrders.map((order) => (
              <OrderCard key={order.id} order={order} />
            ))}
          </div>
        )}
      </section>
    </div>
  );
}
