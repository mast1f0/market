import { NavLink } from "react-router-dom";
import { useAuth } from "../auth/useAuth.ts";

const linkClass = ({ isActive }: { isActive: boolean }) =>
  `text-sm font-medium transition-colors ${
    isActive ? "text-emerald-700" : "text-slate-600 hover:text-emerald-700"
  }`;

const authLinkClass = ({ isActive }: { isActive: boolean }) =>
  `text-sm font-medium rounded-lg px-3 py-1.5 transition-colors ${
    isActive
      ? "bg-emerald-50 text-emerald-700"
      : "bg-emerald-600 text-white hover:bg-emerald-700"
  }`;

export default function Navbar() {
  const { token, profile } = useAuth();
  const role = profile?.role;
  const isSeller = role === "seller" || role === "admin";
  const isAdmin = role === "admin";

  return (
    <header className="sticky top-0 z-20 border-b border-slate-200/80 bg-white/90 backdrop-blur-md">
      <div className="max-w-7xl mx-auto flex items-center justify-between gap-4 px-4 md:px-8 py-3">
        <NavLink to="/" className="flex items-center gap-2 shrink-0" aria-label="На главную">
          <span className="flex h-9 w-9 items-center justify-center rounded-lg bg-emerald-600 text-white text-sm font-bold">
            М
          </span>
          <span className="hidden sm:inline font-semibold text-slate-900">Маркет</span>
        </NavLink>

        <nav className="flex flex-wrap items-center justify-end gap-x-4 md:gap-x-5 gap-y-2">
          <NavLink to="/" end className={linkClass}>
            Главная
          </NavLink>
          <NavLink to="/categories" className={linkClass}>
            Категории
          </NavLink>
          <NavLink to="/about" className={linkClass}>
            О нас
          </NavLink>
          <NavLink to="/cart" className={linkClass}>
            Корзина
          </NavLink>
          {isSeller ? (
            <NavLink to="/seller-panel" className={linkClass}>
              Продавец
            </NavLink>
          ) : null}
          {isAdmin ? (
            <NavLink to="/admin" className={linkClass}>
              Админ
            </NavLink>
          ) : null}
          {token ? (
            <NavLink to="/profiles" className={linkClass}>
              Профиль
            </NavLink>
          ) : (
            <NavLink to="/login" className={authLinkClass}>
              Войти
            </NavLink>
          )}
        </nav>
      </div>
    </header>
  );
}
