import { useState, type FormEvent } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../auth/useAuth.ts";

export default function LoginPage() {
  const { login, register } = useAuth();
  const navigate = useNavigate();
  const [mode, setMode] = useState<"login" | "register">("login");
  const [loginField, setLoginField] = useState("");
  const [password, setPassword] = useState("");
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [info, setInfo] = useState<string | null>(null);

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError(null);
    setInfo(null);
    setBusy(true);
    try {
      if (mode === "register") {
        await register(loginField, password);
        setInfo("Аккаунт создан. Войдите с теми же логином и паролем.");
        setMode("login");
        setPassword("");
      } else {
        await login(loginField, password);
        navigate("/profiles", { replace: true });
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Ошибка");
    } finally {
      setBusy(false);
    }
  };

  return (
    <div className="max-w-md mx-auto px-4 py-12 md:py-16">
      <h1 className="text-2xl font-bold text-slate-900">
        {mode === "login" ? "Вход" : "Регистрация"}
      </h1>
      <div className="mt-6 flex rounded-lg border border-slate-200 p-0.5 bg-slate-50">
        <button
          type="button"
          className={`flex-1 py-2 text-sm font-medium rounded-md transition-colors ${
            mode === "login" ? "bg-white text-slate-900 shadow-sm" : "text-slate-600"
          }`}
          onClick={() => {
            setMode("login");
            setError(null);
            setInfo(null);
          }}
        >
          Вход
        </button>
        <button
          type="button"
          className={`flex-1 py-2 text-sm font-medium rounded-md transition-colors ${
            mode === "register" ? "bg-white text-slate-900 shadow-sm" : "text-slate-600"
          }`}
          onClick={() => {
            setMode("register");
            setError(null);
            setInfo(null);
          }}
        >
          Регистрация
        </button>
      </div>

      {error ? (
        <p className="mt-4 text-sm text-red-700 bg-red-50 border border-red-100 rounded-lg px-3 py-2">{error}</p>
      ) : null}
      {info ? (
        <p className="mt-4 text-sm text-emerald-800 bg-emerald-50 border border-emerald-100 rounded-lg px-3 py-2">
          {info}
        </p>
      ) : null}

      <form onSubmit={onSubmit} className="mt-6 space-y-4 bg-white border border-slate-100 rounded-2xl shadow-sm p-6">
        <label className="block">
          <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Логин</span>
          <input
            required
            autoComplete="username"
            value={loginField}
            onChange={(e) => setLoginField(e.target.value)}
            className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
          />
        </label>
        <label className="block">
          <span className="text-xs font-medium text-slate-500 uppercase tracking-wide">Пароль</span>
          <input
            required
            type="password"
            autoComplete={mode === "login" ? "current-password" : "new-password"}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="mt-1 w-full border border-slate-200 rounded-lg px-3 py-2 text-slate-900 focus:outline-none focus:ring-2 focus:ring-emerald-500/30 focus:border-emerald-500"
          />
        </label>
        <button
          type="submit"
          disabled={busy}
          className="w-full py-2.5 rounded-lg bg-emerald-600 text-white font-medium hover:bg-emerald-700 disabled:opacity-60"
        >
          {busy ? "Подождите…" : mode === "login" ? "Войти" : "Зарегистрироваться"}
        </button>
      </form>

      <p className="mt-6 text-center text-sm text-slate-600">
        <Link to="/" className="text-emerald-700 hover:text-emerald-800 font-medium">
          ← На главную
        </Link>
      </p>
    </div>
  );
}
