import CategoryCard from "../elements/CategoryCard.tsx";

type Category = {
    id: number;
    name: string;
    description: string;
};

type CategoriesGridProps = {
    categories: Category[];
    onCategoryClick?: (category: Category) => void;
};

export default function CategoriesGrid({ categories, onCategoryClick }: CategoriesGridProps) {
    return (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 p-4">
            {categories.map((category) => (
                <CategoryCard
                    key={category.id}
                    name={category.name}
                    description={category.description}
                    onClick={() => onCategoryClick && onCategoryClick(category)}
                />
            ))}
        </div>
    );
}