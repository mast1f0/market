import { Link } from "react-router-dom";
import ResolvedImage from "./ResolvedImage.tsx";
import OrderStatusBadge from "./OrderStatusBadge.tsx";
import { formatRub } from "../lib/format.ts";
import { formatOrderDate, getOrderStatusMeta, orderPreviewImages, type OrderDTO } from "../lib/orders.ts";

type Props = {
  order: OrderDTO;
};

export default function OrderCard({ order }: Props) {
  const itemCount = order.items.reduce((s, i) => s + i.quantity, 0);
  const previews = orderPreviewImages(order);
  const extraCount = Math.max(0, order.items.length - previews.length);
  const meta = getOrderStatusMeta(order.status);

  return (
    <Link
      to={`/orders/${order.id}`}
      className="block rounded-xl border border-slate-100 bg-white p-5 shadow-sm hover:border-slate-200 hover:shadow-md transition-all group"
    >
      <div className="flex flex-wrap items-start justify-between gap-3">
        <div className="min-w-0">
          <p className="font-semibold text-slate-900 group-hover:text-emerald-800 transition-colors">
            Заказ №{order.id}
          </p>
          <p className="mt-1 text-sm text-slate-500">{formatOrderDate(order.created_at)}</p>
        </div>
        <OrderStatusBadge status={order.status} size="sm" />
      </div>

      {order.status !== "cancelled" ? (
        <p className="mt-3 text-xs text-slate-500 line-clamp-1">{meta.description}</p>
      ) : null}

      {previews.length > 0 ? (
        <div className="mt-4 flex items-center gap-2">
          {previews.map((img, i) => (
            <div
              key={`${order.id}-${img}-${i}`}
              className="h-12 w-12 shrink-0 overflow-hidden rounded-lg bg-slate-100 ring-1 ring-slate-100"
            >
              <ResolvedImage imageRef={img} alt="" className="h-full w-full object-cover" />
            </div>
          ))}
          {extraCount > 0 ? (
            <span className="text-xs font-medium text-slate-500">+{extraCount}</span>
          ) : null}
        </div>
      ) : null}

      <p className="mt-4 text-sm text-slate-600">
        {itemCount} {itemCount === 1 ? "товар" : itemCount < 5 ? "товара" : "товаров"}
      </p>
      <p className="mt-2 text-lg font-semibold text-emerald-700">{formatRub(order.total_price)}</p>
    </Link>
  );
}
