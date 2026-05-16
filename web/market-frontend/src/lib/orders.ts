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

const STATUS_LABELS: Record<string, string> = {
  pending: "Ожидает",
  processing: "В обработке",
  shipped: "Отправлен",
  delivered: "Доставлен",
  cancelled: "Отменён",
};

export function orderStatusLabel(status: string): string {
  return STATUS_LABELS[status] ?? status;
}

export function orderStatusTone(status: string): "neutral" | "active" | "success" | "danger" {
  switch (status) {
    case "delivered":
      return "success";
    case "cancelled":
      return "danger";
    case "processing":
    case "shipped":
      return "active";
    default:
      return "neutral";
  }
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
