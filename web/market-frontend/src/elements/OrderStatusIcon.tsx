import type { OrderStatus } from "../lib/orders.ts";

type Props = {
  status: string;
  className?: string;
};

export default function OrderStatusIcon({ status, className = "h-4 w-4" }: Props) {
  const common = `${className} shrink-0`;

  switch (status as OrderStatus) {
    case "pending":
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <circle cx="12" cy="12" r="9" />
          <path d="M12 7v5l3 2" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      );
    case "processing":
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <path d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16V8z" />
          <path d="M3.3 7.7 12 12.5l8.7-4.8M12 22V12.5" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      );
    case "shipped":
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <path d="M14 18V6a2 2 0 0 0-2-2H4v16h10a2 2 0 0 0 2-2z" />
          <path d="M14 8h4l3 4v6h-7V8z" strokeLinecap="round" strokeLinejoin="round" />
          <circle cx="7" cy="18" r="2" />
          <circle cx="17" cy="18" r="2" />
        </svg>
      );
    case "delivered":
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" strokeLinecap="round" />
          <path d="m22 4-10 10.01-3-3" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      );
    case "cancelled":
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <circle cx="12" cy="12" r="9" />
          <path d="m15 9-6 6M9 9l6 6" strokeLinecap="round" />
        </svg>
      );
    default:
      return (
        <svg className={common} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" aria-hidden>
          <circle cx="12" cy="12" r="9" />
          <path d="M12 8v4M12 16h.01" strokeLinecap="round" />
        </svg>
      );
  }
}
