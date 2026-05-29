import type { AuthUser } from "./auth-users.ts";
import { orderStatusLabel, type OrderDTO } from "./orders.ts";

function norm(s: string): string {
  return s.trim().toLowerCase();
}

/** Клиентский поиск по списку заказов. */
export function filterOrders(orders: OrderDTO[], query: string): OrderDTO[] {
  const q = norm(query);
  if (!q) return orders;

  return orders.filter((order) => {
    if (String(order.id).includes(q)) return true;
    if (String(order.user_id).includes(q)) return true;
    if (norm(order.status).includes(q)) return true;
    if (norm(orderStatusLabel(order.status)).includes(q)) return true;
    if (String(order.total_price).includes(q)) return true;

    return order.items.some((item) => norm(item.name_snapshot).includes(q));
  });
}

/** Поиск пользователей по логину, id или роли. */
export function filterAuthUsers(users: AuthUser[], query: string): AuthUser[] {
  const q = norm(query);
  if (!q) return users;

  return users.filter((user) => {
    if (String(user.id).includes(q)) return true;
    if (norm(user.login).includes(q)) return true;
    if (norm(user.role).includes(q)) return true;
    return false;
  });
}
