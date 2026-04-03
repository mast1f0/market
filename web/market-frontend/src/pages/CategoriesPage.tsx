import CategoriesGrid from "../widgets/CategoriesGrid.tsx";

export default function CategoriesPage() {
    const categories = [
        {
            id: 1,
            name: "Электроника",
            description: "Смартфоны, ноутбуки и вся электронная техника",
        },
        {
            id: 2,
            name: "Одежда",
            description: "Футболки, джинсы, куртки и аксессуары",
        },
        {
            id: 3,
            name: "Книги",
            description: "Художественная литература, учебники и комиксы",
        },
        {
            id: 4,
            name: "Игрушки",
            description: "Для детей всех возрастов",
        },
        {
            id: 5,
            name: "Дом и сад",
            description: "Мебель, декор, инструменты и растения",
        },
    ];

    return <CategoriesGrid categories={categories} />;
}