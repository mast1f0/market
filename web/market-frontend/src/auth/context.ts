import { createContext } from "react";

export type AuthProfile = {
  user_id: number;
  role: string;
};

export type AuthContextValue = {
  token: string | null;
  profile: AuthProfile | null;
  profileLoading: boolean;
  login: (login: string, password: string) => Promise<void>;
  register: (login: string, password: string) => Promise<void>;
  logout: () => void;
};

export const AuthContext = createContext<AuthContextValue | null>(null);
