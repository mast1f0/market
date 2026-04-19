import { Link } from "react-router-dom";
import { useAuth } from "../auth/useAuth.ts";

export default function ProfilePage() {
  const { token, profile, profileLoading, logout } = useAuth();

  if (!token) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-2">Профиль</h1>
        <p className="text-slate-600 mb-6">Вы не вошли в аккаунт.</p>
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
        <h1 className="text-2xl font-bold text-slate-900 mb-4">Профиль</h1>
        <p className="text-slate-600 text-sm">Загрузка данных из auth-microservice…</p>
      </div>
    );
  }

  const role = profile?.role ?? "—";
  const isSeller = role === "seller" || role === "admin";

  return (
    <div className="max-w-3xl mx-auto p-6 md:p-8">
      <h1 className="text-2xl md:text-3xl font-bold text-slate-900 mb-2">Профиль</h1>
      <p className="text-slate-600 text-sm mb-8">Данные из GET /profile (auth-microservice).</p>

      <div className="bg-white rounded-2xl border border-slate-100 shadow-sm p-6 md:p-8 mb-6">
        <dl className="space-y-3 text-slate-700">
          <div>
            <dt className="text-xs font-medium text-slate-500 uppercase tracking-wide">User ID</dt>
            <dd className="mt-0.5 font-medium text-slate-900">{profile?.user_id ?? "—"}</dd>
          </div>
          <div>
            <dt className="text-xs font-medium text-slate-500 uppercase tracking-wide">Роль</dt>
            <dd className="mt-0.5">{role}</dd>
          </div>
        </dl>
      </div>

      <div className="flex flex-wrap gap-3">
        {isSeller ? (
          <Link
            to="/seller-panel"
            className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 transition-colors"
          >
            Панель товаров
          </Link>
        ) : null}
        <Link
          to="/add"
          className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-slate-200 text-slate-800 text-sm font-medium hover:bg-slate-50 transition-colors"
        >
          Новый товар
        </Link>
        <button
          type="button"
          onClick={logout}
          className="inline-flex items-center justify-center px-4 py-2.5 rounded-lg border border-red-200 text-red-700 text-sm font-medium hover:bg-red-50"
        >
          Выйти
        </button>
      </div>

      {!isSeller ? (
        <p className="mt-6 text-sm text-slate-600">
          Создание товаров на маркете требует роли <span className="font-medium">seller</span> или{" "}
          <span className="font-medium">admin</span> в JWT (сейчас у вас роль «{role}»).
        </p>
      ) : null}
    </div>
  );
}
