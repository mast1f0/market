/** Парсит цену из поля ввода (поддерживает запятую как разделитель). */
export function parsePrice(value: string): number {
  const normalized = value.trim().replace(",", ".");
  if (!normalized) return 0;
  const n = parseFloat(normalized);
  return Number.isFinite(n) ? n : 0;
}

export function formatRub(amount: number): string {
  const fractionDigits = Math.round(amount * 100) % 100 === 0 ? 0 : 2;
  return new Intl.NumberFormat("ru-RU", {
    style: "currency",
    currency: "RUB",
    minimumFractionDigits: fractionDigits,
    maximumFractionDigits: 2,
  }).format(amount);
}
