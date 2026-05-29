import { useCallback, useEffect, useState } from "react";
import { useAuth } from "../auth/useAuth.ts";
import {
  AUTH_ROLE_LABELS,
  AUTH_ROLES,
  deleteAuthUser,
  fetchAuthUsers,
  formatAuthUserDate,
  updateAuthUserRole,
  type AuthUser,
} from "../lib/auth-users.ts";

export default function AdminUsers() {
  const { profile } = useAuth();
  const [users, setUsers] = useState<AuthUser[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [roleDraft, setRoleDraft] = useState<Record<number, string>>({});
  const [savingId, setSavingId] = useState<number | null>(null);
  const [deletingId, setDeletingId] = useState<number | null>(null);

  const loadUsers = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const list = await fetchAuthUsers();
      setUsers(list);
      setRoleDraft(Object.fromEntries(list.map((u) => [u.id, u.role])));
    } catch (e) {
      setUsers([]);
      setError(e instanceof Error ? e.message : "Не удалось загрузить пользователей");
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    loadUsers();
  }, [loadUsers]);

  const handleSaveRole = async (user: AuthUser) => {
    const role = roleDraft[user.id] ?? user.role;
    if (role === user.role) return;
    setSavingId(user.id);
    try {
      await updateAuthUserRole(user.id, role);
      await loadUsers();
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось обновить роль");
    } finally {
      setSavingId(null);
    }
  };

  const handleDelete = async (user: AuthUser) => {
    if (profile?.user_id === user.id) {
      alert("Нельзя удалить свой аккаунт из этой панели.");
      return;
    }
    if (!confirm(`Удалить пользователя «${user.login}»?`)) return;
    setDeletingId(user.id);
    try {
      await deleteAuthUser(user.id);
      await loadUsers();
    } catch (e) {
      alert(e instanceof Error ? e.message : "Не удалось удалить пользователя");
    } finally {
      setDeletingId(null);
    }
  };

  return (
    <section className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6 mb-8">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-4">
        <div>
          <h2 className="text-lg font-semibold text-slate-900">Пользователи</h2>
          <p className="text-sm text-slate-500 mt-0.5">Учётные записи и роли</p>
        </div>
        <button
          type="button"
          onClick={loadUsers}
          disabled={loading}
          className="px-3 py-1.5 text-sm rounded-lg border border-slate-200 hover:bg-slate-50 disabled:opacity-60"
        >
          Обновить
        </button>
      </div>

      {loading ? (
        <p className="text-slate-600 text-sm">Загрузка…</p>
      ) : error ? (
        <p className="text-red-600 text-sm">{error}</p>
      ) : users.length === 0 ? (
        <p className="text-slate-600 text-sm">Пользователей нет.</p>
      ) : (
        <div className="overflow-x-auto -mx-2">
          <table className="w-full min-w-[32rem] text-sm">
            <thead>
              <tr className="text-left text-xs text-slate-500 border-b border-slate-100">
                <th className="pb-2 px-2 font-medium">ID</th>
                <th className="pb-2 px-2 font-medium">Логин</th>
                <th className="pb-2 px-2 font-medium">Роль</th>
                <th className="pb-2 px-2 font-medium">Создан</th>
                <th className="pb-2 px-2 font-medium text-right">Действия</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-slate-50">
              {users.map((user) => {
                const draft = roleDraft[user.id] ?? user.role;
                const roleChanged = draft !== user.role;
                const isSelf = profile?.user_id === user.id;
                return (
                  <tr key={user.id}>
                    <td className="py-3 px-2 text-slate-500">{user.id}</td>
                    <td className="py-3 px-2 font-medium text-slate-900">
                      {user.login}
                      {isSelf ? (
                        <span className="ml-2 text-xs text-emerald-700">(вы)</span>
                      ) : null}
                    </td>
                    <td className="py-3 px-2">
                      <select
                        value={draft}
                        onChange={(e) =>
                          setRoleDraft((prev) => ({ ...prev, [user.id]: e.target.value }))
                        }
                        className="border border-slate-200 rounded-lg px-2 py-1.5 text-sm min-w-[8rem]"
                      >
                        {AUTH_ROLES.map((r) => (
                          <option key={r} value={r}>
                            {AUTH_ROLE_LABELS[r] ?? r}
                          </option>
                        ))}
                      </select>
                    </td>
                    <td className="py-3 px-2 text-slate-500 whitespace-nowrap">
                      {formatAuthUserDate(user.created_at)}
                    </td>
                    <td className="py-3 px-2">
                      <div className="flex justify-end gap-2">
                        <button
                          type="button"
                          disabled={!roleChanged || savingId === user.id}
                          onClick={() => handleSaveRole(user)}
                          className="px-2.5 py-1.5 text-xs rounded-lg bg-emerald-600 text-white hover:bg-emerald-700 disabled:opacity-40"
                        >
                          {savingId === user.id ? "…" : "Роль"}
                        </button>
                        <button
                          type="button"
                          disabled={isSelf || deletingId === user.id}
                          onClick={() => handleDelete(user)}
                          className="px-2.5 py-1.5 text-xs rounded-lg border border-red-200 text-red-700 hover:bg-red-50 disabled:opacity-40"
                        >
                          {deletingId === user.id ? "…" : "Удалить"}
                        </button>
                      </div>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
