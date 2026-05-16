import { fetchJson } from "./api.ts";

export interface CartItemDTO {
  id: number;
  product_id: number;
  quantity: number;
  price: number;
  name: string;
  image_url: string;
}

export interface CartDTO {
  id: number;
  items: CartItemDTO[];
}

function pick<T>(obj: Record<string, unknown>, ...keys: string[]): T | undefined {
  for (const k of keys) {
    const v = obj[k];
    if (v !== undefined && v !== null) return v as T;
  }
  return undefined;
}

function cartItemPrice(item: Record<string, unknown>, product?: Record<string, unknown>): number {
  const snapshot = pick<number | string>(item, "price_snapshot", "PriceSnapshot");
  const snapNum = snapshot != null ? Number(snapshot) : 0;
  if (snapNum > 0) return snapNum;

  const fromProduct = product ? pick<number | string>(product, "price", "Price") : undefined;
  if (fromProduct != null) {
    const n = Number(fromProduct);
    if (n > 0) return n;
  }

  const direct = pick<number | string>(item, "price", "Price");
  if (direct != null) {
    const n = Number(direct);
    if (n > 0) return n;
  }

  return 0;
}

export function normalizeCart(raw: Record<string, unknown>): CartDTO {
  const items = (pick<unknown[]>(raw, "items", "Items") ?? []) as Record<string, unknown>[];
  return {
    id: Number(pick(raw, "id", "ID") ?? 0),
    items: items.map((item) => {
      const product = pick<Record<string, unknown>>(item, "product", "Product");
      return {
        id: Number(pick(item, "id", "ID") ?? 0),
        product_id: Number(pick(item, "product_id", "ProductID") ?? 0),
        quantity: Number(pick(item, "quantity", "Quantity") ?? 0),
        price: cartItemPrice(item, product),
        name: String(pick(item, "name") ?? product?.name ?? product?.Name ?? "Без названия"),
        image_url: String(
          pick(item, "image_url") ?? product?.image_url ?? product?.ImageURL ?? ""
        ),
      };
    }),
  };
}

export async function addProductToCart(productId: number, quantity = 1): Promise<void> {
  await fetchJson(
    "/cart/items",
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ product_id: productId, quantity }),
    },
    { auth: true }
  );
}
