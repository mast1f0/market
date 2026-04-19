type CategoryCardProps = {
  name: string;
  description?: string;
  onClick?: () => void;
};

export default function CategoryCard({ name, description, onClick }: CategoryCardProps) {
  const subtitle = description?.trim() || "Товары этой категории в общем каталоге на главной.";

  return (
    <div
      role="button"
      tabIndex={0}
      onClick={onClick}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onClick?.();
        }
      }}
      className="cursor-pointer bg-white rounded-xl p-5 border border-slate-100 shadow-sm hover:shadow-md hover:border-slate-200 transition-all duration-200"
    >
      <h2 className="text-lg font-semibold text-slate-900 mb-2">{name}</h2>
      <p className="text-slate-600 text-sm leading-relaxed">{subtitle}</p>
    </div>
  );
}
