import { getOrderStatusMeta } from "../lib/orders.ts";
import OrderStatusIcon from "./OrderStatusIcon.tsx";
import { STATUS_TONE_STYLES } from "./order-status-styles.ts";

type Props = {
  status: string;
  size?: "sm" | "md";
  showIcon?: boolean;
};

export default function OrderStatusBadge({ status, size = "md", showIcon = true }: Props) {
  const meta = getOrderStatusMeta(status);
  const styles = STATUS_TONE_STYLES[meta.tone];

  const sizeClasses =
    size === "sm"
      ? "gap-1.5 px-2.5 py-0.5 text-xs"
      : "gap-2 px-3 py-1 text-sm";

  const iconSize = size === "sm" ? "h-3.5 w-3.5" : "h-4 w-4";

  return (
    <span
      className={`inline-flex items-center font-medium rounded-full ring-1 ring-inset ${sizeClasses} ${styles.badge}`}
      title={meta.description}
    >
      {showIcon ? (
        <span className={styles.icon}>
          <OrderStatusIcon status={status} className={iconSize} />
        </span>
      ) : (
        <span className={`h-1.5 w-1.5 rounded-full ${styles.dot}`} aria-hidden />
      )}
      {meta.label}
    </span>
  );
}
