import { useCallback, useEffect, useMemo, useState } from "react";
import OrderCard from "../elements/OrderCard.tsx";
import OrderSearchBar from "../elements/OrderSearchBar.tsx";
import { AUTH_ROLE_LABELS, fetchAuthUsers, type AuthUser } from "../lib/auth-users.ts";
import { filterAuthUsers, filterOrders } from "../lib/order-search.ts";
import { fetchOrders, type OrderDTO } from "../lib/orders.ts";

export default function AdminOrderSearch() {
  const [users, setUsers] = useState<AuthUser[]>([]);
  const [usersLoading, setUsersLoading] = useState(true);
  const [usersError, setUsersError] = useState<string | null>(null);

  const [userQuery, setUserQuery] = useState("");
  const [selectedUser, setSelectedUser] = useState<AuthUser | null>(null);

  const [orders, setOrders] = useState<OrderDTO[]>([]);
  const [ordersLoading, setOrdersLoading] = useState(false);
  const [ordersError, setOrdersError] = useState<string | null>(null);
  const [orderQuery, setOrderQuery] = useState("");

  const loadUsers = useCallback(async () => {
    setUsersLoading(true);
    setUsersError(null);
    try {
      setUsers(await fetchAuthUsers());
    } catch (e) {
      setUsers([]);
      setUsersError(e instanceof Error ? e.message : "Не удалось загрузить пользователей");
    } finally {
      setUsersLoading(false);
    }
  }, []);

  useEffect(() => {
    loadUsers();
  }, [loadUsers]);

  const matchedUsers = useMemo(() => filterAuthUsers(users, userQuery).slice(0, 8), [users, userQuery]);

  const loadOrdersForUser = useCallback(async (user: AuthUser) => {
    setSelectedUser(user);
    setOrdersLoading(true);
    setOrdersError(null);
    setOrderQuery("");
    try {
      setOrders(await fetchOrders(user.id));
    } catch (e) {
      setOrders([]);
      setOrdersError(e instanceof Error ? e.message : "Не удалось загрузить заказы");
    } finally {
      setOrdersLoading(false);
    }
  }, []);

  const filteredOrders = useMemo(() => filterOrders(orders, orderQuery), [orders, orderQuery]);

  const pickUserById = () => {
    const id = Number(userQuery.trim());
    if (!id || id < 1) return;
    const found = users.find((u) => u.id === id);
    if (found) {
      loadOrdersForUser(found);
      return;
    }
    loadOrdersForUser({ id, login: `#${id}`, role: "", created_at: "" });
  };

  return (
    <section className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6 mb-8">
      <h2 className="text-lg font-semibold text-slate-900 mb-1">Поиск заказов</h2>
      <p className="text-sm text-slate-500 mb-4">Найдите пользователя, затем отфильтруйте его заказы</p>

      <div className="space-y-3">
        <OrderSearchBar
          value={userQuery}
          onChange={setUserQuery}
          placeholder="Логин, ID или роль пользователя…"
        />

        {usersLoading ? (
          <p className="text-sm text-slate-500">Загрузка пользователей…</p>
        ) : usersError ? (
          <p className="text-sm text-red-600">{usersError}</p>
        ) : null}

        {userQuery.trim() && !usersLoading ? (
          <div className="rounded-xl border border-slate-100 bg-slate-50/80 overflow-hidden">
            {matchedUsers.length === 0 ? (
              <div className="p-4 flex flex-wrap items-center justify-between gap-3">
                <p className="text-sm text-slate-600">Пользователь не найден</p>
                {/^\d+$/.test(userQuery.trim()) ? (
                  <button
                    type="button"
                    onClick={pickUserById}
                    className="text-sm font-medium text-emerald-700 hover:text-emerald-800"
                  >
                    Загрузить заказы ID {userQuery.trim()}
                  </button>
                ) : null}
              </div>
            ) : (
              <ul className="divide-y divide-slate-100">
                {matchedUsers.map((user) => (
                  <li key={user.id}>
                    <button
                      type="button"
                      onClick={() => loadOrdersForUser(user)}
                      className={`w-full flex items-center justify-between gap-3 px-4 py-3 text-left hover:bg-white transition-colors ${
                        selectedUser?.id === user.id ? "bg-white ring-1 ring-inset ring-emerald-200" : ""
                      }`}
                    >
                      <span>
                        <span className="font-medium text-slate-900">{user.login}</span>
                        <span className="ml-2 text-xs text-slate-400">id {user.id}</span>
                      </span>
                      <span className="text-xs rounded-full bg-slate-200/80 px-2 py-0.5 text-slate-700">
                        {AUTH_ROLE_LABELS[user.role] ?? user.role}
                      </span>
                    </button>
                  </li>
                ))}
              </ul>
            )}
          </div>
        ) : null}
      </div>

      {selectedUser ? (
        <div className="mt-6 pt-6 border-t border-slate-100">
          <div className="flex flex-wrap items-center justify-between gap-2 mb-4">
            <p className="text-sm text-slate-700">
              Заказы пользователя{" "}
              <span className="font-semibold text-slate-900">{selectedUser.login || `#${selectedUser.id}`}</span>
            </p>
            <button
              type="button"
              onClick={() => {
                setSelectedUser(null);
                setOrders([]);
                setOrderQuery("");
              }}
              className="text-xs text-slate-500 hover:text-slate-800"
            >
              Сбросить
            </button>
          </div>

          <OrderSearchBar
            value={orderQuery}
            onChange={setOrderQuery}
            placeholder="№ заказа, статус, название товара…"
            className="mb-4"
          />

          {ordersLoading ? (
            <div className="space-y-3">
              {[1, 2].map((i) => (
                <div key={i} className="h-28 rounded-xl bg-slate-100 animate-pulse" />
              ))}
            </div>
          ) : ordersError ? (
            <p className="text-sm text-red-700 bg-red-50 border border-red-100 rounded-lg px-4 py-3">{ordersError}</p>
          ) : filteredOrders.length === 0 ? (
            <p className="text-sm text-slate-600 text-center py-8">
              {orders.length === 0 ? "У пользователя нет заказов" : "Ничего не найдено по запросу"}
            </p>
          ) : (
            <>
              <p className="text-xs text-slate-500 mb-3">
                Показано {filteredOrders.length} из {orders.length}
              </p>
              <div className="space-y-3">
                {filteredOrders.map((order) => (
                  <OrderCard key={order.id} order={order} />
                ))}
              </div>
            </>
          )}
        </div>
      ) : null}
    </section>
  );
}
