import { useCallback, useEffect, useMemo, useState, type ReactNode } from "react";
import { authApiUrl } from "../lib/endpoints.ts";
import { ACCESS_TOKEN_KEY, LOGIN_DISPLAY_KEY } from "../lib/token.ts";
import { AuthContext, type AuthProfile } from "./context.ts";

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem(ACCESS_TOKEN_KEY));
  const [profile, setProfile] = useState<AuthProfile | null>(null);
  const [profileLoading, setProfileLoading] = useState(() => !!localStorage.getItem(ACCESS_TOKEN_KEY));

  const logout = useCallback(() => {
    localStorage.removeItem(ACCESS_TOKEN_KEY);
    localStorage.removeItem(LOGIN_DISPLAY_KEY);
    setToken(null);
    setProfile(null);
    setProfileLoading(false);
  }, []);

  const login = useCallback(async (loginStr: string, password: string) => {
    const res = await fetch(authApiUrl("/login"), {
      method: "POST",
      headers: { "Content-Type": "application/json", Accept: "application/json" },
      body: JSON.stringify({ login: loginStr, password }),
    });
    const text = await res.text();
    let access_token: string | undefined;
    try {
      access_token = (JSON.parse(text) as { access_token?: string }).access_token;
    } catch {
      /* ответ не JSON */
    }
    if (!access_token) {
      throw new Error((text || `Ошибка входа (${res.status})`).trim().slice(0, 240));
    }
    localStorage.setItem(ACCESS_TOKEN_KEY, access_token);
    localStorage.setItem(LOGIN_DISPLAY_KEY, loginStr);
    setToken(access_token);
  }, []);

  const register = useCallback(async (loginStr: string, password: string) => {
    const res = await fetch(authApiUrl("/register"), {
      method: "POST",
      headers: { "Content-Type": "application/json", Accept: "application/json" },
      body: JSON.stringify({ login: loginStr, password }),
    });
    const text = await res.text();
    if (!res.ok) {
      throw new Error((text || `Ошибка регистрации (${res.status})`).trim().slice(0, 240));
    }
  }, []);

  useEffect(() => {
    if (!token) {
      setProfile(null);
      setProfileLoading(false);
      return;
    }

    let cancelled = false;
    setProfileLoading(true);
    (async () => {
      try {
        const res = await fetch(authApiUrl("/profile"), {
          headers: { Authorization: `Bearer ${token}`, Accept: "application/json" },
        });
        const text = await res.text();
        if (cancelled) return;
        if (!res.ok) {
          setProfile(null);
          return;
        }
        const p = JSON.parse(text) as AuthProfile;
        if (typeof p.user_id === "number" && typeof p.role === "string") {
          setProfile(p);
        } else {
          setProfile(null);
        }
      } catch {
        if (!cancelled) setProfile(null);
      } finally {
        if (!cancelled) setProfileLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [token]);

  const value = useMemo(
    () => ({
      token,
      profile,
      profileLoading,
      login,
      register,
      logout,
    }),
    [token, profile, profileLoading, login, register, logout]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
