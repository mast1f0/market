import { bearerHeaders } from "./api.ts";
import { marketApiUrl } from "./endpoints.ts";

export interface OrderItemDTO {
  id: number;
  product_id: number;
  quantity: number;
  price_snapshot: number;
  name_snapshot: string;
  image_snapshot: string;
}

export interface OrderDTO {
  id: number;
  user_id: number;
  status: string;
  total_price: number;
  created_at: string;
  items: OrderItemDTO[];
}

export const ORDER_STATUSES = [
  "pending",
  "processing",
  "shipped",
  "delivered",
  "cancelled",
] as const;

export type OrderStatus = (typeof ORDER_STATUSES)[number];

/** Успешный путь доставки (для таймлайна). */
export const ORDER_FLOW_STEPS = ["pending", "processing", "shipped", "delivered"] as const;

export type OrderStatusTone = "amber" | "sky" | "violet" | "emerald" | "rose" | "slate";

export type OrderStatusMeta = {
  label: string;
  description: string;
  tone: OrderStatusTone;
  /** Индекс в ORDER_FLOW_STEPS или null (отмена / неизвестный). */
  flowIndex: number | null;
};

const STATUS_META: Record<string, OrderStatusMeta> = {
  pending: {
    label: "Ожидает",
    description: "Заказ принят и ждёт обработки",
    tone: "amber",
    flowIndex: 0,
  },
  processing: {
    label: "В обработке",
    description: "Собираем и готовим к отправке",
    tone: "sky",
    flowIndex: 1,
  },
  shipped: {
    label: "Отправлен",
    description: "Передан в доставку",
    tone: "violet",
    flowIndex: 2,
  },
  delivered: {
    label: "Доставлен",
    description: "Заказ получен",
    tone: "emerald",
    flowIndex: 3,
  },
  cancelled: {
    label: "Отменён",
    description: "Заказ отменён",
    tone: "rose",
    flowIndex: null,
  },
};

const FALLBACK_META: OrderStatusMeta = {
  label: "Неизвестно",
  description: "",
  tone: "slate",
  flowIndex: null,
};

export function getOrderStatusMeta(status: string): OrderStatusMeta {
  return STATUS_META[status] ?? { ...FALLBACK_META, label: status };
}

export function orderStatusLabel(status: string): string {
  return getOrderStatusMeta(status).label;
}

/** @deprecated Используйте getOrderStatusMeta().tone */
export function orderStatusTone(status: string): "neutral" | "active" | "success" | "danger" {
  const tone = getOrderStatusMeta(status).tone;
  switch (tone) {
    case "emerald":
      return "success";
    case "rose":
      return "danger";
    case "sky":
    case "violet":
      return "active";
    default:
      return "neutral";
  }
}

export function isOrderFlowStatus(status: string): status is (typeof ORDER_FLOW_STEPS)[number] {
  return (ORDER_FLOW_STEPS as readonly string[]).includes(status);
}

function pick<T>(obj: Record<string, unknown>, ...keys: string[]): T | undefined {
  for (const k of keys) {
    if (obj[k] !== undefined && obj[k] !== null) return obj[k] as T;
  }
  return undefined;
}

function orderItemImage(raw: Record<string, unknown>): string {
  const snap = pick<string>(raw, "image_snapshot", "ImageSnapshot");
  if (snap) return snap;
  const product = (pick<Record<string, unknown>>(raw, "product", "Product") ?? {}) as Record<
    string,
    unknown
  >;
  return String(pick(product, "image_url", "ImageURL") ?? "");
}

function orderItemName(raw: Record<string, unknown>): string {
  const snap = pick<string>(raw, "name_snapshot", "NameSnapshot");
  if (snap) return snap;
  const product = (pick<Record<string, unknown>>(raw, "product", "Product") ?? {}) as Record<
    string,
    unknown
  >;
  return String(pick(product, "name", "Name") ?? "Товар");
}

function normalizeOrderItem(raw: Record<string, unknown>): OrderItemDTO {
  return {
    id: Number(pick(raw, "id", "ID") ?? 0),
    product_id: Number(pick(raw, "product_id", "ProductID") ?? 0),
    quantity: Number(pick(raw, "quantity", "Quantity") ?? 0),
    price_snapshot: Number(pick(raw, "price_snapshot", "PriceSnapshot") ?? 0),
    name_snapshot: orderItemName(raw),
    image_snapshot: orderItemImage(raw),
  };
}

/** Уникальные превью позиций заказа (для карточки в списке). */
export function orderPreviewImages(order: OrderDTO, limit = 4): string[] {
  const seen = new Set<string>();
  const out: string[] = [];
  for (const item of order.items) {
    const img = item.image_snapshot.trim();
    if (!img || seen.has(img)) continue;
    seen.add(img);
    out.push(img);
    if (out.length >= limit) break;
  }
  return out;
}

export function normalizeOrder(raw: Record<string, unknown>): OrderDTO {
  const itemsRaw = (pick<unknown[]>(raw, "items", "Items") ?? []) as Record<string, unknown>[];
  const created = pick<string>(raw, "created_at", "CreatedAt") ?? "";
  return {
    id: Number(pick(raw, "id", "ID") ?? 0),
    user_id: Number(pick(raw, "user_id", "UserId") ?? 0),
    status: String(pick(raw, "status", "Status") ?? "pending"),
    total_price: Number(pick(raw, "total_price", "TotalPrice") ?? 0),
    created_at: created,
    items: itemsRaw.map(normalizeOrderItem),
  };
}

export function normalizeOrders(raw: unknown): OrderDTO[] {
  if (!Array.isArray(raw)) return [];
  return raw.map((o) => normalizeOrder(o as Record<string, unknown>));
}

export async function fetchOrders(userId?: number): Promise<OrderDTO[]> {
  const qs = userId != null && userId > 0 ? `?user_id=${userId}` : "";
  const res = await fetch(marketApiUrl(`/orders${qs}`), {
    headers: { Accept: "application/json", ...bearerHeaders() },
  });
  const text = await res.text();
  if (!res.ok) {
    throw new Error(text || `Ошибка ${res.status}`);
  }
  const raw = JSON.parse(text) as unknown;
  return normalizeOrders(raw);
}

export async function updateOrderStatus(orderId: number, status: string): Promise<void> {
  const res = await fetch(marketApiUrl(`/orders/${orderId}`), {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...bearerHeaders(),
    },
    body: JSON.stringify({ status }),
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(text || `Ошибка ${res.status}`);
  }
}

export function formatOrderDate(iso: string): string {
  if (!iso) return "—";
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return "—";
  return new Intl.DateTimeFormat("ru-RU", {
    day: "numeric",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  }).format(d);
}
