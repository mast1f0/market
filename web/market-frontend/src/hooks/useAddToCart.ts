import { useCallback } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../auth/useAuth.ts";
import { addProductToCart } from "../lib/cart.ts";

export function useAddToCart() {
  const { token } = useAuth();
  const navigate = useNavigate();

  return useCallback(
    async (productId: number): Promise<boolean> => {
      if (!token) {
        navigate("/login");
        return false;
      }
      try {
        await addProductToCart(productId);
        return true;
      } catch {
        return false;
      }
    },
    [token, navigate]
  );
}
