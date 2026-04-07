import { useState, useEffect } from "react";
import CategoryCard from "../elements/CategoryCard.tsx";

type Category = {
    id: number;
    name: string;
    description: string;
};

type CategoriesGridProps = {
    onCategoryClick?: (category: Category) => void;
};

export default function CategoriesGrid({ onCategoryClick }: CategoriesGridProps) {
    const [categories, setCategories] = useState<Category[]>([]);

    useEffect(() => {
        const fetchCategories = async () => {
            try {
                const response = await fetch('http://localhost:8080/categories');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setCategories(data);
            } catch (error) {
                console.error('Ошибка загрузки категорий:', error);
                setCategories([]);
            }
        };

        fetchCategories();
    }, []);

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